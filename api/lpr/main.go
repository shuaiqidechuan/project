package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

type Lpr struct {
	infoDate time.Time `json:"date"`
	oneYear  float32   `json:"1Y"`
	fiveYear float32   `json:"5Y`
}

func Insert(db *sql.DB, infoDate time.Time, oneYear float32, fiveYear float32) {
	stms, err := db.Prepare("INSERT INTO lpr(infodate, 1y, 5y) VALUES(?, ?, ?)")
	if err != nil {
		log.Printf("insert data error: %v\n", err)
		return
	}
	_, err = stms.Exec(infoDate, oneYear, fiveYear)
	if err != nil {
		log.Println(err)
	}
}

func main() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("connect fail", err)
		}
	}()
	db, err := sql.Open("mysql", "root:123456@tcp(192.168.0.251:6666)/project")
	if err != nil {
		panic(err)
	}

	lpr := Lpr{}
	defer db.Close()
	r := gin.Default()
	r.POST("/lpr", func(c *gin.Context) {
		Insert(db, lpr.infoDate, lpr.oneYear, lpr.fiveYear)
		c.JSON(http.StatusOK, "lpr")
	})
	r.Run(":8080")
}
