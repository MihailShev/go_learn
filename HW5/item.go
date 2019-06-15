package main

// Item is node of doubly linked list
type Item struct {
	value interface{}
	prev  *Item
	next  *Item
	list  *List
}

//Value returns item value
func (item *Item) Value() interface{} {
	return item.value
}

//Next returns next *item
func (item *Item) Next() *Item {
	return item.next
}

//Prev returns prev *item
func (item *Item) Prev() *Item {
	return item.prev
}

// Remove item from list
func (item *Item) Remove() {
	switch {
	//item is first
	case item.prev == nil && item.next != nil:
		item.list.first = item.next

	//item is last
	case item.prev !=nil && item.next == nil:
		item.list.last = item.prev
		item.list.last.next = nil

	//item is on the list
	case item.list.length == 1:
		item.list.first = nil
		item.list.last = nil

	default:
		item.prev.next = item.next
		item.next.prev = item.prev
	}

	item.list.length--
}
