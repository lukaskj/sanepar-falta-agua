package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"

	"github.com/lukaskj/sanepar-falta-agua/types"
)

func SendSaneparRequest(baseUrl, clientId *string) types.TFaltaAguaResponse {
	fullUrl := *baseUrl + *clientId

	resp, err := http.Get(fullUrl)
	if err != nil {
		log.Fatalln(err)
	}

	rawBody, err := io.ReadAll(resp.Body)

	var body types.TFaltaAguaResponse

	err = json.Unmarshal(rawBody, &body)
	if err != nil {
		log.Println(err)
	}

	return body
}

func IsElegibleToSendNotification(faltaAguaResponse *types.TFaltaAguaResponse) bool {
	return faltaAguaResponse.Mensagem != "NADA CONSTA" && faltaAguaResponse.Mensagem != "SEM ACESSO AO WEB SERVICE";
}
