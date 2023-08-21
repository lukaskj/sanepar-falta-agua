package main

import (
	"fmt"
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
	log.Println()
	log.Println("[+] Starting...")
	for {
		response := SendSaneparRequest(&config.Config.SaneparBaseUrl, &config.Config.SaneparClientId)
		if IsElegibleToSendNotification(&response) {

			messageToSend := fmt.Sprintf("%s\nPrevisão: %s %s\nNormalização: %s %s", config.Config.Env, response.PrevisaoData, response.PrevisaoHora, response.NormalizacaoData, response.NormalizacaoHora)
			log.Printf("[+] Message: %s\n", messageToSend)
			if !config.IsNotificationSent(&response) {
				SendNotificationMessage(messageToSend)
				config.SetNotificationSent(&response, true)
			} else {
				log.Println("[-] Message already sent before...")
			}
		}
		time.Sleep(time.Duration(config.Config.TimeLoopSeconds) * time.Second)
	}
}
