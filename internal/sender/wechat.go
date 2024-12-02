package sender

import (
	"context"
	"github.com/ZoeKyHein/bot/wx"
	"log"
)

func SendMessage(msg string) {
	botClient := wx.BotClient{
		Key:    "ee556a46-a3a7-4978-a186-7e3181f29da9",
		Logger: nil,
	}

	err := botClient.SendMarkDown(context.Background(), msg)
	if err != nil {
		log.Printf("发送消息失败: %v", err)
	}
}
