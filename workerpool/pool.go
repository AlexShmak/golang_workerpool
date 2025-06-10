package workerpool

import (
	"fmt"
	"sync"
)

type WorkerPool struct {
	taskChan     chan Task
	workers      []*Worker
	nextWorkerID int
	wg           sync.WaitGroup
	mu           sync.Mutex
}

func NewWorkerPool(bufferSize int) *WorkerPool {
	return &WorkerPool{
		taskChan: make(chan Task, bufferSize),
		workers:  make([]*Worker, 0),
	}
}

func (wp *WorkerPool) AddTask(task Task) {
	wp.taskChan <- task
}

func (wp *WorkerPool) AddWorker() {
	wp.mu.Lock()
	wp.nextWorkerID++
	worker := NewWorker(wp.nextWorkerID, wp.taskChan, &wp.wg)
	wp.workers = append(wp.workers, worker)
	wp.mu.Unlock()
	worker.Start()
	fmt.Printf("Pool: Added Worker %d (Total: %d)\n", worker.ID, len(wp.workers))
}

func (wp *WorkerPool) RemoveWorker() {
	wp.mu.Lock()
	if len(wp.workers) == 0 {
		wp.mu.Unlock()
		fmt.Println("Pool: No workers to remove")
		return
	}
	worker := wp.workers[len(wp.workers)-1]
	wp.workers = wp.workers[:len(wp.workers)-1]
	wp.mu.Unlock()
	worker.Stop()
	fmt.Printf("Pool: Removed Worker %d (Total: %d)\n", worker.ID, len(wp.workers))
}

func (wp *WorkerPool) StartWorkers(count int) {
	for i := 0; i < count; i++ {
		wp.AddWorker()
	}
}

func (wp *WorkerPool) StopAllWorkers() {
	wp.mu.Lock()
	for _, w := range wp.workers {
		w.Stop()
	}
	wp.workers = nil
	wp.mu.Unlock()
	wp.wg.Wait()
	close(wp.taskChan)
	fmt.Println("Pool: All workers stopped")
}

