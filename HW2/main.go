package main

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

type runeType int

const (
	symbol runeType = iota
	count
	escapeSymbol
	escape
)
const escapeRune rune = '\\'

var input = []string{"a4bc2d5e", "abcd", `qwe\4\5`, `qwe\45`, `qwe\\5`}
var output = []string{"aaaabccddddde", "abcd", `qwe45`, `qwe44444`, `qwe\\\\\`}

func main() {
	// fmt.Println("symbol 0 ->", defineRuneType('\\', '4'))
	// fmt.Println("count 1 -> ", defineRuneType('a', '2'))
	// fmt.Println("escapeSymbol 2 -> ", defineRuneType('\\', '\\'))
	// fmt.Println("escape 3 -> ", defineRuneType('4', '\\'))

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
		currentRune := runes[i]
		if i > 0 {
			prevRune := runes[i-1]
			symbolType := defineRuneType(prevRune, currentRune)
			
			switch {
			case symbolType == symbol || symbolType == escapeSymbol:
				strBuilder.WriteRune(currentRune)
			case symbolType == count:
				countS, _ := strconv.Atoi(string(runes[i]))
				strBuilder.WriteString(multiplyRune(prevRune, countS))	
			}
		} else {
			strBuilder.WriteRune(currentRune)
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

func defineRuneType(prevRune rune, currentRune rune) runeType {
	isNumber := unicode.IsDigit(currentRune)

	switch {
	case prevRune != escapeRune && isNumber:
		return count
	case prevRune == escapeRune && currentRune == escapeRune:
		return escapeSymbol
	case prevRune != escapeRune && currentRune == escapeRune:
		return escape
	default:
		return symbol
	}
}
