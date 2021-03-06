package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"sync"
)

type Counter struct {
	WaitGroupN

	sync.RWMutex
	total int64
}

func NewCounter(k int64) (c *Counter) {
	c = new(Counter)
	c.GroupMax(k)
	return
}

func (c *Counter) Run(address, sep string) {
	if len(address) == 0 {
		return
	}
	_, err := url.Parse(address)
	if err != nil {
		fmt.Printf("Count for %s: 0 (error url address)\r\n", address)
		return
	}
	c.Add()

	go c.run(address, sep)
}

func (c *Counter) run(address, sep string) {
	defer c.Done()

	resp, err := http.Get(address)
	if err != nil {
		fmt.Printf("Count for %s: 0 (error request), err: %s\r\n", address, err.Error())
		return
	}
	defer resp.Body.Close()

	st, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Count for %s: 0 (error response), err: %s\r\n", address, err.Error())
		return
	}

	count := strings.Count(string(st), sep)
	fmt.Printf("Count for %s: %d\r\n", address, count)
	c.AddTotal(int64(count))
}

func (c *Counter) AddTotal(count int64) {
	c.Lock()
	defer c.Unlock()
	c.total += count
}

func (c *Counter) GetTotal() int64 {
	c.RLock()
	defer c.RUnlock()
	return c.total
}
