package main

import (
	"fmt"
	"strconv"
	"strings"
)

func discountPrices(sentence string, discount int) string {
	words := strings.Split(sentence, " ")
	//newString := ""
	for i := 0; i < len(words); i++ {
		p := isPrice(words[i])
		if p == -1 {
			//newString += words[i]
		} else {
			p -= (float64(discount) / 100) * p
			//newString += fmt.Sprintf("$%.2f", p)
			words[i] = fmt.Sprintf("$%.2f", p)
		}
		//newString += " "
	}
	//newString = newString[:len(newString)-1]
	return strings.Join(words, " ")
}

// 传入空格后  $xxx
func isPrice(sentence string) float64 {
	if sentence[0] != '$' {
		return -1
	}
	f, err := strconv.ParseFloat(sentence[1:], 64)
	if (f == 0 && sentence == "$0") || err != nil {
		return -1
	} else {
		return f
	}
}

func main() {
	fmt.Println(discountPrices("there are $1 $2 and 5$ candies in the shop", 100))
}
