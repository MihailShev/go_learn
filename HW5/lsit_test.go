package main

import (
	"fmt"
	"strings"
	"testing"
)

func printListToString(list List) string {
	var strBuilder strings.Builder

	current := list.First()

	for current != nil {
		strBuilder.WriteString(fmt.Sprint(current.Value()))

		if current.next != nil {
			strBuilder.WriteString(" ")
		}

		current = current.next
	}

	return strBuilder.String()
}

func isEqualString(t *testing.T, expected string, got string) {
	if expected != got {
		t.Error("\nExp:\n", expected, "\nGot:\n", got)
	}
}

func printError(t *testing.T, exp interface{}, got interface{}) {
	t.Error("\nExp:\n", exp, "\nGot:\n", got)
}

func populateList() List {
	list := List{}

	for list.length < 10 {
		list.PushFront(list.length + 1)
	}

	return list
}

func TestList_PushBack(t *testing.T) {
	list := List{}

	list.PushBack(1)

	expected := "1"
	got := printListToString(list)
	isEqualString(t, expected, got)


	list.PushBack(2)
	list.PushBack(3)

	expected = "3 2 1"
	got = printListToString(list)
	isEqualString(t, expected, got)
}

func TestList_PushFront(t *testing.T) {
	list := List{}

	list.PushFront(1)

	expected := "1"
	got := printListToString(list)
	isEqualString(t, expected, got)


	list.PushFront(2)
	list.PushFront(3)

	expected = "1 2 3"
	got = printListToString(list)
	isEqualString(t, expected, got)
}

func TestList_First(t *testing.T) {
	list := populateList()

	expected := "1"
	got := fmt.Sprint(list.First().Value())

	isEqualString(t, expected, got)
}

func TestList_Last(t *testing.T) {
	list := populateList()

	expected := "10"
	got := fmt.Sprint(list.Last().Value())

	isEqualString(t, expected, got)
}

func TestList_Len(t *testing.T) {
	list := populateList()

	expected := uint(10)
	got := list.Len()

	if got != expected {
		printError(t, expected, got)
	}
}