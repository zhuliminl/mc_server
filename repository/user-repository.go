package repository

import (
	"database/sql"
	"log"

	"github.com/zhuliminl/mc_server/database"
	"github.com/zhuliminl/mc_server/entity"
)

type UserRepository interface {
	Get(id string) entity.User
	Update(id string) entity.User
	Create(user entity.User) entity.User

	// List() ([]model.NtpServer, error)
	// Save(plan *model.Plan, zones []string) error
	// GetById(id string) (model.Plan, error)
	// Batch(operation string, items []model.Zone) error
	// Batch(operation string, items []model.Zone) error
	// ListByRegionId(id string) ([]model.Zone, error)
	// Page(num, size int) (int, []model.Zone, error)
	// Delete(name string) error
}

type userConnection struct {
	connection *sql.DB
}

func NewUserRepository(db *sql.DB) UserRepository {
	return &userConnection{
		connection: db,
	}
}

func (db *userConnection) Get(id string) entity.User {
	var user entity.User
	user.Username = "saul"
	return user
}

func (db *userConnection) Update(id string) entity.User {
	var user entity.User
	user.Username = "saul"
	return user

}

func (db *userConnection) Create(user entity.User) entity.User {
	stmt, err := db.connection.Prepare(database.CreateUser)
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	_, err1 := stmt.Exec(
		user.UserId,
		user.Username,
		user.Email,
		user.Phone,
		user.Password,
	)
	if err1 != nil {
		log.Fatal("create-user-err1", err1)
	}

	// if _, err := stmt.Exec(id+1, project.mascot, project.release, "open source"); err != nil {
	// 	log.Fatal(err)
	// }
	return user
}
