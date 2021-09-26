package internal

import (
	"context"
	"google.golang.org/api/gmail/v1"
)

type EmailsUseCase interface {
	SendMessage(service *gmail.Service, userID string, message gmail.Message)
	CreateMessage(from string, to string, subject string, content string) gmail.Message
	CreateMessageWithAttachment(from string, to string, subject string, content string, fileDir string, fileName string) gmail.Message
	SendEmail(ctx context.Context, deliveryBody []byte) error
}
