package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"jordiburgos.com/nubolgo/files"

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

	// TODO: check that the React app folder exists
	clientPath, err := filepath.Abs("../dist")
	if err != nil {
		log.Println("Error:", err)
		return
	}
	log.Println("Using client files from: ", clientPath)

	router.Use(static.Serve("/", static.LocalFile(clientPath, true)))

	api := router.Group("/api")
	api.GET("/files", GetFilesHandler)
	api.GET("/download", DownloadFileHandler)
	api.POST("/upload", UploadFileHandler)

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
	requestedFile := c.Query("path")

	fullPath := filepath.Clean(config.RootFolder + "/" + requestedFile)

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

	form, err := c.MultipartForm()
	if err != nil {
		c.String(http.StatusBadRequest, fmt.Sprintf("get form err: %s", err.Error()))
		return
	}
	targetPath := form.Value["path"][0]
	files := form.File["files[]"]
	for _, file := range files {
		targetFileAbsolutePath := filepath.Clean(config.RootFolder + "/" + targetPath + "/" + file.Filename)

		// TODO: check max file size
		fmt.Println(file.Size)

		// Create intermediate folders for file.Filename
		targetFolderAbsolutePath := filepath.Clean(filepath.Dir(targetFileAbsolutePath))
		if _, err := os.Stat(targetFolderAbsolutePath); err == nil {
			// Folder exists
		} else if os.IsNotExist(err) {
			// Folder does not exist
			os.MkdirAll(targetFolderAbsolutePath, 0770)
		} else {
			fmt.Println("Unknown error:", err)
		}

		// Upload the file to specific dst.
		err := c.SaveUploadedFile(file, targetFileAbsolutePath)
		if err != nil {
			// TODO: send error response
			fmt.Println("error!", err)
		}
	}

	c.Request.ParseForm()
	for key, value := range c.Request.PostForm {
		fmt.Println("request key-value", key, value)
	}

	c.JSON(http.StatusOK, gin.H{
		"OK": true,
	})
}
