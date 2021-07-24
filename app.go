package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Println("실행 인수가 2개 이상 필요합니다.")
		return
	}

	word := os.Args[1]
	files := os.Args[2:]

	fmt.Println("찾으려는 단어:", word)
	PrintAllFiles(files)
}

func GetFileList(path string) ([]string, error) {
	return filepath.Glob(path)
}

func PrintAllFiles(filePaths []string) {
	for _, path := range filePaths {
		fileList, err := GetFileList(path)
		if err != nil {
			fmt.Println("파일을 찾을 수 없습니다. err:", err)
			return
		}

		fmt.Println("찾으려는 파일 리스트")
		for _, name := range fileList {
			fmt.Println(name)
		}
	}
}
