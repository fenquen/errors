# errors
personal modification of the project https://github.com/pkg/errors

### Usage
```go
tcpAddr, _ := net.ResolveTCPAddr("tcp", "192.168.0.1:8000")
	_, err := net.ListenTCP("tcp", tcpAddr)
	if nil != err {
		fmt.Print(NewStackTraceableError("net.ListenTCP fails", err))
		return
	}
	
	
```

result
```go
net.ListenTCP fails
errors_.TestA
	/Users/a/github/errors/errors_test.go:13
testing.tRunner
	/usr/local/go/src/testing/testing.go:865
runtime.goexit
	/usr/local/go/src/runtime/asm_amd64.s:1337
listen tcp 192.168.0.1:8000: bind: can't assign requested address--- PASS: TestA (0.00s)
```