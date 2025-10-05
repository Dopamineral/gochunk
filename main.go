package main

import (
	"fmt"
	"os"
	"regexp"
	"strings"
)

const (
	ChunkSize   = 80
	OverlapSize = 10
)

func main() {

	content, err := os.ReadFile("THEFILE.txt")
	if err != nil {
		panic("Error loading file")
	}

	text := string(content)
	chunks := chunk(text)
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

func chunk(inputText string) []string {
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
