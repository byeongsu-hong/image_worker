package main

import (
	"./downloader"
	"log"
	"github.com/valyala/gorpc"
	"strings"
	"path"
)

var epTitle = make(chan string)
var workStack = make(chan string)
var done = make(chan bool)

func main() {
	go downloader.Run(epTitle, workStack, done)

	s := &gorpc.Server{
		Addr: ":5001",
		Handler: func(clientAddr string, request interface{}) interface{} {
			if strings.Contains(request.(string), "folder") {
				epTitle <- path.Base(request.(string))
			} else if request.(string) == "done" {
				done <- true
			} else {
				workStack <- request.(string)
			}

			return true
		},
	}

	if err := s.Serve(); err != nil {
		log.Fatalln(err)
	}
}