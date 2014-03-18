package iterator4go

import (
	"bufio"
	"bytes"
	"errors"
	"io"
	"os"
)

type Iterator struct {
	current int
	data    []interface{}
}

func NewIterator(data []interface{}) *Iterator {
	return &Iterator{data: data, current: -1}
}

func (it *Iterator) Value() interface{} {
	return it.data[it.current]
}
func (it *Iterator) Next() bool {
	if it.current+1 >= len(it.data) {
		return false
	}
	it.current++
	return true
}

type IntIterator struct {
	current int
	data    []int
}

func NewIntIterator(data []int) *IntIterator {
	return &IntIterator{data: data, current: -1}
}

func (it *IntIterator) Value() int {
	return it.data[it.current]
}
func (it *IntIterator) Next() bool {
	if it.current+1 >= len(it.data) {
		return false
	}
	it.current++
	return true
}

type TxtFileIterator struct {
	current int
	data    string
	reader  *bufio.Reader
	file    *os.File
}

var ErrReadFileFailed = errors.New("TxtFileIterator : ErrReadFileFailed")

func NewTxtFileIterator(filePath string) *TxtFileIterator {
	if fr, err := os.Open(filePath); err != nil {
		return nil
	} else {
		return &TxtFileIterator{data: "", reader: bufio.NewReader(fr), file: fr, current: 0}
	}
}

func (it *TxtFileIterator) Value() string {
	return it.data
}

//注意，此方法可能产生panic，分别是bytes.ErrTooLarge、iterator4go.ErrReadFileFailed
func (it *TxtFileIterator) Next() bool {
	var data []byte
	var isPrefix bool
	var err error

	buf := new(bytes.Buffer)

	for {

		if data, isPrefix, err = it.reader.ReadLine(); err == nil {
			buf.Write(data)
			if isPrefix == true {
				continue
			}
			it.data = string(buf.Bytes())
			it.current++
			return true
		} else {
			if err == io.EOF {
				it.data = ""
				return false
			} else {
				panic(ErrReadFileFailed)
			}
		}

	}

	if it.current+1 >= len(it.data) {
		return false
	}
	it.current++
	return true
}

func (it *TxtFileIterator) Close() {
	it.reader = nil
	if it.file != nil {
		it.file.Close()
		it.file = nil
	}
}
