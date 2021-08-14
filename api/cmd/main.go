package main

import (
	"database/sql"
	"log"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/shuaiqidechuan/project/api/controller/tables"
)

func main() {
	r := gin.Default()
	dbConn, err := sql.Open("mysql", "root:123456@tcp(192.168.0.251:6666)/project")
	if err != nil {
		panic(err)
	}
	Conn := tables.New(dbConn)
	Conn.RegisterRouter(r.Group("api"))
	log.Fatal(r.Run(":8080"))
}
