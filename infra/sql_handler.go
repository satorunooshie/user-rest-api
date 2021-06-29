package infra

import (
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/jinzhu/gorm"

	"github.com/satorunooshie/user-rest-api/infra/database"
)

type sqlHandler struct {
	conn *gorm.DB
}

func NewMySQLDB() database.SQLHandler {
	conn, err := open(
		fmt.Sprintf(
			"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=true&loc=Local",
			os.Getenv("MYSQL_USER"),
			os.Getenv("MYSQL_PASSWORD"),
			os.Getenv("MYSQL_HOST"),
			os.Getenv("MYSQL_PORT"),
			os.Getenv("MYSQL_DATABASE"),
		),
		30,
	)
	if err != nil {
		panic(err)
	}
	if err := conn.DB().Ping(); err != nil {
		panic(err)
	}
	conn.LogMode(true)
	conn.Set("gorm:table_options", "ENGIN=InnoDB")
	sqlHandler := new(sqlHandler)
	sqlHandler.conn = conn
	return sqlHandler
}

// MySQLが立ち上がってからAPIコンテナが立ち上がるようにする
func open(path string, count uint) (*gorm.DB, error) {
	db, err := gorm.Open("mysql", path)
	if err != nil {
		if count == 0 {
			return nil, errors.New("retry count over")
		}
		time.Sleep(time.Second)
		count--
		return open(path, count)
	}
	return db, nil
}

func (h *sqlHandler) Find(out interface{}, where ...interface{}) *gorm.DB {
	return h.conn.Find(out, where...)
}

func (h *sqlHandler) Create(v interface{}) *gorm.DB {
	return h.conn.Create(v)
}

func (h *sqlHandler) Save(v interface{}) *gorm.DB {
	return h.conn.Save(v)
}

func (h *sqlHandler) Delete(v interface{}) *gorm.DB {
	return h.conn.Delete(v)
}
