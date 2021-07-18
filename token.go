package token

import (
	"bytes"
	"encoding/json"
	"errors"
	"log"
	"net/http"
)

// RequestBody Input to the request token function
type RequestBody struct {
	AccessKey    string      `json:"accessKey"`
	AccessSecret string      `json:"accessSecret"`
	Data         RequestData `json:"data"`
}

type RequestData struct {
	Name   string `json:"name"`
	Room   string `json:"room"`
	Type   string `json:"type"`
	Record bool   `json:"record"`
}

type responseBody struct {
	Token string `json:"token"`
}

type errorResponseBody struct {
	Message string `json:"msg"`
}

type HttpClient interface {
	Do(req *http.Request) (*http.Response, error)
}

var (
	Client HttpClient
)

func init() {
	Client = &http.Client{}
}

const url = "https://token.acadcare.com/api/token"

// GenerateToken Function to generate token
func GenerateToken(r *RequestBody) (string, error) {
	jsonValue, err := json.Marshal(r)
	if err != nil {
		log.Fatal("Error while decoding json request object")
	}

	request, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(jsonValue))
	if err != nil {
		return "", err
	}
	res, e := Client.Do(request)

	if e != nil || res.StatusCode == http.StatusForbidden {
		decoder := json.NewDecoder(res.Body)
		var response errorResponseBody
		decodeErr := decoder.Decode(&response)
		if decodeErr != nil {
			return "", errors.New("Error while fetching response from server")
		}
		return "", errors.New(response.Message)
	}

	decoder := json.NewDecoder(res.Body)
	var response responseBody
	decodeErr := decoder.Decode(&response)
	if decodeErr != nil {
		return "", errors.New("Error while decoding response")
	}

	defer res.Body.Close()
	return response.Token, nil
}
