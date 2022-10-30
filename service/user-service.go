package service

import (
	"github.com/zhuliminl/mc_server/entity"
	"github.com/zhuliminl/mc_server/repository"
)

type UserService interface {
	Profile(id string) entity.User
}

type userService struct {
	userRepository repository.UserRepository
}

func NewUserService(userRepo repository.UserRepository) UserService {
	return &userService{
		userRepository: userRepo,
	}
}

func (service *userService ) Profile (id string) entity.User {
	return service.userRepository.ProfileUser(id)
}