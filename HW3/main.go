package main

import (
	"fmt"
	"sort"
	"strings"
	"unicode"
)

const text string = `Все гости совершали обряд приветствования никому не известной, никому не интересной и не нужной тетушки. Анна Павловна с грустным, торжественным участием следила за их приветствиями, молчаливо одобряя их. Ma tante каждому говорила в одних и тех же выражениях о его здоровье, о своем здоровье и о здоровье ее величества, которое нынче было, слава Богу, лучше. Все подходившие, из приличия не выказывая поспешности, с чувством облегчения исполненной тяжелой обязанности отходили от старушки, чтоб уж весь вечер ни разу не подойти к ней.
Молодая княгиня Болконская приехала с работой в шитом золотом бархатном мешке. Ее хорошенькая, с чуть черневшимися усиками верхняя губка была коротка по зубам, но тем милее она открывалась и тем еще милее вытягивалась иногда и опускалась на нижнюю. Как это бывает у вполне привлекательных женщин, недостаток ее — короткость губы и полуоткрытый рот — казались ее особенною, собственно ее красотой. Всем было весело смотреть на эту полную здоровья и живости хорошенькую будущую мать, так легко переносившую свое положение. Старикам и скучающим, мрачным молодым людям казалось, что они сами делаются похожи на нее, побыв и поговорив несколько времени с ней. Кто говорил с ней и видел при каждом слове ее светлую улыбочку и блестящие белые зубы, которые виднелись беспрестанно, тот думал, что он особенно нынче любезен. И это думал каждый.
Маленькая княгиня, переваливаясь, маленькими быстрыми шажками обошла стол с рабочею сумочкой на руке и, весело оправляя платье, села на диван, около серебряного самовара, как будто все, что она ни делала, было partie de plaisir 3 для нее и для всех ее окружавших.`

type word struct {
	word  string
	count uint
}

func main() {
	wordCache := calcWords(strings.FieldsFunc(text, isRuneLetter))
	printTopWords(makeWordList(wordCache), 10)
}

func isRuneLetter(r rune) bool {
	return !unicode.IsLetter(r)
}

func calcWords(words []string) map[string]word {
	wordCache := map[string]word{}

	for _, v := range words {
		value, ok := wordCache[v]

		if ok {
			value.count++
			wordCache[v] = value
		} else {
			wordCache[v] = word{word: v, count: 1}
		}
	}

	return wordCache
}

func makeWordList(cache map[string]word) []word {
	wordList := make([]word, 0, len(cache))

	for _, value := range cache {
		wordList = append(wordList, value)
	}

	return wordList
}

func printTopWords(wordList []word, top int) {
	copiedWords := make([]word, len(wordList))

	copy(copiedWords, wordList)

	sort.Slice(copiedWords, func(i, j int) bool {
		return copiedWords[i].count > copiedWords[j].count
	})

	for i := 0; i < top && i < len(copiedWords); i++ {
		fmt.Println(i+1, "\t: ", copiedWords[i])
	}
}
