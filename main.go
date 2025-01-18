package main

import (
	"net/http"

	// import the extractions
	"TestApi/configs"
	"TestApi/routes"

	"github.com/gin-gonic/gin"
)

var db = make(map[string]string)

func setupRouter() *gin.Engine {
	// Disable Console Color
	// gin.DisableConsoleColor()
	r := gin.Default()

	// / endpoint
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"data": "Hello World",
		})
	})

	authorized := r.Group("/", gin.BasicAuth(gin.Accounts{
		"foo":  "bar", // user:foo password:bar
		"manu": "123", // user:manu password:123
	}))

	authorized.POST("admin", func(c *gin.Context) {
		user := c.MustGet(gin.AuthUserKey).(string)

		// Parse JSON
		var json struct {
			Value string `json:"value" binding:"required"`
		}

		if c.Bind(&json) == nil {
			db[user] = json.Value
			c.JSON(http.StatusOK, gin.H{"status": "ok"})
		}
	})

	return r
}

func main() {
	// Set the router as the default one provided by Gin
	r := setupRouter()

	//run database
	configs.ConnectDB()

	// routes
	routes.UserRoute(r)

	// Listen and Server in 0.0.0.0:8080
	r.Run(":8080")

}
