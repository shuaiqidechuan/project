package tables

import (
	"database/sql"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/shuaiqidechuan/project/api/controller/mysql"
)

type Controller struct {
	db *sql.DB
}

func New(db *sql.DB) *Controller {
	return &Controller{
		db: db,
	}
}

func (s *Controller) RegisterRouter(r gin.IRouter) {
	if r == nil {
		log.Fatal("[InitRouter]: server is nil")
	}
	r.POST("/lpr", s.insertLpr)
	r.POST("/shibor", s.insertShibor)
	r.GET("/shibor", s.queryShibor)
	r.GET("lpr", s.queryLpr)
}

func (s *Controller) insertShibor(c *gin.Context) {
	var (
		req struct {
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
	)
	err := c.ShouldBind(&req)
	if err != nil {
		c.Error(err)
		c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest})
		return
	}

	id, err := mysql.InsertShibor(s.db, req.InfoDate, req.OverNight, req.OneWeek, req.TwoWeek,
		req.OneMonth, req.ThreeMonth, req.SixMonth, req.NineMonth, req.OneYear)
	if err != nil {
		c.Error(err)
		c.JSON(http.StatusBadGateway, gin.H{"status": http.StatusBadGateway})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "ID": id})

}

func (s *Controller) insertLpr(c *gin.Context) {
	var (
		req struct {
			InfoDate time.Time `json:"date"`
			OneYear  float32   `json:"1Y"`
			FiveYear float32   `json:"5Y"`
		}
	)
	err := c.ShouldBind(&req)
	if err != nil {
		c.Error(err)
		c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest})
		return
	}

	id, err := mysql.InsertLpr(s.db, req.InfoDate, req.OneYear, req.FiveYear)
	if err != nil {
		c.Error(err)
		c.JSON(http.StatusBadGateway, gin.H{"status": http.StatusBadGateway})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "ID": id})

}

func (s *Controller) queryShibor(c *gin.Context) {
	var (
		req struct {
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
	)

	err := c.ShouldBind(&req)
	if err != nil {
		c.Error(err)
		c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest})
		return
	}

	shibor, err := mysql.QueryShibor(s.db)
	if err != nil {
		c.Error(err)
		c.JSON(http.StatusBadGateway, gin.H{"status": http.StatusBadGateway})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "shibor": shibor})
}

func (s *Controller) queryLpr(c *gin.Context) {
	var (
		req struct {
			InfoDate time.Time `json:"date"`
			OneYear  float32   `json:"1Y"`
			FiveYear float32   `json:"5Y"`
		}
	)

	err := c.ShouldBind(&req)
	if err != nil {
		c.Error(err)
		c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest})
		return
	}

	lpr, err := mysql.QueryLpr(s.db)
	if err != nil {
		c.Error(err)
		c.JSON(http.StatusBadGateway, gin.H{"status": http.StatusBadGateway})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "lpr": lpr})
}
