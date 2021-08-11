package queue

import "fmt"

type queue struct {
	array  []int
	length int
}

func (q *queue) Enqueue(elem int) {
	q.array = append(q.array, elem)
	q.length++
}

func (q *queue) Dequeue() {
	q.array = q.array[1:]
	q.length--
}

func (q *queue) isEmpty() bool {
	return q.length == 0
}

func (q *queue) Length() int {
	return q.length
}

func Queue() {
	var q queue
	q.Enqueue(8)
	q.Enqueue(1)
	q.Enqueue(6)
	q.Enqueue(10)
	q.Enqueue(3)
	q.Enqueue(7)
	q.Enqueue(2)
	q.Enqueue(5)
	fmt.Println(q.array)
	q.Dequeue()
	fmt.Println(q.array)
	q.Dequeue()
	fmt.Println(q.array)
	fmt.Println(q.isEmpty())
	fmt.Println(q.length)
}
