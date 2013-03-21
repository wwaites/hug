package tgutil

import (
	"bufio"
	"bytes"
	"os"
	"strconv"
	"vect"
)

// given a file with floating point numbers, space separated, several
// per line, return each line, parsed, into a channel
func Numbers(f *os.File) (ch chan vect.Vector) {
	ch = make(chan vect.Vector)
	go func() {
		b := bufio.NewReader(f)
		for { 
			line, _, err := b.ReadLine()
			if err != nil {
				break
			}

			sp := bytes.Split(line, []byte(" "))
			v := make(vect.Vector, 0, len(sp))
			for _, s := range sp {
				x, err := strconv.ParseFloat(string(s), 64)
				if err != nil {
					continue
				}
				v = append(v, x)
			}
			if len(v) == len(sp) {
				ch <- v
			}
		}
	close(ch)
	}()
	return ch
}
