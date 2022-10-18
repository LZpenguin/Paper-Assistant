package controller

import (
	"fmt"
	"git.bingyan.net/doc-aid-re-go/model"
	"git.bingyan.net/doc-aid-re-go/util"
	"git.bingyan.net/doc-aid-re-go/util/response"
	"git.bingyan.net/doc-aid-re-go/util/wx"
	"github.com/labstack/echo/v4"
	"net/http"
)

// /api/login
func Login(ctx echo.Context) error {
	loginParam := new(LoginParam)
	if err := ctx.Bind(loginParam); err != nil {
		return ctx.String(http.StatusBadRequest, err.Error())
	}
	wxResp, err := wx.WXLogin(loginParam.Code)
	if err != nil {
		return ctx.String(http.StatusBadRequest, err.Error())
	}
	fmt.Println("Openid:", wxResp.OpenId, " 登录成功！")

	// 查数据库
	m := model.GetUserModel()
	isnew, err := m.QueryUserByOpenid(wxResp.OpenId)
	if err != nil {
		return ctx.String(http.StatusBadRequest, err.Error())
	}

	if isnew {
		user := &model.User{
			OpenID: wxResp.OpenId,
		}
		user, err = m.CreateUser(user)
		if err != nil {
			return ctx.JSON(http.StatusBadRequest, response.Response{
				Status: http.StatusBadRequest,
				Data:   echo.Map{"err": err.Error()}},
			)
		}
	}

	// 发放JWTToken
	// 完成

	claims := util.JWTClaims{
		Openid: wxResp.OpenId,
	}
	token, err := util.GenerateJWTToken(claims)
	if err != nil {
		return ctx.String(http.StatusBadRequest, err.Error())
	}

	return ctx.JSON(http.StatusOK, response.Response{
		Status: http.StatusOK,
		Data: echo.Map{
			"token": token,
			"isNew": isnew,
		},
	})
}

type GetJWTParam struct {
	OpenID string `json:"openID"`
}

// GET /user/debug
func GetJWT(ctx echo.Context) error {
	param := new(GetJWTParam)
	if err := ctx.Bind(param); err != nil {
		return ctx.JSON(http.StatusBadRequest, response.Response{
			Status: http.StatusBadRequest,
			Data: echo.Map{
				"error": "参数错误，请检查参数",
			},
		})
	}
	m := model.GetUserModel()
	u, err := m.GetUserByOpenID(param.OpenID)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, response.Response{
			Status: http.StatusBadRequest,
			Data: echo.Map{
				"error": "获取user失败",
			},
		})
	}

	claims := util.JWTClaims{
		Openid: u.OpenID,
	}
	token, err := util.GenerateJWTToken(claims)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, response.Response{
			Status: http.StatusBadRequest,
			Data: echo.Map{
				"error": "服务器错误，请重试",
			},
		})
	}

	return ctx.JSON(http.StatusOK, response.Response{
		Status: http.StatusOK,
		Data: echo.Map{
			"token": token,
		},
	})
}

type CreateUserParam struct {
	OpenId string `json:"openid"`
}

// debug用
func CreateUser(ctx echo.Context) error {
	p := new(CreateUserParam)
	if err := ctx.Bind(p); err != nil {
		return ctx.JSON(http.StatusBadRequest, response.Response{
			Status: http.StatusBadRequest,
			Data: echo.Map{
				"err": err.Error(),
			},
		})
	}
	m := model.GetUserModel()

	u, err := m.CreateUser(&model.User{OpenID: p.OpenId, Role: model.CustomerRole})
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, response.Response{
			Status: http.StatusBadRequest,
			Data: echo.Map{
				"err": "Create User 失败" + err.Error(),
			},
		})
	}

	claims := util.JWTClaims{
		Openid: u.OpenID,
	}
	token, err := util.GenerateJWTToken(claims)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, response.Response{
			Status: http.StatusInternalServerError,
			Data: echo.Map{
				"error": "服务器错误，请重试",
			},
		})
	}

	return ctx.JSON(http.StatusOK, response.Response{
		Status: http.StatusOK,
		Data: echo.Map{
			"msg":   "Create User 成功",
			"token": token,
		},
	})
}

// /user/all
func ListAllUsers(ctx echo.Context) error {
	m := model.GetUserModel()
	users, err := m.GetAllUsers()
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, response.Response{
			Status: http.StatusBadRequest,
			Data: echo.Map{
				"err": err.Error(),
			},
		})
	}
	return ctx.JSON(http.StatusOK, response.Response{
		Status: http.StatusOK,
		Data: echo.Map{
			"users": users,
		},
	})
}
