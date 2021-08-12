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

type Shibor struct {
	infoDate   time.Time `json:"date"`
	overNight  float32   `json:"O/N`
	oneWeek    float32   `json:"1W"`
	twoWeek    float32   `json:"2W"`
	oneMonth   float32   `json:"1M"`
	threeMonth float32   `json:"3M"`
	sixMonth   float32   `json:"6M"`
	nineMonth  float32   `json:"9M"`
	oneYear    float32   `json:"1Y"`
}

func Insert(db *sql.DB, infoDate time.Time, overNight float32, oneWeek float32, twoWeek float32, oneMonth float32,
	threeMonth float32, sixMonth float32, nineMonth float32, oneYear float32) {
	stms, err := db.Prepare("INSERT INTO shibor(infodate, 1n, 1w, 2w, 1m, 3m, 6m, 9m, 1y) VALUES(?, ?, ?, ?, ?, ?, ?, ?, ?)")
	if err != nil {
		log.Printf("insert data error: %v\n", err)
		return
	}
	_, err = stms.Exec(infoDate, overNight, oneWeek, twoWeek, oneMonth, threeMonth, sixMonth, nineMonth, oneYear)
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

	shibor := Shibor{}
	defer db.Close()
	r := gin.Default()
	r.POST("/shibor", func(c *gin.Context) {
		Insert(db, shibor.infoDate, shibor.overNight, shibor.oneWeek, shibor.twoWeek, shibor.oneMonth,
			shibor.threeMonth, shibor.sixMonth, shibor.nineMonth, shibor.oneYear)
		c.JSON(http.StatusOK, "shibor")
	})
	r.Run(":8080")
}
