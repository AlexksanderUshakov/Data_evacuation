package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type PERSON struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {

	people := []PERSON{
		{Name: "Alex", Age: 25},
		{Name: "Vika", Age: 18},
	}

	file, err := os.Create("data.json")

	if err != nil {
		fmt.Println("Файл не создан!")
		return
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	err = encoder.Encode(people)

	if err != nil {
		fmt.Println("Эвакуация не произашла")
	}

	fmt.Println("Эвакуация данных прошла успешно")
}
