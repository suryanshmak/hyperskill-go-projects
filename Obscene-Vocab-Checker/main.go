package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode/utf8"
)

func main() {
	var fileName string
	fmt.Println("Enter the file name: ")
	fmt.Scan(&fileName)

	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	scanner.Split(bufio.ScanWords)
	cussWords := make(map[string]bool)

	for scanner.Scan() {
		cussWords[scanner.Text()] = true
	}

	var sentence string

	for sentence != "exit" {
		fmt.Scan(&sentence)
		words := strings.Split(sentence, " ")
		for i, word := range words {
			if _, ok := cussWords[strings.ToLower(word)]; ok {
				words[i] = strings.Repeat("*", utf8.RuneCountInString(word))
			}
		}
		fmt.Println(strings.Join(words, " "))
	}
	fmt.Println("Bye!")
}
