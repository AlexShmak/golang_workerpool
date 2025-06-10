package main

import (
	"fmt"
	"time"

	"github.com/AlexShmak/golang_workerpool/workerpool"
)

func main() {
	pool := workerpool.NewWorkerPool(10)

	fmt.Println("--- Initializing with 2 workers ---")
	pool.StartWorkers(2)

	fmt.Println("\n--- Adding initial tasks ---")
	for i := 1; i <= 5; i++ {
		pool.AddTask(workerpool.Task{ID: i, Data: fmt.Sprintf("Initial data %d", i)})
	}

	time.Sleep(2 * time.Second)

	fmt.Println("\n--- Dynamically adding a worker ---")
	pool.AddWorker()

	fmt.Println("\n--- Adding more tasks ---")
	for i := 6; i <= 10; i++ {
		pool.AddTask(workerpool.Task{ID: i, Data: fmt.Sprintf("More data %d", i)})
	}

	time.Sleep(2 * time.Second)

	fmt.Println("\n--- Dynamically removing a worker ---")
	pool.RemoveWorker()

	fmt.Println("\n--- Adding final tasks ---")
	for i := 11; i <= 15; i++ {
		pool.AddTask(workerpool.Task{ID: i, Data: fmt.Sprintf("Final data %d", i)})
	}

	time.Sleep(3 * time.Second)

	fmt.Println("\n--- Stopping all workers ---")
	pool.StopAllWorkers()

	fmt.Println("--- Program Finished ---")
}

