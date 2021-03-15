package main

import (
	"fmt"
	"os"

	"github.com/alexandear/scalc/cmd/scalc"
)

func main() {
	res, err := scalc.Execute(os.Args)
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "%v\n%v", err, res)
		os.Exit(1)
	}

	println(res)
}
