package main

import "fmt"

func prepare_sentence(text string) []int {
	text = clear_text(text)
	sentence_int := []int{}

	for j := 0; j < len(text); j++ {
		for key, value := range mapping {
			if string(text[j]) == key {

				sentence_int = append(sentence_int, int(value))
			}
		}
	}

	if len(sentence_int) < n_columns {
		diff := n_columns - len(sentence_int)
		for i := 0; i < diff; i++ {
			sentence_int = append(sentence_int, 0)
		}
	}

	fmt.Println(sentence_int)

	return sentence_int
}

func sentence_to_onehot() {}

/* func sentence_to_avg(sentence []int) float64 {
	var avg float64

	for i := 0; i < len(sentence); i++ {
		avg += float64(sentence[i])
	}
	avg = avg / float64(len(sentence))

	return avg
} */
