package repository

import (
	"database/sql"
	"log"

	"github.com/zhuliminl/mc_server/database"
	"github.com/zhuliminl/mc_server/dto"
	"github.com/zhuliminl/mc_server/entity"
)

type UserRepository interface {
	Get(id string) (entity.User, error)
	GetAll() ([]dto.UserAll, error)
	Update(id string) (entity.User, error)
	Create(user entity.User) (entity.User, error)
	Delete(userId string) error

	// GetAll() []interface{}
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

func (db *userConnection) Get(userId string) (entity.User, error) {
	var user entity.User
	err := db.connection.QueryRow(database.FindUserByUserId, userId).Scan(
		&user.UserId,
		&user.Username,
		&user.Email,
		&user.Phone,
		&user.WechatNickname,
		&user.WechatNumber,
	)

	switch {
	case err == sql.ErrNoRows:
		log.Println("saul ------------->>>> no user with id", userId)
		return user, err
	case err != nil:
		log.Println("query Error", err)
		return user, err
	default:
		return user, nil
	}
}

func (db *userConnection) GetAll() ([]dto.UserAll, error) {
	var allUsers []dto.UserAll

	rows, err := db.connection.Query(database.FindUserAll)
	if err != nil {
		// log.Panicln("db-find-all-user-err", err)
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var (
			userId         sql.NullString
			username       sql.NullString
			email          sql.NullString
			phone          sql.NullString
			wechatNickname sql.NullString
		)
		if err := rows.Scan(
			&userId,
			&username,
			&email,
			&phone,
			&wechatNickname,
		); err != nil {
			log.Println("db-scan-all-user-err", err)
			return nil, err
		}
		user := dto.UserAll{
			UserId:         userId.String,
			Username:       username.String,
			Email:          email.String,
			Phone:          phone.String,
			WechatNickname: wechatNickname.String,
		}
		allUsers = append(allUsers, user)
	}
	return allUsers, nil
}

func (db *userConnection) Update(id string) (entity.User, error) {
	var user entity.User
	user.Username = "saul"
	return user, nil
}

func (db *userConnection) Create(user entity.User) (entity.User, error) {
	stmt, err := db.connection.Prepare(database.CreateUser)
	if err != nil {
		return user, err
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
		return user, err
	}
	return user, nil
}

func (db *userConnection) Delete(userId string) error {
	var (
		id       int
		_userId  string
		username string
	)
	err := db.connection.QueryRow(database.FindUserByUserId, userId).Scan(&id, &_userId, &username)
	switch {
	case err == sql.ErrNoRows:
		log.Println("no user with id", userId)
		return err
	case err != nil:
		log.Println("query Error", err)
		return err
	default:
		return err
	}
}
