package dining

import (
	"context"
	"fmt"
	"golang.org/x/sync/semaphore"
	"sync"
)

type Dining struct {
	n int
	wg *sync.WaitGroup
	forkMutex []sync.Mutex
	peopleSem *semaphore.Weighted
	ctx context.Context
}

func (d *Dining) WantsToEat(person int) {
	d.wg.Add(1)
	go func() {
		defer d.wg.Done()
		d.peopleSem.Acquire(d.ctx, 1)
		d.pickLeftFork(person)
		d.pickRightFork(person)
		d.eat(person)
		d.putLeftFork(person)
		d.putRightFork(person)
		d.peopleSem.Release(1)
	}()
}

func (d *Dining) pickLeftFork(person int){
	leftFork := person
	d.forkMutex[leftFork].Lock()
	fmt.Printf("[%d,%d,%d]", person, 1, 1)
}

func (d *Dining) pickRightFork(person int){
	rightFork := (person+d.n-1)%d.n
	d.forkMutex[rightFork].Lock()
	fmt.Printf("[%d,%d,%d]", person, 2, 1)
}

func (d *Dining) eat(person int){
	fmt.Printf("[%d,%d,%d]", person, 0, 3)
}

func (d *Dining) putLeftFork(person int){
	leftFork := person
	d.forkMutex[leftFork].Unlock()
	fmt.Printf("[%d,%d,%d]", person, 1, 2)
}

func (d *Dining) putRightFork(person int){
	rightFork := (person+4)%5
	d.forkMutex[rightFork].Unlock()
	fmt.Printf("[%d,%d,%d]", person, 2, 2)
}

func (d *Dining) init(){
	d.wg = &sync.WaitGroup{}
	d.ctx = context.Background()
	d.peopleSem = semaphore.NewWeighted(int64(d.n-1))
	d.forkMutex = make([]sync.Mutex, d.n)
}

func Start(){
	d := Dining{n: 5}
	d.init()
	for i:=0; i<1; i++{
		for p:=0; p<d.n; p++{
			d.WantsToEat(p)
		}
	}
	d.wg.Wait()
}



