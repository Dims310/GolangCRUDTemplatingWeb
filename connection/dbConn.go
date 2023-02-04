package connection

import (
	"database/sql"
	"rest-api-sim/helper"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func Connection() (DB *sql.DB) {
	DB, err := sql.Open("mysql", "root:@tcp(localhost:3306)/kehadiran?parseTime=true")
	helper.Output(err)

	DB.SetConnMaxIdleTime(15 * time.Minute)
	DB.SetConnMaxLifetime(30 * time.Minute)
	DB.SetMaxIdleConns(5)
	DB.SetMaxOpenConns(10)

	return
}

func Status() bool {
	err := Connection().Ping()
	if err != nil {
		return false
	} else {
		return true
	}
}
