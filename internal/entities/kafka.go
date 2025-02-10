package entities

import (
	"github.com/google/uuid"
)

const (
	AdminRoot        = "a"
	ServerRoot       = "b"
	PaymentRoot      = "c"
	InstructionsRoot = "f"
	ReviewsRoot      = "h"
	AddDemoRoot      = "i"
)

type CommandAsCallbackData struct {
	ChatID   int64  `json:"chat_id"`
	Callback string `json:"callback"`
}

type ParentReward struct {
	TargetUserID     uuid.UUID `json:"target_user_id"`
	ParentUserChatID int64     `json:"parent_user_chat_id"`
}
