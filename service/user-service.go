package service

import (
	uuid "github.com/satori/go.uuid"
	"github.com/zhuliminl/mc_server/dto"
	"github.com/zhuliminl/mc_server/entity"
	"github.com/zhuliminl/mc_server/helper"
	"github.com/zhuliminl/mc_server/repository"
)

type UserService interface {
	Get(id string) (entity.User, error)
	GetAll() ([]dto.UserAll, error)
	Create(userPayload dto.UserCreate) (entity.User, error)
	Delete(userPayload dto.UserDelete) error
	GenerateUsers(amount int) error

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

func (service *userService) Get(id string) (entity.User, error) {
	return service.userRepository.Get(id)
}

func (service *userService) GetAll() ([]dto.UserAll, error) {
	return service.userRepository.GetAll()
}

func (service *userService) Create(userPayload dto.UserCreate) (entity.User, error) {
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

func (service *userService) Foo(id string) (entity.User, error) {
	return service.userRepository.Get(id)
}

func (service *userService) Delete(userPayload dto.UserDelete) error {
	return service.userRepository.Delete(userPayload.UserId)
}

func (service *userService) GenerateUsers(length int) error {
	for i := 1; i <= length; i++ {
		user := helper.FakerAUser()
		service.userRepository.Create(user)
	}
	return nil
}
