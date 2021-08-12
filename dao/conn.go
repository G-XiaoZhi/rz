package dao

import (
	"database/sql"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	DB *gorm.DB
	err error
)

func init() {
	connStr := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local", user, password, server, port, database)
	sd, err := sql.Open("mysql", connStr)
	DB, err = gorm.Open(mysql.New(mysql.Config{Conn: sd}), &gorm.Config{})
	if err != nil {
		panic(err.Error())
	}
	//DB = DB.Debug()
}