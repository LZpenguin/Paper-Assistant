package controller

import (
	"git.bingyan.net/doc-aid-re-go/model"
	"git.bingyan.net/doc-aid-re-go/util/response"
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
)

type CreateFavoriteParam struct {
	OpenID  string `json:"openID"`
	PaperID string `json:"paperID"`
}

// with JWT
func CreateFavorite(ctx echo.Context) error {
	param := new(CreateFavoriteParam)
	if err := ctx.Bind(param); err != nil {
		return ctx.JSON(http.StatusBadRequest, response.Response{
			Status: http.StatusBadRequest,
			Data: echo.Map{
				"error": "参数错误",
			},
		})
	}

	m := model.GetFavoriteModel()
	err := m.CreateFav(&model.Favorite{
		UserRefer: param.OpenID,
		PaperID:   param.PaperID,
	})
	if err != nil {
		log.Println("[CreateFavorite] : 创建Fav失败,err:" + err.Error())
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

type DeleteFavoriteParam struct {
	OpenID  string `json:"openID"`
	PaperID string `json:"paperID"`
}

func DeleteFavorite(ctx echo.Context) error {
	param := new(DeleteFavoriteParam)
	if err := ctx.Bind(param); err != nil {
		return ctx.JSON(http.StatusBadRequest, response.Response{
			Status: http.StatusBadRequest,
			Data: echo.Map{
				"error": "参数错误",
			},
		})
	}

	m := model.GetFavoriteModel()
	err := m.DeleteFavWithID(param.OpenID, param.PaperID)
	if err != nil {
		log.Println("[CreateFavorite] : 删除Fav失败,err:" + err.Error())
		return ctx.JSON(http.StatusBadRequest, response.Response{
			Status: http.StatusBadRequest,
			Data: echo.Map{
				"error": "删除失败",
			},
		})
	}

	return ctx.JSON(http.StatusAccepted, response.Response{
		Status: http.StatusAccepted,
		Data: echo.Map{
			"msg": "删除成功",
		},
	})
}
