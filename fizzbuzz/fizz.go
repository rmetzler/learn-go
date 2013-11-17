package fizzbuzz

import "strconv"

func fizzbuzz(i int) string {
	if i%15 == 0 {
		return "fizzbuzz"
	}
	if i%5 == 0 {
		return "buzz"
	}
	if i%3 == 0 {
		return "fizz"
	}
	return strconv.Itoa(i) // convert int to string
}
