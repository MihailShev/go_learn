package logger

import (
	"io/ioutil"
	"os"
	"strings"
	"testing"
	"time"
)

func TestHwAccepted_Log(t *testing.T) {
	expected := time.Now().Format(dateFormat) + " accepted 3456 4"
	got := HwAccepted{3456, 4}.Log()

	if got != expected {
		t.Error("\nExp:\n", expected, "\nGot:\n", got)
	}
}

func TestHwSubmitted_Log(t *testing.T) {
	expected := time.Now().Format(dateFormat) + " submitted 3456 \"please take a look at my homework\""
	got := HwSubmitted{3456, "code", "please take a look at my homework"}.Log()

	if got != expected {
		t.Error("\nExp:\n", expected, "\nGot:\n", got)
	}
}

func TestLogOtusEvent(t *testing.T) {
	fileName := "test_otus_event.log"
	file, err := os.Create(fileName)

	if err != nil {
		t.Error("Unable to create file:", err)
		os.Exit(1)
	}

	defer file.Close()

	submitted := HwSubmitted{3456, "code", "please take a look at my homework"}
	accepted := HwAccepted{3456, 4}

	err = LogOtusEvent(submitted, file)

	if err != nil {
		t.Error("Logging error caused by submit event", err)
	}

	err = LogOtusEvent(accepted, file)

	if err != nil {
		t.Error("Logging error caused by accept event", err)
	}

	data, err := ioutil.ReadFile(fileName)

	if err != nil {
		t.Error("Error reading file:", fileName, err)
	}

	logLines := strings.Split(string(data), lineBreak)

	submitLog := time.Now().Format(dateFormat) + " submitted 3456 \"please take a look at my homework\""
	acceptLog := time.Now().Format(dateFormat) + " accepted 3456 4"

	if logLines[0] != submitLog {
		t.Error("\nExp:\n", submitLog, "\nGot:\n", logLines[0])
	}

	if logLines[1] != acceptLog {
		t.Error("\nExp:\n", acceptLog, "\nGot:\n", logLines[1])
	}

	err = os.Remove(fileName)

	if err != nil {
		t.Error("Error removing file:", fileName, err)
	}
}
