package main

import (
	"fmt"
	"os"
)

func main() {
	res, err := Execute(os.Args)
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "%v\n%v", err, res)
		os.Exit(1)
	}

	println(res)
}
