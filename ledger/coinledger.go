package ledger

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func Open(name string) (*bufio.Reader, func() error, error) {
	f, err := os.Open(name)
	if err != nil {
		return nil, nil, err
	}
	r := bufio.NewReader(f)
	closeFn := func() error {
		err := f.Close()
		if err != nil {
			return err
		}
		return nil
	}
	return r, closeFn, nil
}

func Read(r *bufio.Reader) (index uint, line string) {
	var i uint = 0
	for {
		l, erx := r.ReadString('\n')

		if l != "" {
			// version-based parse and return index and line
			// if version is 0, continue
			i++
			return i, l
		}

		switch erx {
		case io.EOF:
			return 0, "EOF"
		case nil:
			continue
		default:
			return 0, ""
		}
	}
}

func parseCoin(line string)

func ReadAll(file string) {
	f, err := os.Open(file)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	r := bufio.NewReader(f)
	for {
		// read from buffer
		line, err := r.ReadString('\n')

		if line != "" {
			fmt.Println(line)
		}

		switch err {
		case io.EOF:
			break
		case nil:
			continue
		default:
			panic(err)
		}
	}
}
