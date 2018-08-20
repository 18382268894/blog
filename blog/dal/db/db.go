/**
*FileName: db
*Create on 2018/8/14 下午12:07
*Create by mok
*/

package db

import (
	_"github.com/go-sql-driver/mysql"
	"time"
	"github.com/jmoiron/sqlx"
)

var DB *sqlx.DB
func Init(dsn string)error{

	var err error
	DB, err = sqlx.Open("mysql", dsn)
	if err != nil {
		return err
	}
	err = DB.Ping()
	if err != nil {
		return err
	}
	DB.SetConnMaxLifetime(100 * time.Second)
	DB.SetMaxOpenConns(100)
	DB.SetMaxIdleConns(16)
	return nil
}


