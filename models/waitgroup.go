package models

import (
	"fmt"
	"runtime/debug"
	"sync"
)

type WaitGroupWrapper struct {
	sync.WaitGroup
}

func (w *WaitGroupWrapper) Wrap(cb func()) {
	w.Add(1)
	go func() {
		defer func() {
			if err := recover(); err != nil {
				fmt.Printf("WaitGroupWrapper: %v", err)
				debug.PrintStack()
			}
		}()
		cb()
		w.Done()
	}()
}
