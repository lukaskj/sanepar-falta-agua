package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
)

func SendSaneparRequest(baseUrl, clientId *string) TFaltaAguaResponse {
	fullUrl := *baseUrl + *clientId

	resp, err := http.Get(fullUrl)
	if err != nil {
		log.Fatalln(err)
	}

	rawBody, err := io.ReadAll(resp.Body)

	var body TFaltaAguaResponse

	err = json.Unmarshal(rawBody, &body)
	if err != nil {
		log.Println(err)
	}

	return body
}
