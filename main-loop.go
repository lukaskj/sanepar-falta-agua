package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/lukaskj/sanepar-falta-agua/config"
)

var quitChannel = make(chan os.Signal, 1)

func Start() {
	errLoadConfig := config.Load()
	if errLoadConfig != nil {
		panic(errLoadConfig)
	}

	go mainLoop()

	//
	signal.Notify(quitChannel, syscall.SIGINT, syscall.SIGTERM)
	<-quitChannel

	config.SaveConfigJson()

	log.Println("Exiting...")
}

func mainLoop() {
	for {
		SendSaneparRequest(&config.Config.SaneparBaseUrl, &config.Config.SaneparClientId)
		time.Sleep(time.Duration(config.Config.TimeLoopSeconds) * time.Second)
	}
}
