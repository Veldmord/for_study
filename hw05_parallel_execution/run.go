package hw05parallelexecution

import (
	"errors"
	"sync"
	"sync/atomic"
)

var ErrErrorsLimitExceeded = errors.New("errors limit exceeded")

type Task func() error

var syncOnce sync.Once

// Run starts tasks in n goroutines and stops its work when receiving m errors from tasks.

func Run(tasks []Task, n, m int) error {
	wg := sync.WaitGroup{}
	chTask := make(chan Task)
	//chError := make(chan error)
	chStop := make(chan interface{})
	var countError int32

	for i := 0; i < n; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for {
				select {
				case t, ok := <-chTask:
					if !ok {
						return
					}
					if err := t(); err != nil {
						atomic.AddInt32(&countError, 1)
						if countError >= int32(m) {
							syncOnce.Do(func() { close(chStop) })
							return
						}
					}
				case <-chStop:
					return
				}
			}
		}()
	}

	go func() {
		defer close(chTask)
		for _, task := range tasks {
			select {
			case chTask <- task:
			case <-chStop:
				return
			}
		}
	}()

	// go func() {
	// 	defer close(chError)
	// 	for err := range chError {
	// 		countError++
	// 		if countError >= int32(m) {
	// 			syncOnce.Do(func() { close(chStop) })
	// 		}
	// 		_ = err
	// 	}
	// }()

	wg.Wait()

	if countError >= int32(m) {
		return ErrErrorsLimitExceeded
	} else {
		return nil
	}
}
