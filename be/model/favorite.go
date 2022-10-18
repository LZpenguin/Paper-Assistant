package model

import (
	"gorm.io/gorm"
)

// 收藏Paper 和 收藏夹

type Favorite struct {
	gorm.Model `json:"-"`

	// 属于User, User have many Favorites
	UserRefer string `json:"-"`

	// Favorite belongs to paper
	PaperID string `gorm:"primarykey;uniqueIndex;not null;" json:"-"`
	Paper   Paper  `gorm:"foreignKey:PaperID;"`
}

// 限制唯一
func (f *Favorite) TableUnique() [][]string {
	return [][]string{
		{"user_refer", "paper_id"},
	}
}

type FavoriteModel struct {
	db *gorm.DB
}

func GetFavoriteModel() *FavoriteModel {
	return &FavoriteModel{db_gorm}
}

// CreateFav
// with UserRefer and PaperID
func (f *FavoriteModel) CreateFav(fav *Favorite) error {
	result := f.db.Create(fav)
	return result.Error
}

// DeleteFav
// with UserRefer and PaperID
func (f *FavoriteModel) DeleteFavWithID(UserRefer, PaperID string) error {
	result := f.db.Where("user_refer = ? AND paper_id = ?", UserRefer, PaperID).Delete(&Favorite{})
	return result.Error
}

// DeleteFavs
func (f *FavoriteModel) DeleteFavsWithID(UserRefer string, PaperIDs []string) error {
	result := f.db.Where("user_refer = ? AND paper_id IN ?", UserRefer, PaperIDs).Delete(&Favorite{})
	return result.Error
}
