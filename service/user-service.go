package service

import (
	uuid "github.com/satori/go.uuid"
	"github.com/zhuliminl/mc_server/entity"
	"github.com/zhuliminl/mc_server/forms"
	"github.com/zhuliminl/mc_server/repository"
)

type UserService interface {
	Profile(id string) entity.User
	CreateUser(userPayload forms.UserCreate) entity.User
}

type userService struct {
	userRepository repository.UserRepository
}

func NewUserService(userRepo repository.UserRepository) UserService {
	return &userService{
		userRepository: userRepo,
	}
}

func (service *userService) Profile(id string) entity.User {
	return service.userRepository.ProfileUser(id)
}

func (service *userService) CreateUser(userPayload forms.UserCreate) entity.User {
	id := uuid.NewV4()
	user := entity.User{
		UserId: id.String(),
		Username: userPayload.Username,
		Email: userPayload.Email,
		Phone: userPayload.Phone,
		WechatNickname: "",
		WechatNumber: "",
	}
	return service.userRepository.CreateUser(user)
}
