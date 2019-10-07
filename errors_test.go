package errors_

import (
	"fmt"
	"net"
	"testing"
)

func TestA(t *testing.T) {
	tcpAddr, _ := net.ResolveTCPAddr("tcp", "192.168.0.1:8000")
	_, err := net.ListenTCP("tcp", tcpAddr)
	if nil != err {
		fmt.Print(NewStackTraceableError("net.ListenTCP fails", err))
		return
	}

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
