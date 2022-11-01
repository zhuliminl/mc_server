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

	// Get(name string) (*dto.User, error)
	// List(user dto.SessionUser, conditions condition.Conditions) ([]dto.User, error)
	// Create(isSuper bool, creation dto.UserCreate) (*dto.User, error)
	// Page(num, size int, user dto.SessionUser, conditions condition.Conditions) (*page.Page, error)
	// Delete(name string) error
	// Update(name string, isSuper bool, update dto.UserUpdate) (*dto.User, error)
	// Batch(op dto.UserOp) error
	// ChangePassword(isSuper bool, ch dto.UserChangePassword) error
	// UserAuth(name string, password string) (user *model.User, err error)
	// ResetPassword(fp dto.UserForgotPassword) error
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
	return service.userRepository.Get(id)
}

func (service *userService) CreateUser(userPayload forms.UserCreate) entity.User {
	id := uuid.NewV4()
	user := entity.User{
		UserId:         id.String(),
		Username:       userPayload.Username,
		Email:          userPayload.Email,
		Phone:          userPayload.Phone,
		WechatNickname: "",
		WechatNumber:   "",
	}
	return service.userRepository.Create(user)
}

func (service *userService) Foo(id string) entity.User {
	return service.userRepository.Get(id)
}
