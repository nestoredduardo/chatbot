package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"
	"unicode"

	"golang.org/x/text/runes"
	"golang.org/x/text/transform"
	"golang.org/x/text/unicode/norm"
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

func clear_text(text string) string {
	t := transform.Chain(norm.NFD, runes.Remove(runes.In(unicode.Mn)), norm.NFC)
	output, _, e := transform.String(t, text)
	if e != nil {
		panic(e)
	}

	reg, err := regexp.Compile("[^a-zA-Z0-9!?]+")

	if err != nil {
		log.Fatal(err)
	}

	text2 := reg.ReplaceAllString(output, "")
	text2 = strings.Replace(text2, "?", "", -1)
	return text2
}

func data_preparation() {
	var X_prepared []string
	for i := 0; i < len(X); i++ {
		X_prepared = append(X_prepared, clear_text(X[i]))
	}
	fmt.Println(X_prepared)
}
