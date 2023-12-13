package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type dog struct {
	Name string `json:"name"`
	Age  uint   `json:"age"`
}

func main() {
	dogInJSON := `{"age":3,"name":"Patricia"}`
	var d dog

	if err := json.Unmarshal([]byte(dogInJSON), &d); err != nil {
		log.Fatal(err)
	}

	fmt.Println(d)

}
