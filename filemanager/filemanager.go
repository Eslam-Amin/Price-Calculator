package filemanager

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"time"
)


type FileManager struct{
	InputFilePath string
	OutputFilePath string
}

func (fm FileManager) ReadLines()([]string, error){
	file, err := os.Open(fm.InputFilePath)
	if err != nil {
		return nil, errors.New("failed to open file")
	}

	defer file.Close()

	
	scanner := bufio.NewScanner(file)
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	err = scanner.Err()
	if err != nil{
		fmt.Println("Couldn't Read file content!")
		fmt.Println(err)
		return nil, errors.New("failed to read lines file")
	}

	return lines, nil
}

func (fm FileManager) WriteResult(data interface{})error{
	file, err := os.Create(fm.OutputFilePath)
	if err != nil{
		return errors.New("failed to create file")
	}
	
	defer file.Close()

	time.Sleep(3* time.Second)

	encoder:=json.NewEncoder(file)
	err = encoder.Encode(data)
	if err != nil {
		file.Close()
		return errors.New("failed to write to file")
	}
	file.Close()
	return nil
}

func New(inputFilepath, outputFilepath string) *FileManager{
	return &FileManager{
		InputFilePath: inputFilepath,
		OutputFilePath: outputFilepath,
	}

}