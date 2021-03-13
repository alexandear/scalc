package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"strconv"
)

func ScanInt(scanner *bufio.Scanner) (int, error) {
	cont := scanner.Scan()

	if err := scanner.Err(); err != nil {
		return 0, fmt.Errorf("scanner: %w", err)
	}

	if scanner.Text() == "" {
		return 0, io.EOF
	}

	number, err := strconv.Atoi(scanner.Text())
	if err != nil {
		return 0, fmt.Errorf("not an integer: %w", err)
	}

	if !cont {
		return number, io.EOF
	}

	return number, nil
}

func ReadInts(r io.Reader) ([]int, error) {
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanLines)

	var res []int

	for {
		n, err := ScanInt(scanner)
		if errors.Is(err, io.EOF) {
			return res, nil
		} else if err != nil {
			return []int{}, fmt.Errorf("scan: %w", err)
		}

		res = append(res, n)
	}
}
