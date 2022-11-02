package helper

import (
	fakerV4 "github.com/bxcodec/faker/v4"
	"github.com/zhuliminl/mc_server/entity"
)

func FakerAUser() entity.User {
	var user entity.User
	user.UserId = fakerV4.UUIDHyphenated()
	user.Email = fakerV4.Email()
	user.Phone = fakerV4.E164PhoneNumber()
	user.Password = fakerV4.Password()
	user.WechatNickname = fakerV4.Name()
	user.WechatNumber = fakerV4.CCNumber()
	return user
}
