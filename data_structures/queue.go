package data_structures

import "errors"

type Queue struct {
	First  *Node
	Last   *Node
	length int
}

func NewQueue() *Queue {
	return &Queue{
		length: 0,
	}
}

func (q *Queue) Add(data int) {
	item := Node{
		Data: data,
	}

	if q.First == nil {
		q.First = &item
	}

	if q.Last != nil {
		q.Last.Next = &item
	}

	q.Last = &item
	q.length++
}

func (q *Queue) Remove() (int, error) {
	if q.IsEmpty() {
		return 0, errors.New("the queue is empty")
	}

	item := q.First
	q.First = q.First.Next

	if q.Last == item {
		q.Last = q.First
	}

	q.length--

	return item.Data, nil
}

func (q *Queue) Peek() (int, error) {
	if q.IsEmpty() {
		return 0, errors.New("the queue is empty")
	}

	return q.First.Data, nil
}

func (q *Queue) IsEmpty() bool {
	return q.First == nil
}

func (q *Queue) Length() int {
	return q.length
}
