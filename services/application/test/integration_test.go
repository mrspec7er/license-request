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
	"github.com/mrspec7er/license-request/services/application/internal/db"
)

type GetApplicationsResponse struct {
	Status   bool        `json:"status"`
	Message  string      `json:"message"`
	Data     []*dto.Form `json:"data"`
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

	util := &db.RedisRepository{
		Cache: Memcache,
	}

	util.Store(context.Background(), "mock-auth", &dto.User{
		ID:            "108062360841627093205",
		Email:         "test@Email.com",
		VerifiedEmail: true,
		Role:          string(dto.RoleAdmin),
	})

	reqBody := `{
        "number": "DRV-TESTING",
        "formID": 1,
        "responses": [
          {
            "field_id": 1,
            "value": "Testing User"
          },
          {
            "field_id": 2,
            "value": "Testing Description: Little bit too aggressive in corner but has very good handling and fast response"
          }
        ]
     }`

	serverURL := "http://localhost" + os.Getenv("PORT") + "/applications"

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

	util := &db.RedisRepository{
		Cache: Memcache,
	}

	util.Store(context.Background(), "mock-auth", &dto.User{
		ID:            "108062360841627093205",
		Email:         "test@Email.com",
		VerifiedEmail: true,
		Role:          string(dto.RoleAdmin),
	})

	serverURL := "http://localhost" + os.Getenv("PORT") + "/applications"

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

	var expectedResponse GetApplicationsResponse
	if err := json.NewDecoder(resp.Body).Decode(&expectedResponse); err != nil {
		t.Fatalf("failed to unmarshal response body: %v", err)
	}

	if expected := http.StatusOK; resp.StatusCode != expected {
		t.Errorf("unexpected response status code: expected %d, got %d", expected, resp.StatusCode)
	}

}
