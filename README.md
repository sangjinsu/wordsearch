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

### 3. recvCnt를 sync.WaitGroup 으로 변경

`var wg sync.WaitGroup` 전역 변수에 추가

```go
// FindWordInFiles 에서 waitGroup 하나씩 추가
for _, name := range fileList {
	wg.Add(1)
	go FindWordInFile(word, name, fileDataChannel)
}

/////////////////////////////////////
// FindWordInFile 에서 waitGroup 종료 신호 전달
func FindWordInFile(word, fileName string, fileDataChannel chan FileData) {
	defer wg.Done()
    
//////////////////////////////////////
// FindWordInFiles 에서 goroutine 종료 대기 후 채널 닫기
    
// unbuffered channel
go func() {
    wg.Wait()
    close(fileDataChannel)
}()
    
// buffered channel
wg.Wait()
close(fileDataChannel)

```

#### 추가자료: [buffered channel, unbuffered channel 사용시 wg.Wait() 주의사항](https://stackoverflow.com/questions/46560204/why-does-my-code-work-correctly-when-i-run-wg-wait-inside-a-goroutine)

