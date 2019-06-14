package main

import "fmt"

// Item is node of doubly linked list
type Item struct {
	value interface{}
	prev  *Item
	next  *Item
	list *List
}

// Remove item from list
func (item *Item) Remove() {
	if item.prev == nil && item.next != nil {
		item.list.first = item.next
		item = nil
	}
}

//Value returns item value
func (item *Item) Value() interface{} {
	return item.value
}

//List is doubly linked list
type List struct {
	length uint
	first  *Item
	last   *Item
}

//PushFront adds an item to the end of the list
func (list *List) PushFront(value interface{}) {
	switch list.length {
	case 0:
		list.first = &Item{value, nil, nil, list}
		list.last = list.first
	case 1:
		list.last = &Item{value, list.first, nil, list}
		list.first.next = list.last
	default:
		temp := list.last
		list.last = &Item{value, temp, nil, list}
		temp.next = list.last
	}

	list.length++
}

//PushBack adds an item to the start of the list
func (list *List) PushBack(value interface{}) {
	switch list.length {
	case 0:
		list.first = &Item{value, nil, nil, list}
		list.last = list.first
	case 1:
		list.first = &Item{value, nil, list.first, list}
	default:
		temp := list.first
		list.first = &Item{value, nil, temp, list}
		temp.prev = list.first
	}

	list.length++
}

//Len returns size of doubly linked list
func (list *List) Len() uint {
	return list.length
}

//First returns first Item
func (list *List) First() *Item {
	return list.first
}

//Last returns last Item
func (list *List) Last() *Item {
	return list.last
}

func main() {
	list := List{}

	list.PushBack(1)
	list.PushBack(2)
	list.PushBack(3)
	list.PushBack(4)

	fmt.Println(list.Len())
	current := list.First()

	for current != nil {
		fmt.Println(current.Value())
		current = current.next
	}
}
