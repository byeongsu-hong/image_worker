package downloader

import (
	"os"
	"log"
)

func generateFolder(path string) {
	var err error

	err = os.MkdirAll(path, 0777)
	if err != nil {
		log.Fatalln(err)
	}
}