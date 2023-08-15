package main

import (
	"fmt"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sns"
	"github.com/lukaskj/sanepar-falta-agua/config"
)

var messageChan = make(chan string)

func SendNotificationMessage(message string) {
	go sendSNSNotification(message)
}

func sendSNSNotification(message string) {
	fmt.Println("[-] Sending notification at " + time.Now().Format(time.DateTime))

	creds := credentials.NewEnvCredentials()

	// Retrieve the credentials value
	_, err := creds.Get()
	if err != nil {
		panic(err)
	}

	sess, _ := session.NewSession(&aws.Config{
		Region:      aws.String(config.Config.AwsRegion),
		Credentials: creds,
	})

	svc := sns.New(sess)

	_, err = svc.Publish(&sns.PublishInput{
		Message:  &message,
		TopicArn: &config.Config.AwsSnsTopicArn,
	})

	if err != nil {
		fmt.Println(err.Error())
	}
}
