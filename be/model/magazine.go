package model

import (
	"database/sql/driver"
	"gorm.io/gorm"
	"strings"
	"time"
)

// 期刊

type Magazine struct {
	ID        uint
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`

	Name      string `gorm:"primarykey;uniqueIndex;not null;"` // 期刊名
	EnName    string // 期刊英文名
	Unit      string // 主办单位
	Cycle     string // 出版周期
	Issn      string // ISSN
	Cn        string // CN
	IssueName string // 专辑名称 or 大类
	TopicName string // 专题名称 or 小类

	Cif string // 复合影响因子
	Zif string // 综合影响因子

	Url string // 期刊详情页链接
	Img string // 期刊图片链接

	PaperCount int // 出版文献量
}

type MagazineModel struct {
	db *gorm.DB
}

func GetMagazineModel() *MagazineModel {
	return &MagazineModel{db_gorm}
}

// CreateMagazine 添加一个期刊
func (m *MagazineModel) CreateMagazine(magazine Magazine) error {
	result := m.db.Create(&magazine)
	return result.Error
}

// CreateMagazines 添加多个期刊
func (m *MagazineModel) CreateMagazines(magazine []Magazine) error {
	result := m.db.Create(&magazine)
	return result.Error
}

// GetMagazineById 通过id得到期刊
func (m *MagazineModel) GetMagazineById(id uint) (Magazine, error) {
	magazine := Magazine{}
	result := m.db.First(&magazine, id)
	return magazine, result.Error
}

// GetMagazineByTitle 通过title得到期刊
func (m *MagazineModel) GetMagazinesByName(name string) (Magazine, error) {
	magazine := Magazine{}

	result := m.db.Where("name = ?", name).First(&magazine)
	return magazine, result.Error
}

// DeleteMagazineByTitle 通过title删除期刊
func (m *MagazineModel) DeleteMagazineByTitle(title string) error {
	result := m.db.Where("title = ?", title).Delete(&Magazine{})
	return result.Error
}

// DeleteMagazineById 通过id删除期刊
func (m *MagazineModel) DeleteMagazineById(id uint) error {
	result := m.db.Where("id = ?", id).Delete(&Magazine{})
	return result.Error
}

// DeleteMagazineByDoi 通过doi删除期刊
func (m *MagazineModel) DeleteMagazineByDoi(Doi string) error {
	result := m.db.Where("doi = ?", Doi).Delete(&Magazine{})
	return result.Error
}

// GetMagazineAndSub 给出user的关注状态
func (m *MagazineModel) GetMagazineAndSub(openid, name string) (sub bool, magazine Magazine, err error) {
	result := m.db.Where("name = ?", name).First(&magazine)
	count := m.db.Model(&User{OpenID: openid}).Where("name = ?", name).Association("Magazines").Count()
	if count == 1 {
		sub = true
	}
	err = result.Error
	return
}

func (m *MagazineModel) GetMagsInIssueAndTopic(issue, topic string) ([]Magazine, error) {
	var ms []Magazine
	result := m.db.Where("issue_name = ? AND topic_name = ?", issue, topic).Find(&ms)
	if result.RowsAffected == 0 {
		return []Magazine{}, nil
	}
	return ms, result.Error
}

type Content []string

func (c *Content) Scan(value interface{}) error {
	src, _ := value.(string)
	src = src[1 : len(src)-1]
	*c = strings.Split(src, ",")
	return nil
}
func (c Content) Value() (driver.Value, error) {
	return "{" + strings.Join(c, ",") + "}", nil
}

type ResultGetIssueAndTopic struct {
	Name    string
	Content Content
}

func (m *MagazineModel) GetIssueAndTopic() ([]ResultGetIssueAndTopic, error) {
	var results []ResultGetIssueAndTopic // 待测试
	err := m.db.Model(&Magazine{}).Select("issue_name as name, array_agg(topic_name) as content").Group("issue_name").Scan(&results).Error
	return results, err
}
