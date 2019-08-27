package multitask

import (
	"fmt"
	"sync"
	"sync/atomic"
)

type ErrCounter struct {
	count int
	mutex sync.Mutex
}

func Run(tasks []func() error, tasksInParallelCount int, maxErrorsCount int) {

	tasksChan := make(chan func() error, 100)

	var errCounter int32

	var wg sync.WaitGroup
	wg.Add(tasksInParallelCount)

	// start workers
	for i := 1; i <= tasksInParallelCount; i++ {
		go runner(i, tasksChan, &errCounter, maxErrorsCount, &wg)
	}

	go func() {
		defer close(tasksChan)
		for _, task := range tasks {
			if errCounter >= int32(maxErrorsCount) {
				return
			}
			tasksChan <- task
		}
	}()

	wg.Wait()

}

func runner(id int, tasksChan chan func() error, errCounter *int32, maxErrors int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("%d) runner start....\n", id)
	for task := range tasksChan {

		if *errCounter >= int32(maxErrors) {
			return
		}

		fmt.Printf("%d) run task...\n", id)
		err := task()

		if err != nil {
			fmt.Printf("%d) Error!\n", id)
			atomic.AddInt32(errCounter, 1)
		} else {
			fmt.Printf("%d) OK!\n", id)
		}

	}
}
