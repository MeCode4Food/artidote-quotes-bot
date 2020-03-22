package main

import (
	"artidote-quote/constants"
	"artidote-quote/instagram"
	"artidote-quote/telegram"
	"artidote-quote/util"
	"fmt"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/service/ssm"
)

func handler(_ events.CloudWatchEvent) (int, error) {
	fmt.Println("hi")
	session := util.StartAWSSession("ap-southeast-1")
	ssmsvc := ssm.New(session)

	telegramBotKey := util.GetSSMParams(ssmsvc, constants.PSKTelegramBotKey)

	messageToSend := instagram.GetInstagramMessage()
	roomIDs := util.GetRoomIDsFromDB(session)
	telegram.SendMessagesToRooms(messageToSend, roomIDs, telegramBotKey)
	println(messageToSend)
	fmt.Println(roomIDs)
	return 0, nil
}

func main() {
	if len(os.Args) >= 2 && os.Args[1] == "debug" {
		os.Setenv("AWS_SDK_LOAD_CONFIG", "1")
		os.Setenv("AWS_PROFILE", os.Args[2])
		handler(events.CloudWatchEvent{})
	} else {
		lambda.Start(handler)
	}
}
