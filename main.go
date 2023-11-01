package main

import (
	//...
	"fmt"
	"net/http"

	"github.com/corazawaf/coraza/v3"
	"github.com/corazawaf/coraza/v3/seclang"
	"github.com/gin-gonic/gin"
)

// Initialize Coraza
var waf = coraza.NewWAF()

func main() {

	// Gin Router
	router := gin.Default()

	// TODO : Connect to QuestDB

	// TODO : Connect to SQLite

	// Load Rules
    parser := seclang.NewParser(waf)
	parser.FromFile("default.conf")
	fmt.Println(waf.Rules)

	router.Use(middlewareTest())

	router.GET("/ping", func(c *gin.Context) {
		c.HTML(200,"index.html", gin.H{
			"title": "National Api Day",
		})
	})

	router.POST("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong(POST)",
		})
	})

	router.LoadHTMLGlob("frontend/*.html") 

	router.Run() // localhost:8080
}

func middlewareTest() gin.HandlerFunc{
	return func(c *gin.Context){
		tx := waf.NewTransaction(c.Request.Context())
		fmt.Println(tx.Logdata)
		fmt.Println(tx.ProcessRequestBody())

		// rule := coraza.NewRule()
		// rule.AddAction("sqlijection")

		// waf.Rules.Add()/

		it := tx.ProcessRequestHeaders()
		fmt.Println(it)

		fmt.Printf("%s \n",c.Request.Proto)
		fmt.Printf("%s   : %s\n", c.Request.Method, c.Request.URL.Path)
		fmt.Printf("IP   : %s\n", c.ClientIP())
		fmt.Printf("Body : %s\n", c.Request.Body)
	}
}