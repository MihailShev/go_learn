package main

import (
	"fmt"
	"strings"
	"testing"
)

//Test Item

func TestItem_Value(t *testing.T) {
	list := List{}

	list.PushFront(1)
	list.PushFront("some test value")

	expVal1 := 1
	expVal2 := "some test value"
	val1 := list.First().Value()
	val2 := list.Last().Value()

	if expVal1 != val1 {
		printError(t, expVal1, val1)
	}

	if val2 != expVal2 {
		printError(t, expVal2, val2)
	}
}

func TestItem_Prev(t *testing.T) {
	list := List{}

	list.PushFront(1)
	list.PushFront("some test value")

	expected := 1
	got := list.Last().Prev().Value()

	if expected != got {
		printError(t, expected, got)
	}
}

func TestItem_Next(t *testing.T) {
	list := List{}

	list.PushFront(1)
	list.PushFront("some test value")

	expected := "some test value"
	got := list.First().Next().Value()

	if expected != got {
		printError(t, expected, got)
	}
}

func TestItem_RemoveFirstItem(t *testing.T) {
	list := populateList(10)

	_ = list.first.Remove()

	expected := "2 3 4 5 6 7 8 9 10"
	got := printListToString(*list)

	if expected != got {
		printError(t, expected, got)
	}
}

func TestItem_RemoveLast(t *testing.T) {
	list := populateList(10)

	_ = list.Last().Remove()

	expected := "1 2 3 4 5 6 7 8 9"
	got := printListToString(*list)

	if expected != got {
		printError(t, expected, got)
	}
}

func TestItem_Remove(t *testing.T) {
	list := populateList(10)

	_ = list.First().Next().Next().Remove()

	expected := "1 2 4 5 6 7 8 9 10"
	got := printListToString(*list)

	if expected != got {
		printError(t, expected, got)
	}
}

func TestItem_RemoveError(t *testing.T) {
	list := populateList(10)
	 _ = list.First().Remove()

	got := list.First().Remove()
	expected := removeErrorMessage

	if expected != got.Error() {
		printError(t, expected, got)
	}
}

// Test List

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
	list := populateList(10)

	expected := "1"
	got := fmt.Sprint(list.First().Value())

	isEqualString(t, expected, got)
}

func TestList_Last(t *testing.T) {
	list := populateList(10)

	expected := "10"
	got := fmt.Sprint(list.Last().Value())

	isEqualString(t, expected, got)
}

func TestList_Len(t *testing.T) {
	list := populateList(5)

	expected := uint(5)
	got := list.Len()

	if got != expected {
		printError(t, expected, got)
	}
}

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

//Utils

func isEqualString(t *testing.T, expected string, got string) {
	if expected != got {
		t.Error("\nExp:\n", expected, "\nGot:\n", got)
	}
}

func printError(t *testing.T, exp interface{}, got interface{}) {
	t.Error("\nExp:\n", exp, "\nGot:\n", got)
}

func populateList(count uint) *List {
	list := List{}

	for list.length < count {
		list.PushFront(list.length + 1)
	}

	return &list
}