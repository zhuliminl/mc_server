package service

import (
	"github.com/zhuliminl/mc_server/dto"
	"github.com/zhuliminl/mc_server/repository"
)

type AuthService interface {
	VerifyCredential(email string, password string) (dto.UserVerify, error)
	//CreateUser(user dto.UserRegister) (dto.User, error)
	//FindByEmail(email string) (dto.User, error)
	//isDuplicateEmail(email string) (bool, error)
}

type authService struct {
	userRepository repository.UserRepository
}

func (service *authService) VerifyCredential(email string, password string) (dto.UserVerify, error) {
	return dto.UserVerify{
		IsValid: true,
	}, nil
}

// func (service *authService) CreateUser(user dto.UserRegister) (dto.User, error) {
// }
//
// func (service *authService) FindByEmail(email string) (dto.User, error) {
// }
//
// func (service *authService) isDuplicateEmail(email string) (bool, error) {
// }

func NewAuthService(userRepo repository.UserRepository) AuthService {
	return &authService{
		userRepository: userRepo,
	}
}
