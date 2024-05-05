package test

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"os"
	"testing"

	"github.com/joho/godotenv"
	"github.com/mrspec7er/license-request-utility/dto"
	"github.com/mrspec7er/license-request/services/form/internal/db"
	"github.com/mrspec7er/license-request/services/form/internal/src"
)

type GetFormResponse struct {
	Status   bool        `json:"status"`
	Message  string      `json:"message"`
	Data     *dto.Form   `json:"data"`
	Metadata interface{} `json:"metadata"`
}

func init() {
	err := godotenv.Load("../.env")
	if err != nil {
		panic(err)
	}
}

func TestCreateForm(t *testing.T) {

	Memcache := db.MemcacheConnection()

	util := &src.Util{
		Memcache: Memcache,
	}

	util.MemcacheStore(context.Background(), "mock-auth", dto.User{
		ID:            "1101",
		Email:         "test@Email.com",
		VerifiedEmail: true,
		Role:          string(dto.RoleAdmin),
	})

	reqBody := `{
        "name": "Register Form Testing",
        "category": "Sample Category",
        "sections": [
          {
            "name": "First Section Testing",
            "fields": [
              {
                "label": "Label Testing",
                "type": "Text",
                "order": 1
              },
              {
                "label": "Label Testing II",
                "type": "Number",
                "order": 2
              }
            ]
          }
       ]
    }`

	serverURL := "http://localhost" + os.Getenv("PORT") + "/forms"

	req, err := http.NewRequest(http.MethodPost, serverURL, bytes.NewBufferString(reqBody))
	if err != nil {
		t.Fatalf("failed to create HTTP request: %v", err)
	}
	req.Header.Set("Content-Type", "application/json")

	cookie := http.Cookie{Name: "auth", Value: "mock-auth"}
	req.AddCookie(&cookie)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		t.Fatalf("failed to send HTTP request: %v", err)
	}
	defer resp.Body.Close()

	if expected := http.StatusCreated; resp.StatusCode != expected {
		t.Errorf("unexpected response status code: expected %d, got %d", expected, resp.StatusCode)
	}

}

func TestGetForm(t *testing.T) {

	Memcache := db.MemcacheConnection()

	util := &src.Util{
		Memcache: Memcache,
	}

	util.MemcacheStore(context.Background(), "mock-auth", dto.User{
		ID:            "1101",
		Email:         "test@Email.com",
		VerifiedEmail: true,
		Role:          string(dto.RoleAdmin),
	})

	serverURL := "http://localhost" + os.Getenv("PORT") + "/forms/1"

	req, err := http.NewRequest(http.MethodGet, serverURL, nil)
	if err != nil {
		t.Fatalf("failed to create HTTP request: %v", err)
	}
	req.Header.Set("Content-Type", "application/json")

	cookie := http.Cookie{Name: "auth", Value: "mock-auth"}
	req.AddCookie(&cookie)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		t.Fatalf("failed to send HTTP request: %v", err)
	}
	defer resp.Body.Close()

	var expectedResponse GetFormResponse
	if err := json.NewDecoder(resp.Body).Decode(&expectedResponse); err != nil {
		t.Fatalf("failed to unmarshal response body: %v", err)
	}

	if expected := http.StatusOK; resp.StatusCode != expected {
		t.Errorf("unexpected response status code: expected %d, got %d", expected, resp.StatusCode)
	}

}
