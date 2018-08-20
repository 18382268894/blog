/**
*FileName: db
*Create on 2018/8/20 上午11:30
*Create by mok
*/

package db

import "fmt"

const (
	USERNAME = "root"
	PASSWORD = "chen19950210"
	NETWORK  = "tcp"
	SERVER   = "localhost"
	PORT     = 3306
	DATABASE = "blog"
)

func init(){
	dsn := fmt.Sprintf("%s:%s@%s(%s:%d)/%s?parseTime=true", USERNAME, PASSWORD, NETWORK, SERVER, PORT, DATABASE)
	err := Init(dsn)
	if err != nil{
		panic(err)
	}
}