package fizzbuzz

import "sync"

type FizzBuzz struct {
	n  int

	wg *sync.WaitGroup
}

func (fb *FizzBuzz) fizz(){

}

func (fb *FizzBuzz) buzz(){

}

func (fb *FizzBuzz) fizzbuzz(){

}

func (fb *FizzBuzz) number(){

}

func (fb *FizzBuzz) init(){

}

func Start(){
	fb := FizzBuzz{
		n:  15,
		wg: &sync.WaitGroup{},
	}

	fb.wg.Wait()
}
