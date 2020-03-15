package main

import (
	"artidote-quote/constants"
	"artidote-quote/telegram"
	"artidote-quote/util"
	"fmt"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/service/ssm"
)

func handler(_ events.CloudWatchEvent) int {
	fmt.Println("hi")
	session := util.StartAWSSession("ap-southeast-1")
	ssmsvc := ssm.New(session)

	telegramBotKey := util.GetSSMParams(ssmsvc, constants.PSKTelegramBotKey)
	messageToSend := []string{"test"}
	a := messageToSend[0]
	roomIDs := util.GetRoomIDsFromDB(session)
	telegram.SendMessagesToRooms(a, roomIDs, telegramBotKey)
	println(a)
	fmt.Println(roomIDs)
	println(telegramBotKey)

	return 0
}

func main() {
	if os.Args[1] == "debug" {
		os.Setenv("AWS_SDK_LOAD_CONFIG", "1")
		os.Setenv("AWS_PROFILE", os.Args[2])
		handler(events.CloudWatchEvent{})
	} else {
		lambda.Start(handler)
	}
}
