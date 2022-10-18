package controller

import (
	"git.bingyan.net/doc-aid-re-go/model"
	"git.bingyan.net/doc-aid-re-go/util/response"
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
	"strconv"
)

type CreateOneKeywordParam struct {
	Explain string `json:"explain" form:"explain"`
}

func CreateOneKeyword(ctx echo.Context) error {

	keyword := new(CreateOneKeywordParam)
	if err := ctx.Bind(keyword); err != nil {
		return ctx.JSON(http.StatusBadRequest, response.Response{
			Status: http.StatusBadRequest,
			Data: echo.Map{
				"error": "参数错误",
			},
		})
	}
	if len(keyword.Explain) <= 0 {
		return ctx.JSON(http.StatusBadRequest, response.Response{
			Status: http.StatusBadRequest,
			Data: echo.Map{
				"error": "keyword 应该有描述！",
			},
		})
	}

	m := model.GetKeywordModel()
	err := m.CreateKeyword(model.Keyword{Explain: keyword.Explain})
	if err != nil {
		log.Println("[CreateKeyword]" + err.Error())
		return ctx.JSON(http.StatusInternalServerError, response.Response{
			Status: http.StatusInternalServerError,
			Data: echo.Map{
				"error": "服务器错误",
			},
		})
	}

	return ctx.JSON(http.StatusOK, response.Response{
		Status: http.StatusOK,
		Data: echo.Map{
			"msg": "Create Keyword Success",
		},
	})
}

type CreateManyKeyWordsParam struct {
	Explains []string `json:"explains" form:"explains"`
}

func CreateManyKeyWords(ctx echo.Context) error {

	keywords := new(CreateManyKeyWordsParam)

	if err := ctx.Bind(keywords); err != nil {
		return ctx.JSON(http.StatusBadRequest, response.Response{
			Status: http.StatusBadRequest,
			Data: echo.Map{
				"error": "参数错误",
			},
		})
	}

	kws := make([]model.Keyword, 1)
	f := func(s string) model.Keyword {
		return model.Keyword{Explain: s}
	}
	for _, s := range keywords.Explains {
		if len(s) > 0 {
			kws = append(kws, f(s))
		}
	}
	if len(kws) <= 0 {
		return ctx.JSON(http.StatusBadRequest, response.Response{
			Status: http.StatusBadRequest,
			Data: echo.Map{
				"error": "参数错误",
			},
		})
	}

	m := model.GetKeywordModel()
	err := m.CreateManyKeyword(kws)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, response.Response{
			Status: http.StatusInternalServerError,
			Data: echo.Map{
				"error": "服务器错误",
			},
		})
	}
	return ctx.JSON(http.StatusOK, echo.Map{
		"status": http.StatusOK,
		"msg":    "Add Keywords Success",
	})
}

// 这个方法估计用不到
func DeleteOneKeyword(ctx echo.Context) error {
	id := ctx.Param("keywordid") // 这个是对的

	m := model.GetKeywordModel()

	idInt, err := strconv.Atoi(id)
	if err != nil {
		return ctx.String(http.StatusInternalServerError, err.Error())
	}
	err = m.DeleteKeywordById(uint(idInt))
	if err != nil {
		return ctx.String(http.StatusInternalServerError, err.Error())
	}
	return ctx.JSON(http.StatusOK, echo.Map{
		"status": http.StatusOK,
		"msg":    "Delete Keyword Success",
	})
}

func GetOneKeyword(ctx echo.Context) error {
	id := ctx.Param("keywordid")

	m := model.GetKeywordModel()
	idInt, err := strconv.Atoi(id)
	if err != nil {
		return ctx.String(http.StatusInternalServerError, err.Error())
	}
	keyword, err := m.GetKeywordById(uint(idInt))
	if err != nil {
		return ctx.String(http.StatusInternalServerError, err.Error())
	}
	return ctx.JSON(http.StatusOK, response.Response{
		Status: 200,
		Data:   keyword,
	})
}

func GetAllKeywords(ctx echo.Context) error {
	m := model.GetKeywordModel()

	kws, err := m.GetAllKeywords()
	if err != nil {
		return ctx.String(http.StatusInternalServerError, err.Error())
	}
	return ctx.JSON(http.StatusOK, response.Response{
		Status: 200,
		Data:   kws,
	})
}
