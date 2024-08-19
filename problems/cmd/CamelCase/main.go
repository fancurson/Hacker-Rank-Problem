package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {

	file, err := os.Open("camel.txt")
	if err != nil {
		log.Fatalf("Open file error: %v", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		answer := 1
		for _, char := range line {
			if strings.ToUpper(string(char)) == string(char) {
				answer++
			}
		}
		fmt.Printf("line: %s, res: %d\n", line, answer)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Ошибка при сканировании файла:", err)
	}
}
