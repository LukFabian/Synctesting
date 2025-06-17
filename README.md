Run the TestAfterFunc example test (this assumes the shell is currently in the project directory):
```shell
export GOEXPERIMENT=synctest
export GODEBUG=asynctimerchan=0
go test -run TestAfterFunc
```

