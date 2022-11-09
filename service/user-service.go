package service

import (
	"log"

	"github.com/zhuliminl/mc_server/constError"

	uuid "github.com/satori/go.uuid"
	"github.com/zhuliminl/mc_server/dto"
	"github.com/zhuliminl/mc_server/entity"
	"github.com/zhuliminl/mc_server/helper"
	"github.com/zhuliminl/mc_server/repository"
)

type UserService interface {
	Get(id string) (dto.User, error)
	GetAll() ([]dto.User, error)
	Create(userCreate dto.UserCreate) (dto.User, error)
	Delete(userDelete dto.UserDelete) error
	GenerateUsers(amount int) ([]dto.User, error)
}

type userService struct {
	userRepository repository.UserRepository
}

func NewUserService(userRepo repository.UserRepository) UserService {
	return &userService{
		userRepository: userRepo,
	}
}

func (service *userService) Get(id string) (dto.User, error) {
	user, err := service.userRepository.Get(id)
	if err != nil {
		return dto.User{}, err
	}
	return MapEntityUserToUser(user), nil
}

func (service *userService) GetAll() ([]dto.User, error) {
	var users []dto.User
	allUsers, err := service.userRepository.GetAll()
	if err != nil {
		return users, err
	}
	for _, item := range allUsers {
		dtoItem := MapEntityUserToUser(item)
		users = append(users, dtoItem)
	}
	return users, nil
}

func (service *userService) Create(userCreate dto.UserCreate) (dto.User, error) {
	userId := uuid.NewV4()
	user := entity.User{
		UserId:   userId.String(),
		Username: userCreate.Username,
		Email:    userCreate.Email,
		Phone:    userCreate.Phone,
		Password: userCreate.Password,
		// WechatNickname: "",
		// WechatNumber:   "",
	}
	newUser, err := service.userRepository.Create(user)
	if err != nil {
		return dto.User{}, err
	}
	return MapEntityUserToUser(newUser), err
}

func (service *userService) Delete(userDelete dto.UserDelete) error {
	exist, err := service.userRepository.Exist(userDelete.UserId)
	if err != nil {
		return err
	}
	if !exist {
		return constError.NewUserNotFound(err, "msg:上下文校对")
	}

	return service.userRepository.Delete(userDelete.UserId)
}

func (service *userService) GenerateUsers(length int) ([]dto.User, error) {
	var users []dto.User
	for i := 1; i <= length; i++ {
		fakeUser := helper.FakerAUser()
		user, err := service.userRepository.Create(fakeUser)
		if err != nil {
			log.Println("GenerateUsersError", err)
		}
		users = append(users, MapEntityUserToUser(user))
	}
	return users, nil
}

func MapEntityUserToUser(user entity.User) dto.User {
	return dto.User{
		UserId:         user.UserId,
		Username:       user.Username,
		Email:          user.Email,
		Phone:          user.Phone,
		WechatNickname: user.WechatNickname,
	}
}
