package executor

type Task = func() error

func Executor(taskList []Task, maxExecution, maxErr int) {
	tasksNumber := len(taskList)
	resultCount := 0

	limiter := make(chan struct{}, maxExecution)
	resultCh := make(chan interface{}, tasksNumber)
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
			close(resultCh)
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
