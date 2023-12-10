package adresses

import "testing"

type testScenario struct {
	testAddress     string
	expectedAddress string
}

func TestTypeOfAddress(t *testing.T) {

	testScenarios := []testScenario{
		{"Rua Georgina", "Rua"},
		{"Avenida Paulista", "Avenida"},
		{"Rodovia Ayrton Senna", "Rodovia"},
		{"Nakka", "Invalid type"},
		{"", "Invalid type"},
		{"Estrada Patricia", "Estrada"},
	}

	for _, testScenatio := range testScenarios {

		receivedAddress := TypeOfAddress(testScenatio.testAddress)

		if receivedAddress != testScenatio.expectedAddress {
			t.Errorf("The received type [%s] is not equal expected type [%s]", testScenatio.expectedAddress, receivedAddress)
		}
	}
}
