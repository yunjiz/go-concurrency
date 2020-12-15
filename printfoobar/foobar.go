package printfoobar

import (
	"context"
	"fmt"
	"golang.org/x/sync/semaphore"
	"sync"
)

type FooBar struct {
	n int
	wg *sync.WaitGroup

	//fooChan chan int
	//barChan chan int
	semFoo *semaphore.Weighted
	semBar *semaphore.Weighted
	ctx context.Context
}

func (fb *FooBar) foo(printFunc func()){
	for i:=0; i<fb.n; i++{
		fb.wg.Add(1)
		go func() {
			defer fb.wg.Done()

			fb.semFoo.Acquire(fb.ctx, 1)
			//<- fb.barChan
			printFunc()
			//fb.fooChan <- 1
			fb.semBar.Release(1)
		}()
	}
}

func (fb *FooBar) bar(printFunc func()){
	for i:=0; i<fb.n; i++{
		fb.wg.Add(1)
		go func() {
			defer fb.wg.Done()
			fb.semBar.Acquire(fb.ctx, 1)
			//<- fb.fooChan
			printFunc()
			//fb.barChan <- 1
			fb.semFoo.Release(1)
		}()
	}
}

func (fb *FooBar) init(){
	fb.semBar.Acquire(fb.ctx,1)
}

func Start(){
	fb := FooBar{
		n:       20,
		wg:      &sync.WaitGroup{},
		//fooChan: make(chan int),
		//barChan: make(chan int),
		semFoo:  semaphore.NewWeighted(1),
		semBar:  semaphore.NewWeighted(1),
		ctx: context.Background(),
	}

	fb.init()

	fb.foo(func() {
		fmt.Print("foo")
	})
	fb.bar(func(){
		fmt.Print("bar")
	})

	fb.wg.Wait()
}
