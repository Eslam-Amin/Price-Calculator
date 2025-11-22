package filemanager

import (
	"bufio"
	"errors"
	"fmt"
	"os"
)

func ReadLines(filepath string)([]string, error){
	file, err := os.Open(filepath)
	if err != nil {
		return nil, errors.New("failed to open file")
	}
	
	scanner := bufio.NewScanner(file)
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	err = scanner.Err()
	if err != nil{
		fmt.Println("Couldn't Read file content!")
		fmt.Println(err)
		file.Close()
		return nil, errors.New("failed to read lines file")
	}

	file.Close()
	return lines, nil
}