package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

var db = make(map[string]string)

func setupRouter() *gin.Engine {
	// Disable Console Color
	// gin.DisableConsoleColor()
	r := gin.Default()

	// Ping test
	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	r.GET("/test", func(c *gin.Context) {
		code := `package main
			
		import (
			"fmt"
			"math/rand"
		)
		func main() {
			var ran int	
			rand.Seed(1)
			ran = rand.Intn(10)
			fmt.Println(ran)
		}`
		AnalysisAst(code)
		ok := true
		if ok {
			c.JSON(http.StatusOK, gin.H{"user": "user", "value": "value"})
		} else {
			c.JSON(http.StatusOK, gin.H{"user": "user", "status": "no value"})
		}
	})

	r.POST("/upload", func(c *gin.Context) {
		c.Request.ParseMultipartForm(1000)
		c.Header("Access-Control-Allow-Headers", "Content-Type, Authorization, Origin")
		c.Header("Access-Control-Allow-Credentials", "true")
		c.Header("Access-Control-Allow-Origin", "http://127.0.0.8:8080")
		c.Header("Access-Control-Allow-Methods", "POST")

		code := c.Request.PostForm["code"][0]
		//fmt.Println(code)
		c.String(http.StatusOK, AnalysisAst(code))
	})

	return r
}

func main() {
	r := setupRouter()
	// Listen and Server in 0.0.0.0:8080
	r.Run(":8080")
}
