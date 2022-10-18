package transfer

import (
	"git.bingyan.net/doc-aid-re-go/controller/resp"
	"git.bingyan.net/doc-aid-re-go/model"
)

func KeywordsTransfer(kws []model.Keyword) []resp.Kw {
	k := make([]resp.Kw, 0)
	for _, kw := range kws {
		k = append(k, resp.Kw{Explain: kw.Explain})
	}
	return k
}

func KeywordsTransferWithSub(kws []model.Keyword, subs []bool) []resp.KwSub {
	k := make([]resp.KwSub, 0)
	for i, _ := range kws {
		k = append(k, resp.KwSub{Explain: kws[i].Explain, Sub: subs[i]})
	}
	return k
}

func KeywordsTransferString(kws []model.Keyword) []string {
	k := make([]string, 0)
	for _, kw := range kws {
		k = append(k, kw.Explain)
	}
	return k
}

func TransfertoKeywords(kws []string) []model.Keyword {
	k := make([]model.Keyword, 0)
	for _, kw := range kws {
		k = append(k, model.Keyword{Explain: kw})
	}
	return k
}
