package repository

import (
	"context"
	"database/sql"
	"log"

	"github.com/zhuliminl/mc_server/database"
	"github.com/zhuliminl/mc_server/entity"
)

var (
	ctx context.Context
)

type UserRepository interface {
	Get(id string) entity.User
	Update(id string) entity.User
	Create(user entity.User) entity.User
	Delete(userId string)

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

func (db *userConnection) Delete(userId string) {
	/*
		stmtFind, err := db.connection.Prepare(database.FindUserByUserId)
		if err != nil {
			log.Fatal("prepare-find-user-err", err)
		}
		defer stmtFind.Close()

		user, err := stmtFind.Exec(userId)
		if err != nil {
			log.Fatal("exec-find-user-err", err)
		}

		fmt.Println("uuuuuuu", user)
	*/

	/*
		rows, err := db.connection.Query(database.FindUserByUserId, userId)
		if err != nil {
			log.Fatal("prepare-delete-user-err", err)
		}
		defer rows.Close()
		for rows.Next() {
			var (
				id       int
				userId   string
				username string
			)
			if err := rows.Scan(&id, &userId, &username); err != nil {
				log.Fatal(err)
			}
			log.Printf("id %d name is %s\n", id, username)
		}
	*/

	// var ctx context.Context
	var (
		id       int
		_userId  string
		username string
	)

	err := db.connection.QueryRowContext(ctx, database.FindUserByUserId, userId).Scan(&id, &_userId, &username)
	switch {
	case err == sql.ErrNoRows:
		log.Println("no user with id", userId)
	case err != nil:
		log.Println("query Error", err)
	default:
		log.Println("-------------------->>", username)
	}

	/*
	stmtDelete, err := db.connection.Prepare(database.DeleteUserByUserId)
	if err != nil {
		log.Fatal("prepare-delete-user-err", err)
	}
	defer stmtDelete.Close()

	if _, err := stmtDelete.Exec(userId); err != nil {
		log.Fatal("exec-delete-user-err", err)
	}
	*/
}
