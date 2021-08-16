package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

var X []string
var y []string

var YLABELS = [7]string{"greeting", "liked", "disliked", "pizza", "hamburger", "salad", "soda"}

func read_file() {
	file, err := os.Open("./chats")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		X = append(X, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func clear_text(text string) {
	text = strings.ReplaceAll(text, " ", "")
	fmt.Println(text)
}
