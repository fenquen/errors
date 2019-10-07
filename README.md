# errors
personal modification of the project https://github.com/pkg/errors

### Usage
```go
tcpListener, err := net.ListenTCP(config.TCP_MODE, localTcpSvrAddr)
	if nil != err {
		fmt.Print(NewStackTraceableError("net.ListenTCP fails",err))
		return
	}
```
