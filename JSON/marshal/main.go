package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
)

type dog struct {
	Name string `json:"name"`
	Age  uint   `json:"age"`
}

// func toJson(obj any) *bytes.Buffer {

// 	JSONobj, err := json.Marshal(obj)

// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	return bytes.NewBuffer(JSONobj)
// }

func main() {
	d := dog{"Patricia", 3}

	JSONDog, err := json.Marshal(d)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(JSONDog)
	fmt.Println(bytes.NewBuffer(JSONDog))

	d2 := map[string]string{
		"name": "Patricia",
		"age":  "3",
	}

	JSONDog2, err := json.Marshal(d2)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(JSONDog2)
	fmt.Println(bytes.NewBuffer(JSONDog2))

}
