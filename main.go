package main

import (
	"fmt"
	"os"
	"regexp"
	"strings"
)

const (
	ChunkSize   = 10
	OverlapSize = 3
)

func main() {

	content, err := os.ReadFile("THEFILE.txt")
	if err != nil {
		panic("Error loading file")
	}

	text := string(content)
	// chunks := chunkTextOnSize(text)
	chunks := chunkTextOnDelimiter(text, " ")

	chunks = filterLower(chunks)
	chunks = filterAlphaNumeric(chunks)

	for _, c := range chunks {
		fmt.Println(c)
	}

}

func filterAlphaNumeric(chunks []string) []string {
	re := regexp.MustCompile(`[^a-zA-z0-9 ]`)
	var filteredChunks []string
	for _, chunk := range chunks {
		filteredChunks = append(filteredChunks, re.ReplaceAllString(chunk, ""))
	}
	return filteredChunks
}

func filterLower(chunks []string) []string {
	var filteredChunks []string
	for _, chunk := range chunks {
		filteredChunks = append(filteredChunks, strings.ToLower(chunk))
	}
	return filteredChunks
}

func chunkTextOnDelimiter(inputText string, delimiter string) []string {
	re := regexp.MustCompile(delimiter)
	chunks := re.Split(inputText, -1)
	var buffer []string

	startIndex := 0
	step := ChunkSize - OverlapSize
	endIndex := startIndex + OverlapSize

	for {
		buffer = append(buffer, strings.Join(chunks[startIndex:endIndex], delimiter))
		startIndex += step
		if startIndex > len(chunks) {
			break
		}

		if endIndex == len(chunks) {
			break
		}

		endIndex = startIndex + ChunkSize
		if endIndex > len(chunks) {
			endIndex = len(chunks)
		}

	}
	return buffer
}

func chunkTextOnSize(inputText string) []string {
	var chunks []string
	var buffer []rune
	var overlapBuffer []rune

	overlapBufferLength := 0
	bufferLength := 0

	for _, r := range inputText {
		if overlapBufferLength > bufferLength {
			buffer = append(overlapBuffer, buffer...)
			bufferLength = OverlapSize
			overlapBuffer = []rune{}
			overlapBufferLength = 0
		}

		if bufferLength == ChunkSize {
			chunks = append(chunks, string(buffer))
			bufferTail := buffer[len(buffer)-OverlapSize:]
			overlapBuffer = append(overlapBuffer, bufferTail...)
			buffer = []rune{}
			bufferLength = 0
			overlapBufferLength = OverlapSize
		}

		buffer = append(buffer, r)
		bufferLength += 1
	}

	if bufferLength > 0 {
		chunks = append(chunks, string(buffer))
	}
	return chunks

}
