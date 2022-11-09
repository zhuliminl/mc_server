package service

import (
	"context"
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
	"github.com/zhuliminl/mc_server/repository"
)

const (
	code2sessionURL = "https://api.weixin.qq.com/sns/jscode2session?appid=%s&secret=%s&js_code=%s&grant_type=authorization_code"
	appID           = "wxfd67f0c2f607440b"
	appSecret       = "a18ab85b749acb11c421cc96df3318da"
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
	// 存入 redis
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

func NewWechatService(userRepo repository.UserRepository, userService UserService, rdb *redis.Client) WechatService {
	return &wechatService{
		userRepository: userRepo,
		userService:    userService,
		rdb:            *rdb,
	}
}
