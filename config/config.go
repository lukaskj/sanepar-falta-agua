package config

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/lukaskj/sanepar-falta-agua/utils"
)

type TConfig struct {
	emailSentAt     map[string]bool
	saneparBaseUrl  string
	saneparClientId string
	jsonFileName    string
}

var Config TConfig

func Load() {
	fmt.Println("Load config")

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	Config = TConfig{emailSentAt: make(map[string]bool)}
	Config.saneparBaseUrl = os.Getenv("SANEPAR_BASE_URL")
	Config.saneparClientId = os.Getenv("SANEPAR_CLIENT_ID")
	Config.jsonFileName = os.Getenv("SENT_EMAILS_JSON_FILENAME")
	if Config.jsonFileName == "" {
		Config.jsonFileName = "email_sent_at.json"
	}

	loadEmailSentJson()
}

func IsEmailSentAt(dateStr string) bool {
	return Config.emailSentAt[dateStr]
}

func SetEmailSentAt(dateStr string, val bool) {
	Config.emailSentAt[dateStr] = val
}

func SetEmailSentToday(val bool) {
	dateStr := utils.CurrentDateStr()

	Config.emailSentAt[dateStr] = val
}

func IsEmailSentToday() bool {
	date := utils.CurrentDateStr()

	return Config.emailSentAt[date]
}

func SaveConfigJson() {
	fmt.Println("Saving json file")

	jsonStr, err := json.Marshal(Config.emailSentAt)
	if err != nil {
		print("Error: %s\n", err.Error())
	} else {
		os.WriteFile(Config.jsonFileName, jsonStr, 0644)
	}
}

func loadEmailSentJson() {
	jsonFile, err := os.Open(Config.jsonFileName)
	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
		return
	}
	defer jsonFile.Close()

	contents, _ := io.ReadAll(jsonFile)

	json.Unmarshal(contents, &Config.emailSentAt)

	fmt.Println(Config.emailSentAt)
}
