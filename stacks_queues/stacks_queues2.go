package stacksqueues

import "fmt"

// Queues represent a queue that holds a slice
type Queue struct {
	items []int
}

// Enqueue
func (q *Queue) Enqueue(i int) {
	q.items = append(q.items, i)
}

// Deqeue
func (q *Queue) Deqeue() int {
	toRemove := q.items[0]
	q.items = q.items[1:]
	return toRemove
}

func Demo2() {
	myQueue := Queue{}
	myQueue.Enqueue(25)
	myQueue.Enqueue(50)
	myQueue.Enqueue(75)

	fmt.Println(myQueue)

	myQueue.Deqeue()
	fmt.Println(myQueue)
}
