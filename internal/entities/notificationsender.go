package entities

import (
	"github.com/google/uuid"
)

type PayloadChunk struct {
	Users  []uuid.UUID   `json:"users"`
	Items  []PayloadItem `json:"items"`
	Offset int64         `json:"offset"`
}

type PayloadItem struct {
	DataType     UserNotificationType `json:"dataType"`
	TelegramData TelegramData         `json:"telegramData"`
	EmailData    EmailData            `json:"emailData"`
}

type TelegramData struct {
	MessageOrCallback string       `json:"messageOrCallback"`
	File              TelegramFile `json:"file"`
}

type TelegramFile struct {
	Name string `json:"name"`
	Data []byte `json:"data"`
	Key  string `json:"key"`
}

type EmailData struct {
	Data []byte `json:"data"`
}
