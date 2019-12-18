package datastruct

// StackNode => stack node
type StackNode struct {
	data int
	next *StackNode
}

// Stack => stack
type Stack struct {
	top *StackNode
}

// Push => push value in stack
func (stack *Stack) Push(i int) {
	data := &StackNode{data: i}
	if stack.top != nil {
		data.next = stack.top
	}

	stack.top = data
}

// Pop => pop value in stack
func (stack *Stack) Pop() (int, bool) {
	if stack.top == nil {
		return 0, false
	}

	i := stack.top.data
	stack.top = stack.top.next

	return i, true
}

// Peek => get the stack top value
func (stack *Stack) Peek() (int, bool) {
	if stack.top == nil {
		return 0, false
	}

	return stack.top.data, true
}

// Get => get values in stack
func (stack *Stack) Get() []int {
	var items []int

	current := stack.top

	for current != nil {
		items = append(items, current.data)
		current = current.next
	}

	return items
}

// IsEmpty => judge the stack is not it empty
func (stack *Stack) IsEmpty() bool {
	return stack.top == nil
}

// Empty => empty the stack
func (stack *Stack) Empty() {
	stack.top = nil
}
