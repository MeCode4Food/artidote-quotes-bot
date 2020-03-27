# Daily Artidote Quote

## Who is The Artidote?

> [The Artidote is] a virtual space where to story-tell, empathize, bond and heal through art. 
>
> \- Jovanny Varela Ferreyra, Creator of The Artidote

The Artidote is an account on Facebook/Instagram/Twitter where short stories/quotes are shared alongside a piece of artwork, with a focus on mental health

## Purpose

This repository hosts an AWS Lambda that sends you the latest story (or one of The Artidote's older posts) on invocation.

## How it works

![DynamoDB Setup](/docs/hld.png?raw=true "DynamoDB Setup")

1. The invoker (CloudWatch Events, AWS APIGateway) calls the AWS Lambda
2. The Lambda gets the Bot Key from AWS Parameter Store
3. The Lambda gets posts from Instagram
4. The Lambda gets the list of people to send to from the DynamoDB
5. After processing the posts, sends the posts using the Telegram Bot Key to Telegram

## How to use

Pre-requisite: Fork the repository

### 1. Setup parameters

You will need the following parameters:

- AWS Access Key
- AWS Secret Key
- AWS Region
- Lambda Invocation Role that satisfies the following:
  - Internet Access
  - Read Access to DynamoDB
- Telegram Bot Key // Get from the BotFather https://core.telegram.org/bots

### 2. Set up AWS Parameter Store

The Telegram Bot Key is stored inside the AWS Parameter Store (Runtime secrets are stored in the Parameter Store) using the key as described inside `constants/constants.go`

```go
const (
	PSKTelegramBotKey = "/PROJECT/DAILY_ARTIDOTE_QUOTE/TELEGRAM_BOT_KEY" // PSK stands for Parameter Store Key
)
```

### 3. Setup DynamoDB constants

Set up DynamoDB to be as shown in the NoSQL workbench, or you can tweak it under `constants/constants.go`

![DynamoDB Setup](/docs/dynamodb.png?raw=true "DynamoDB Setup")

```go
const (
	DBName            = "DailyArtidoteQuote" // table name
)

var DocID = DocIDList{
	ChatIDList: "TELEGRAM_CHAT_ID_LIST", // DocIDList is an array of strings
}
```

#### 4. You're done!

If you have any questions, do raise an issue and I'll look into them if I have the time (I promise!)


### FAQ

1. How do I get my chat ID?

Refer to this article
https://docs.influxdata.com/kapacitor/v1.5/event_handlers/telegram/#get-your-telegram-chat-id