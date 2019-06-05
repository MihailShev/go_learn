package main

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

var input = []string{"a4bc2d5e", "abcd", `qwe\4\5`, `qwe\45`, `qwe\\5`}
var output = []string{"aaaabccddddde", "abcd", `qwe45`, `qwe44444`, `qwe\\\\\`}
const escape rune = '\\'

func main() {
	for i, v := range input {
		result := strUnpacker(v)

		if result == output[i] {
			fmt.Println("Ok ", v, " ➜ ", result)
		} else {
			fmt.Println("Fail ", v, " ❌ ", result)
		}
	}
}

func strUnpacker(str string) string {
	var strBuilder strings.Builder

	runes := []rune(str)

	for i := 0; i < len(runes); i++ {
		if i > 0 && runes[i] == escape && runes[i - 1] != escape  {
			continue
		}

		if unicode.IsDigit(runes[i]) && i > 0 && runes[i - 1] != escape {

			count, _ := strconv.Atoi(string(runes[i]))

			strBuilder.WriteString(multiplyRune(runes[i-1], count))
		} else {
			strBuilder.WriteRune(runes[i])
		}
	}

	return strBuilder.String()
}

func multiplyRune(r rune, count int) string {
	var strBuilder strings.Builder

	for i := 1; i < count; i++ {
		strBuilder.WriteRune(r)
	}

	return strBuilder.String()
}
