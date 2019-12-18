package datastruct

// LinkedListNode ...
type LinkedListNode struct {
	data int
	next *LinkedListNode
}

// LinkList => link list
type LinkList struct {
	head *LinkedListNode
}

// InsertFirst => insert in first
func (list *LinkList) InsertFirst(i int) {
	data := &LinkedListNode{data: i}
	if list.head != nil {
		data.next = list.head
	}

	list.head = data
}

// InsertLast => insert in last
func (list *LinkList) InsertLast(i int) {
	data := &LinkedListNode{data: i}
	if list.head == nil {
		list.head = data
		return
	}

	current := list.head

	for current.next != nil {
		current = current.next
	}

	current.next = data
}

// RemoveByValue => remove item by value
func (list *LinkList) RemoveByValue(i int) bool {
	if list.head == nil {
		return false
	}

	if list.head.data == i {
		list.head = list.head.next
		return true
	}

	current := list.head

	for current.next != nil {
		if current.next.data == i {
			current.next = current.next.next
			return true
		}

		current = current.next
	}

	return false
}

// RemoveByIndex => remove item by index
func (list *LinkList) RemoveByIndex(i int) bool {
	if list.head == nil {
		return false
	}

	if i < 0 {
		return false
	}

	if i == 0 {
		list.head = list.head.next
		return true
	}

	current := list.head

	for u := 1; i < i; u++ {
		if current.next.next == nil {
			return false
		}
		current = current.next
	}

	current.next = current.next.next

	return true
}

// SearchValue => search value in linklist
func (list *LinkList) SearchValue(target int) bool {
	if list.head == nil {
		return false
	}

	current := list.head

	for current != nil {
		if current.data == target {
			return true
		}
		current = current.next
	}

	return false
}

// GetFirst => get first value in linklist
func (list *LinkList) GetFirst() (int, bool) {
	if list.head == nil {
		return 0, false
	}

	return list.head.data, true
}

// GetLast => get last value in linklist
func (list *LinkList) GetLast() (int, bool) {
	if list.head == nil {
		return 0, false
	}

	current := list.head
	for current.next != nil {
		current = current.next
	}

	return current.data, true
}

// GetSize => get linklist size
func (list *LinkList) GetSize() int {
	count := 0
	current := list.head
	for current != nil {
		count++
		current = current.next
	}

	return count
}

// GetItems => get linklist items
func (list *LinkList) GetItems() []int {
	var items []int
	current := list.head
	for current != nil {
		items = append(items, current.data)
		current = current.next
	}

	return items
}
