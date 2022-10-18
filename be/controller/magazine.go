package controller

import (
	"git.bingyan.net/doc-aid-re-go/model"
	"git.bingyan.net/doc-aid-re-go/util/response"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"net/http"
	"strconv"
)

type AddOneMagazineParam struct {
	Name      string `json:"name"`      // 期刊名
	EnName    string `json:"enName"`    // 期刊英文名
	Unit      string `json:"unit"`      // 主办单位
	Cycle     string `json:"cycle"`     // 出版周期
	Issn      string `json:"ISSN"`      // ISSN
	Cn        string `json:"CN"`        // CN
	IssueName string `json:"issueName"` // 专辑名称 or 大类
	TopicName string `json:"topicName"` // 专题名称 or 小类

	Cif string `json:"FIF"` // 复合影响因子
	Zif string `json:"ZIF"` // 综合影响因子

	Url        string `json:"URL"` // 期刊详情页链接
	Img        string `json:"img"` // 期刊图片链接
	PaperCount int    `json:"paperCount"`
}

// POST /magazine/add
func AddOneMagazine(ctx echo.Context) error {
	magazine := new(AddOneMagazineParam)

	m := model.GetMagazineModel()

	if err := ctx.Bind(magazine); err != nil {
		return ctx.JSON(http.StatusBadRequest, response.Response{
			Status: http.StatusBadRequest,
			Data:   echo.Map{"err": "参数错误:" + err.Error()},
		})
	}
	err := m.CreateMagazine(model.Magazine{
		Name:       magazine.Name,
		EnName:     magazine.EnName,
		Unit:       magazine.Unit,
		Cycle:      magazine.Cycle,
		Issn:       magazine.Issn,
		Cn:         magazine.Cn,
		IssueName:  magazine.IssueName,
		TopicName:  magazine.TopicName,
		Cif:        magazine.Cif,
		Zif:        magazine.Zif,
		Url:        magazine.Url,
		Img:        magazine.Img,
		PaperCount: magazine.PaperCount,
	})
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, response.Response{
			Status: http.StatusInternalServerError,
			Data:   echo.Map{"err": "Add Magazine失败:" + err.Error()},
		})
	}
	return ctx.JSON(http.StatusOK, response.Response{
		Status: http.StatusOK,
		Data: echo.Map{
			"msg": "Add Magazine 成功！",
		},
	})
}

// GET /magazine/delete/:magazineid
func DeleteOneMagazine(ctx echo.Context) error {
	id := ctx.Param("magazineid")

	m := model.GetMagazineModel()

	idint, err := strconv.Atoi(id)
	if err != nil {
		return ctx.String(http.StatusInternalServerError, err.Error())
	}
	err = m.DeleteMagazineById(uint(idint))
	if err != nil {
		return ctx.String(http.StatusInternalServerError, "Delete Magazine 失败:"+err.Error())
	}
	return ctx.JSON(http.StatusOK, response.Response{
		Status: http.StatusOK,
		Data: echo.Map{
			"msg": "Delete Magazine 成功！",
		},
	})
}

type GetOneMagazineResp struct {
	Name      string `json:"name"`      // 期刊名
	EnName    string `json:"enName"`    // 期刊英文名
	Unit      string `json:"unit"`      // 主办单位
	Cycle     string `json:"cycle"`     // 出版周期
	Issn      string `json:"ISSN"`      // ISSN
	Cn        string `json:"CN"`        // CN
	IssueName string `json:"issueName"` // 专辑名称 or 大类
	TopicName string `json:"topicName"` // 专题名称 or 小类

	Cif string `json:"FIF"` // 复合影响因子
	Zif string `json:"ZIF"` // 综合影响因子

	Url        string `json:"URL"` // 期刊详情页链接
	Img        string `json:"img"` // 期刊图片链接
	PaperCount int    `json:"paperCount"`
}

// GET /magazine/:magazineid
func GetOneMagazine(ctx echo.Context) error {
	id := ctx.Param("magazineid")

	m := model.GetMagazineModel()

	idint, err := strconv.Atoi(id)
	if err != nil {
		return ctx.String(http.StatusInternalServerError, err.Error())
	}
	magazine, err := m.GetMagazineById(uint(idint))
	if err == gorm.ErrRecordNotFound {
		return ctx.JSON(http.StatusNotFound, response.Response{
			Status: http.StatusNotFound,
			Data: echo.Map{
				"error": "no such magazine",
			},
		})
	}
	if err != nil {
		return ctx.String(http.StatusInternalServerError, err.Error())
	}
	return ctx.JSON(http.StatusOK, response.Response{
		Status: http.StatusOK,
		Data: echo.Map{
			"magazine": GetOneMagazineResp{

				Name:       magazine.Name,
				EnName:     magazine.Name,
				Unit:       magazine.Unit,
				Cycle:      magazine.Cycle,
				Issn:       magazine.Issn,
				Cn:         magazine.Cn,
				IssueName:  magazine.IssueName,
				TopicName:  magazine.TopicName,
				Cif:        magazine.Cif,
				Zif:        magazine.Zif,
				Url:        magazine.Url,
				Img:        magazine.Img,
				PaperCount: magazine.PaperCount,
			},
		},
	})
}
