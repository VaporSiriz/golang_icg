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

	r.GET("/upload", func(c *gin.Context) {
		c.Request.ParseMultipartForm(1000)
		c.Header("Access-Control-Allow-Headers", "Content-Type, Authorization, Origin")
		c.Header("Access-Control-Allow-Credentials", "true")
		c.Header("Access-Control-Allow-Origin", "http://127.0.0.1:8080")
		c.Header("Access-Control-Allow-Methods", "GET")
		code := c.Request.Form["code"][0]
		c.String(http.StatusOK, AnalysisAst(code))
	})

	r.GET("/index", func(c *gin.Context) {
		// OK 이면 index.html파일에 JSON데이터를 넘겨서 보여줌
		c.HTML(http.StatusOK, "index.html", gin.H{
			"title": "Home Page",
		},
		)
	})

	return r
}

func main() {
	r := setupRouter()
	//r.LoadHTMLFiles("templates/")
	r.LoadHTMLGlob("template/*")
	r.Static("static", "./static")
	// Listen and Server in 0.0.0.0:8080
	r.Run(":8080")
}
