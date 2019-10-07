package errors

import (
	"errors"
	"fmt"
	"testing"
)

func TestA(t *testing.T) {
	fmt.Print(NewStackTraceableError("该error本身的msg",errors.New("cause的msg")))

}

func TestSwitch(t *testing.T) {
	a := 0
	switch a {
	case 0:
		fallthrough
	case 1:
		fmt.Println("a")
	}
}
