package mysql

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

var (
	MYSQL_HOST = "MYSQL_HOST"
	MYSQL_PORT = "MYSQL_PORT"
	MYSQL_USER = "MYSQL_USER"
	MYSQL_PASS = "MYSQL_PASS"
	MYSQL_DB   = "MYSQL_DB"
)

func NewMysqlConnection() (*sql.DB, error) {
	mysql_user := os.Getenv(MYSQL_USER)
	mysql_pass := os.Getenv(MYSQL_PASS)
	mysql_host := os.Getenv(MYSQL_HOST)
	mysql_port := os.Getenv(MYSQL_PORT)
	mysql_db := os.Getenv(MYSQL_DB)

	conn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", mysql_user, mysql_pass, mysql_host, mysql_port, mysql_db)

	db, err := sql.Open("mysql", conn)
	if err != nil {
		return nil, err
	}

	return db, nil
}
