package cproj

import (
	"bytes"
	"errors"
	"fmt"
	"log"
	"math"
)

func pj_init_plus(*byte) int __asm__("pj_init_plus")
func pj_free(int) __asm__("pj_free")
func pj_transform(int, int, int, int, *float64, *float64, *float64) int __asm__("pj_transform")
func pj_is_latlong(int) int __asm__("pj_is_latlong")

type Proj struct {
	pj int
}

func InitPlus(spec string) (proj *Proj, err error) {
	buf := bytes.NewBufferString(spec)
	buf.WriteByte(0)
	p := pj_init_plus(&(buf.Bytes()[0]))
	if p == 0 {
		err = errors.New(fmt.Sprintf("Invalid spec: %s", spec))
	} else {
		proj = &Proj{p}
	}
	return 
}

func (p* Proj) Free() {
	pj_free(p.pj)
}

func (p* Proj) IsLatLong() bool {
	if pj_is_latlong(p.pj) == 0 {
		return false
	}
	return true
}

func Transform(from_pj, to_pj *Proj, coord []float64) (result []float64, err error) {
	if len(coord) < 2 || len(coord) > 3 {
		err = errors.New("Invalid coordinates to transform")
		return
	}
	inplace := make([]float64, len(coord), 3)
	copy(inplace, coord)

	if from_pj.IsLatLong() {
		for i, v := range inplace {
			inplace[i] = v * math.Pi / 180
		}
		log.Print("source radians: ", inplace)
	}

	var ret int
	if len(coord) == 2 {
		ret = pj_transform(from_pj.pj, to_pj.pj, 1, 1,
			&inplace[0], &inplace[1], nil)
	} else {
		ret = pj_transform(from_pj.pj, to_pj.pj, 1, 1,
			&inplace[0], &inplace[1], &inplace[2])
	}
	if ret != 0 {
		err = errors.New("Transformation failed")
	} else {
		if to_pj.IsLatLong() {
			log.Print("result radians: ", inplace)
			for i, v := range inplace {
				inplace[i] = v * 180 / math.Pi
			}
		}
		result = inplace
	}
	return
}
