package cron

import (
	"context"
	"fmt"
	"sync"
	"time"
)

type Task struct {
	F          func(context.Context)
	D          time.Duration
	processing bool
	mux        sync.RWMutex
}

func (t *Task) isProcessing() bool {
	t.mux.RLock()
	defer t.mux.RUnlock()
	return t.processing
}

func (t *Task) SetProcessing(b bool) {
	t.mux.Lock()
	defer t.mux.Unlock()
	t.processing = b
}

type Cron struct {
	tickerDuration time.Duration
	tasks          []*Task
	ctx            context.Context
	cancel         context.CancelFunc
}

func NewCron(tickerDuration time.Duration) *Cron {
	if tickerDuration == 0 {
		tickerDuration = time.Second
	}
	c := &Cron{
		tickerDuration: tickerDuration,
		tasks:          make([]*Task, 0),
	}

	c.ctx, c.cancel = context.WithCancel(context.TODO())
	return c
}

func (c *Cron) AddTask(t *Task) {
	c.tasks = append(c.tasks, t)
}

func (c *Cron) Run() {
	tiker := time.NewTicker(c.tickerDuration)
	sum := time.Duration(0)
	for {
		select {
		case <-tiker.C:
			sum += c.tickerDuration
			for _, task := range c.tasks {
				if sum%task.D == 0 {
					if task.isProcessing(){
						continue
					}
					task.SetProcessing(true)
					go func(t *Task) {
						defer func() {
							t.SetProcessing(false)
							if err := recover(); err != nil {
								fmt.Println("panic:", err)
							}
						}()
						t.F(c.ctx)
					}(task)
				}
			}
		case <-c.ctx.Done():
			return
		}
	}
}

func (c *Cron) Stop() {
	c.cancel()
}
