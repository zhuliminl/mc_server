package database

import (
	"database/sql"
	"log"
	// _ "github.com/go-sql-driver/mysql"
	// "github.com/zhuliminl/mc_server/config"
)

var DB *sql.DB

func InitDB() {
	// c := config.GetConfig()

	// reRreate := c.GetBool("database.enableCreate")

	// fmt.Println("初始化数据库", reRreate)
}

func ConnectDB() *sql.DB {

	// DB, err := sql.Open("mysql", GetDbUrl())

	// // execSQL(createAdHead)

	// var name string
	// err = DB.QueryRow("select headId from adHead where id = ?", 1).Scan(&name)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// fmt.Println(name)

	// if err != nil {
	// 	log.Println("database => connect-db-error:", err)
	// }

	// _err := DB.Ping()
	// if _err != nil {
	// 	fmt.Println("db ===>> ", _err)
	// 	log.Println("database => connect-db-error:", _err)
	// }
	// fmt.Println("db ===>> ", DB)
	return DB
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
