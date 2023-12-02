package main

import (
	"fmt"
	"log"
	"os"
	"yole/app"
)

func main() {
	fmt.Println("Initialize cli")

	application := app.Generate()
	err := application.Run(os.Args)

	if err != nil {
		log.Fatal(err)
	}
}
