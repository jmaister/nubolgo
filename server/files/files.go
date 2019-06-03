package files

import (
	"errors"
	"io/ioutil"
	"log"
	"path/filepath"
	"strings"
	"time"
)

// DirInfo stores folder data recursively
type DirInfo struct {
	Name     string    `json:"name"`
	IsFolder bool      `json:"isFolder"`
	Size     int64     `json:"size"`
	Time     time.Time `json:"time"`
	Files    []DirInfo `json:"files"`
}

// ListFolder lists the contents of a folder (not recursively)
func ListFolder(rootFolder string, folder string) (DirInfo, error) {

	rootFolder, err := filepath.Abs(rootFolder)
	if err != nil {
		return DirInfo{}, err
	}

	fullPath, err := filepath.Abs(rootFolder + folder)
	if err != nil {
		return DirInfo{}, err
	}

	// TODO: check if folder is under rootFolder
	log.Println("ListFolder root:", rootFolder)
	log.Println("ListFolder full: ", fullPath)

	if !strings.HasPrefix(fullPath, rootFolder) {
		return DirInfo{}, errors.New("do not try to hack")
	}

	files, err := ioutil.ReadDir(fullPath)
	if err != nil {
		// log.Fatal(err)
		return DirInfo{}, err
	}

	Files := []DirInfo{}

	for _, f := range files {
		Files = append(Files, DirInfo{
			Name:     f.Name(),
			IsFolder: f.IsDir(),
			Size:     f.Size(),
			Time:     f.ModTime(),
		})
	}

	info := DirInfo{
		Name:     folder,
		IsFolder: true,
		Size:     0,
		Files:    Files,
	}

	return info, nil
}

/*
// ListFolder lists the contents of a folder (not recursively)
func SaveFile(path string, file File) (bool, error) {

}
*/
