package main

import (
	"fmt"
	"os"

	"github.com/orenvadi/ITM-Task/masking"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go path/to/file")
		return
	}

	filePath := os.Args[1]

	prod := masking.NewFileProducer(filePath)
	pres := masking.NewFileWriterPresenter("output.txt") // Путь к файлу вывода результата

	service := masking.NewService(prod, pres)
	if err := service.Run(); err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("\nLinks masked succesfully!")
}
