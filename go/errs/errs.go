package errs

import (
	"fmt"
	"runtime"
)

// Error. カスタムエラー。errorインターフェースを満たしている。
type Error struct {
	err        error
	location   string
	stacktrace string
}

func New(err error) *Error {
	location := "unknown"

	// MEMO: 呼び出し元のfile, lineを取得
	_, file, line, ok := runtime.Caller(1)
	if ok {
		location = fmt.Sprintf("%s:%d", file, line)
	}

	buf := make([]byte, 4096)
	n := runtime.Stack(buf, false)
	stacktrace := string(buf[:n])

	return &Error{
		err:        err,
		location:   location,
		stacktrace: stacktrace,
	}
}

func (e *Error) Error() string {
	return e.err.Error()
}

func (e *Error) Location() string {
	return e.location
}

func (e *Error) Stacktrace() string {
	return e.stacktrace
}
