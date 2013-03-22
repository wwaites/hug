package misc

import (
	"strconv"
	"strings"
	"gallows.inf.ed.ac.uk/hug/alg"
)

func ParseCoord(coord string) (v alg.Vector, err error) {
	xs := strings.Split(coord, ",")
	v = make(alg.Vector, 0, len(xs))
	for _, s := range xs {
		x, e := strconv.ParseFloat(s, 64)
		if e != nil {
			err = e
			break
		}
		v = append(v, x)
	}
	return
}

