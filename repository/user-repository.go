package repository

import (
	"database/sql"
	"log"

	"github.com/zhuliminl/mc_server/database"
	"github.com/zhuliminl/mc_server/entity"
)

type UserRepository interface {
	ProfileUser(id string) entity.User
	UpdateUser(id string) entity.User
	CreateUser(user entity.User) entity.User
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
	user.Username = "saul"
	return user
}

func (db *userConnection) UpdateUser(id string) entity.User {
	var user entity.User
	user.Username = "saul"
	return user

}

func (db *userConnection) CreateUser(user entity.User) entity.User {
	stmt, err := db.connection.Prepare(database.CreateUser)
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	_, err1 :=  stmt.Exec(
		user.UserId, 
		user.Username, 
		user.Email, 
		user.Phone, 
		user.Password,
	)
	if err1 != nil {
		log.Fatal("create-user-err1",err1)
	}

	// if _, err := stmt.Exec(id+1, project.mascot, project.release, "open source"); err != nil {
	// 	log.Fatal(err)
	// }
	return user
}
