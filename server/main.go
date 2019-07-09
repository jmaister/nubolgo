package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"./files"

	"github.com/gin-contrib/static"
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
	rootPath, err := filepath.Abs(config.RootFolder)
	if err != nil {
		log.Println("Error:", err)
		return
	}
	log.Println("Using root path: ", rootPath)

	router := gin.Default()

	clientPath, err := filepath.Abs("dist")
	if err != nil {
		log.Println("Error:", err)
		return
	}
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
	api.GET("/download", DownloadFileHandler)

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

// DownloadFileHandler is used to download a file from server
func DownloadFileHandler(c *gin.Context) {
	path := c.Query("path")

	fullPath := filepath.Clean(config.RootFolder + "/" + path)

	file, err := os.Open(fullPath)
	if err != nil {
		fmt.Println("File reading error", err)
		return
	}
	fileInfo, err := file.Stat()
	if err != nil {
		fmt.Println("File stat reading error", err)
		return
	}

	log.Println("--->", fileInfo.Name())

	reader := file
	contentLength := fileInfo.Size()
	contentType := "text/plain" //response.Header.Get("Content-Type")

	extraHeaders := map[string]string{
		"Content-Disposition": "attachment; filename=\"" + fileInfo.Name() + "\"",
	}

	c.DataFromReader(http.StatusOK, contentLength, contentType, reader, extraHeaders)
}
