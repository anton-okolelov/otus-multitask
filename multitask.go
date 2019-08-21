package multitask

import (
	"fmt"
	"sync"
)

type ErrCounter struct {
	count int
	mutex sync.Mutex
}

func Run(tasks []func() error, tasksInParallelCount int, maxErrorsCount int) {

	tasksChan := make(chan func() error, 100)

	errCounter := ErrCounter{}

	var wg sync.WaitGroup
	wg.Add(tasksInParallelCount)

	// start workers
	for i := 1; i <= tasksInParallelCount; i++ {
		go runner(i, tasksChan, &errCounter, maxErrorsCount, &wg)
	}

	go func() {
		defer close(tasksChan)
		for _, task := range tasks {
			if errCounter.count >= maxErrorsCount {
				return
			}
			tasksChan <- task
		}
	}()

	wg.Wait()

}

func runner(id int, tasksChan chan func() error, errCounter *ErrCounter, maxErrors int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("%d) runner start....\n", id)
	for {

		if errCounter.count >= maxErrors {
			return
		}

		task, ok := <-tasksChan
		if ok {
			fmt.Printf("%d) run task...\n", id)
			err := task()
			if err != nil {
				fmt.Printf("%d) Error!\n", id)
				errCounter.mutex.Lock()
				errCounter.count++
				errCounter.mutex.Unlock()
				return
			} else {
				fmt.Printf("%d) OK!\n", id)
			}
		} else {
			return
		}

	}
}
