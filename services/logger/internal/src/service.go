package src

import (
	"encoding/json"
	"net/http"
	"net/smtp"
	"os"

	"github.com/mrspec7er/license-request-utility/dto"
	"gorm.io/gorm"
)

type Service struct {
	DB   *gorm.DB
	Util *Util
}

type GetApplicationsResponse struct {
	Status   bool        `json:"status"`
	Message  string      `json:"message"`
	Data     *dto.User   `json:"data"`
	Metadata interface{} `json:"metadata"`
}

func (s Service) SendNotification(log *dto.Logger) (int, error) {

	resp, err := http.Get("http://user:8080/auth/" + log.UID)
	if err != nil {
		return 400, err
	}

	user := &GetApplicationsResponse{}
	if err := json.NewDecoder(resp.Body).Decode(&user); err != nil {
		return 400, err
	}

	message := []byte(log.Message)

	to := []string{
		user.Data.Email,
	}

	smtpAuth := smtp.PlainAuth("", os.Getenv("SMTP_EMAIL"), os.Getenv("SMTP_PASSWORD"), os.Getenv("SMTP_HOST"))

	err = smtp.SendMail(os.Getenv("SMTP_HOST")+":"+os.Getenv("SMTP_PORT"), smtpAuth, os.Getenv("SMTP_EMAIL"), to, message)
	if err != nil {
		return 400, err
	}
	return 200, nil
}
