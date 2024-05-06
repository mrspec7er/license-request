package src

import (
	"encoding/json"
	"fmt"
	"net/http"

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

	fmt.Println("RESULT: ", *user.Data)
	return 200, nil
}
