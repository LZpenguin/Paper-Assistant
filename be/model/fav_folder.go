package model

import (
	"errors"
	"gorm.io/gorm"
)

type FavoritesFolder struct {
	gorm.Model `json:"-"`

	Order uint

	Name string `gorm:"uniqueIndex:name_refer;"`

	Count int

	// 属于User, User have many FavoritesFolders
	UserRefer string `gorm:"uniqueIndex:name_refer;" json:"-"`

	// m2m
	Favorites []Favorite `gorm:"many2many:folder_favorites;"` // 收藏夹里放的用户的收藏
}

type FavoritesFolderModel struct {
	db *gorm.DB
}

func GetFavoritesFolderModel() *FavoritesFolderModel {
	return &FavoritesFolderModel{db_gorm}
}

// CreateFavoritesFolder
// order = 0
func (f *FavoritesFolderModel) CreateFavoritesFolder(folder *FavoritesFolder) error {
	result := f.db.Create(folder)
	return result.Error
}

// DeleteFavoritesFolder
func (f *FavoritesFolderModel) DeleteFavoritesFolder(folderID uint) error {
	result := f.db.Unscoped().Delete(&FavoritesFolder{}, folderID)
	return result.Error
}

// DeleteFoldersWithID
func (f *FavoritesFolderModel) DeleteFoldersWithID(folderIDs []uint) error {
	result := f.db.Unscoped().Delete(&FavoritesFolder{}, folderIDs)
	return result.Error
}

// AddFavoriteToFolder 在收藏夹上添加收藏
func (f *FavoritesFolderModel) AddFavoriteToFolder(folderID uint, favID uint) error {
	ff := FavoritesFolder{}
	result := f.db.First(&ff, folderID)
	if result.Error != nil {
		return result.Error
	}
	err := f.db.Model(&ff).Association("Favorites").Append(&Favorite{
		Model: gorm.Model{ID: favID},
	})
	return err
}

// ReOrderFolders
// id -> order
func (f *FavoritesFolderModel) ReOrderFolders(folderIDs []uint) error {
	err := f.db.Transaction(func(tx *gorm.DB) error {
		for i, ID := range folderIDs {
			if err := tx.Model(&FavoritesFolder{}).Where("id = ?", ID).Update("order", i+1).Error; err != nil {
				return err
			}
		}
		return nil
	})

	return err
}

func (f *FavoritesFolderModel) GetFavoritesFolder(openid, name string) (FavoritesFolder, error) {
	var res FavoritesFolder
	result := f.db.Preload("Favorites").Where("user_refer = ? AND name = ?", openid, name).First(&res)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return res, nil
	}
	if result.Error != nil {
		return res, result.Error
	}
	result = f.db.Preload("Paper").Find(&res.Favorites)
	return res, result.Error
}

func (f *FavoritesFolderModel) DeleteFavoritesFolderByName(openid, name string) error {
	ff := FavoritesFolder{}
	err := f.db.Where("user_refer = ? AND name = ?", openid, name).First(&ff).Error
	if err != nil {
		return err
	}
	err = f.db.Model(&ff).Association("Favorites").Clear()
	if err != nil {
		return err
	}
	result := f.db.Unscoped().Where("user_refer = ? AND name = ?", openid, name).Delete(&FavoritesFolder{})
	return result.Error
}

func (f *FavoritesFolderModel) Order(openid string, names []string) {
	i := 1
	for _, name := range names {
		f.db.Model(&FavoritesFolder{}).Where("user_refer = ? AND name = ?", openid, name).Update("order", i)
		i += 1
	}
}

// AddFavoritesToFolder 在收藏夹上添加收藏
func (f *FavoritesFolderModel) AddFavoritesToFolder(openid, folderName string, favNames []string) error {
	favs := make([]Favorite, len(favNames))
	for i, id := range favNames {
		favs[i] = Favorite{
			UserRefer: openid,
			PaperID:   id,
		}
		err := f.db.FirstOrCreate(&favs[i]).Error
		if err != nil {
			return err
		}
	}
	folder := &FavoritesFolder{}
	err := f.db.Where("name = ? AND user_refer = ?", folderName, openid).First(&folder).Error
	if err != nil {
		return err
	}
	err = f.db.Model(&folder).Association("Favorites").Append(&favs)
	return err
}

func (f *FavoritesFolderModel) DeleteFavoritesInFolder(openid string, folderName string, favNames []string) error {
	var favs []Favorite
	for _, id := range favNames {
		fav := Favorite{}
		err := f.db.Where("user_refer = ? AND paper_id = ?", openid, id).First(&fav).Error
		if err != nil {
			return err
		}
		favs = append(favs, fav)
	}
	ff := FavoritesFolder{}
	err := f.db.Where("user_refer = ? AND name = ?", openid, folderName).First(&ff).Error
	if err != nil {
		return err
	}
	err = f.db.Model(&ff).Association("Favorites").Delete(favs)
	if err != nil {
		return err
	}
	err = f.db.Unscoped().Delete(&favs).Error
	return err
}
