package executor

import (
	"errors"
	"sync"
	"testing"
	"time"
)

var mu sync.Mutex
var executionCount = 0
var totalExecution = 0
var maxParallelExecution = 0

var testCase1 = []Task{
	taskGenerator(1, false),
	taskGenerator(3, false),
	taskGenerator(2, false),
	taskGenerator(5, false),
	taskGenerator(4, false),
	taskGenerator(2, false),
	taskGenerator(1, false),
}

func taskGenerator(t time.Duration, isError bool) func() error {
	return func() error {
		var err error

		mu.Lock()
		executionCount++
		totalExecution++

		if executionCount > maxParallelExecution {
			maxParallelExecution = executionCount
		}

		mu.Unlock()

		time.Sleep(time.Second * t)

		if isError {
			err = errors.New("error")
		}

		mu.Lock()
		executionCount--
		mu.Unlock()
		return err
	}
}

func TestExecutor(t *testing.T) {
	maxExecution := 3
	expecExecution := 7
	Executor(testCase1, maxExecution, 1)

	if maxExecution != maxParallelExecution {
		t.Error("failed max parallel execution:", maxParallelExecution, "expected", maxExecution)
	} else {
		t.Log("successful expected max parallel execution:", maxExecution, "got", maxParallelExecution)
	}

	if expecExecution != totalExecution {
		t.Error("failed total execution:", totalExecution, "expected", expecExecution)
	} else {
		t.Log("successful total execution: ", totalExecution)
	}
}
