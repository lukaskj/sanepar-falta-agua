package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/lukaskj/sanepar-falta-agua/utils"
)

type TConfig struct {
	emailSentAt     map[string]bool
	saneparBaseUrl  string
	saneparClientId string
}

var Config TConfig

func Load() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	Config = TConfig{emailSentAt: make(map[string]bool)}
	Config.saneparBaseUrl = os.Getenv("SANEPAR_BASE_URL")
	Config.saneparClientId = os.Getenv("SANEPAR_CLIENT_ID")

	loadEmailSentJson()

	println("Load config")
}

func IsEmailSentAt(dateStr string) bool {
	return Config.emailSentAt[dateStr]
}

func SetEmailSentAt(dateStr string, val bool) {
	Config.emailSentAt[dateStr] = val
}

func IsEmailSentToday() bool {
	date := utils.CurrentDateStr()

	return Config.emailSentAt[date]
}

func SaveConfigJson() {
	
}


func loadEmailSentJson() {}