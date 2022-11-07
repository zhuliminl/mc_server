package service

import (
	"github.com/zhuliminl/mc_server/helper"
	"log"

	"github.com/zhuliminl/mc_server/constError"
	"github.com/zhuliminl/mc_server/dto"
	"github.com/zhuliminl/mc_server/repository"
)

type AuthService interface {
	VerifyCredential(email string, password string) (dto.UserVerify, error)
	VerifyRegisterByEmail(user dto.UserRegisterByEmail) error
	CreateUser(user dto.UserRegisterByEmail) (dto.User, error)
	//FindByEmail(email string) (dto.User, error)
	//isDuplicateEmail(email string) (bool, error)
}

type authService struct {
	userRepository repository.UserRepository
	userService    UserService
}

func (service *authService) VerifyCredential(email string, password string) (dto.UserVerify, error) {
	user, err := service.userRepository.GetByEmail(email)
	if err != nil {
		return dto.UserVerify{}, err
	}

	log.Println("saul >>>>", user, password)
	matchPwd := user.Password == password
	if !matchPwd {
		return dto.UserVerify{}, constError.NewPasswordNotMatch(nil, "密码匹配错误")
	}

	return dto.UserVerify{
		IsValid: true,
		User:    MapEntityUserToUser(user),
	}, nil

}
func (service *authService) VerifyRegisterByEmail(user dto.UserRegisterByEmail) error {
	if !helper.IsEmailValid(user.Email) {
		return constError.NewEmailNotValid(nil, "邮箱格式错误")
	}
	if !helper.IsPasswordValid(user.Password) {
		return constError.NewPasswordNotValid(nil, "密码格式错误")
	}
	userFind, err := service.userRepository.GetByEmail(user.Email)
	if constError.Is(err, constError.UserNotFound) {
		return nil
	}
	if err != nil {
		return err
	}
	if userFind.UserId != "" {
		return constError.NewUserDuplicated(nil, "用户已注册")
	}
	return nil
}

func (service *authService) CreateUser(userRegister dto.UserRegisterByEmail) (dto.User, error) {
	username := userRegister.Username
	if username == "" {
		username = helper.GenerateDefaultUserName()
	}
	user, err := service.userService.Create(dto.UserCreate{
		Username: username,
		Email:    userRegister.Email,
		Password: userRegister.Password,
		Phone:    "",
	})
	if err != nil {
		return dto.User{}, err
	}
	return user, nil
}

// func (service *authService) FindByEmail(email string) (dto.User, error) {
// }
//
// func (service *authService) isDuplicateEmail(email string) (bool, error) {
// }

func NewAuthService(userRepo repository.UserRepository, userService UserService) AuthService {
	return &authService{
		userRepository: userRepo,
		userService:    userService,
	}
}
