package requests

import (
	"context"
	"fmt"
	"github.com/kensparta/fullcycle-multithreading/dto"
	"log"
)

type ApiCep struct {
	Cep     string
	request Request[dto.ApiCepDto]
}

func (a *ApiCep) Execute(ctx context.Context, result chan dto.ResultChannel[dto.ApiCepDto]) {
	a.request.Url = fmt.Sprintf("https://cdn.apicep.com/file/apicep/%s.json", a.Cep)
	r, err := a.request.Do(ctx)
	if err != nil {
		log.Println(fmt.Errorf("API-CEP error: %v, going to another option", err))
		return
	}

	if r.Code == "" {
		log.Println(fmt.Errorf("API-CEP error: empty response, going to another option"))
		return
	}

	r.ApiName = "ApiCep"
	result <- dto.ResultChannel[dto.ApiCepDto]{
		Error: nil,
		Data:  r,
	}
}
