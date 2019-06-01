package files

import (
	"io/ioutil"
	"log"
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
func ListFolder(folder string) DirInfo {
	files, err := ioutil.ReadDir(folder)
	if err != nil {
		log.Fatal(err)
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

	return info
}
