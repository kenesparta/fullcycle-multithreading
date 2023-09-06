package dto

type ApiCepDto struct {
	ApiName    string `json:"api_name"`
	Code       string `json:"code"`
	State      string `json:"state"`
	City       string `json:"city"`
	District   string `json:"district"`
	Address    string `json:"address"`
	Status     int    `json:"status"`
	Ok         bool   `json:"ok"`
	StatusText string `json:"statusText"`
}

type ViaCepDto struct {
	ApiName     string `json:"api_name"`
	Cep         string `json:"cep"`
	Logradouro  string `json:"logradouro"`
	Complemento string `json:"complemento"`
	Bairro      string `json:"bairro"`
	Localidade  string `json:"localidade"`
	Uf          string `json:"uf"`
	Ibge        string `json:"ibge"`
	Gia         string `json:"gia"`
	Ddd         string `json:"ddd"`
	Siafi       string `json:"siafi"`
}

type ResultChannel[T any] struct {
	Error error
	Data  *T
}

type ErrorResponse struct {
	Message string
}
