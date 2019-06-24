package logger

import (
	"fmt"
	"io"
	"time"
)

const (
	dateFormat = "2006.01.02"
	lineBreak  = "\n"
)

type OtusEvent interface {
	Log() string
}

type HwAccepted struct {
	ID    int
	Grade int
}

func (a HwAccepted) Log() string {
	dt := time.Now()

	return fmt.Sprint(dt.Format(dateFormat), " accepted ", a.ID, " ", a.Grade)
}

type HwSubmitted struct {
	ID      int
	Code    string
	Comment string
}

func (s HwSubmitted) Log() string {
	dt := time.Now()

	return fmt.Sprint(dt.Format(dateFormat), " submitted ", s.ID, " \"", s.Comment, "\"")
}

func LogOtusEvent(e OtusEvent, w io.Writer) error {
	_, err := w.Write([]byte(e.Log() + lineBreak))

	return err
}
