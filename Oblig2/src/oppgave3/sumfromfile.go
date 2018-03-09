package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
)

type Lines struct {
	data  []byte
	index []int // line start, end pairs for data[start:end]
}

func NewLines(data []byte, nLines int) *Lines {
	bom := []byte{0xEF, 0xBB, 0xBF}
	if bytes.HasPrefix(data, bom) {
		data = data[len(bom):]
	}
	lines := Lines{data: data, index: make([]int, 0, 2*nLines)}
	for i := 0; ; {
		j := bytes.IndexByte(lines.data[i:], '\n')
		if j < 0 {
			if len(lines.data[i:]) > 0 {
				lines.index = append(lines.index, i)
				lines.index = append(lines.index, len(lines.data))
			}
			break
		}
		lines.index = append(lines.index, i)
		j += i
		i = j + 1
		if j > 0 && lines.data[j-1] == '\r' {
			j--
		}
		lines.index = append(lines.index, j)
	}
	if len(lines.index) != cap(lines.index) {
		lines.index = append([]int(nil), lines.index...)
	}
	return &lines
}

func (l *Lines) N() int {
	return len(l.index) / 2
}

func (l *Lines) At(n int) (string, error) {
	if 1 > n || n > l.N() {
		err := fmt.Errorf(
			"data has %d lines: at %d out of range",
			l.N(), n,
		)
		return "", err
	}
	m := 2 * (n - 1)
	return string(l.data[l.index[m]:l.index[m+1]]), nil
}

var (
	fName  = `test.txt`
	nLines = 2
)

func sumValues() {

	fName  := os.Args[1]
	data, err := ioutil.ReadFile(fName)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}

	lines := NewLines(data, nLines)

	for _, at := range []int{1} {
		line1, err := lines.At(at)
		if err != nil {
			fmt.Fprintf(os.Stderr, "%d\t%v\n", at, err)
			continue
		}

		for _, at := range []int{2} {
			line2, err := lines.At(at)
			if err != nil {
				fmt.Fprintf(os.Stderr, "%d\t%v\n", at, err)
				continue
			}
			l1, err := strconv.Atoi(line1)
			if err == nil {
				fmt.Println(l1)
			}

			l2, err := strconv.Atoi(line2)
			if err == nil {
				fmt.Println(l2)
			}

			file, err := os.OpenFile(fName, os.O_WRONLY|os.O_APPEND, 0777)
			if err != nil {
				panic(err)
			}
			defer file.Close()
			sumUp := l1 + l2
			fmt.Println(sumUp)
			bv, err := file.WriteString(fmt.Sprintf("\n%d\n", sumUp))
			fmt.Println(bv)
			return
		}
		}
	}

func main () {
	sumValues()
}
