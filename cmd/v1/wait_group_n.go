package main

import (
	"sync"
)

type WaitGroupN struct {
	wg sync.WaitGroup

	sync.Mutex
	k int64
	i int64

	ch []chan struct{}
}

func NewWaitGroupN(k int64) (wgn *WaitGroupN) {
	wgn = new(WaitGroupN)
	wgn.k = k
	return
}

func (wgn *WaitGroupN) GroupMax(k int64) {
	wgn.Lock()
	defer wgn.Unlock()

	wgn.k = k
}

func (wgn *WaitGroupN) add() (ch chan struct{}, ok bool) {
	wgn.Lock()
	defer wgn.Unlock()

	if wgn.i < wgn.k {
		wgn.i++
		ok = true
	} else {
		ch = make(chan struct{}, 1)
		wgn.ch = append(wgn.ch, ch)
	}
	return
}

func (wgn *WaitGroupN) Add() {
	for {
		ch, ok := wgn.add()
		if ok {
			wgn.wg.Add(1)
			return
		}
		<-ch
	}
}

func (wgn *WaitGroupN) Done() {
	var ch chan struct{}

	wgn.Lock()
	wgn.i--
	if len(wgn.ch) > 0 {
		ch = wgn.ch[0]
		wgn.ch = wgn.ch[1:]
	}
	wgn.Unlock()

	wgn.wg.Done()

	if ch != nil {
		ch <- struct{}{}
	}
}

func (wgn *WaitGroupN) Wait() {
	wgn.wg.Wait()
}
