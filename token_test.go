package token

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"
)

const token = "TOKEN"

type MockDoType func(req *http.Request) (*http.Response, error)

type MockClient struct {
	MockDo MockDoType
}

func (m *MockClient) Do(req *http.Request) (*http.Response, error) {
	return m.MockDo(req)
}

func TestToken(t *testing.T) {
	const TOKEN = "your-token"
	jsonResponse := `{
		"token": "your-token"	
	}`
	r := ioutil.NopCloser(bytes.NewReader([]byte(jsonResponse)))
	Client = &MockClient{
		MockDo: func(req *http.Request) (*http.Response, error) {
			return &http.Response{
				StatusCode: 200,
				Body:       r,
			}, nil
		},
	}

	data := RequestData{
		Name:   "Harkirat",
		Room:   "123",
		Type:   "producer",
		Record: true,
	}
	body := RequestBody{
		Data:         data,
		AccessKey:    "accessKey",
		AccessSecret: "accessSecret",
	}
	result, err := GenerateToken(&body)
	if err != nil {
		t.Error("Received error response from GenerateToken function")
		return
	}
	if result != TOKEN {
		t.Error("Token is not equal to expected value")
		return
	}
}

func TestTokenErr(t *testing.T) {
	jsonResponse := `{
		"msg": "Please provide accessKey and accessSecret"
	}
	`
	r := ioutil.NopCloser(bytes.NewReader([]byte(jsonResponse)))
	Client = &MockClient{
		MockDo: func(req *http.Request) (*http.Response, error) {
			return &http.Response{
				StatusCode: http.StatusForbidden,
				Body:       r,
			}, nil
		},
	}

	data := RequestData{
		Name:   "Harkirat",
		Room:   "123",
		Type:   "producer",
		Record: true,
	}
	body := RequestBody{
		Data:         data,
		AccessSecret: "accessSecret",
	}
	_, err := GenerateToken(&body)
	if err == nil {
		fmt.Println(err)
		t.Error("Expected error but found none")
	}
	if err.Error() != "Please provide accessKey and accessSecret" {
		t.Error("Wrong error message received")
	}
}
