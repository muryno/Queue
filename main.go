package main

import "fmt"




	type Queue struct {
		slice []int
	}

	// Enqueue add integer provided to the back
	//of the queue
	func (q *Queue) Enqueue(i int)  {
		q.slice = append(q.slice,i)
	}
	// Dequeue return the first item in the queue
	// return first item from the queue
	// or panic if there isnt one
	func (q *Queue) Dequeue() int {

		req := q.slice[0]

		q.slice = q.slice[1:len(q.slice)]

		return req
	}

	
	func (q *Queue) string() string  {
		return fmt.Sprint(q.slice)
	}

	func main() {

		var	q  = new (Queue)

		q.Enqueue(98765)
		q.Enqueue(129876)
		q.Enqueue(8765)
		fmt.Println(q.string())

		fmt.Println(q.Dequeue())
		fmt.Println(q.Dequeue())
	}

