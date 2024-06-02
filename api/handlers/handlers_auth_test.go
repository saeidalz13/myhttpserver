package handlers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"myserver/api/models"
	"net/http"
	"net/http/httptest"
	"testing"
)

type AuthTest struct {
	name           string
	body           models.User
	contentType    string
	expectedStatus int
}

func TestHandlePostSignUp(t *testing.T) {
	validUser := models.User{
		Email:    "saeid@hotmail.com",
		Password: "SomePassword12345",
	}
	validContentType := "application/json"

	tests := []AuthTest{
		{
			name:           "valid email password",
			body:           validUser,
			expectedStatus: http.StatusCreated,
			contentType:    validContentType,
		},
		{
			name:           "invalid email",
			body:           models.User{Password: "SomePassword1234"},
			expectedStatus: http.StatusBadRequest,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			body, err := json.Marshal(test.body)
			if err != nil {
				t.Fatal(err)
			}
			rr := httptest.NewRecorder()
			req := httptest.NewRequest(http.MethodPost, "/signup", bytes.NewBuffer(body))
			req.Header.Set("Content-Type", test.contentType)
			AuthHandlerTest.HandlePostSignUp(rr, req)

			resp := rr.Result()

			if resp.StatusCode != test.expectedStatus {
				t.Fatalf("expected status: %d\treceived: %d", test.expectedStatus, resp.StatusCode)
			}
			bufBytes, err := io.ReadAll(resp.Body)
			if err != nil {
				t.Fatal(err)
			}

			fmt.Printf("response: %s\n", string(bufBytes))
		})
	}

}

func TestHandlePostLogin(t *testing.T) {
	validUser := models.User{
		Email:    "saeid@hotmail.com",
		Password: "SomePassword12345",
	}

	validContentType := "application/json"

	tests := []AuthTest{
		{
			name:           "valid existing email",
			body:           validUser,
			contentType:    validContentType,
			expectedStatus: http.StatusOK,
		},
		{
			name:           "non existing email",
			body:           models.User{Email: "invalidemail", Password: "somepass"},
			contentType:    validContentType,
			expectedStatus: http.StatusNotFound,
		},
		{
			name:           "invalid password",
			body:           models.User{Email: "saeid@hotmail.com", Password: "invalidpass"},
			contentType:    validContentType,
			expectedStatus: http.StatusNotFound,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			body, err := json.Marshal(test.body)
			if err != nil {
				t.Fatal(err)
			}

			w := httptest.NewRecorder()
			req := httptest.NewRequest(http.MethodPost, "/login", bytes.NewBuffer(body))
			req.Header.Set("Content-Type", test.contentType)

			AuthHandlerTest.HandlePostLogin(w, req)
			resp := w.Result()

			if resp.StatusCode != test.expectedStatus {
				t.Fatalf("expected status: %d\treceived: %d", test.expectedStatus, resp.StatusCode)
			}
			bufBytes, err := io.ReadAll(resp.Body)
			if err != nil {
				t.Fatal(err)
			}

			fmt.Printf("response: %s\n", string(bufBytes))
		})
	}
}
