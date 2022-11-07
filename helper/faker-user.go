package helper

import (
	fakerV4 "github.com/bxcodec/faker/v4"
	"github.com/zhuliminl/mc_server/entity"
	"math/rand"
)

func FakerAUser() entity.User {
	var user entity.User
	user.UserId = fakerV4.UUIDHyphenated()
	user.Username = fakerV4.FirstNameMale()
	user.Email = fakerV4.Email()
	user.Phone = fakerV4.E164PhoneNumber()
	user.Password = "Test1234"
	user.WechatNickname = fakerV4.Name()
	user.WechatNumber = fakerV4.CCNumber()
	return user
}

func GenerateDefaultUserName() string {
	amount := 10
	index := rand.Intn(amount)
	names := []string{"saul", "jemo", "tang", "qingle", "lee", "xiaoshitou", "brance", "jack", "zhuliminl", "janlin"}
	return names[index]
}
