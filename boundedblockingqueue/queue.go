package boundedblockingqueue

import "fmt"

type Queue struct {
	queue []int
	capacity int
}

func (q *Queue) size() int{
	return len(q.queue)
}

func (q *Queue) enqueue(element int){
	q.queue = append(q.queue, element)
}

func (q *Queue) dequeue() int {
	var element int
	if len(q.queue) > 0{
		element = q.queue[0]
	}
	q.queue = q.queue[1:]
	return element
}

func Start(){
	q := Queue{
		queue:    nil,
		capacity: 5,
	}

	q.enqueue(1)
	q.enqueue(2)
	q.enqueue(3)
	fmt.Print(q.queue)
	q.dequeue()
	q.dequeue()
	fmt.Print(q.queue)
}