# gRPC Tutorial

> Simple gRPC service written in Go with steps to reproduce

## 1. Create "Hello World!" go application

```
go mod init
touch main.go
```

`main.go` contents:

```go
package main

import "fmt"

func main() {
	fmt.Println("Welcome to the gRPC Tutorial!")
}
```
