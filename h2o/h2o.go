package h2o

import (
	"fmt"
	"sync"
)

type H2O struct {
	n int
	wg *sync.WaitGroup
}

func (h *H2O) hydrogen(){
	for i:=0; i<h.n; i++{
		h.wg.Add(1)
		go produceH(h.wg)
		h.wg.Add(1)
		go produceH(h.wg)
	}
}

func (h *H2O) oxygen(){
	for i:=0; i<h.n; i++{
		h.wg.Add(1)
		go produceO(h.wg)
	}
}

func produceH(wg *sync.WaitGroup){
	defer wg.Done()
	fmt.Print("H")
}

func produceO(wg *sync.WaitGroup){
	defer wg.Done()
	fmt.Print("O")
}

func Start(){
	h := H2O{
		n:  10,
		wg: &sync.WaitGroup{},
	}

	h.hydrogen()
	h.oxygen()

	h.wg.Wait()
}
