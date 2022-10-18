package transfer

import (
	"git.bingyan.net/doc-aid-re-go/controller/resp"
	"git.bingyan.net/doc-aid-re-go/model"
	"strconv"
)

func PaperResp(papers []model.Paper, fav []bool) []resp.PaperResp {

	rs := make([]resp.PaperResp, len(papers))

	for i, paper := range papers {
		year, _ := strconv.Atoi(paper.Year)
		issue, _ := strconv.Atoi(paper.Issue)
		rs[i] = resp.PaperResp{
			Url:      paper.Url,
			Img:      paper.Magazine.Img,
			Intro:    paper.Introduction,
			Authors:  paper.Authors,
			Year:     year,
			Magazine: paper.MagazineName,
			Id:       paper.ID,
			Title:    paper.Title,
			Issue:    issue,
			Fav:      fav[i],
		}
	}
	return rs
}
