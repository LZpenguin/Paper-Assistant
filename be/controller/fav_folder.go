package controller

import (
	"git.bingyan.net/doc-aid-re-go/model"
	"git.bingyan.net/doc-aid-re-go/util/response"
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
)

type CreateFavFolderParam struct {
	Order uint   `json:"order"`
	Name  string `json:"name"`
	//Count  int    `json:"count"`
	OpenID string `json:"openID"`
}

func CreateFavFolder(ctx echo.Context) error {
	param := new(CreateFavFolderParam)
	if err := ctx.Bind(param); err != nil {
		return ctx.JSON(http.StatusBadRequest, response.Response{
			Status: http.StatusBadRequest,
			Data: echo.Map{
				"error": "参数错误",
			},
		})
	}

	m := model.GetFavoritesFolderModel()
	err := m.CreateFavoritesFolder(&model.FavoritesFolder{
		Order: param.Order,
		Name:  param.Name,
		//Count:     0,
		UserRefer: param.OpenID,
	})
	if err != nil {
		log.Println("[CreateFavFolder] : CreateFavoritesFolder失败,err:" + err.Error())
		return ctx.JSON(http.StatusBadRequest, response.Response{
			Status: http.StatusBadRequest,
			Data: echo.Map{
				"error": "创建失败",
			},
		})
	}

	return ctx.JSON(http.StatusCreated, response.Response{
		Status: http.StatusCreated,
		Data: echo.Map{
			"msg": "创建Favorite成功",
		},
	})
}

type DeleteFavFolderParam struct {
	FolderID uint `json:"folderID"`
}

func DeleteFavFolder(ctx echo.Context) error {
	param := new(DeleteFavFolderParam)
	if err := ctx.Bind(param); err != nil {
		return ctx.JSON(http.StatusBadRequest, response.Response{
			Status: http.StatusBadRequest,
			Data: echo.Map{
				"error": "参数错误",
			},
		})
	}

	m := model.GetFavoritesFolderModel()
	err := m.DeleteFavoritesFolder(param.FolderID)
	if err != nil {
		log.Println("[DeleteFavFolder] : DeleteFavFolder失败,err:" + err.Error())
		return ctx.JSON(http.StatusBadRequest, response.Response{
			Status: http.StatusBadRequest,
			Data: echo.Map{
				"error": "Delete失败",
			},
		})
	}

	return ctx.JSON(http.StatusCreated, response.Response{
		Status: http.StatusCreated,
		Data: echo.Map{
			"msg": "Delete成功",
		},
	})
}

type AddFavToFolderParam struct {
	FavID    uint `json:"favID"`
	FolderID uint `json:"folderID"`
}

func AddFavToFolder(ctx echo.Context) error {
	param := new(AddFavToFolderParam)
	if err := ctx.Bind(param); err != nil {
		return ctx.JSON(http.StatusBadRequest, response.Response{
			Status: http.StatusBadRequest,
			Data: echo.Map{
				"error": "参数错误",
			},
		})
	}

	m := model.GetFavoritesFolderModel()
	err := m.AddFavoriteToFolder(param.FolderID, param.FavID)
	if err != nil {
		log.Println("[AddFavToFolder] : 失败,err:" + err.Error())
		return ctx.JSON(http.StatusBadRequest, response.Response{
			Status: http.StatusBadRequest,
			Data: echo.Map{
				"error": "收藏失败",
			},
		})
	}

	return ctx.JSON(http.StatusCreated, response.Response{
		Status: http.StatusCreated,
		Data: echo.Map{
			"msg":      "收藏成功",
			"favID":    param.FavID,
			"folderID": param.FolderID,
		},
	})
}

type AddFavsToFolderParam struct {
	FavIDs   []uint `json:"favIDs"`
	FolderID uint   `json:"folderID"`
}
