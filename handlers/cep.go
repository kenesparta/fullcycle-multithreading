package handlers

import (
	"regexp"
)

func isValidCep(cep string) bool {
	r, err := regexp.Compile("^\\d{5}-\\d{3}$")
	if err != nil {
		return false
	}

	return r.MatchString(cep)
}
