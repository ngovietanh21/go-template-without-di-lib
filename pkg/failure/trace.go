package failure

import (
	"fmt"
	"path/filepath"
	"runtime"
)

const maxTraceback = 10

func ErrorWithTrace(err error) error {
	return fmt.Errorf("%w \n at %s", err, trace())
}

func trace() string {
	pc := make([]uintptr, maxTraceback)
	n := runtime.Callers(3, pc)
	frames := runtime.CallersFrames(pc[:n])
	frame, _ := frames.Next()
	return fmt.Sprintf("\t%s:%d | %s", filepath.Base(frame.File), frame.Line, frame.Function)
}
