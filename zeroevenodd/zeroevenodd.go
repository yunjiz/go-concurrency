package zeroevenodd

import (
	"fmt"
	"sync"
)

type ZeroEvenOdd struct {
	n int
	wg *sync.WaitGroup
}

func (zeo *ZeroEvenOdd) zero() {
	zeo.wg.Add(1)
	go func() {
		defer zeo.wg.Done()
		for i:=0; i<zeo.n; i++{
			fmt.Print(0)
		}

	}()
}

func (zeo *ZeroEvenOdd) even(){
	zeo.wg.Add(1)
	go func() {
		defer zeo.wg.Done()
		count := zeo.n/2
		number := 2
		for i:= 0; i<count; i++ {
			fmt.Print(number)
			number += 2
		}
	}()
}

func (zeo *ZeroEvenOdd) odd(){
	zeo.wg.Add(1)
	go func() {
		defer zeo.wg.Done()
		count := zeo.n/2
		if zeo.n %2 == 1{
			count++
		}
		number := 1

		for i:= 0; i<count; i++ {
			fmt.Print(number)
			number += 2
		}
	}()
}

func Start(){
	zeo := ZeroEvenOdd{
		n:  15,
		wg: &sync.WaitGroup{},
	}

	zeo.zero()
	zeo.even()
	zeo.odd()

	zeo.wg.Wait()
}