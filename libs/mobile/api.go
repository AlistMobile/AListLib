package mobile

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func RunAPIServer() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"code":    0,
			"message": "pong",
		})
	})
	r.GET("/set-config-data", func(c *gin.Context) {
		path := c.DefaultQuery("path", "")
		SetConfigData(path)
		c.JSON(http.StatusOK, gin.H{
			"code":    0,
			"message": "success",
		})
	})
	r.GET("/set-admin-password", func(c *gin.Context) {
		pwd := c.DefaultQuery("password", "")
		SetAdminPassword(pwd)
		c.JSON(http.StatusOK, gin.H{
			"code":    0,
			"message": "success",
		})
	})
	r.GET("/init", func(c *gin.Context) {
		err := Init(eventEntity, callback)
		if err != nil {
			log.Println(err.Error())
			c.JSON(http.StatusOK, gin.H{
				"code":    1,
				"message": err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"code":    0,
			"message": "success",
		})
	})
	r.GET("/status", func(c *gin.Context) {
		Start()
		c.JSON(http.StatusOK, gin.H{
			"code":         0,
			"message":      "success",
			"http_status":  IsRunning("http"),
			"https_status": IsRunning("https"),
		})
	})
	r.GET("/start", func(c *gin.Context) {
		Start()
		c.JSON(http.StatusOK, gin.H{
			"code":    0,
			"message": "success",
		})
	})
	r.GET("/shutdown", func(c *gin.Context) {
		ms := c.GetInt64("ms")
		if ms == 0 {
			ms = 1000
		}
		err := Shutdown(ms)
		if err != nil {
			log.Println(err.Error())
			c.JSON(http.StatusOK, gin.H{
				"code":    1,
				"message": err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"code":    0,
			"message": "success",
		})
	})
	err := r.Run("0.0.0.0:15244")
	if err != nil {
		log.Println(err.Error())
	}
}
