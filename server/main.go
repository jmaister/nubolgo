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
	api.GET("/list", GetFilesHandler)
	api.GET("/file", DownloadFileHandler)
	api.POST("/file", UploadFileHandler)
	api.DELETE("/file", DeleteFileHandler)

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

// DownloadFileHandler is used to download a file from server
func DownloadFileHandler(c *gin.Context) {
	path := c.Query("path")

	fullPath := filepath.Clean(config.RootFolder + "/" + path)

	file, err := os.Open(fullPath)
	defer file.Close()
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

// UploadFileHandler is used to handle file uploads
func UploadFileHandler(c *gin.Context) {

	form, _ := c.MultipartForm()
	path := form.Value["path"][0]
	files := form.File["files[]"]
	for _, file := range files {
		fullPath := filepath.Clean(config.RootFolder + "/" + path + "/" + file.Filename)

		// TODO: create intermediate folders for file.Filename

		fmt.Println("filename:", file.Filename)
		fmt.Println("uploading to:", fullPath)

		// Upload the file to specific dst.
		err := c.SaveUploadedFile(file, fullPath)
		if err != nil {
			fmt.Println("error!", err)
		}
	}

	fmt.Println("path:", path)

	c.Request.ParseForm()
	for key, value := range c.Request.PostForm {
		fmt.Println(key, value)
	}

	c.JSON(http.StatusOK, gin.H{
		"OK": true,
	})
}

// DeleteFileHandler is used to delete files
func DeleteFileHandler(c *gin.Context) {

	path := c.Query("path")

	fullPath := filepath.Clean(config.RootFolder + "/" + path)
	log.Println("---> deleting...", fullPath)

	err := os.Remove(fullPath)
	if err != nil {
		fmt.Println("File delete error", err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"OK": true,
	})
}
