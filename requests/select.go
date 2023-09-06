package requests

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/kensparta/fullcycle-multithreading/constants"
	"github.com/kensparta/fullcycle-multithreading/dto"
	"log"
	"net/http"
	"time"
)

func SelectResponse(ctx context.Context, cep string, w http.ResponseWriter) error {
	viaCepChan := make(chan dto.ResultChannel[dto.ViaCepDto])
	apiCepChan := make(chan dto.ResultChannel[dto.ApiCepDto])

	viacep := &ViaCep{Cep: cep}
	apicep := &ApiCep{Cep: cep}

	go viacep.Execute(ctx, viaCepChan)
	go apicep.Execute(ctx, apiCepChan)

	select {
	case vc := <-viaCepChan:
		log.Print("VIA-CEP success response")
		err := json.NewEncoder(w).Encode(vc.Data)
		if err != nil {
			return err
		}
	case ac := <-apiCepChan:
		log.Print("API-CEP success response")
		err := json.NewEncoder(w).Encode(ac.Data)
		if err != nil {
			return err
		}
	case <-time.After(constants.MaxRequestTimeout):
		log.Print("timeout error!")
		return fmt.Errorf("timeout error")
	}
	return nil
}
