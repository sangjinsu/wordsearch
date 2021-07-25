package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"sync"
)

type Line struct {
	number int
	text   string
}

type FileData struct {
	fileName string
	lines    []Line
}

var wg sync.WaitGroup

func main() {
	if len(os.Args) < 3 {
		fmt.Println("실행 인수가 2개 이상 필요합니다.")
		return
	}

	word := os.Args[1]
	files := os.Args[2:]
	var fileDataList []FileData

	for _, path := range files {
		fileDataList = append(fileDataList, FindWordInFiles(word, path)...)
	}

	for _, fileData := range fileDataList {
		fmt.Println(fileData.fileName)
		fmt.Println("=============================================")
		for _, line := range fileData.lines {
			fmt.Println("\t", line.number, "\t", line.text)
		}
		fmt.Println("=============================================")
		fmt.Println()
	}
}

func GetFileList(path string) ([]string, error) {
	return filepath.Glob(path)
}

func FindWordInFiles(word string, filePath string) []FileData {
	var fileDataList []FileData

	fileList, err := GetFileList(filePath)
	if err != nil {
		fmt.Println("파일을 찾을 수 없습니다. err:", err)
		return fileDataList
	}

	fileDataChannel := make(chan FileData)

	for _, name := range fileList {
		wg.Add(1)
		go FindWordInFile(word, name, fileDataChannel)
	}

	for fileData := range fileDataChannel {
		fileDataList = append(fileDataList, fileData)
	}

	wg.Wait()
	close(fileDataChannel)

	return fileDataList
}

func FindWordInFile(word, fileName string, fileDataChannel chan FileData) {

	defer wg.Done()

	fileData := FileData{fileName, []Line{}}
	file, err := os.Open(fileName)
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			return
		}
	}(file)
	if err != nil {
		fmt.Println(fileName, "파일을 찾을 수 없습니다. err:", err)
		fileDataChannel <- fileData
	}

	lineNumber := 1
	scanner := bufio.NewScanner(file)

	wordToLower := strings.ToLower(word)
	for scanner.Scan() {
		text := scanner.Text()
		textToLower := strings.ToLower(text)
		if strings.Contains(textToLower, wordToLower) {
			fileData.lines = append(fileData.lines, Line{lineNumber, text})
		}
		lineNumber++
	}

	fileDataChannel <- fileData
}
