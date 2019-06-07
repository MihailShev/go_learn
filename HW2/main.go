package main

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

type runeType int

const (
	symbolType runeType = iota
	countType
	escapeSymbolType
	escapeType
	errorType
	emptyType
)

const escapeRune rune = '\\'

func main() {
	input := [6]string{"a4bc2d5e", "abcd", "45", `qwe\4\5`, `qwe\45`, `qwe\\5`}
	output := [6]string{"aaaabccddddde", "abcd", "", `qwe45`, `qwe44444`, `qwe\\\\\`}

	//test
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
	prevType := emptyType

	strBuilder.WriteRune(runes[0])

	for i := 1; i < len(runes); i++ {

		currentRune := runes[i]
		prevRune := runes[i-1]
		currentType := defineRuneType(prevRune, currentRune, prevType)
		prevType = currentType

		switch {
		case currentType == symbolType || currentType == escapeSymbolType:
			strBuilder.WriteRune(currentRune)

		case currentType == countType:
			strBuilder.WriteString(multiplyRune(prevRune, currentRune))

		case currentType == errorType:
			strBuilder.Reset()
			break
		}
	}

	return strBuilder.String()
}

func multiplyRune(runeToMultiply rune, countRune rune) string {
	var strBuilder strings.Builder
	count, err := strconv.Atoi(string(countRune))

	if err != nil {
		panic("Can't convert countRune to int")
	}

	for i := 1; i < count; i++ {
		strBuilder.WriteRune(runeToMultiply)
	}

	return strBuilder.String()
}

func defineRuneType(prevRune rune, currentRune rune, prevType runeType) runeType {
	isNumber := unicode.IsDigit(currentRune)
	isPrevNumber := unicode.IsDigit(prevRune)

	switch {
	case prevType != symbolType && isPrevNumber && isNumber:
		return errorType
	case (prevRune != escapeRune && isNumber) || (prevType == escapeSymbolType && isNumber):
		return countType
	case prevRune == escapeRune && currentRune == escapeRune:
		return escapeSymbolType
	case prevRune != escapeRune && currentRune == escapeRune:
		return escapeType
	default:
		return symbolType
	}
}
