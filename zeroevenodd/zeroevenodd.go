package zeroevenodd

import (
	"context"
	"fmt"
	"golang.org/x/sync/semaphore"
	"sync"
)

type ZeroEvenOdd struct {
	n int
	wg *sync.WaitGroup
	sm0 *semaphore.Weighted
	sm1 *semaphore.Weighted
	sm2 *semaphore.Weighted
	ctx context.Context
}

func (zeo *ZeroEvenOdd) zero() {
	zeo.wg.Add(1)
	go func() {
		defer zeo.wg.Done()
		for i:=0; i<zeo.n; i++{
			zeo.sm0.Acquire(zeo.ctx, 1)
			fmt.Print(0)

			if i & 1 == 1{
				zeo.sm2.Release(1)
			} else {
				zeo.sm1.Release(1)
			}
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
			zeo.sm2.Acquire(zeo.ctx, 1)
			fmt.Print(number)
			number += 2
			zeo.sm0.Release(1)
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
			zeo.sm1.Acquire(zeo.ctx, 1)
			fmt.Print(number)
			number += 2
			zeo.sm0.Release(1)
		}
	}()
}

func (zeo *ZeroEvenOdd) init(){
	zeo.sm1.Acquire(zeo.ctx,1)
	zeo.sm2.Acquire(zeo.ctx, 1)
}

func Start(){
	zeo := ZeroEvenOdd{
		n:   15,
		wg:  &sync.WaitGroup{},
		sm0: semaphore.NewWeighted(1),
		sm1: semaphore.NewWeighted(1),
		sm2: semaphore.NewWeighted(1),
		ctx: context.Background(),
	}

	zeo.init()
	zeo.zero()
	zeo.even()
	zeo.odd()

	zeo.wg.Wait()
}