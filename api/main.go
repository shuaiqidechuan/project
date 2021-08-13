package main

import (
	"database/sql"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

type Shibor struct {
	InfoDate   time.Time `json:"date"`
	OverNight  float32   `json:"O/N"`
	OneWeek    float32   `json:"1W"`
	TwoWeek    float32   `json:"2W"`
	OneMonth   float32   `json:"1M"`
	ThreeMonth float32   `json:"3M"`
	SixMonth   float32   `json:"6M"`
	NineMonth  float32   `json:"9M"`
	OneYear    float32   `json:"1Y"`
}

type Lpr struct {
	InfoDate time.Time `json:"date"`
	OneYear  float32   `json:"1Y"`
	FiveYear float32   `json:"5Y"`
}

func InsertShibor(db *sql.DB, infoDate time.Time, overNight float32, oneWeek float32, twoWeek float32, oneMonth float32,
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

func InsertLpr(db *sql.DB, infoDate time.Time, oneYear float32, fiveYear float32) {
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
	db, err := sql.Open("mysql", "root:123456@tcp(192.168.0.1:6666)/project")
	if err != nil {
		log.Println("connect fail", err)
	}

	shibor := Shibor{}
	lpr := Lpr{}
	defer db.Close()
	r := gin.Default()
	r.POST("/shibor", func(c *gin.Context) {
		InsertShibor(db, shibor.InfoDate, shibor.OverNight, shibor.OneWeek, shibor.TwoWeek, shibor.OneMonth,
			shibor.ThreeMonth, shibor.SixMonth, shibor.NineMonth, shibor.OneYear)
		c.JSON(http.StatusOK, "shibor")
	})
	r.POST("/lpr", func(c *gin.Context) {
		InsertLpr(db, lpr.InfoDate, lpr.OneYear, lpr.FiveYear)
		c.JSON(http.StatusOK, "lpr")
	})
	r.Run(":8080")
}
