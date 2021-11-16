# PHP Interpreter written in Go

```go
src := []byte(`<? echo "Hello world";`)
err := interpreter.Run(src, "7.4", os.Getenv("DEBUG") != "")
if err != nil {
	panic(err)
}
```
