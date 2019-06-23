package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

const format = "2006.01.02"

// OtusEvent is the interface that describes event of otus education platform
type OtusEvent interface {
	Log() string
}

// HwAccepted is the struct that contains inforamtion about accepted event
type HwAccepted struct {
	ID    int
	Grade int
}

func (a HwAccepted) Log() string {
	dt := time.Now()

	return fmt.Sprintln(dt.Format(format), "acepted", a.ID, a.Grade)
}

type HwSubmitted struct {
	ID      int
	Code    string
	Comment string
}

func (s HwSubmitted) Log() string {
	dt := time.Now()

	return fmt.Sprintln(dt.Format(format), "submitted", s.ID, s.Comment)
}

func LogOtusEvent(e OtusEvent, w io.Writer) {
	w.Write([]byte(e.Log()))
}

func main() {
	accept := HwAccepted{3346, 4}
	submit := HwSubmitted{3346, "code", "please take a look at my homework"}
	LogOtusEvent(accept, os.Stdout)
	LogOtusEvent(submit, os.Stdout)
}
