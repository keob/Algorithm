package datastruct

// QueueNode => node
type QueueNode struct {
	data int
	next *QueueNode
}

// Queue => queue
type Queue struct {
	rear *QueueNode
}

// Enqueue => insert item in queue
func (queue *Queue) Enqueue(i int) {
	data := &QueueNode{data: i}
	if queue.rear != nil {
		data.next = queue.rear
	}

	queue.rear = data
}

// Dequeue => remove item in queue
func (queue *Queue) Dequeue() (int, bool) {
	if queue.rear == nil {
		return 0, false
	}

	if queue.rear.next == nil {
		i := queue.rear.data
		queue.rear = nil
		return i, true
	}
	current := queue.rear

	for {
		if current.next.next == nil {
			i := current.next.data
			current.next = nil
			return i, true
		}
		current = current.next
	}
}

// Peek => return head value in queue
func (queue *Queue) Peek() (int, bool) {
	if queue.rear == nil {
		return 0, false
	}
	return queue.rear.data, true
}

// Get => get values in queue
func (queue *Queue) Get() []int {
	var items []int
	current := queue.rear
	for current != nil {
		items = append(items, current.data)
		current = current.next
	}

	return items
}

// IsEmpty => judge the queue is not it empty
func (queue *Queue) IsEmpty() bool {
	return queue.rear == nil
}

// Empty => empty the queue
func (queue *Queue) Empty() {
	queue.rear = nil
}
