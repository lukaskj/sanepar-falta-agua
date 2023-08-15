package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/lukaskj/sanepar-falta-agua/config"
)

var quitChannel = make(chan os.Signal, 1)

func Start() {
	config.Load()
	go mainLoop()

	//
	signal.Notify(quitChannel, syscall.SIGINT, syscall.SIGTERM)
	<-quitChannel

	config.SaveConfigJson()

	fmt.Println("Exiting...")
}

func mainLoop() {
	for {
		// fmt.Printf("%v+\n", time.Now())
		time.Sleep(time.Second)
	}
}
