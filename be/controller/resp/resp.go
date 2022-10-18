package resp

import "git.bingyan.net/doc-aid-re-go/model"

type MagazineResp struct {
	Sub bool `json:"sub"`

	Name      string `json:"name"`      // 期刊名
	EnName    string `json:"enName"`    // 期刊英文名
	Unit      string `json:"unit"`      // 主办单位
	Cycle     string `json:"cycle"`     // 出版周期
	Issn      string `json:"issn"`      // ISSN
	Cn        string `json:"cn"`        // CN
	IssueName string `json:"issueName"` // 专辑名称 or 大类
	TopicName string `json:"topicName"` // 专题名称 or 小类

	Cif string `json:"cif"` // 复合影响因子
	Zif string `json:"zif"` // 综合影响因子

	Url string `json:"url"` // 期刊详情页链接
	Img string `json:"img"` // 期刊图片链接

	PaperCount int `json:"paperCount"` // 出版文献量
}

func MagazineRespTransferWithSub(magazine model.Magazine, sub bool) MagazineResp {
	return MagazineResp{
		Sub:        sub,
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
	}
}

func MagazineRespsTransferWithSub(magazines []model.Magazine, sub []bool) []MagazineResp {
	resps := make([]MagazineResp, len(sub))
	for i, magazine := range magazines {
		resps[i] = MagazineRespTransferWithSub(magazine, sub[i])
	}
	return resps
}

func MagazineRespTransfer(magazine model.Magazine) MagazineResp {
	return MagazineResp{
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
	}
}

func MagazineRespsTransfer(magazines []model.Magazine) []MagazineResp {
	resps := make([]MagazineResp, 0)
	for _, magazine := range magazines {
		resps = append(resps, MagazineRespTransfer(magazine))
	}
	return resps
}

type Kw struct {
	Explain string `json:"explain"`
}

type KwSub struct {
	Explain string `json:"explain"`
	Sub     bool   `json:"sub"`
}

type PaperResp struct {
	Url      string `json:"url"`
	Img      string `json:"img"`
	Intro    string `json:"intro"`
	Authors  string `json:"authors"`
	Year     int    `json:"year"`
	Magazine string `json:"magazine"`
	Id       string `json:"id"`
	Title    string `json:"title"`
	Issue    int    `json:"issue"`
	Fav      bool   `json:"fav"`
}
