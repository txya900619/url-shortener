package queue

import "errors"

type StringQueue struct {
	capacity int
	q        chan string
}

func (q *StringQueue) Insert(s string) error {
	if len(q.q) < q.capacity {
		q.q <- s
		return nil
	}

	return errors.New("queue is full")
}

func (q *StringQueue) Remove() (string, error) {
	if len(q.q) == 0 {
		return "", errors.New("queue is empty")
	}

	s := <-q.q
	return s, nil
}

func NewStringQueue(capacity int) *StringQueue {
	return &StringQueue{
		capacity: capacity,
		q:        make(chan string, capacity),
	}
}
