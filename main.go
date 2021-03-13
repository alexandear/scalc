package main

import (
	"log"
	"os"
)

func main() {
	expr, err := Parse(`[ LE 3 a.txt b.txt c.txt [ GR 1 d.txt e.txt ] [ EQ 2 f.txt [ LE 1 g.txt h.txt ] ] ]`)
	if err != nil {
		log.Panic(err)
	}

	log.Printf("%+v", expr)

	fa, err := os.Open("test/a.txt")
	if err != nil {
		log.Panicf("open file: %v", err)
	}

	defer func() {
		if errClose := fa.Close(); errClose != nil {
			log.Printf("failed to close file: %v", errClose)
		}
	}()

	intsA, err := ReadInts(fa)
	if err != nil {
		log.Panic(err)
	}

	log.Println(intsA)
}
