package database

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func InitDB() {
	fmt.Println("init db")
	// ConnectDB()
}

func ConnectDB() sql.DB {
	DB, err := sql.Open("mysql", GetDbUrl())

	// execSQL(createAdHead)

	var name string
	err = DB.QueryRow("select headId from adHead where id = ?", 1).Scan(&name)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(name)

	if err != nil {
		log.Println("database => connect-db-error:", err)
	}

	_err := DB.Ping()
	if _err != nil {
		fmt.Println("db ===>> ", _err)
		log.Println("database => connect-db-error:", _err)
	}
	fmt.Println("db ===>> ", DB)
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

const createAdHead = `
  CREATE TABLE IF NOT EXISTS adHead (
    id INTEGER UNSIGNED NOT NULL PRIMARY KEY AUTO_INCREMENT,
    headId INTEGER UNSIGNED COMMENT '团长id',
    headName VARCHAR(32) COMMENT '团长名称',
    dodokCommission INTEGER UNSIGNED COMMENT '多多客佣金%',
    headCommission INTEGER UNSIGNED COMMENT '团长佣金%',
    coupon INTEGER UNSIGNED COMMENT '优惠券金额',
    wechatNickname VARCHAR(32) COMMENT '微信昵称',
    wechatNumber VARCHAR(20) COMMENT '微信号',
    pddNickname VARCHAR(32) COMMENT '拼多多昵称'
  );
`
