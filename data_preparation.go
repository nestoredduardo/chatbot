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
	"gonum.org/v1/gonum/mat"
)

var X []string
var y = [28]int{1, 1, 1, 1, 1, 1, 2, 2, 2, 2, 2, 3, 3, 3, 3, 3, 4, 4, 4, 5, 5, 5, 6, 6, 6, 7, 7, 7}

//var YLABELS = [7]string{"greeting", "liked", "disliked", "pizza", "hamburger", "salad", "soda"}
//var VOCABULARY = [27]string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z"}

var y_mapping = map[string]int{
	"greeting":  1,
	"liked":     2,
	"disliked":  3,
	"pizza":     4,
	"hamburger": 5,
	"salad":     6,
	"soda":      7,
}

var mapping = map[string]int{
	"a": 1,
	"b": 2,
	"c": 3,
	"d": 4,
	"e": 5,
	"f": 6,
	"g": 7,
	"h": 8,
	"i": 9,
	"j": 10,
	"k": 11,
	"l": 12,
	"m": 13,
	"n": 14,
	"ñ": 15,
	"o": 16,
	"p": 17,
	"q": 18,
	"r": 19,
	"s": 20,
	"t": 21,
	"u": 22,
	"v": 23,
	"w": 24,
	"x": 25,
	"y": 26,
	"z": 27,
}

const n_columns = 28

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

	reg, err := regexp.Compile("[^a-zA-Z!?]+")

	if err != nil {
		log.Fatal(err)
	}

	text2 := reg.ReplaceAllString(output, "")
	text2 = strings.Replace(text2, "?", "", -1)
	text2 = strings.ToLower(text2)
	return text2
}

func matPrint(X mat.Matrix) {
	fa := mat.Formatted(X, mat.Prefix(""), mat.Squeeze())
	fmt.Printf("%v\n", fa)
}

func data_preparation() {
	var X_prepared []string

	for i := 0; i < len(X); i++ {
		X_prepared = append(X_prepared, clear_text(X[i]))
	}

	X := mat.NewDense(len(X_prepared), n_columns, nil)

	for i := 0; i < len(X_prepared); i++ {
		var sentence string = X_prepared[i]
		var sentence_int []float64
		for j := 0; j < len(sentence); j++ {
			for key, value := range mapping {
				if string(sentence[j]) == key {
					sentence_int = append(sentence_int, float64(value))
				}
			}
		}
		if len(sentence_int) < n_columns {
			diff := n_columns - len(sentence_int)
			for i := 0; i < diff; i++ {
				sentence_int = append(sentence_int, 0)
			}
		}
		X.SetRow(i, sentence_int)
	}

	println("X:")
	matPrint(X)
}
