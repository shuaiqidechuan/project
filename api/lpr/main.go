package main

import (
	"database/sql"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

type Lpr struct {
	InfoDate time.Time `json:"date"`
	OneYear  float32   `json:"1Y"`
	FiveYear float32   `json:"5Y"`
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
	db, err := sql.Open("mysql", "root:123456@tcp(192.168.0.251:6666)/project")
	if err != nil {
		log.Println("connect fail", err)
	}

	lpr := Lpr{}
	defer db.Close()
	r := gin.Default()
	r.POST("/lpr", func(c *gin.Context) {
		Insert(db, lpr.InfoDate, lpr.OneYear, lpr.FiveYear)
		c.JSON(http.StatusOK, "lpr")
	})
	r.Run(":8080")
}
