package requests

import (
	"context"
	"fmt"
	"log"

	"github.com/kensparta/fullcycle-multithreading/dto"
)

type ViaCep struct {
	Cep     string
	request Request[dto.ViaCepDto]
}

func (v *ViaCep) Execute(ctx context.Context, result chan dto.ResultChannel[dto.ViaCepDto]) {
	v.request.Url = fmt.Sprintf("http://viacep.com.br/ws/%s/json/", v.Cep)
	r, err := v.request.Do(ctx)
	if err != nil {
		log.Println(fmt.Errorf("VIA-CEP error: %v, going to another option", err))
		return
	}

	if r.Cep == "" {
		log.Println(fmt.Errorf("VIA-CEP error: empty response, going to another option"))
		return
	}

	r.ApiName = "ViaCep"
	result <- dto.ResultChannel[dto.ViaCepDto]{
		Error: nil,
		Data:  r,
	}
}
