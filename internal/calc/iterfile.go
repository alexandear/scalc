package calc

import (
	"fmt"
	"io"
	"os"

	"github.com/alexandear/scalc/pkg/scalc"
)

type FileIteratorImpl struct{}

func NewFileIterator() *FileIteratorImpl {
	return &FileIteratorImpl{}
}

func (f *FileIteratorImpl) Iterator(file string) (scalc.Iterator, io.Closer, error) {
	fi, err := os.Open(file)
	if err != nil {
		return nil, nil, fmt.Errorf("open file: %w", err)
	}

	return newIterableReader(fi), fi, nil
}
