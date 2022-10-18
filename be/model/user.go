package model

import (
	"errors"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"log"
	"time"
)

// gorm的数据库 表命名是小写的蛇形命名法， 且为复数

const (
	CustomerRole = "customer"
	AdminRole    = "admin"
)

// 用户信息
type User struct {
	ID        uint
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`

	OpenID string `gorm:"primarykey;uniqueIndex;not null"`

	//Username  string //
	//Gender    int    // 性别
	//AvatarUrl string // 头像url
	Role string // "admin" 为管理员， “customer”为客户

	// manytomany
	Magazines []Magazine `gorm:"many2many:user_magazines;"` //订阅期刊
	Keywords  []Keyword  `gorm:"many2many:user_keywords;"`  // 订阅关键词

	// has-many
	FeedBacks        []Feedback        `gorm:"foreignKey:UserRefer;references:OpenID"`
	SharePapers      []SharePaper      `gorm:"foreignKey:UserRefer;references:OpenID"` // 历史文章分享
	Favorites        []Favorite        `gorm:"foreignKey:UserRefer;references:OpenID"` //收藏
	FavoritesFolders []FavoritesFolder `gorm:"foreignKey:UserRefer;references:OpenID"` //收藏夹
}

// 用户的分享记录
type SharePaper struct {
	gorm.Model

	// 属于User, User have many SharePapers
	UserRefer string

	// 外键关联Paper，belongs to关系
	PaperID uint
	Paper   Paper `gorm:"foreignKey:PaperID;"`
}

// 新的写法
type UserModel struct {
	db *gorm.DB
}

func GetUserModel() *UserModel {
	return &UserModel{db_gorm}
}

// 以下是一些数据库操作

//
// 				DEBUG 用
//
func (m *UserModel) GetAllUsers() ([]User, error) {
	var users []User
	result := m.db.Preload(clause.Associations).Find(&users)
	if result.Error != nil {
		return nil, result.Error
	}

	if result.RowsAffected == 0 {
		return []User{}, nil
	}
	return users, nil
}

//
//
//

// CreateUser create a user
func (m *UserModel) CreateUser(user *User) (*User, error) {
	result := m.db.Create(user)
	return user, result.Error
}

func (m *UserModel) GetOrCreateUser(user *User) error {
	result := m.db.Preload(clause.Associations).FirstOrCreate(user)
	return result.Error
}

// DeleteUserById Delete User By Id
func (m *UserModel) DeleteUserById(id uint) error {
	result := m.db.Where("id = ?", id).Delete(&User{})

	return result.Error
}

// GetUserByUsername 通过username获取普通的用户数据
func (m *UserModel) GetUserByUsername(username string) (User, error) {
	user := User{}

	result := m.db.Where("username = ?", username).First(&user)
	return user, result.Error
}

// QueryUserByOpenid if exists return true
func (m *UserModel) QueryUserByOpenid(openid string) (bool, error) {

	result := m.db.Model(&User{}).Where("open_id = ?", openid).First(&User{})
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return false, nil
	} else if result.Error != nil {
		return false, result.Error
	}
	return true, nil
}

// GetUserByID 获取普通的用户数据
func (m *UserModel) GetUserByID(id uint) (User, error) {
	user := User{}
	result := m.db.First(&user, id)
	return user, result.Error
}

// GetUserByOpenID 获取普通的用户数据
func (m *UserModel) GetUserByOpenID(openid string) (User, error) {
	user := User{}
	result := m.db.Where("open_id = ?", openid).First(&user)
	return user, result.Error
}

// GetUserByOpenidWithAllInfo 获取完整的User信息
func (m *UserModel) GetUserByOpenidWithAllInfo(openid string) (User, error) {
	user := User{}

	result := m.db.Where("open_id = ?", openid).Preload(clause.Associations).Find(&user)

	return user, result.Error
}

// GetUserByIDWithAllInfo 获取完整的User信息
func (m *UserModel) GetUserByIDWithAllInfo(id uint) (User, error) {
	user := User{}
	result := m.db.Where("id = ?", id).Preload(clause.Associations).Find(&user)
	return user, result.Error
}

// GetKeywords get keywords 用
func (m *UserModel) GetKeywords(openID string) (User, error) {
	user := User{}
	result := m.db.Where("open_id = ?", openID).Preload("Keywords").Find(&user)
	return user, result.Error
}

// GetMagazines get Magazines 用
func (m *UserModel) GetMagazines(openID string) (User, error) {
	user := User{}
	result := m.db.Where("open_id = ?", openID).Preload("Magazines").Find(&user)
	return user, result.Error
}

// GetSharePapers get SharePapers 用
func (m *UserModel) GetSharePapers(OpenID string) (User, error) {
	user := User{}
	result := m.db.Where("open_id = ?", OpenID).Preload("SharePapers").Find(&user)
	return user, result.Error
}

// GetFavorites get Favorites 用
func (m *UserModel) GetFavorites(OpenID string) (User, error) {
	user := User{}
	result := m.db.Where("open_id = ?", OpenID).Preload("Favorites").Find(&user)
	return user, result.Error
}

// GetFavoritesFolders get FavoritesFolders 用
func (m *UserModel) GetFavoritesFolders(OpenID string) ([]FavoritesFolder, error) {
	user := User{}
	result := m.db.Where("open_id = ?", OpenID).Preload("FavoritesFolders.Favorites.Paper.Magazine").Preload("FavoritesFolders").Find(&user)
	if result.Error != nil {
		return []FavoritesFolder{}, result.Error
	}
	fs := user.FavoritesFolders
	return fs, nil
}

// UpdateUserBySave update
func (m *UserModel) UpdateUserBySave(user *User) error {
	result := m.db.Save(user)
	return result.Error
}

func (m *UserModel) SubMagazine(openid, MagazineName string) error {
	err := m.db.Model(&User{OpenID: openid}).Association("Magazines").Append(&Magazine{Name: MagazineName})
	return err
}

func (m *UserModel) DeleteSubMagazine(openid, MagazineName string) error {
	err := m.db.Model(&User{OpenID: openid}).Association("Magazines").Delete(&Magazine{Name: MagazineName})
	return err
}

func (m *UserModel) DeleteSubMagazines(openid string, MagazineNames []string) error {
	var ms []Magazine
	for _, name := range MagazineNames {
		ms = append(ms, Magazine{
			Name: name,
		})
	}
	err := m.db.Model(&User{OpenID: openid}).Association("Magazines").Delete(ms)
	return err
}

func (u *UserModel) IfSubMagazines(openID string, ms []Magazine) []bool {
	yes := make([]bool, len(ms))
	if openID == "" {
		return yes
	}
	for i, m := range ms {
		var result User_Magazines
		rs := u.db.Table("user_magazines").Select("*").Where("user_open_id = ? AND magazine_name = ?", openID, m.Name).First(&result)
		if errors.Is(rs.Error, gorm.ErrRecordNotFound) {
			yes[i] = false
			continue
		}
		c := rs.RowsAffected
		if c == 1 {
			yes[i] = true
		}
	}
	return yes
}

func (u *UserModel) IfSubMagazine(openID string, ms string) bool {
	if openID == "" {
		return false
	}
	var result User_Magazines
	c := u.db.Table("user_magazines").Select("*").Where("user_open_id = ? AND magazine_name = ?", openID, ms).First(&result).RowsAffected
	if c >= 1 {
		return true
	}
	return false
}

// papers
func (u *UserModel) IfSubPaper(openID string, paperID string) bool {
	if openID == "" {
		return false
	}
	result := u.db.Model(&Favorite{}).Select("*").Where("user_refer = ? AND paper_id = ?", openID, paperID).First(&Favorite{})
	c := result.RowsAffected
	if c >= 1 {
		log.Println("[IfSubPaper]true")
		return true
	}
	log.Println("[IfSubPaper]false")
	return false
}
func (u *UserModel) IfSubPapers(openID string, paperIDs []string) []bool {
	fav := make([]bool, len(paperIDs))
	if openID == "" {
		return fav
	}
	for i, paperID := range paperIDs {
		result := u.db.Model(&Favorite{}).Select("*").Where("user_refer = ? AND paper_id = ?", openID, paperID).First(&Favorite{})
		c := result.RowsAffected
		if c >= 1 {
			//	log.Println("[IfSubPaper]true")
			fav[i] = true
			continue
		}
		//log.Println("[IfSubPaper]false")
		fav[i] = false
	}
	return fav
}

// keywords

func (u *UserModel) SubKeywords(openid string, kws []Keyword) error {
	return u.db.Model(&User{OpenID: openid}).Association("Keywords").Append(kws)
}

func (u *UserModel) IfSubKeywords(openid string, kws []string) []bool {
	subs := make([]bool, len(kws))
	if openid == "" {
		return subs
	}
	for i, expalin := range kws {
		var result User_Keywords
		rs := u.db.Table("user_keywords").Select("*").Where("user_open_id = ? AND keyword_explain = ?", openid, expalin).First(&result)
		if errors.Is(rs.Error, gorm.ErrRecordNotFound) {
			subs[i] = false
			continue
		}
		subs[i] = true
	}
	return subs
}

func (u *UserModel) DeleteKeywords(openid string, kws []Keyword) error {
	err := u.db.Model(&User{OpenID: openid}).Association("Keywords").Delete(kws)
	return err
}
