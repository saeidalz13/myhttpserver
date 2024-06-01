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

type Test struct {
	name           string
	body           models.Item
	contentType    string
	expectedStatus int
	urlKey         string
	// expectedResp   string
}

func TestHandlePostItem(t *testing.T) {
	// recorder records all the writes to w (http.ResponseWriter)
	// on the server side. So this is useful for checking the headers!
	// rr := httptest.NewRecorder()

	// server := httptest.NewServer(http.HandlerFunc(StoreHandlerTest.HandlePostItem))
	// defer server.Close()

	validBodyReq := models.Item{
		Key:   "somekey",
		Value: "somevalue",
	}
	validContentType := "application/json"

	tests := []Test{
		{
			name:           "valid request",
			body:           validBodyReq,
			contentType:    validContentType,
			expectedStatus: http.StatusCreated,
		},
		{
			name:           "wrong header",
			body:           validBodyReq,
			contentType:    "invalid",
			expectedStatus: http.StatusBadRequest,
		},
		{
			name:           "no key in request body",
			body:           models.Item{Value: "somevalue"},
			contentType:    validContentType,
			expectedStatus: http.StatusBadRequest,
		},
		{
			name:           "no value in request body",
			body:           models.Item{Key: "somekey"},
			contentType:    validContentType,
			expectedStatus: http.StatusBadRequest,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			rr := httptest.NewRecorder()
			body, err := json.Marshal(test.body)
			if err != nil {
				t.Fatal(err)
			}
			req := httptest.NewRequest(http.MethodPost, "/store", bytes.NewBuffer(body))
			req.Header.Set("Content-Type", test.contentType)
			StoreHandlerTest.HandlePostItem(rr, req)
			resp := rr.Result()
			// resp, err := HttpClientTest.Post(server.URL+"/store", test.header, bytes.NewBuffer(body))
			// if err != nil {
			// 	t.Fatal(err)
			// }

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

func TestHandleGetItem(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(StoreHandlerTest.HandleGetItem))
	defer server.Close()

	tests := []Test{
		{
			name:           "get item w valid key",
			expectedStatus: http.StatusOK,
			urlKey:         "somekey",
		},
		{
			name:           "invalid key",
			expectedStatus: http.StatusNotFound,
			urlKey:         "nonexistentkey",
		},
		{
			name:           "invalid url no key",
			expectedStatus: http.StatusBadRequest,
			urlKey:         "",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			url := fmt.Sprintf("%s/retrieve/%s", server.URL, test.urlKey)
			resp, err := HttpClientTest.Get(url)
			if err != nil {
				t.Fatal(err)
			}

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

func TestHandleDeleteItem(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(StoreHandlerTest.HandleDeleteItem))
	defer server.Close()

	tests := []Test{
		{
			name:           "invalid url no key",
			expectedStatus: http.StatusBadRequest,
			urlKey:         "",
		},
		{
			name:           "invalid key deletion",
			urlKey:         "nonexistentkey",
			expectedStatus: http.StatusNotFound,
		},
		{
			name:           "valid key deletion",
			urlKey:         "somekey",
			expectedStatus: http.StatusNoContent,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			url := fmt.Sprintf("%s/delete-item/%s", server.URL, test.urlKey)
			req, err := http.NewRequest(http.MethodDelete, url, nil)
			if err != nil {
				t.Fatal(err)
			}

			resp, err := HttpClientTest.Do(req)
			if err != nil {
				t.Fatal(err)
			}

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
