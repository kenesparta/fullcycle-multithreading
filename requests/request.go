package requests

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/kensparta/fullcycle-multithreading/constants"
	"io"
	"log"
	"net/http"
)

type RequestInterface[R interface{}] interface {
	Do(ctx context.Context) (R, error)
}

type Request[R interface{}] struct {
	Url string
}

func (rt *Request[R]) Do(ctx context.Context) (*R, error) {
	ctx, cancel := context.WithTimeout(ctx, constants.MaxRequestTimeout)
	defer cancel()
	client := http.Client{Timeout: constants.MaxRequestTimeout}
	request, reqErr := http.NewRequestWithContext(
		ctx,
		http.MethodGet,
		rt.Url,
		nil)
	if reqErr != nil {
		return nil, reqErr
	}

	request.Header.Set("Accept", "application/json")
	response, respErr := client.Do(request)
	if respErr != nil {
		log.Printf("timeout error %v", respErr)
		return nil, respErr
	}

	if response.StatusCode < http.StatusOK || response.StatusCode >= 300 {
		return nil, fmt.Errorf("bad response, error: %v", response.StatusCode)
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Printf("close stream error %v", err)
			return
		}
	}(response.Body)

	readBody, readErr := io.ReadAll(response.Body)
	if readErr != nil {
		log.Printf("read error %v", readErr)
		return nil, readErr
	}

	var conv R
	unmErr := json.Unmarshal(readBody, &conv)
	if unmErr != nil {
		return nil, unmErr
	}

	return &conv, nil
}
