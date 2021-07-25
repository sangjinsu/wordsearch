# wordsearch
## 단어 검색 프로그램

### 1. struct 구조체 이름 변경
```go
type Line struct {
	number int
	text   string
}

type FileData struct {
	fileName string
	lines    []Line
}
```

### 2. 소문자 대문자 상관 없이 단어 검색
```go
// FindWordInFile function 
wordToLower := strings.ToLower(word)
for scanner.Scan() {
	text := scanner.Text()
	textToLower := strings.ToLower(text)
	if strings.Contains(textToLower, wordToLower) {
		fileData.lines = append(fileData.lines, Line{lineNumber, text})
	}
	lineNumber++
}
```
