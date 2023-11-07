package common

import (
	"errors"
	"fmt"
	"sync"
	"time"
)

type PanicHandleFunc func(err any)

func NewGoroutineWithPanicHandler(f func(), panicHandlers ...PanicHandleFunc) {
	if len(panicHandlers) == 0 {
		panicHandlers = []PanicHandleFunc{
			func(err any) {
				//TODO 默认日志打印错误信息
				fmt.Println(err)
			},
		}
	}
	go func() {
		defer func() {
			if err := recover(); err != nil {
				for _, fn := range panicHandlers {
					fn(err)
				}
			}
		}()
		f()
	}()
}

func Retry(f func() error, times int32, wait int32) error {
	if times <= 0 {
		return errors.New("retry failed")
	}
	err := f()
	if err != nil {
		time.Sleep(time.Duration(wait) * time.Millisecond)
		times -= 1
		return Retry(f, times, wait)
	}
	return nil
}

type TimerWaitGroup struct {
	group *sync.WaitGroup
	ch    chan bool
}

func (t *TimerWaitGroup) Add(delta int) {
	t.group.Add(delta)
}

func (t *TimerWaitGroup) Done() {
	t.group.Done()
}

func (t *TimerWaitGroup) Wait(duration time.Duration) bool {
	go time.AfterFunc(duration, func() {
		t.ch <- true
	})
	go func() {
		t.group.Wait()
		t.ch <- false
	}()
	return <-t.ch
}

func NewTimerWaitGroup() *TimerWaitGroup {
	return &TimerWaitGroup{ch: make(chan bool), group: &sync.WaitGroup{}}
}
