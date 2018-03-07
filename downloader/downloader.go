package downloader

import (
	"os/exec"
	"io/ioutil"
	"log"
	"time"
	"github.com/Pallinder/go-randomdata"
	"path"
	"os"
)

var p = ""

func Run(epTitle <-chan string, workStack <-chan string, done <-chan bool) {
	for {
		select {
		case title := <-epTitle:
			p = getHomeDir() + title
			generateFolder(p)
		case work := <-workStack:
			downloadImage(work)
			time.Sleep(time.Duration(randomdata.Number(1000, 2000)) * time.Millisecond)
		case <-done:
			return
		}
	}
}

func downloadImage(url string) {
	var err error

	err = os.Chdir(p)
	if err != nil {
		log.Panicln(err)
	}

	var data []byte

	data, err = exec.Command("curl", url).Output()
	if err != nil {
		log.Panicln(err)
	}

	err = ioutil.WriteFile(path.Base(url), data, 0700)
	if err != nil {
		log.Panicln(err)
	}

	err = os.Chdir("..")
	if err != nil {
		log.Panicln(err)
	}
}

