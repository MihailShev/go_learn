package executor

type Task = func() error

func Executor(taskList []Task, maxExecution, maxErr int) {
	tasksNumber := len(taskList)
	resultCount := 0

	// limit parallel execution
	limiter := make(chan struct{}, maxExecution)
	// all task result
	resultCh := make(chan interface{}, tasksNumber)
	// only error task result
	errCh := make(chan error, maxErr)

	for _, v := range taskList {
		go func(t Task) {
			limiter <- struct{}{}

			if len(errCh) < maxErr {
				executeTask(t, resultCh, errCh)
			}

			<-limiter
		}(v)
	}

	for range resultCh {
		resultCount++

		if resultCount == tasksNumber || len(errCh) == maxErr {
			break
		}
	}
}

func executeTask(t Task, resultCh chan<- interface{}, errCh chan<- error) {
	err := t()
	resultCh <- err

	if err != nil {
		errCh <- err
	}
}
