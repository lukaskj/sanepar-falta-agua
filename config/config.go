package config

import (
	"encoding/json"
	"errors"
	"io"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/joho/godotenv"
	"github.com/lukaskj/sanepar-falta-agua/utils"
)

type TConfig struct {
	NotificationSentAt map[string]bool
	SaneparBaseUrl     string
	SaneparClientId    string
	JsonFileName       string
	TimeLoopSeconds    int
	AwsSnsTopicArn     string
	AwsRegion          string
}

var Config TConfig

func Load() error {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	Config = TConfig{NotificationSentAt: make(map[string]bool)}

	Config.SaneparBaseUrl = strings.TrimSpace(os.Getenv("SANEPAR_BASE_URL"))
	Config.SaneparClientId = strings.TrimSpace(os.Getenv("SANEPAR_CLIENT_ID"))
	if Config.SaneparBaseUrl == "" || Config.SaneparClientId == "" {
		return errors.New("EnvAccessKeyNotFound: SANEPAR_BASE_URL or SANEPAR_CLIENT_ID not found in environment")
	}

	Config.JsonFileName = strings.TrimSpace(os.Getenv("SENT_NOTIFICATIONS_JSON_FILENAME"))
	if Config.JsonFileName == "" {
		Config.JsonFileName = "notifications_sent_at.json"
	}

	timeSecondsStr := strings.TrimSpace(os.Getenv("TIME_LOOP_SECONDS"))
	if timeSecondsStr == "" {
		timeSecondsStr = "60"
	}

	Config.TimeLoopSeconds, err = strconv.Atoi(timeSecondsStr)
	if err != nil {
		Config.TimeLoopSeconds = 60
	}

	// Aws config
	Config.AwsRegion = strings.TrimSpace(os.Getenv("AWS_REGION"))
	if Config.AwsRegion == "" {
		return errors.New("EnvAccessKeyNotFound: AWS_REGION not found in environment")
	}

	Config.AwsSnsTopicArn = strings.TrimSpace(os.Getenv("AWS_SNS_TOPIC_ARN"))
	if Config.AwsSnsTopicArn == "" {
		return errors.New("EnvAccessKeyNotFound: AWS_SNS_TOPIC_ARN not found in environment")
	}

	loadNotificationSentJson()

	log.Printf("Time Loop Seconds: %d\n", Config.TimeLoopSeconds)
	log.Printf("Sent notifications JSON filename: %s\n", Config.JsonFileName)

	return nil
}

func IsNotificationSentAt(dateStr string) bool {
	return Config.NotificationSentAt[dateStr]
}

func SetNotificationSentAt(dateStr string, val bool) {
	Config.NotificationSentAt[dateStr] = val
}

func SetNotificationSentToday(val bool) {
	dateStr := utils.CurrentDateStr()

	Config.NotificationSentAt[dateStr] = val
}

func IsNotificationSentToday() bool {
	date := utils.CurrentDateStr()

	return Config.NotificationSentAt[date]
}

func SaveConfigJson() {
	log.Println("Saving json file...")

	jsonStr, err := json.Marshal(Config.NotificationSentAt)
	if err != nil {
		log.Printf("Error: %s\n", err.Error())
	} else {
		os.WriteFile(Config.JsonFileName, jsonStr, 0644)
	}
}

func loadNotificationSentJson() {
	jsonFile, err := os.Open(Config.JsonFileName)
	// if we os.Open returns an error then handle it
	if err != nil {
		log.Println(err)
		return
	}
	defer jsonFile.Close()

	contents, _ := io.ReadAll(jsonFile)

	json.Unmarshal(contents, &Config.NotificationSentAt)
}
