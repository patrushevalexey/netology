package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func ReadProcessWrite(
	inputPath string,
	outputPath string,
	process func(string) (string, error),
) error {
	data, err := os.ReadFile(inputPath)
	if err != nil {
		return fmt.Errorf("ошибка чтения файла: %w", err)
	}

	result, err := process(string(data))
	if err != nil {
		return fmt.Errorf("ошибка обработки текста: %w", err)
	}

	err = os.WriteFile(outputPath, []byte(result), 0644)
	if err != nil {
		return fmt.Errorf("ошибка записи файла: %w", err)
	}

	return nil
}

func main() {
	err := ReadProcessWrite(
		"input.txt",
		// "input2.txt",
		"output.txt",
		func(input string) (string, error) {
			return strings.ToUpper(input), nil
		},
	)

	if err != nil {
		log.Println(err)
		return
	}

	fmt.Println("Готово")
}
