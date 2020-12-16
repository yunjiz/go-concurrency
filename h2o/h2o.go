package h2o

import (
	"context"
	"fmt"
	"golang.org/x/sync/semaphore"
	"sync"
)

type H2O struct {
	n int
	wg *sync.WaitGroup

	semH *semaphore.Weighted
	semO *semaphore.Weighted
	ctx context.Context
}

func (h *H2O) hydrogen(){
	for i:=0; i<h.n; i++{
		h.wg.Add(1)
		go h.produceH()
		h.wg.Add(1)
		go h.produceH()
	}
}

func (h *H2O) oxygen(){
	for i:=0; i<h.n; i++{
		h.wg.Add(1)
		go h.produceO()
	}
}

func (h *H2O) produceH(){
	defer h.wg.Done()

	h.semH.Acquire(h.ctx, 1)
	fmt.Print("H")
	h.semO.Release(1)
}

func (h *H2O) produceO(){
	defer h.wg.Done()
	h.semO.Acquire(h.ctx, 2)
	fmt.Print("O")
	h.semH.Release(2)
}

func (h *H2O) init(){
	h.semO.Acquire(h.ctx, 2)
}

func Start(){
	h := H2O{
		n:    10,
		wg:   &sync.WaitGroup{},
		semH: semaphore.NewWeighted(2),
		semO: semaphore.NewWeighted(2),
		ctx: context.Background(),
	}

	h.init()
	h.hydrogen()
	h.oxygen()

	h.wg.Wait()
}
