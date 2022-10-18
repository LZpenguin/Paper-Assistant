package controller

import (
	"errors"
	"fmt"
	"git.bingyan.net/doc-aid-re-go/controller/resp"
	"git.bingyan.net/doc-aid-re-go/controller/transfer"
	"git.bingyan.net/doc-aid-re-go/model"
	"git.bingyan.net/doc-aid-re-go/util"
	"git.bingyan.net/doc-aid-re-go/util/response"
	"git.bingyan.net/doc-aid-re-go/util/wx"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"log"
	"net/http"
	"sort"
	"strconv"
	"strings"
)

// GET /document
func GetDocument(ctx echo.Context) error {
	openID := util.GetOpenID(ctx)
	m := model.GetUserModel()

	papers, err := m.Any20() // TODO: 更改方式

	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, response.Response{
			Status: http.StatusInternalServerError,
			Data:   err.Error(),
		})
	}
	paperIDs := make([]string, len(papers))
	for i, paper := range papers {
		paperIDs[i] = paper.ID
	}

	favs := m.IfSubPapers(openID, paperIDs)

	return ctx.JSON(http.StatusOK, response.Response{
		Status: http.StatusOK,
		Data:   transfer.PaperResp(papers, favs),
	})
}

type PostDocumentParam struct {
	Url      string   `json:"url"`
	Intro    string   `json:"intro"`
	Authors  []string `json:"authors"`
	Year     int      `json:"year"`
	Keywords []string `json:"keywords"`
	Magazine string   `json:"magazine"`
	Title    string   `json:"title"`
	Issue    int      `json:"issue"`
}

func PostDocumentParamToPaper(params []PostDocumentParam) []model.Paper {
	res := make([]model.Paper, len(params))
	for i, param := range params {
		res[i].Url = param.Url
		res[i].Introduction = param.Intro
		res[i].Year = strconv.Itoa(param.Year)
		res[i].Keywords = transfer.TransfertoKeywords(param.Keywords)
		res[i].MagazineName = param.Magazine
		res[i].Title = param.Title
		res[i].Issue = strconv.Itoa(param.Issue)
		res[i].Authors = strings.Join(param.Authors, " ")
	}
	return res
}

// POST /document
func PostDocument(ctx echo.Context) error {
	param := new([]PostDocumentParam)

	if err := ctx.Bind(&param); err != nil {
		return ctx.JSON(http.StatusBadRequest, response.Response{
			Status: http.StatusBadRequest,
			Data:   "参数错误",
		})
	}
	m := model.GetPaperModel()
	err := m.CreatePapers(PostDocumentParamToPaper(*param))
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, response.Response{
			Status: http.StatusBadRequest,
			Data:   "上传失败",
		})
	}
	return ctx.JSON(http.StatusCreated, response.Response{
		Status: http.StatusCreated,
		Data:   "上传成功",
	})
}

type GetOnePaperResp struct {
	Id      string `json:"id"`
	Title   string `json:"title"`
	Authors string `json:"authors"`

	Chapter int    `json:"chapter"` // 卷
	Phase   int    `json:"phase"`   // 期
	Doi     string `json:"doi"`     // DOI
	Year    string `json:"year"`    // 年
	Issue   string `json:"issue"`   // 发表时间

	Url string `json:"url"` // 文献详情页链接
	Img string `json:"img"`

	Introduction string   `json:"intro"` // 摘要
	Keywords     []string `json:"keywords"`
	MagazineName string   `json:"magazine"`

	Fav bool `json:"fav"`

	MagazineSub bool `json:"magazineSub"`

	Suggest []PaperContent `json:"suggest"`
}

// GET /document/:documentId
func GetOneDocumentById(ctx echo.Context) error {
	var openid string
	user := ctx.Get("user")
	if user != nil {
		openid = util.GetOpenID(ctx)
	}

	id := ctx.Param("documentId")
	m := model.GetPaperModel()
	p, err := m.GetPaperById(id)
	if err != nil {
		return ctx.String(http.StatusInternalServerError, err.Error())
	}
	u := model.GetUserModel()
	fav := u.IfSubPaper(openid, id)
	sub := u.IfSubMagazine(openid, p.MagazineName)
	var suggest []PaperContent
	return ctx.JSON(http.StatusOK, response.Response{
		Status: 200,
		Data: GetOnePaperResp{
			Id:           p.ID,
			Title:        p.Title,
			Authors:      p.Authors,
			Chapter:      p.Chapter,
			Phase:        p.Phase,
			Doi:          p.Doi,
			Year:         p.Year,
			Issue:        p.Issue,
			Url:          p.Url,
			Img:          p.Magazine.Img,
			Introduction: p.Introduction,
			Keywords:     transfer.KeywordsTransferString(p.Keywords),
			MagazineName: p.MagazineName,
			Fav:          fav,
			MagazineSub:  sub,
			Suggest:      suggest,
		},
	})
}

// GET /document/search/:keywords
// 根据标题搜索文献
func SearchDocumentByKeywords(ctx echo.Context) error {
	//openid := util.GetOpenID(ctx)
	//keywords := ctx.Param("keywords")
	//kws := strings.Split(keywords, " ")

	return nil
}

type PaperContent struct {
	Authors  string `json:"authors"`
	Year     string `json:"year"`
	Magazine string `json:"magazine"`
	Id       string `json:"id"`
	Title    string `json:"title"`
	Issue    string `json:"issue"`
}

type GetUserFavoriteResp struct {
	UpdTime string         `json:"updTime"`
	Content []PaperContent `json:"content"`
	Name    string         `json:"name"`
}

func PaperToContent(p model.Paper) PaperContent {
	return PaperContent{
		Authors:  p.Authors,
		Year:     p.Year,
		Magazine: p.MagazineName,
		Id:       p.ID,
		Title:    p.Title,
		Issue:    p.Issue,
	}
}

func GetUserFavoriteRespTransfer(fs []model.FavoritesFolder) []GetUserFavoriteResp {
	res := make([]GetUserFavoriteResp, len(fs))

	for i, f := range fs {
		res[i].UpdTime = f.UpdatedAt.String()
		res[i].Name = f.Name
		res[i].Content = []PaperContent{}
		for _, p := range f.Favorites {
			res[i].Content = append(res[i].Content, PaperToContent(p.Paper))
		}
	}
	return res
}

type Slice []model.FavoritesFolder

func (p Slice) Len() int           { return len(p) }
func (p Slice) Less(i, j int) bool { return p[i].Order <= p[j].Order }
func (p Slice) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

func ReOrderFolder(fs []model.FavoritesFolder) []model.FavoritesFolder {

	sort.Sort(Slice(fs))
	return fs
}

// GET /fav
// 获取全部收藏夹内容
// with JWT
func GetUserFavorite(ctx echo.Context) error {
	openid := util.GetOpenID(ctx)
	u := model.GetUserModel()
	fs, err := u.GetFavoritesFolders(openid)
	fs = ReOrderFolder(fs)
	if err != nil {
		log.Println("[GetUserFavorite]GetFavoritesFolders:" + err.Error())
		return ctx.JSON(http.StatusBadRequest, response.Response{
			Status: http.StatusBadRequest,
			Data:   nil,
		})
	}
	return ctx.JSON(http.StatusOK, response.Response{
		Status: http.StatusOK,
		Data:   GetUserFavoriteRespTransfer(fs),
	})
}

// DELETE /fav/:folderName
// 批量从指定收藏夹取消收藏文献
// with JWT
func DeleteFavoritesFromFolder(ctx echo.Context) error {
	openid := util.GetOpenID(ctx)
	fName := ctx.Param("folderName")
	param := new([]string)
	if err := ctx.Bind(&param); err != nil {
		return ctx.JSON(http.StatusBadRequest, response.Response{
			Status: http.StatusBadRequest,
			Data: echo.Map{
				"err": "参数错误",
			},
		})
	}
	f := model.GetFavoritesFolderModel()
	err := f.DeleteFavoritesInFolder(openid, fName, *param)
	if err != nil {
		log.Println("[PostFavoritesToFolder]DeleteFavoritesInFolder :" + err.Error())
		return ctx.JSON(http.StatusBadRequest, response.Response{
			Status: http.StatusBadRequest,
			Data:   nil,
		})
	}

	return ctx.JSON(http.StatusOK, response.Response{
		Status: http.StatusOK,
		Data:   nil,
	})
}

//POST /fav/:folderName
//批量收藏文献到指定收藏夹
// with JWT
func PostFavoritesToFolder(ctx echo.Context) error {
	openid := util.GetOpenID(ctx)
	fName := ctx.Param("folderName")
	param := new([]string)
	if err := ctx.Bind(&param); err != nil {
		return ctx.JSON(http.StatusBadRequest, response.Response{
			Status: http.StatusBadRequest,
			Data: echo.Map{
				"err": "参数错误",
			},
		})
	}
	log.Println("[PostFavoritesToFolder]param:", param)
	f := model.GetFavoritesFolderModel()
	err := f.AddFavoritesToFolder(openid, fName, *param)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		log.Println("[PostFavoritesToFolder]AddFavoritesToFolder :" + err.Error())
		return ctx.JSON(http.StatusNotFound, response.Response{
			Status: http.StatusNotFound,
			Data:   nil,
		})
	}
	if err != nil {
		log.Println("[PostFavoritesToFolder]AddFavoritesToFolder :" + err.Error())
		return ctx.JSON(http.StatusBadRequest, response.Response{
			Status: http.StatusBadRequest,
			Data:   nil,
		})
	}

	return ctx.JSON(http.StatusOK, response.Response{
		Status: http.StatusOK,
		Data:   nil,
	})
}

//PUT /folder
// 排序收藏夹
// with JWT
func SortFolder(ctx echo.Context) error {
	openid := util.GetOpenID(ctx)
	param := new([]string)
	if err := ctx.Bind(&param); err != nil {
		return ctx.JSON(http.StatusBadRequest, response.Response{
			Status: http.StatusBadRequest,
			Data: echo.Map{
				"err": "参数错误",
			},
		})
	}
	f := model.GetFavoritesFolderModel()
	f.Order(openid, *param)
	return ctx.JSON(http.StatusOK, response.Response{
		Status: http.StatusOK,
		Data:   nil,
	})
}

//DELETE /folder/:folderName
// with JWT
func DeleteFolder(ctx echo.Context) error {
	openid := util.GetOpenID(ctx)
	fName := ctx.Param("folderName")

	f := model.GetFavoritesFolderModel()

	err := f.DeleteFavoritesFolderByName(openid, fName)
	if err != nil {
		log.Println("[DeleteFolder]DeleteFavoritesFolderByName err:" + err.Error())
		return ctx.JSON(http.StatusBadRequest, response.Response{
			Status: http.StatusBadRequest,
			Data:   nil,
		})
	}
	return ctx.JSON(http.StatusOK, response.Response{
		Status: http.StatusOK,
		Data:   nil,
	})
}

// GET /folder/:folderName
// 导出收藏夹全部链接
// with JWT
func GetAllURLFromFolder(ctx echo.Context) error {
	openid := util.GetOpenID(ctx)
	fName := ctx.Param("folderName")

	f := model.GetFavoritesFolderModel()

	folder, err := f.GetFavoritesFolder(openid, fName)
	if err != nil {
		log.Println("[GetAllURLFromFolder]GetFavoritesFolder err:" + err.Error())
		return ctx.JSON(http.StatusBadRequest, response.Response{
			Status: http.StatusBadRequest,
			Data:   nil,
		})
	}
	resp := make([]string, 0)
	for _, k := range folder.Favorites {
		resp = append(resp, k.Paper.Url)
	}
	return ctx.JSON(http.StatusOK, response.Response{
		Status: http.StatusOK,
		Data:   resp,
	})
}

// POST /folder/:folderName
// 添加收藏夹
// with JWT
func PostFolder(ctx echo.Context) error {
	openid := util.GetOpenID(ctx)
	fName := ctx.Param("folderName")

	f := model.GetFavoritesFolderModel()

	err := f.CreateFavoritesFolder(&model.FavoritesFolder{
		Order:     0,
		Name:      fName,
		Count:     0,
		UserRefer: openid,
	})
	if err != nil {
		log.Println("[PostFolder]CreateFavoritesFolder err:" + err.Error())
		return ctx.JSON(http.StatusBadRequest, response.Response{
			Status: http.StatusBadRequest,
			Data:   nil,
		})
	}

	return ctx.JSON(http.StatusOK, response.Response{
		Status: http.StatusOK,
		Data:   nil,
	})
}

// DELETE /keyword
// 删除订阅的关键词
// with JWT
func DeleteKeyword(ctx echo.Context) error {
	openid := util.GetOpenID(ctx)
	param := new([]string)
	if err := ctx.Bind(&param); err != nil {
		log.Println("[DeleteKeyword]err:" + err.Error())
		return ctx.JSON(http.StatusBadRequest, response.Response{
			Status: http.StatusBadRequest,
			Data: echo.Map{
				"err": "参数错误",
			},
		})
	}

	u := model.GetUserModel()
	err := u.DeleteKeywords(openid, transfer.TransfertoKeywords(*param))
	if err != nil {
		log.Println("[DeleteKeyword]DeleteKeywords err:" + err.Error())
		return ctx.JSON(http.StatusBadRequest, response.Response{
			Status: http.StatusBadRequest,
			Data:   nil,
		})
	}

	return ctx.JSON(http.StatusOK, response.Response{
		Status: http.StatusOK,
		Data:   nil,
	})
}

// GET /keyword
// 查看订阅的关键词
// with JWT
func GetKeywords(ctx echo.Context) error {
	openid := util.GetOpenID(ctx)

	u := model.GetUserModel()
	user, err := u.GetKeywords(openid)
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
		Data:   transfer.KeywordsTransferString(user.Keywords),
	})
}

// POST /keyword
// 订阅关键词
// with JWT
func PostKeywords(ctx echo.Context) error {
	openid := util.GetOpenID(ctx)
	param := new([]string)
	if err := ctx.Bind(&param); err != nil {
		log.Println("[PostKeywords]err:" + err.Error())
		return ctx.JSON(http.StatusBadRequest, response.Response{
			Status: http.StatusBadRequest,
			Data: echo.Map{
				"err": "参数错误",
			},
		})
	}

	kws := transfer.TransfertoKeywords(*param)

	u := model.GetUserModel()
	err := u.SubKeywords(openid, kws)
	if err != nil {
		log.Println("[PostKeywords]subkws :" + err.Error())
		return ctx.JSON(http.StatusInternalServerError, response.Response{
			Status: http.StatusInternalServerError,
			Data:   nil,
		})
	}

	return ctx.JSON(http.StatusOK, response.Response{
		Status: http.StatusOK,
		Data:   nil,
	})
}

// GET /keyword/search/:keywords
// 搜索关键词
// with JWT
func SearchKeywords(ctx echo.Context) error {
	openid := util.GetOpenID(ctx)
	keyword := ctx.Param("keywords")
	log.Println("key:", keyword)

	m := model.GetKeywordModel()
	var kws []model.Keyword
	var err error
	if keyword == ":keywords" {
		log.Println("SearchAll")
		kws, err = m.SearchAll()
	} else {
		keywords := strings.Split(keyword, " ")
		log.Println("Search by ", keywords)
		kws, err = m.SearchSome(keywords)
	}

	if err != nil {
		log.Println("[SearchKeywords]:" + err.Error())
		return ctx.JSON(http.StatusBadRequest, response.Response{
			Status: http.StatusBadRequest,
			Data:   nil,
		})
	}

	u := model.GetUserModel()
	subs := u.IfSubKeywords(openid, transfer.KeywordsTransferString(kws))

	resp := transfer.KeywordsTransferWithSub(kws, subs)

	return ctx.JSON(http.StatusOK, response.Response{
		Status: http.StatusOK,
		Data:   resp,
	})
}

// GET /magazine
// 获取全部期刊分类
func GetMagazine(ctx echo.Context) error {
	m := model.GetMagazineModel()
	rs, err := m.GetIssueAndTopic()
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
		Data:   rs,
	})
}

type CreateMagazineParam struct {
	Name      string `json:"name"`
	EnName    string `json:"enName"`
	Issn      string `json:"issn"`
	Zif       string `json:"zif"`
	Url       string `json:"url"`
	Img       string `json:"img"`
	TopicName string `json:"topicName"`
	Fif       string `json:"fif"`
	Cn        string `json:"cn"`
	Cycle     string `json:"cycle"`
	IssueName string `json:"issueName"`
	Unit      string `json:"unit"`
}

func mTransfer(param []CreateMagazineParam) []model.Magazine {
	res := make([]model.Magazine, len(param))
	for i, p := range param {
		res[i] = model.Magazine{
			Name:      p.Name,
			EnName:    p.EnName,
			Unit:      p.Unit,
			Cycle:     p.Cycle,
			Issn:      p.Issn,
			Cn:        p.Cn,
			IssueName: p.IssueName,
			TopicName: p.TopicName,
			Cif:       p.Fif,
			Zif:       p.Zif,
			Url:       p.Url,
			Img:       p.Img,
		}
	}
	return res
}

// POST /magazine
// **爬虫用**添加期刊。
func CreateMagazine(ctx echo.Context) error {
	param := new([]CreateMagazineParam)

	if err := ctx.Bind(&param); err != nil {
		return ctx.JSON(http.StatusBadRequest, response.Response{
			Status: http.StatusBadRequest,
			Data: echo.Map{
				"err": "参数错误",
			},
		})
	}

	m := model.GetMagazineModel()
	err := m.CreateMagazines(mTransfer(*param))
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, response.Response{
			Status: http.StatusBadRequest,
			Data: echo.Map{
				"err": "上传失败",
			},
		})
	}

	return ctx.JSON(http.StatusCreated, response.Response{
		Status: http.StatusCreated,
		Data:   nil,
	})
}

// GET /magazine/:issueName/:topicName
// 获取分类内全部期刊
// With JWT
func GetMagazinesWithTopic(ctx echo.Context) error {
	openid := util.GetOpenID(ctx)

	issue := ctx.Param("issueName")
	topic := ctx.Param("topicName")
	m := model.GetMagazineModel()
	ms, err := m.GetMagsInIssueAndTopic(issue, topic)
	if err != nil {
		log.Println("[GetMagazinesWithTopic]err:" + err.Error())
		return ctx.JSON(http.StatusBadRequest, response.Response{
			Status: http.StatusBadRequest,
			Data: echo.Map{
				"err": "查询失败",
			},
		})
	}

	u := model.GetUserModel()
	yes := u.IfSubMagazines(openid, ms)
	//log.Println("len ms = ", len(ms))
	//log.Println("ms=", ms)
	//log.Println("len yes = ", len(yes))
	return ctx.JSON(http.StatusOK, response.Response{
		Status: http.StatusOK,
		Data:   resp.MagazineRespsTransferWithSub(ms, yes),
	})
}

// GET /magazine/:magazineName
// 获取某期刊详细信息
// With JWT
func GetMagazineInfoWithMagazineName(ctx echo.Context) error {
	openid := util.GetOpenID(ctx)
	magazineName := ctx.Param("magazineName")
	m := model.GetMagazineModel()
	sub, magazine, err := m.GetMagazineAndSub(openid, magazineName)
	if err != nil {
		return ctx.JSON(http.StatusNotFound, response.Response{
			Status: http.StatusNotFound,
			Data: echo.Map{
				"err": "404 Not Found",
			},
		})
	}
	return ctx.JSON(http.StatusOK, response.Response{
		Status: http.StatusOK,
		Data:   resp.MagazineRespTransferWithSub(magazine, sub),
	})
}

type DeleteSubParam struct {
	Names []string `json:"names"`
}

// DELETE /sub
// 批量取消订阅期刊
// With JWT
func DeleteSub(ctx echo.Context) error {
	openid := util.GetOpenID(ctx)
	param := new(DeleteSubParam)
	if err := ctx.Bind(param); err != nil {
		log.Println("[DeleteSub]err:" + err.Error())
		return ctx.JSON(http.StatusBadRequest, response.Response{
			Status: http.StatusBadRequest,
			Data: echo.Map{
				"err": "参数错误",
			},
		})
	}

	m := model.GetUserModel()
	err := m.DeleteSubMagazines(openid, param.Names)
	if err != nil {
		log.Println("[DeleteSub]err:" + err.Error())
		return ctx.JSON(http.StatusBadRequest, response.Response{
			Status: http.StatusBadRequest,
			Data: echo.Map{
				"err": "获取订阅失败",
			},
		})
	}
	return ctx.JSON(http.StatusOK, response.Response{
		Status: http.StatusOK,
	})
}

// GET /sub
// 查看订阅期刊
// With JWT
func GetSub(ctx echo.Context) error {
	openid := util.GetOpenID(ctx)
	m := model.GetUserModel()
	u, err := m.GetMagazines(openid)
	if err != nil {
		log.Println("[GetSub]err:" + err.Error())
		return ctx.JSON(http.StatusBadRequest, response.Response{
			Status: http.StatusBadRequest,
			Data: echo.Map{
				"err": "获取订阅失败",
			},
		})
	}
	mResp := resp.MagazineRespsTransfer(u.Magazines)

	return ctx.JSON(http.StatusOK, response.Response{
		Status: http.StatusOK,
		Data:   mResp,
	})
}

// POST /sub/:magazineName
// 订阅期刊
// With JWT
func SubMagazineWithMagazineName(ctx echo.Context) error {
	openid := util.GetOpenID(ctx)
	magazineName := ctx.Param("magazineName")
	m := model.GetUserModel()
	err := m.SubMagazine(openid, magazineName)
	if err != nil {
		log.Println("[SubMagazineWithMagazineName]err:" + err.Error())
		return ctx.JSON(http.StatusBadRequest, response.Response{
			Status: http.StatusBadRequest,
			Data: echo.Map{
				"err": "订阅失败",
			},
		})
	}
	return ctx.JSON(http.StatusOK, response.Response{
		Status: http.StatusOK,
		Data: echo.Map{
			"msg": "订阅成功",
		},
	})
}

type GatherFeedbackParam struct {
	Content string `json:"content"`
}

// POST /user/feedback
// 上传反馈
// With JWT
func GatherFeedback(ctx echo.Context) error {
	openid := util.GetOpenID(ctx)
	param := new(GatherFeedbackParam)
	if err := ctx.Bind(param); err != nil {
		return ctx.JSON(http.StatusBadRequest, response.Response{
			Status: http.StatusBadRequest,
			Data: echo.Map{
				"err": "参数错误",
			},
		})
	}
	m := model.GetUserModel()
	err := m.FeedBack(openid, param.Content)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, response.Response{
			Status: http.StatusInternalServerError,
			Data: echo.Map{
				"err": "上传反馈失败",
			},
		})
	}

	return ctx.JSON(http.StatusOK, response.Response{
		Status: http.StatusOK,
		Data: echo.Map{
			"msg": "上传反馈成功",
		},
	})
}

type LoginParam struct {
	Code string `json:"code" yaml:"code"`
}

// POST /user/login
func UserLogin(ctx echo.Context) error {
	loginParam := new(LoginParam)
	if err := ctx.Bind(loginParam); err != nil {
		return ctx.JSON(http.StatusBadRequest, response.Response{
			Status: http.StatusBadRequest,
			Data: echo.Map{
				"err": "参数错误",
			},
		})
	}
	wxResp, err := wx.WXLogin(loginParam.Code)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, response.Response{
			Status: http.StatusBadRequest,
			Data: echo.Map{
				"err": "登录失败,请重试",
			},
		})
	}
	fmt.Println("Openid:", wxResp.OpenId, " 登录成功！")

	// 查数据库
	m := model.GetUserModel()
	user := model.User{OpenID: wxResp.OpenId}
	err = m.GetOrCreateUser(&user)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, response.Response{
			Status: http.StatusInternalServerError,
			Data: echo.Map{
				"err": "登录失败,请重试",
			},
		})
	}
	// 发放JWTToken
	// 完成
	claims := util.JWTClaims{
		Openid: wxResp.OpenId,
	}
	token, err := util.GenerateJWTToken(claims)
	if err != nil {
		log.Println("[UserLogin]err:" + err.Error())
		return ctx.JSON(http.StatusInternalServerError, response.Response{
			Status: http.StatusInternalServerError,
			Data:   echo.Map{},
		})
	}

	return ctx.JSON(http.StatusOK, response.Response{
		Status: http.StatusOK,
		Data: echo.Map{
			"token": token,
		},
	})
}
