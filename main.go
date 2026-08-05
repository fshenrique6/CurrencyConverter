package main

import (
	"currency-converter/converter"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	if len(os.Args) < 4 {
		fmt.Println("uso: converter <valor> <moeda_origem> <moeda_destino>")
		fmt.Println("exemplo: converter 100 USD BRL")
		return
	}

	amount, err := strconv.ParseFloat(os.Args[1], 64)
	if err != nil {
		fmt.Println("valor inválido:", os.Args[1])
		return
	}

	from := strings.ToUpper(os.Args[2])
	to := strings.ToUpper(os.Args[3])

	result, err := converter.Convert(amount, from, to)
	if err != nil {
		fmt.Println("erro:", err)
		return
	}

	fmt.Printf("%.2f %s = %.2f %s\n", amount, from, result, to)
}