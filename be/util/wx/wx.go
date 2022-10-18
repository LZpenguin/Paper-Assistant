package wx

import (
	"encoding/json"
	"errors"
	"fmt"
	"git.bingyan.net/doc-aid-re-go/config"
	"net/http"
)

// 	首先调用wx.login(前端) ->auth.code2Session(后端)
// 	返回以下内容
//	属性		类型		说明
//	openid	string	用户唯一标识
//	session_key	string	会话密钥
//	unionid	string	用户在开放平台的唯一标识符，若当前小程序已绑定到微信开放平台帐号下会返回，详见 UnionID 机制说明。
//	errcode	number	错误码
//	errmsg	string	错误信息

type WXLoginResp struct {
	OpenId     string `json:"openid"`
	SessionKey string `json:"session_key"`
	UnionId    string `json:"unionid"`
	ErrCode    int    `json:"errcode"`
	ErrMsg     string `json:"errmsg"`
}

// GET https://api.weixin.qq.com/sns/jscode2session?appid=APPID&secret=SECRET&js_code=JSCODE&grant_type=authorization_code
func WXLogin(js_code string) (*WXLoginResp, error) {
	url := "https://api.weixin.qq.com/sns/jscode2session?appid=%s&secret=%s&js_code=%s&grant_type=authorization_code"

	// 合成url, 这里的appId和secret是在微信公众平台上获取的
	url = fmt.Sprintf(url, config.C.Wx.AppId, config.C.Wx.AppSecret, js_code)

	// 创建http get请求
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	fmt.Println("[WXLogin]resp body :", resp.Body)

	// 解析http请求中body 数据到我们定义的结构体中
	wxResp := WXLoginResp{}
	decoder := json.NewDecoder(resp.Body)
	if err := decoder.Decode(&wxResp); err != nil {
		return nil, err
	}

	// 判断微信接口返回的是否是一个异常情况
	if wxResp.ErrCode != 0 {
		return nil, errors.New(fmt.Sprintf("ErrCode:%s  ErrMsg:%s", wxResp.ErrCode, wxResp.ErrMsg))
	}

	return &wxResp, nil
}

// 然后直接解密encryptedData就行

func DecryptData(encryptedData, sessionKey, iv string) (string, error) {
	decrypt, err := aesDecrypt(encryptedData, sessionKey, iv)
	if err != nil {
		//log
		return "", err
	}
	return string(decrypt), nil
}
