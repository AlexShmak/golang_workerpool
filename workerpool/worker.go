package workerpool

import (
	"fmt"
	"sync"
	"time"
)

type Worker struct {
	ID       int
	taskChan <-chan Task
	quitChan chan struct{}
	wg       *sync.WaitGroup
}

func NewWorker(id int, taskChan <-chan Task, wg *sync.WaitGroup) *Worker {
	return &Worker{
		ID:       id,
		taskChan: taskChan,
		quitChan: make(chan struct{}),
		wg:       wg,
	}
}

func (w *Worker) Start() {
	w.wg.Add(1)
	go func() {
		defer w.wg.Done()
		for {
			select {
			case task, ok := <-w.taskChan:
				if !ok {
					return
				}
				fmt.Printf("Worker %d: Processing Task %d: %s\n", w.ID, task.ID, task.Data)
				time.Sleep(100 * time.Millisecond)
			case <-w.quitChan:
				return
			}
		}
	}()
}

func (w *Worker) Stop() {
	close(w.quitChan)
}

