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
		if response.Mensagem != "NADA CONSTA" {

			messageToSend := fmt.Sprintf("\nPrevisão: %s %s\nNormalização: %s %s", response.PrevisaoData, response.PrevisaoHora, response.NormalizacaoData, response.NormalizacaoHora)
			log.Printf("[+] Message: %s\n", messageToSend)
			if !config.IsNotificationSentToday() {
				SendNotificationMessage(messageToSend)
				config.SetNotificationSentToday(true)
			} else {
				log.Println("[-] Message already sent before...")
			}
		}
		time.Sleep(time.Duration(config.Config.TimeLoopSeconds) * time.Second)
	}
}
