package main

import (
	"errors"
	"fmt"
	"log"
	"reflect"
)

type element struct {
	name  string
	value int
}

var testData = []element{
	{name: "a", value: 1},
	{name: "b", value: 10},
	{name: "c", value: 22},
	{name: "d", value: 2},
	{name: "e", value: 3},
}

func main() {
	testFindMax()
	testFindMaxV2()
}

func testFindMax() {
	data := make([]interface{}, len(testData))

	// Convert test data to []interface{}
	for i, v := range testData {
		data[i] = v
	}

	maxElement, err := findMax(data, func(currentIndex, maxElementIndex int) bool {
		return testData[currentIndex].value > testData[maxElementIndex].value
	})

	if err != nil {
		log.Print(err)
	} else {
		fmt.Println("max element: ", maxElement)
	}
}

func testFindMaxV2() {
	maxElement, err := findMaxV2(testData, func(currentIndex, maxElementIndex int) bool {
		return testData[currentIndex].value > testData[maxElementIndex].value
	})

	if err != nil {
		log.Println(err)
	} else {
		fmt.Println("max element: ", maxElement)
	}
}

// Find maximum element without using reflection
func findMax(slice []interface{}, comparator func(elementIndex, maxElementIndex int) bool) (interface{}, error) {
	var err error
	var maxElement interface{}

	maxElementIndex := 0

	if len(slice) == 0 {
		err = errors.New("slice is empty")
	} else {
		maxElementIndex = 0

		for i := 1; i < len(slice); i++ {
			if comparator(i, maxElementIndex) {
				maxElementIndex = i
			}
		}

		maxElement = slice[maxElementIndex]
	}

	return maxElement, err
}

// Find maximum element with using reflection
func findMaxV2(slice interface{}, comparator func(elementIndex, maxElementIndex int) bool) (interface{}, error) {
	s, err := convertValueToSlice(reflect.ValueOf(slice))

	if err != nil {
		return s, err
	} else {
		return findMax(s, comparator)
	}
}

// Convert reflect value to []interface{}
func convertValueToSlice(rv reflect.Value) ([]interface{}, error) {
	var err error
	var slice []interface{}

	if rv.Kind() != reflect.Slice {
		err = errors.New("argument slice has incorrect type")
	} else if rv.Len() == 0 {
		err = errors.New("slice is empty")
	}

	if err == nil {
		slice = make([]interface{}, rv.Len())

		for i := 0; i < rv.Len(); i++ {
			slice[i] = rv.Index(i).Interface()
		}
	}

	return slice, err
}
