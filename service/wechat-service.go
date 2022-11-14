package service

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/zhuliminl/mc_server/constError"
	"github.com/zhuliminl/mc_server/helper"
	"log"
	"net/http"
	"time"

	"github.com/go-redis/redis/v9"
	uuid "github.com/satori/go.uuid"
	"github.com/zhuliminl/mc_server/constant"
	"github.com/zhuliminl/mc_server/dto"
	"github.com/zhuliminl/mc_server/repository"
)

const (
	code2sessionURL = "https://api.weixin.qq.com/sns/jscode2session?appid=%s&secret=%s&js_code=%s&grant_type=authorization_code"
	appID           = "wx333789da24c1f349"
	appSecret       = "62c790cced311e81907eea2d0b3a6310"
	appBaseLink     = "http://www.baidu.com"
)

var ctx = context.Background()

type WechatService interface {
	GetOpenId(wechatCode dto.WechatCodeDto) (dto.ResJsCode2session, error)
	GenerateAppLink() (dto.WechatAppLink, error)
	ScanOver(loginSessionId string) error
	GetMiniLinkStatus(loginSessionId string) (dto.MiniLinkStatus, error)
	LoginWithEncryptedPhoneData(wxLoginData dto.WxLoginData) (dto.ResWxLogin, error)
}

type wechatService struct {
	userRepository repository.UserRepository
	userService    UserService
	rdb            redis.Client
}

func (service wechatService) GetMiniLinkStatus(loginSessionId string) (dto.MiniLinkStatus, error) {
	var miniLinkStatus dto.MiniLinkStatus
	value, err := service.rdb.Get(ctx, loginSessionId+constant.PrefixLogin).Result()
	if err == redis.Nil {
		return miniLinkStatus, constError.NewWechatLoginUidNotFound(err, "uid key 不存在")
	} else if err != nil {
		return miniLinkStatus, err
	}
	miniLinkStatus.Status = value
	return miniLinkStatus, nil
}

func (service wechatService) GenerateAppLink() (dto.WechatAppLink, error) {
	loginSessionId := uuid.NewV4().String()
	var linkDto dto.WechatAppLink
	linkDto.Link = appBaseLink + "?login_session_id=" + loginSessionId
	linkDto.LoginSessionId = loginSessionId
	err := service.rdb.Set(ctx, loginSessionId+constant.PrefixLogin, constant.WechatLoginScanReady, constant.MiniLoginExpiredMinute*time.Minute).Err()
	if err != nil {
		// better panic
		return linkDto, err
	}

	return linkDto, nil
}

func (service wechatService) ScanOver(loginSessionId string) error {
	_, err := service.rdb.Get(ctx, loginSessionId+constant.PrefixLogin).Result()
	if err == redis.Nil {
		return constError.NewWechatLoginUidNotFound(err, "loginSessionId key 不存在")
	} else if err != nil {
		return err
	}

	err = service.rdb.Set(ctx, loginSessionId+constant.PrefixLogin, constant.WechatLoginScanOver, constant.MiniLoginExpiredMinute*time.Minute).Err()
	if err != nil {
		return err
	}

	return nil
}

func (service wechatService) GetOpenId(wechatCodeDto dto.WechatCodeDto) (dto.ResJsCode2session, error) {
	// 测试解密
	/*
		aesKey, err := base64.StdEncoding.DecodeString("q8vJomuvMB6QISZznoXSDw==")
		aesIv, err := base64.StdEncoding.DecodeString("Iim2Edy+qLm3bTT9tsQ13A==")

		base64Ciphertext :=
			"mF6qBn0ixpZspD3VKzie0tfI1g7uVSgJAK2PLOg3i3QpM78+i+sP81J2qchYu6u9jwpmWtkKXQ7kkOdSAeOefKKKEI3Y8tkMtz0Qz/MGqtuwFIwsbAiTU2htWZWyOnTL45LuPuihw3t3mx874gXCcJi9ZiDscaKBcPxHuVMYqaTiYVKHvyBlsui6l/l5v7+/6eNyi2jHvD3QLClCf98UxQ=="
		ciphertext, err := base64.StdEncoding.DecodeString(base64Ciphertext)
		if err != nil {
			log.Println("saul >>>>>>>>>>>>>>kjll", err)
		}

		raw, err := helper.AESDecryptData(ciphertext, aesKey, aesIv)
		if err != nil {
			log.Println("saul >>>>nnnnnnnnnn", err)
		}
		log.Println("saul AESDecryptData", string(raw))
		//
	*/

	var session dto.ResJsCode2session
	url := fmt.Sprintf(code2sessionURL, appID, appSecret, wechatCodeDto.Code)
	log.Println("code2sessionURL", url)

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

	log.Println("code2sessionWechatMap", wxMap)

	var errorCode int
	if _, ok := wxMap["errcode"].(int); ok {
		errorCode = wxMap["errcode"].(int)
	} else {
		errorCode = int(wxMap["errcode"].(float64))
	}

	session.OpenId = wxMap["openid"].(string)
	session.SessionKey = wxMap["session_key"].(string)
	session.Errcode = errorCode
	session.Errmsg = wxMap["errmsg"].(string)

	if session.SessionKey != "" {
		// 绑定 sessionKey 到 loginSessionId
		err := service.rdb.Set(ctx, wechatCodeDto.LoginSessionId+constant.PrefixWechatSessionKey, session.SessionKey, constant.MiniLoginExpiredMinute*time.Minute).Err()
		if err != nil {
			// better panic
			return session, err
		}
	}

	return session, nil
}

func (service wechatService) LoginWithEncryptedPhoneData(wxLoginData dto.WxLoginData) (dto.ResWxLogin, error) {
	var resWxLogin dto.ResWxLogin
	_, err := service.rdb.Get(ctx, wxLoginData.LoginSessionId+constant.PrefixLogin).Result()
	if err == redis.Nil {
		return resWxLogin, constError.NewWechatLoginUidNotFound(err, "loginSessionId key 不存在")
	} else if err != nil {
		return resWxLogin, err
	}

	sessionKey, err := service.rdb.Get(ctx, wxLoginData.LoginSessionId+constant.PrefixWechatSessionKey).Result()
	if err == redis.Nil {
		return resWxLogin, errors.New("wechat sessionKey 不存在，可能已过期")
	} else if err != nil {
		return resWxLogin, err
	}

	aesKey, err := base64.StdEncoding.DecodeString(sessionKey)
	aesIv, err := base64.StdEncoding.DecodeString(wxLoginData.Iv)

	base64Ciphertext := wxLoginData.EncryptedData
	ciphertext, err := base64.StdEncoding.DecodeString(base64Ciphertext)
	if err != nil {
		log.Println("WechatDecodeStringError", err)
	}
	raw, err := helper.AESDecryptData(ciphertext, aesKey, aesIv)
	if err != nil {
		log.Println("WechatAESDecryptDataError", err)
	}

	var resOfNumber dto.WxGetPhoneNumberRes
	log.Println("AESDecryptData", string(raw))
	err = json.Unmarshal(raw, &resOfNumber)
	if err != nil {
		return resWxLogin, err
	}

	resWxLogin.Phone = resOfNumber.PurePhoneNumber
	return resWxLogin, nil
}

func NewWechatService(userRepo repository.UserRepository, userService UserService, rdb *redis.Client) WechatService {
	return &wechatService{
		userRepository: userRepo,
		userService:    userService,
		rdb:            *rdb,
	}
}
