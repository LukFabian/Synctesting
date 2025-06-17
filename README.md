This project includes tests that use Go’s testing/synctest package to validate concurrent behavior in a deterministic, race-free manner. These tests are especially useful when verifying interactions between goroutines and ensuring no unintended scheduling issues occur.
Prerequisites

Ensure you are using Go 1.22 or later, which introduces the synctest package and synthetic scheduling.

Check your version with:
```shell
go version
```


This example demonstrates how to test context.AfterFunc deterministically using synctest.

Run it like this (assuming your shell is in the project directory):
```shell
export GOEXPERIMENT=synctest
export GODEBUG=asynctimerchan=0
go test -run TestAfterFunc
```

You can also use the go race condition checker:
```shell
go test -race -run TestHandleEchoAck
```

These environment variables are required because:

```GOEXPERIMENT=synctest``` enables the synthetic scheduler.

```GODEBUG=asynctimerchan=0``` disables asynchronous timers that would interfere with deterministic execution.

Further reading:

[Go dev blog](https://go.dev/blog/synctest)
[synctest documentation](https://pkg.go.dev/testing/synctest)
