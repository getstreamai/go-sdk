package token

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
)

// RequestBody Input to the request token function
type RequestBody struct {
	AccessKey    string      `json:"accessKey"`
	AccessSecret string      `json:"accessSecret"`
	Data         requestData `json:"data"`
}

type requestData struct {
	Name   string `json:"name"`
	Room   string `json:"room"`
	Type   string `json:"type"`
	Record bool   `json:"record"`
}

type responseBody struct {
	Token string `json:"token"`
}

const url = "https://token.acadcare.com/api/token"

// GenerateToken Function to generate token
func GenerateToken(r *RequestBody) (string, error) {
	jsonValue, err := json.Marshal(r)
	if err != nil {
		log.Fatal("Error while decoding json request object")
	}
	res, e := http.Post(url, "application/json", bytes.NewBuffer(jsonValue))

	if e != nil {
		log.Fatal("Error while sending out request to Stream")
	}
	decoder := json.NewDecoder(res.Body)
	var response responseBody
	decodeErr := decoder.Decode(&response)
	if decodeErr != nil {
		log.Fatal("Error while decoding response")
	}
	io.Copy(os.Stdout, res.Body)
	defer res.Body.Close()
	return response.Token, nil
}
