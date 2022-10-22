package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func OpenFile(fileName string) (*os.File, error) {
	return os.Open(fileName)
}

func CloseFile(file *os.File) {
	file.Close()
	fmt.Println("File was closed")
}

func GetFloats(fileName string) ([]float64, error) {
	var numbers []float64
	file, err := OpenFile(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer CloseFile(file)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		number, err := strconv.ParseFloat(scanner.Text(), 64)
		if err != nil {
			log.Fatal(err)
		}
		numbers = append(numbers, number)
	}
	if scanner.Err() != nil {
		return nil, err
	}
	return numbers, nil
}

func main() {
	fileName := os.Args[1]
	fileName = strings.TrimSpace(fileName)
	nums, err := GetFloats(fileName)
	if err != nil {
		log.Fatal(err)
	}
	for _, num := range nums {
		fmt.Println(num)
	}

}
