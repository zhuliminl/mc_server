package repository

import (
	"database/sql"

	"github.com/zhuliminl/mc_server/entity"
)

type UserRepository interface {
	ProfileUser(id string) entity.User
	UpdateUser(id string) entity.User
}

type userConnection struct {
	connection *sql.DB
}

func NewUserRepository(db *sql.DB) UserRepository {
	return &userConnection{
		connection: db,
	}
}

func (db *userConnection) ProfileUser(id string) entity.User {
	var user entity.User
	user.Name = "saul"
	return user
}


func (db *userConnection) UpdateUser(id string) entity.User {
	var user entity.User
	user.Name = "saul"
	return user

}