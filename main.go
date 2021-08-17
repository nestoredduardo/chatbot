package main

import "fmt"

func main() {
	/* read_file()
	var X_train, y = data_preparation() */

	var prepared []int = prepare_sentence("Morrocan couscous is my favorite dish")
	avg := sentence_to_avg(prepared)

	fmt.Println(avg)
}
