package parse

import (
	"bufio"
	"log"
	"os"
)

func GetLinesFromFile(filePath string) []string {
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var (
		lines    []string
		bytes    []byte
		isPrefix bool = false
	)

	reader := bufio.NewReader(file)
	for !isPrefix && err == nil {
		bytes, isPrefix, err = reader.ReadLine()
		lines = append(lines, string(bytes))
	}

	return lines
}
