package database

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/zhuliminl/mc_server/config"
)

var DB *sql.DB

func GetDB() *sql.DB {
	return DB
}

func InitDB() {
	connectDB()

	c := config.GetConfig()
	needCreate := c.GetBool("database.enableCreate")
	fmt.Println("初始化数据库", needCreate)
	if needCreate {
		createDB()
	}
}

func connectDB() {
	var err error
	DB, err = sql.Open("mysql", GetDbUrl())
	if err != nil {
		log.Fatal("db-connect-db-fail", err)
	}
}

func createDB() {
	execSQL(createUserTable)
}

func execSQL(sqlStmt string) {
	log.Print("db inited execSQL sqlStmt: ", sqlStmt)
	stmt, err := DB.Prepare(sqlStmt)
	if err != nil {
		log.Println("db-execSQL-prepare-error: ", err)
	}
	defer stmt.Close()
	_, err = stmt.Exec()
	if err != nil {
		log.Println("db-execSQL-exec-error: ", err)
	}
}
