package controller

import (
	"git.bingyan.net/doc-aid-re-go/controller/transfer"
	"git.bingyan.net/doc-aid-re-go/model"
	"git.bingyan.net/doc-aid-re-go/util/response"
	"github.com/labstack/echo/v4"
	"net/http"
)

type AddOnePaperParam struct {
	Title   string `json:"title"`
	Authors string `json:"authors"`

	Chapter int    `json:"chapter"` // 卷
	Phase   int    `json:"phase"`   // 期
	Doi     string `json:"doi"`     // DOI
	Year    string `json:"year"`    // 年
	Issue   string `json:"issue"`   // 发表时间

	Url          string   `json:"url"`   // 文献详情页链接
	Introduction string   `json:"intro"` // 摘要
	Keywords     []string `json:"keywords"`

	MagazineName string `json:"magazineName"`
}

func createKeywords(kws []string) []model.Keyword {
	k := []model.Keyword{}
	for _, s := range kws {
		k = append(k, model.Keyword{Explain: s})
	}
	return k
}

// POST /document/add
func AddOnePaper(ctx echo.Context) error {
	paper := new(AddOnePaperParam)

	m := model.GetPaperModel()

	if err := ctx.Bind(paper); err != nil {
		return ctx.String(http.StatusBadRequest, "参数错误:"+err.Error())
	}

	// TODO: 验证Paper重复性
	paperM := model.Paper{
		Title:        paper.Title,
		Authors:      paper.Authors,
		Chapter:      paper.Chapter,
		Phase:        paper.Phase,
		Doi:          paper.Doi,
		Year:         paper.Year,
		Issue:        paper.Issue,
		Url:          paper.Url,
		Introduction: paper.Introduction,
		Keywords:     createKeywords(paper.Keywords),
		// 处理归属Magazine问题 ，使用MagazineName定位
		MagazineName: paper.MagazineName,
	}
	err := m.CreatePaper(&paperM)

	if err != nil {
		return ctx.String(http.StatusInternalServerError, "add paper fail:"+err.Error())
	}
	return ctx.JSON(http.StatusOK, response.Response{
		Status: http.StatusOK,
		Data: echo.Map{
			"msg": "Add paper succeed!",
			"id":  paperM.ID,
		},
	})
}

// GET /document/delete/:documentid
func DeleteOnePaper(ctx echo.Context) error {
	id := ctx.Param("documentid")

	m := model.GetPaperModel()

	err := m.DeletePaperById(id)
	if err != nil {
		return ctx.String(http.StatusInternalServerError, "删除失败！"+err.Error())
	}
	return ctx.JSON(http.StatusOK, response.Response{
		Status: http.StatusOK,
		Data: echo.Map{
			"msg": "Delete paper succeed!",
		},
	})
}

// GET /document/:documentid
func GetOnePaper(ctx echo.Context) error {
	id := ctx.Param("documentid")

	m := model.GetPaperModel()

	p, err := m.GetPaperById(id)
	if err != nil {
		return ctx.String(http.StatusInternalServerError, err.Error())
	}

	return ctx.JSON(http.StatusOK, echo.Map{
		"Status": 200,
		"Data": GetOnePaperResp{
			Id:           p.ID,
			Title:        p.Title,
			Authors:      p.Authors,
			Chapter:      p.Chapter,
			Phase:        p.Phase,
			Doi:          p.Doi,
			Year:         p.Year,
			Issue:        p.Issue,
			Url:          p.Url,
			Introduction: p.Introduction,
			Keywords:     transfer.KeywordsTransferString(p.Keywords),
		},
	})
}
