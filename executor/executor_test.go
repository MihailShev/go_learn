package executor

import (
	"errors"
	"sync"
	"testing"
	"time"
)

var mu sync.Mutex

// Global counters
var executionCount = 0
var totalExecution = 0
var maxParallelExecution = 0
var errorCount = 0

var testCase1 = []Task{
	taskGenerator(1, false),
	taskGenerator(3, false),
	taskGenerator(2, false),
	taskGenerator(5, false),
	taskGenerator(4, false),
	taskGenerator(2, false),
	taskGenerator(1, false),
}

var testCase2 = []Task{
	taskGenerator(3, true),
	taskGenerator(1, true),
	taskGenerator(2, false),
	taskGenerator(1, false),
	taskGenerator(2, true),
	taskGenerator(2, false),
	taskGenerator(4, true),
}

// Test max parallel execution without error
func TestExecutor(t *testing.T) {
	maxExecution := 3
	expecExecution := 7
	Executor(testCase1, maxExecution, 1)

	if maxExecution != maxParallelExecution {
		t.Error("Failed: max parallel execution", maxParallelExecution, "got", maxExecution)
	} else {
		t.Log("Successful: max parallel execution:", maxExecution, "got", maxParallelExecution)
	}

	if expecExecution != totalExecution {
		t.Error("Failed: total execution:", totalExecution, "got", expecExecution)
	} else {
		t.Log("Successful: expected total execution: ", totalExecution, "got", expecExecution)
	}

	resetGlobalCounters()
}

// Test parallel execution with error
func TestExecutor2(t *testing.T) {
	maxExecution := 3
	expectErrorCount := 2

	Executor(testCase2, maxExecution, expectErrorCount)

	if expectErrorCount != errorCount {
		t.Error("Filed: error count", errorCount, "expected", expectErrorCount)
	} else {
		t.Log("Successful: error count", errorCount, "expected", expectErrorCount)
	}

	resetGlobalCounters()
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

			mu.Lock()
			errorCount++
			mu.Unlock()
		}

		mu.Lock()
		executionCount--
		mu.Unlock()

		return err
	}
}

func resetGlobalCounters() {
	totalExecution = 0
	maxParallelExecution = 0
	executionCount = 0
	errorCount = 0
}
