package main

import (
	"net/http"

	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
)

// DirInfo stores folder data recursively
type DirInfo struct {
	Name     string    `json:"name"`
	IsFolder bool      `json:"isFolder"`
	Files    []DirInfo `json:"files"`
}

func main() {
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
	// api.POST("/files", LikeJoke)

	router.Run(":3000")
}

// GetFilesHandler returns a list of files
func GetFilesHandler(c *gin.Context) {

	data := new(DirInfo)
	data.Name = "root"
	data.IsFolder = true
	data.Files = []DirInfo{}

	// c.BindJSON(&data)
	c.Header("Content-Type", "application/json")
	c.JSON(http.StatusOK, data)
}
