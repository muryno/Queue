package main

import (
	"Queue/queues"
	"fmt"
)





	func main() {

		var	q  = new (queues.Queue2)

		q.Enqueue(98765)
		q.Enqueue(129876)
		q.Enqueue(8765)
		fmt.Println(q.String())

		fmt.Println(q.Dequeue())
		fmt.Println(q.Dequeue())
	}

