package main

import "fmt"

func main() {
	fmt.Println("Maps")

	user := map[string]string{
		"name": "Filipe",
	}

	fmt.Println(user["name"])

	user2 := map[string]map[string]string{
		"skills": {
			"language":        "js",
			"second_language": "go",
		},
	}

	fmt.Println(user2["skills"])
	fmt.Println(user2["skills"]["language"])

	delete(user2, "skills")
	fmt.Println(user2)
}
