package utils

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

var (
	Db  *sql.DB
	err error
)

func init() {
	Db, err = sql.Open(
		"mysql",
		"for_test:Hd000000@tcp(rm-2ze4fd5ht58seg8tdpo.mysql.rds.aliyuncs.com)/for_test",
	)
	if err != nil {
		panic(err.Error())
	}
}
