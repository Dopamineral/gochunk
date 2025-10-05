package main

import (
	"fmt"
	"os"
)

const (
	ChunkSize   = 20
	OverlapSize = 5
)

func main() {

	content, err := os.ReadFile("THEFILE.txt")
	if err != nil {
		panic("Error loading file")
	}

	text := string(content)
	chunks := chunk(text)

	for _, c := range chunks {
		fmt.Println(c)
	}

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
