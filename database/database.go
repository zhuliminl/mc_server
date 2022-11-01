package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/rs/zerolog"
	sqldblogger "github.com/simukti/sqldb-logger"
	"github.com/simukti/sqldb-logger/logadapter/zerologadapter"
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

	loggerAdapter := zerologadapter.New(zerolog.New(os.Stdout))
	DB = sqldblogger.OpenDriver(GetDbUrl(), DB.Driver(),
		loggerAdapter,
		// sqldblogger.WithErrorFieldname("sql_error"),                  // default: error
		// sqldblogger.WithDurationFieldname("query_duration"),          // default: duration
		// sqldblogger.WithTimeFieldname("log_time"),                    // default: time
		// sqldblogger.WithSQLQueryFieldname("sql_query"),               // default: query
		// sqldblogger.WithSQLArgsFieldname("sql_args"),                 // default: args
		// sqldblogger.WithMinimumLevel(sqldblogger.LevelTrace),         // default: LevelDebug
		// sqldblogger.WithLogArguments(false),                          // default: true
		// sqldblogger.WithDurationUnit(sqldblogger.DurationNanosecond), // default: DurationMillisecond
		// sqldblogger.WithTimeFormat(sqldblogger.TimeFormatRFC3339),    // default: TimeFormatUnix
		// sqldblogger.WithLogDriverErrorSkip(true),                     // default: false
		// sqldblogger.WithSQLQueryAsMessage(true),                      // default: false
		// // sqldblogger.WithUIDGenerator(sqldblogger.UIDGenerator),       // default: *defaultUID
		// sqldblogger.WithConnectionIDFieldname("con_id"),       // default: conn_id
		// sqldblogger.WithStatementIDFieldname("stm_id"),        // default: stmt_id
		// sqldblogger.WithTransactionIDFieldname("trx_id"),      // default: tx_id
		// sqldblogger.WithWrapResult(false),                     // default: true
		// sqldblogger.WithIncludeStartTime(true),                // default: false
		// sqldblogger.WithStartTimeFieldname("start_time"),      // default: start
		// sqldblogger.WithPreparerLevel(sqldblogger.LevelDebug), // default: LevelInfo
		// sqldblogger.WithQueryerLevel(sqldblogger.LevelDebug),  // default: LevelInfo
		// sqldblogger.WithExecerLevel(sqldblogger.LevelDebug),   // default: LevelInfo
		/*, using_default_options*/) // db is STILL *sql.DB

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
