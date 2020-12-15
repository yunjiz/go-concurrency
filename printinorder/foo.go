package printinorder

import (
	"fmt"
	"sync"
)

type Foo struct {
	wg *sync.WaitGroup
	firstChan chan int
	secondChan chan int
}

func (f *Foo) first(printFunc func()) {
	f.wg.Add(1)
	go func() {
		defer f.wg.Done()

		printFunc()

		f.firstChan <- 1
	}()
}

func (f *Foo) second(printFunc func()) {
	f.wg.Add(1)
	go func() {
		defer f.wg.Done()

		<- f.firstChan

		printFunc()

		f.secondChan <- 1
	}()
}

func (f *Foo) third(printFunc func()) {
	f.wg.Add(1)
	go func() {
		defer f.wg.Done()

		<- f.secondChan

		printFunc()
	}()
}

func Start(){
	f := Foo{
		wg:         &sync.WaitGroup{},
		firstChan:  make(chan int),
		secondChan: make(chan int),
	}

	f.first(func() {
		fmt.Println("first")
	})
	f.second(func() {
		fmt.Println("second")
	})
	f.third(func() {
		fmt.Println("third")
	})

	f.wg.Wait()
}
