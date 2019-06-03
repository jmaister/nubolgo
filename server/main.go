package main

import (
	"flag"
	"log"
	"net/http"
	"path/filepath"

	"./files"

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

	// TODO: check if folder exists, or create

	router := gin.Default()

	clientPath, _ := filepath.Abs("client/dist")
	log.Println("Using client files from: ", clientPath)

	router.Use(static.Serve("/", static.LocalFile(clientPath, true)))

	api := router.Group("/api")
	{
		api.GET("/", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "pong",
			})
		})
	}

	api.GET("/files", GetFilesHandler)
	api.POST("/files", PostFilesHandler)

	router.Run(":3000")
}

// GetFilesHandler returns a list of files
func GetFilesHandler(c *gin.Context) {
	path := c.Query("path")

	if len(path) == 0 {
		path = "/"
	}

	data, err := files.ListFolder(config.RootFolder, path)
	if err != nil {
		c.Header("Content-Type", "application/json")
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	} else {
		c.Header("Content-Type", "application/json")
		c.JSON(http.StatusOK, data)
	}
}

// PostFilesHandler stores files from clients
func PostFilesHandler(c *gin.Context) {
	form, _ := c.MultipartForm()
	files := form.File["upload[]"]
	path := form.Value["path"]

	log.Println("path:", path)

	for _, file := range files {
		log.Println(file.Filename)
		log.Println(config.RootFolder + "/" + file.Filename)

		//c.SaveUploadedFile(file, config.RootFolder+path+'/'+file.Filename)
	}

	c.JSON(http.StatusOK, gin.H{
		"OK": true,
	})
}
