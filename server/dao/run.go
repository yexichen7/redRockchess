
package dao

import (
	"database/sql"
	"github.com/go-redis/redis"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

// Db 全局变量
var Db *sql.DB
var Redisdb *redis.Client


func RUNDB() {
	//启用数据库
	db, err := sql.Open("mysql", "root:277187@tcp(127.0.0.1:3306)/chess?charset=utf8mb4&loc=Local&parseTime=True")
	if err != nil {
		log.Fatal(err)
	}
	Db = db

	//initRedis()
}
