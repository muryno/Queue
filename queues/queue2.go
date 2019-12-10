package queues

import "fmt"

type Queue2 struct {
	slice []int
}

// Enqueue add integer provided to the back
//of the queue
func (q *Queue2) Enqueue(i int)  {
	q.slice = append(q.slice,i)
}
// Dequeue return the first item in the queue
// return first item from the queue
// or panic if there isnt one
func (q *Queue2) Dequeue() int {

	req := q.slice[0]

	q.slice = q.slice[1:len(q.slice)]

	return req
}


func (q *Queue2) String() string  {
	return fmt.Sprint(q.slice)
}
