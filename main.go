package main

import (
	"./files"

	"flag"
	"net/http"

	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
)

// Config stores the app configuration
type Config struct {
	RootFolder string
}

var config Config

func main() {

	rootFolder := flag.String("root", "./tmp", "Storage root folder.")
	flag.Parse()

	config.RootFolder = *rootFolder

	router := gin.Default()

	router.Use(static.Serve("/", static.LocalFile("./views", true)))

	api := router.Group("/api")
	{
		api.GET("/", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "pong",
			})
		})
	}

	api.GET("/files", GetFilesHandler)

	router.Run(":3000")
}

// GetFilesHandler returns a list of files
func GetFilesHandler(c *gin.Context) {
	path := c.Query("path")

	if len(path) == 0 {
		path = "/"
	}

	data := files.ListFolder(path)

	// c.BindJSON(&data)
	c.Header("Content-Type", "application/json")
	c.JSON(http.StatusOK, data)
}
