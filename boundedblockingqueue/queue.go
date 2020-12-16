package boundedblockingqueue

import (
	"context"
	"fmt"
	"golang.org/x/sync/semaphore"
	"sync"
)

type Queue struct {
	queue []int
	capacity int

	wg *sync.WaitGroup
	enqueueSem *semaphore.Weighted
	dequeueSem *semaphore.Weighted
	mutex *sync.Mutex
	ctx context.Context
}

func (q *Queue) size() int{
	return len(q.queue)
}

func (q *Queue) enqueue(element int){
	q.wg.Add(1)
	go func() {
		defer q.wg.Done()
		q.enqueueSem.Acquire(q.ctx, 1)
		q.mutex.Lock()
		q.queue = append(q.queue, element)
		q.mutex.Unlock()
		q.dequeueSem.Release(1)
	}()
}

func (q *Queue) dequeue() int {
	q.wg.Add(1)
	element := make(chan int)
	go func() {
		defer q.wg.Done()
		q.dequeueSem.Acquire(q.ctx, 1)
		q.mutex.Lock()
		if len(q.queue) > 0{
			element <- q.queue[0]
		}
		q.queue = q.queue[1:]
		q.mutex.Unlock()
		q.enqueueSem.Release(1)
	}()
	return <-element
}

func (q *Queue) init(capacity int){
	q.capacity = capacity
	q.enqueueSem = semaphore.NewWeighted(int64(capacity))
	q.dequeueSem = semaphore.NewWeighted(int64(capacity))
	q.ctx = context.Background()
	q.dequeueSem.Acquire(q.ctx, int64(q.capacity))
}

func Start(){
	q := Queue{
		wg: &sync.WaitGroup{},
		mutex: &sync.Mutex{},
	}

	q.init(5)
	q.enqueue(1)
	q.enqueue(2)
	q.enqueue(3)
	q.enqueue(1)
	q.enqueue(2)
	q.enqueue(3)
	q.dequeue()
	q.dequeue()

	q.wg.Wait()
	fmt.Print(q.queue)
}