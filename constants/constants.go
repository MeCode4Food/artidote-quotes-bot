package constants

// PSKTelegramBotKey Parameter Store Name for Telegram Bot Key
const (
	PSKTelegramBotKey = "/PROJECT/DAILY_ARTIDOTE_QUOTE/TELEGRAM_BOT_KEY"
	DBName            = "DailyArtidoteQuote"
)

type DocIDList struct {
	ChatIDList string
}

var DocID = DocIDList{
	ChatIDList: "TELEGRAM_CHAT_ID_LIST",
}
