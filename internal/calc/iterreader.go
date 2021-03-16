package calc

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"log"
	"strconv"
)

type iterableReader struct {
	scanner *bufio.Scanner
}

func newIterableReader(r io.Reader) *iterableReader {
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanLines)

	return &iterableReader{
		scanner: scanner,
	}
}

func (r *iterableReader) Next() (value int, ok bool) {
	n, err := r.readInt()

	if errors.Is(err, io.EOF) {
		return n, false
	} else if err != nil {
		log.Panicf("scan fail: %v", err)
	}

	return n, true
}

func (r *iterableReader) readInt() (int, error) {
	cont := r.scanner.Scan()

	if err := r.scanner.Err(); err != nil {
		return 0, fmt.Errorf("scanner: %w", err)
	}

	if r.scanner.Text() == "" {
		return 0, io.EOF
	}

	number, err := strconv.Atoi(r.scanner.Text())
	if err != nil {
		return 0, fmt.Errorf("not an integer: %w", err)
	}

	if !cont {
		return number, io.EOF
	}

	return number, nil
}
