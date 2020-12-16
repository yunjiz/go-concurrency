package fizzbuzz

import (
	"fmt"
	"sync"
)

type FizzBuzz struct {
	n  int

	fizzChan chan int
	buzzChan chan int
	fizzbuzzChan chan int
	numberChan chan int
	closeChan chan int
	wg *sync.WaitGroup
}

func (fb *FizzBuzz) fizz(){
	fb.wg.Add(1)
	go func() {
		defer fb.wg.Done()
		for {
			select {
			case value := <- fb.fizzChan:
				fmt.Println("Fizz")
				fb.setNextChan(value)
			case <-fb.closeChan:
				return
			}
		}
	}()
}

func (fb *FizzBuzz) buzz(){
	fb.wg.Add(1)
	go func() {
		defer fb.wg.Done()
		for {
			select {
			case value := <- fb.buzzChan:
				fmt.Println("Buzz")
				fb.setNextChan(value)
			case <-fb.closeChan:
				return
			}
		}
	}()

}

func (fb *FizzBuzz) fizzbuzz(){
	fb.wg.Add(1)
	go func() {
		defer fb.wg.Done()
		for {
			select {
			case value := <- fb.fizzbuzzChan:
				fmt.Println("FizzBuzz")
				fb.setNextChan(value)
			case <-fb.closeChan:
				return
			}
		}
	}()
}

func (fb *FizzBuzz) number(){
	fb.wg.Add(1)
	go func() {
		defer fb.wg.Done()
		for {
			select {
			case value := <- fb.numberChan:
				fmt.Println(value)
				//Need use "go" to set back number chan for select again
				go fb.setNextChan(value)
			case <-fb.closeChan:
				return
			}
		}
	}()
}

func (fb *FizzBuzz) init(){
	go func() {
		fb.numberChan<-1
	}()
}

func (fb *FizzBuzz) setClose(){
	fb.closeChan<-1
	fb.closeChan<-1
	fb.closeChan<-1
	fb.closeChan<-1
}

func (fb *FizzBuzz) setNextChan(value int){
	value++
	if value > fb.n{
		fb.setClose()
	} else if value%3==0 && value%5==0{
		fb.fizzbuzzChan<-value
	} else if value%3 == 0{
		fb.fizzChan <- value
	} else if value%5 == 0{
		fb.buzzChan <- value
	} else {
		fb.numberChan <- value
	}
}

func Start(){
	fb := FizzBuzz{
		n:            21,
		fizzChan:     make(chan int),
		buzzChan:     make(chan int),
		fizzbuzzChan: make(chan int),
		numberChan:   make(chan int),
		closeChan:    make(chan int, 4),
		wg:           &sync.WaitGroup{},
	}

	fb.init()
	fb.fizz()
	fb.buzz()
	fb.fizzbuzz()
	fb.number()

	fb.wg.Wait()
}
