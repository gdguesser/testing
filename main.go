package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type CEPInfo struct {
	Logradouro  string `json:"logradouro"`
	Complemento string `json:"complemento"`
	Bairro      string `json:"bairro"`
	Localidade  string `json:"localidade"`
	UF          string `json:"uf"`
}

func main() {
	someValue := SomeValue()

	addition := Add(3, 2)

	multiplication := Multiply(5, -6)

	value, err := GetCEPInfo("01001000")
	if err != nil {
		log.Println("test: ", err)
	}

	fmt.Println(value)

	fmt.Printf("Some value: %s, Addition: %d, Multiplication: %d\n",
		someValue, addition, multiplication)
}

func SomeValue() string {
	return "expected value"
}

func Add(a, b int) int {
	return a + b
}

func Multiply(a, b int) int {
	return a * b
}

func GetCEPInfo(cep string) (*CEPInfo, error) {
	resp, err := http.Get(fmt.Sprintf("https://viacep.com.br/ws/%s/json/", cep))
	if err != nil {
		return nil, fmt.Errorf("failed to perform GET request: %w", err)
	}
	defer resp.Body.Close()

	var info CEPInfo
	if err := json.NewDecoder(resp.Body).Decode(&info); err != nil {
		return nil, fmt.Errorf("failed to decode response body: %w", err)
	}

	return &info, nil
}
