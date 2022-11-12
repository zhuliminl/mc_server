package service

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/zhuliminl/mc_server/constError"
	"log"
	"net/http"
	"time"

	"github.com/go-redis/redis/v9"
	uuid "github.com/satori/go.uuid"
	"github.com/zhuliminl/mc_server/constant"
	"github.com/zhuliminl/mc_server/dto"
	"github.com/zhuliminl/mc_server/helper"
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
	GetOpenId(wechatCode dto.WechatCode) (dto.ResJsCode2session, error)
	GenerateAppLink() (dto.WechatAppLink, error)
	ScanOver(uid string) error
	GetMiniLinkStatus(uid string) (dto.MiniLinkUidStatus, error)
}

type wechatService struct {
	userRepository repository.UserRepository
	userService    UserService
	rdb            redis.Client
}

func (service wechatService) GetMiniLinkStatus(uid string) (dto.MiniLinkUidStatus, error) {
	var miniLinkStatus dto.MiniLinkUidStatus
	value, err := service.rdb.Get(ctx, uid).Result()
	if err == redis.Nil {
		return miniLinkStatus, constError.NewWechatLoginUidNotFound(err, "uid key 不存在")
	} else if err != nil {
		return miniLinkStatus, err
	}
	miniLinkStatus.Status = value
	return miniLinkStatus, nil
}

func (service wechatService) GenerateAppLink() (dto.WechatAppLink, error) {
	uid := uuid.NewV4().String()
	var linkDto dto.WechatAppLink
	linkDto.Link = appBaseLink + uid
	linkDto.Uid = uid
	err := service.rdb.Set(ctx, uid, constant.WechatLoginScanReady, 1*time.Minute).Err()
	if err != nil {
		// better panic
		return linkDto, err
	}

	return linkDto, nil
}

func (service wechatService) ScanOver(uid string) error {
	_, err := service.rdb.Get(ctx, uid).Result()
	if err == redis.Nil {
		return constError.NewWechatLoginUidNotFound(err, "uid key 不存在")
	} else if err != nil {
		return err
	}

	err = service.rdb.Set(ctx, uid, constant.WechatLoginScanOver, 1*time.Minute).Err()
	if err != nil {
		return err
	}

	return nil
}

func (service wechatService) GetOpenId(wechatCode dto.WechatCode) (dto.ResJsCode2session, error) {
	// 测试解密
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

	var session dto.ResJsCode2session
	url := fmt.Sprintf(code2sessionURL, appID, appSecret, wechatCode.Code)
	log.Println("saul URL of getOpenId >>>", url)

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

	return session, nil
}

func NewWechatService(userRepo repository.UserRepository, userService UserService, rdb *redis.Client) WechatService {
	return &wechatService{
		userRepository: userRepo,
		userService:    userService,
		rdb:            *rdb,
	}
}
