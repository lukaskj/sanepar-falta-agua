package config

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
	"github.com/lukaskj/sanepar-falta-agua/utils"
)

type TConfig struct {
	EmailSentAt     map[string]bool
	SaneparBaseUrl  string
	SaneparClientId string
	JsonFileName    string
	TimeLoopSeconds int
}

var Config TConfig

func Load() {
	fmt.Println("Load config")

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	Config = TConfig{EmailSentAt: make(map[string]bool)}
	Config.SaneparBaseUrl = os.Getenv("SANEPAR_BASE_URL")
	Config.SaneparClientId = os.Getenv("SANEPAR_CLIENT_ID")
	Config.JsonFileName = os.Getenv("SENT_EMAILS_JSON_FILENAME")
	if Config.JsonFileName == "" {
		Config.JsonFileName = "email_sent_at.json"
	}
	timeSecondsStr := os.Getenv("TIME_LOOP_SECONDS")
	if timeSecondsStr == "" {
		timeSecondsStr = "60"
	}
	Config.TimeLoopSeconds, err = strconv.Atoi(timeSecondsStr)
	if err != nil {
		Config.TimeLoopSeconds = 60
	}

	loadEmailSentJson()
}

func IsEmailSentAt(dateStr string) bool {
	return Config.EmailSentAt[dateStr]
}

func SetEmailSentAt(dateStr string, val bool) {
	Config.EmailSentAt[dateStr] = val
}

func SetEmailSentToday(val bool) {
	dateStr := utils.CurrentDateStr()

	Config.EmailSentAt[dateStr] = val
}

func IsEmailSentToday() bool {
	date := utils.CurrentDateStr()

	return Config.EmailSentAt[date]
}

func SaveConfigJson() {
	fmt.Println("Saving json file")

	jsonStr, err := json.Marshal(Config.EmailSentAt)
	if err != nil {
		print("Error: %s\n", err.Error())
	} else {
		os.WriteFile(Config.JsonFileName, jsonStr, 0644)
	}
}

func loadEmailSentJson() {
	jsonFile, err := os.Open(Config.JsonFileName)
	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
		return
	}
	defer jsonFile.Close()

	contents, _ := io.ReadAll(jsonFile)

	json.Unmarshal(contents, &Config.EmailSentAt)
}
