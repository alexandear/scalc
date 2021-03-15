package calc

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"log"
	"strconv"
)

type IterableReader struct {
	scanner *bufio.Scanner
}

func NewIterableReader(r io.Reader) *IterableReader {
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanLines)

	return &IterableReader{
		scanner: scanner,
	}
}

func (r *IterableReader) Next() (value int, ok bool) {
	n, err := r.readInt()

	if errors.Is(err, io.EOF) {
		return n, false
	} else if err != nil {
		log.Panicf("scan fail: %v", err)
	}

	return n, true
}

func (r *IterableReader) readInt() (int, error) {
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
