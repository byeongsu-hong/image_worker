package downloader

import (
	"os/user"
	"log"
)

func getHomeDir() (desktop string) {
	myself, err := user.Current()
	if err != nil {
		log.Fatalln(err)
	}
	return myself.HomeDir + "/"
}