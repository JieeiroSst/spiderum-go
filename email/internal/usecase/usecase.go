package usecase

import (
	"context"
	"encoding/base64"
	"encoding/json"
	email"gitlab.com/Spide_IT/spide_it/email/internal"
	"gitlab.com/Spide_IT/spide_it/email/internal/model"
	EmailPkg "gitlab.com/Spide_IT/spide_it/email/pkg/email"
	"gitlab.com/Spide_IT/spide_it/email/utils"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/gmail/v1"
	"io/ioutil"
	"log"
	"net/http"
)

type EmailUseCase struct {
	emailsRepo email.EmailsRepository
	publisher  email.EmailsPublisher
	emailUtils utils.EmailUtils
	EmailPkg   EmailPkg.ServicePkg
}

func NewEmailUseCase(emailsRepo email.EmailsRepository, publisher email.EmailsPublisher,emailUtils utils.EmailUtils,EmailPkg EmailPkg.ServicePkg) *EmailUseCase {
	return &EmailUseCase{emailsRepo: emailsRepo, publisher: publisher,emailUtils:emailUtils,EmailPkg:EmailPkg}
}

func (e *EmailUseCase) SendMessage(service *gmail.Service, userID string, message gmail.Message){
	_, err := service.Users.Messages.Send(userID, &message).Do()
	if err != nil {
		log.Fatalf("Unable to send message: %v", err)
	} else {
		log.Println("Email message sent!")
	}
}

func (e *EmailUseCase) CreateMessage(from string, to string, subject string, content string) gmail.Message{
	var message gmail.Message
	messageBody := []byte("From: " + from + "\r\n" +
		"To: " + to + "\r\n" +
		"Subject: " + subject + "\r\n\r\n" +
		content)
	message.Raw = base64.StdEncoding.EncodeToString(messageBody)
	return message
}

func (e *EmailUseCase) CreateMessageWithAttachment(from string, to string, subject string, content string, fileDir string, fileName string) gmail.Message {
	var message gmail.Message
	fileBytes, err := ioutil.ReadFile(fileDir + fileName)
	if err != nil {
		log.Fatalf("Unable to read file for attachment: %v", err)
	}

	fileMIMEType := http.DetectContentType(fileBytes)
	fileData := base64.StdEncoding.EncodeToString(fileBytes)
	boundary := e.emailUtils.RandStr(32, "alphanum")
	messageBody := []byte("Content-Type: multipart/mixed; boundary=" + boundary + " \n" +
		"MIME-Version: 1.0\n" +
		"to: " + to + "\n" +
		"from: " + from + "\n" +
		"subject: " + subject + "\n\n" +

		"--" + boundary + "\n" +
		"Content-Type: text/plain; charset=" + string('"') + "UTF-8" + string('"') + "\n" +
		"MIME-Version: 1.0\n" +
		"Content-Transfer-Encoding: 7bit\n\n" +
		content + "\n\n" +
		"--" + boundary + "\n" +

		"Content-Type: " + fileMIMEType + "; name=" + string('"') + fileName + string('"') + " \n" +
		"MIME-Version: 1.0\n" +
		"Content-Transfer-Encoding: base64\n" +
		"Content-Disposition: attachment; filename=" + string('"') + fileName + string('"') + " \n\n" +
		e.emailUtils.ChunkSplit(fileData, 76, "\n") +
		"--" + boundary + "--")
	message.Raw = base64.URLEncoding.EncodeToString(messageBody)
	emailData:=model.Email{
		Email:       from,
		To:          to,
		From:        from,
		Body:        content,
		Subject:     subject,
		ContentType: "",
	}
	mailBytes, err := json.Marshal(emailData)
	err=e.publisher.Publish(mailBytes,emailData.ContentType)
	if err!=nil{
		log.Println(err)
	}
	return message
}

func (e *EmailUseCase) SendEmail(ctx context.Context, deliveryBody []byte) error{
	credential, err := ioutil.ReadFile("client_secret.json")
	if err != nil {
		log.Fatalf("Unable to read client secret file: %v", err)
	}
	config, err := google.ConfigFromJSON(credential, gmail.GmailSendScope)
	if err != nil {
		log.Fatalf("Unable to parse client secret file to config: %v", err)
	}

	client := e.EmailPkg.GetClient(ctx, config)
	gmailClientService, err := gmail.New(client)
	if err != nil {
		log.Fatalf("Unable to initiate new gmail client: %v", err)
	}
	emailData:=model.Email{}
	err = json.Unmarshal(deliveryBody, &emailData)
	if err!=nil{
		log.Println(err)
	}
	emailResult, err := e.emailsRepo.CreateEmail(emailData)
	if err!=nil{
		log.Println(err)
	}
	messageWithAttachment := e.CreateMessageWithAttachment(emailData.From, emailData.To, emailData.Subject, emailData.Body, "./", "img.pdf")
	user := emailResult.Id
	e.SendMessage(gmailClientService, user, messageWithAttachment)
	return nil
}
