package service

import (
	"encoding/json"
	"fmt"
	"github.com/zhuliminl/mc_server/dto"
	"github.com/zhuliminl/mc_server/repository"
	"log"
	"net/http"
)

const (
	code2sessionURL = "https://api.weixin.qq.com/sns/jscode2session?appid=%s&secret=%s&js_code=%s&grant_type=authorization_code"
	appID           = "wxfd67f0c2f607440b"
	appSecret       = "a18ab85b749acb11c421cc96df3318da"
)

type WechatService interface {
	GetOpenId(wechatCode dto.WechatCode) (dto.ResJsCode2session, error)
}

type wechatService struct {
	userRepository repository.UserRepository
	userService    UserService
}

func (service wechatService) GetOpenId(wechatCode dto.WechatCode) (dto.ResJsCode2session, error) {
	var session dto.ResJsCode2session
	url := fmt.Sprintf(code2sessionURL, appID, appSecret, wechatCode.Code)
	log.Println("URL of getOpenId >>>", url)
	//return session, nil

	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("GetOpenIdHttpGetError", err)
		return session, err
	}

	wxMap := map[string]interface{}{"openid": "", "session_key": "", "errcode": 0, "errmsg": ""}

	err = json.NewDecoder(resp.Body).Decode(&wxMap)
	if err != nil {
		fmt.Println("GetOpenIdHttpDecodeError", err)
		return session, err
	}
	defer resp.Body.Close()

	log.Println("saul ==============>>> wxMap", wxMap)

	session.OpenId = wxMap["openid"].(string)
	session.SessionKey = wxMap["session_key"].(string)
	session.Errcode = int(wxMap["errcode"].(float64))
	session.Errmsg = wxMap["errmsg"].(string)

	return session, nil
}

func NewWechatService(userRepo repository.UserRepository, userService UserService) WechatService {
	return &wechatService{
		userRepository: userRepo,
		userService:    userService,
	}
}
