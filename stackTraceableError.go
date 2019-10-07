package errors_

import (
	"fmt"
	"io"
)

type stackTraceableError struct {
	msg   string
	cause error
	*stack
}

func (s *stackTraceableError) Error() string { return s.msg }

func (s *stackTraceableError) Cause() error { return s.cause }

func (s *stackTraceableError) Format(fmtState fmt.State, verb rune) {
	switch verb {
	case 'v': // %v	打印值的默认格式。当打印结构体时，“加号”标记（%+v）会添加字段名
		//if fmtState.Flag('+') {
		_, _ = io.WriteString(fmtState, s.msg)
		s.stack.Format(fmtState, verb)

		if nil == s.cause {
			return
		}
		_, _ = fmt.Fprintf(fmtState, "\n%+v", s.cause)

		//}
		//fallthrough
	case 's':
		_, _ = io.WriteString(fmtState, s.Error())
	case 'q':
		_, _ = fmt.Fprintf(fmtState, "%q", s.msg)
	}

}
