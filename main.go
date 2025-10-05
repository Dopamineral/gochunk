package main

import (
	"flag"
	"fmt"
	"os"
	"regexp"
	"strings"
)

func main() {
	size := flag.Int("size", 50, "Size of the chunk")
	overlap := flag.Int("overlap", 10, "Overlap between chunks")

	flag.Parse()
	fileName := flag.Args()[0]
	// strange: in go you have to place the flags before the ARGs,
	// -> ./main --size=5 --overlap=1 THEFILE.txt

	content, err := os.ReadFile(fileName)
	if err != nil {
		panic("Error loading file")
	}

	text := string(content)
	chunks := chunkTextOnDelimiter(text, " ", *size, *overlap)
	chunks = filterLower(chunks)
	chunks = filterAlphaNumeric(chunks)

	for _, c := range chunks {
		fmt.Println(c)
	}
	fmt.Println(*size, *overlap)

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

func chunkTextOnDelimiter(inputText string, delimiter string, size int, overlap int) []string {
	re := regexp.MustCompile(delimiter)
	chunks := re.Split(inputText, -1)
	var buffer []string

	startIndex := 0
	step := size - overlap
	endIndex := startIndex + overlap

	for {
		buffer = append(buffer, strings.Join(chunks[startIndex:endIndex], delimiter))
		startIndex += step
		if startIndex > len(chunks) {
			break
		}

		if endIndex == len(chunks) {
			break
		}

		endIndex = startIndex + size
		if endIndex > len(chunks) {
			endIndex = len(chunks)
		}

	}
	return buffer
}

func chunkTextOnSize(inputText string, overlap int, size int) []string {
	var chunks []string
	var buffer []rune
	var overlapBuffer []rune

	overlapBufferLength := 0
	bufferLength := 0

	for _, r := range inputText {
		if overlapBufferLength > bufferLength {
			buffer = append(overlapBuffer, buffer...)
			bufferLength = overlap
			overlapBuffer = []rune{}
			overlapBufferLength = 0
		}

		if bufferLength == size {
			chunks = append(chunks, string(buffer))
			bufferTail := buffer[len(buffer)-overlap:]
			overlapBuffer = append(overlapBuffer, bufferTail...)
			buffer = []rune{}
			bufferLength = 0
			overlapBufferLength = overlap
		}

		buffer = append(buffer, r)
		bufferLength += 1
	}

	if bufferLength > 0 {
		chunks = append(chunks, string(buffer))
	}
	return chunks

}
