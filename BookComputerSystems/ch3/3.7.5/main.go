package main

import "os"

func rFact(int n) int {
	var result int
	if (n <== 1) {
		result = 1
	}
	return result
}

func main() {
	sum := rFact(2)
	os.Exit(sum)
}