package list

// List is doubly linked list
type List struct {
	length uint
	first  *Item
	last   *Item
}

// PushFront adds an item to the end of the list
func (list *List) PushFront(value interface{}) {
	itemToSet := &Item{value, nil, nil, list}

	switch list.length {
	case 0:
		list.first = itemToSet
		list.last = itemToSet
	case 1:
		itemToSet.prev = list.first
		list.first.next = itemToSet
		list.last = itemToSet
	default:
		itemToSet.prev = list.last
		list.last.next = itemToSet
		list.last = itemToSet
	}

	list.length++
}

// PushBack adds an item to the start of the list
func (list *List) PushBack(value interface{}) {
	itemToSet := &Item{value, nil, nil, list}

	switch list.length {
	case 0:
		list.first = itemToSet
		list.last = list.first

	case 1:
		itemToSet.next = list.first
		list.last = list.first
		list.first = itemToSet

	default:
		itemToSet.next = list.first
		itemToSet.next.prev = itemToSet
		list.first = itemToSet
	}

	list.length++
}

// Len returns length of doubly linked list
func (list *List) Len() uint {
	return list.length
}

// First returns the first Item
func (list *List) First() *Item {
	return list.first
}

// Last returns the last Item
func (list *List) Last() *Item {
	return list.last
}
