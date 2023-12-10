package adresses

import (
	"errors"
	"strings"
)

// typeOfAdres returns a valid type of address or an error
func TypeOfAddress(address string) string {
	validTypes := []string{"Rua", "Avenida", "Estrada", "Rodovia"}

	firstWord := strings.Split(address, " ")[0]

	var isValid bool

	for _, addressType := range validTypes {
		if strings.ToLower(addressType) == strings.ToLower(firstWord) {
			isValid = true
		}
	}

	if isValid {
		return firstWord
	}

	return errors.New("Invalid type").Error()
}
