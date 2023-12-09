package utils

import (
	"bufio"
	"fmt"
	"os"
)

func ReadFile(filename string) ([]string, error) {
	fmt.Println(os.Getwd())
	readFile, err := os.Open(filename)
	defer readFile.Close()
	if err != nil {
		return nil, err
	}
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	var lines []string
	for fileScanner.Scan() {
		lines = append(lines, fileScanner.Text())
	}

	return lines, nil
}
