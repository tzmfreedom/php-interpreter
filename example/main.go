package main

import (
	"github.com/tzmfreedom/php-interpreter"
	"os"
)

func main() {
	src := []byte(`<? echo "Hello world";echo 2*3+4+6*7;`)
	err := interpreter.Run(src, "7.4", os.Getenv("DEBUG") != "")
	if err != nil {
		panic(err)
	}
}

