package errors

import (
	"fmt"
	"io"
	"path"
	"runtime"
	"strconv"
	"strings"
)

// Frame represents a program counter inside a stack frame.
// For historical reasons if Frame is interpreted as a uintptr
// its value represents the program counter + 1.
type Frame uintptr

// pc returns the program counter for this frame;
// multiple frames may have the same PC value.
// 程序计数器是用于存放下一条指令所在单元的地址的地点https://www.zhihu.com/question/22609253
func (frame Frame) pc() uintptr { return uintptr(frame) - 1 }

// targetFilePath() returns the full path to the targetFilePath that contains the function for this Frame's pc.
func (frame Frame) targetFilePath() string {
	fn := runtime.FuncForPC(frame.pc())
	if fn == nil {
		return "unknown"
	}
	file, _ := fn.FileLine(frame.pc())
	return file
}

// lineNum() returns the lineNum number of source code of the function for this Frame's pc.
func (frame Frame) lineNum() int {
	fn := runtime.FuncForPC(frame.pc())
	if fn == nil {
		return 0
	}
	_, line := fn.FileLine(frame.pc())
	return line
}

// funcName() returns the funcName of this function, if known.
func (frame Frame) funcName() string {
	fn := runtime.FuncForPC(frame.pc())
	if fn == nil {
		return "unknown"
	}
	return fn.Name()
}

// funcSimpleName removes the path prefix component of a function's funcName reported by func.Name().
func funcSimpleName(funcName string) string {
	i := strings.LastIndex(funcName, "/")
	funcName = funcName[i+1:]
	i = strings.Index(funcName, ".")
	return funcName[i+1:]
}

// Format formats the frame according to the fmt.Formatter interface.
//
//    %s    source targetFilePath
//    %d    source lineNum
//    %n    function funcName
//    %v    equivalent to %s:%d
//
// Format accepts flags that alter the printing of some verbs, as follows:
//
//    %+s   function funcName and path of source targetFilePath relative to the compile time
//          GOPATH separated by \n\t (<funcSimpleName>\n\t<path>)
//    %+v   equivalent to %+s:%d
func (frame Frame) Format(fmtState fmt.State, verb rune) {
	switch verb {
	case 's':
		switch {
		case fmtState.Flag('+'):
			_, _ = io.WriteString(fmtState, frame.funcName())
			_, _ = io.WriteString(fmtState, "\n\t")
			_, _ = io.WriteString(fmtState, frame.targetFilePath())
		default:
			_, _ = io.WriteString(fmtState, path.Base(frame.targetFilePath()))
		}
	case 'd':
		_, _ = io.WriteString(fmtState, strconv.Itoa(frame.lineNum()))
	case 'n':
		_, _ = io.WriteString(fmtState, funcSimpleName(frame.funcName()))
	case 'v':// fmt.Println()
		frame.Format(fmtState, 's')
		_, _ = io.WriteString(fmtState, ":")
		frame.Format(fmtState, 'd')
	}
}

// MarshalText formats a stacktrace Frame as a text string. The output is the
// same as that of fmt.Sprintf("%+v", f), but without newlines or tabs.
func (frame Frame) MarshalText() ([]byte, error) {
	name := frame.funcName()
	if name == "unknown" {
		return []byte(name), nil
	}
	return []byte(fmt.Sprintf("%s %s:%d", name, frame.targetFilePath(), frame.lineNum())), nil
}


