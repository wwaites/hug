package jproj

import (
	"cproj"
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"webx"
)

type xformReq struct {
	from_pj, to_pj *cproj.Proj
	coords []float64
}

func (x *xformReq) Free() {
	if x.from_pj != nil {
		x.from_pj.Free()
	}
	if x.to_pj != nil {
		x.to_pj.Free()
	}
}

func parseForm(req *http.Request) (xfr *xformReq, err error) {
	var xr xformReq
	var e error

	xr.coords = make([]float64, 0, 3)

	xs := req.FormValue("x")
	x, e := strconv.ParseFloat(xs, 64)
	if e != nil {
		err = e
		return
	}
	xr.coords = append(xr.coords, x)

	ys := req.FormValue("y")
	y, e := strconv.ParseFloat(ys, 64)
	if e != nil {
		err = e
		return
	}
	xr.coords = append(xr.coords, y)

	zs := req.FormValue("z")
	z, e := strconv.ParseFloat(zs, 64)
	if e == nil {
		xr.coords = append(xr.coords, z)
	}

	from_srid := req.FormValue("from_srid")
	from_epsg := bytes.NewBufferString("+init=epsg:")
	from_epsg.WriteString(from_srid)

	xr.from_pj, e = cproj.InitPlus(from_epsg.String())
	if e != nil {
		err = e
		return
	}

	to_srid := req.FormValue("to_srid")
	to_epsg := bytes.NewBufferString("+init=epsg:")
	to_epsg.WriteString(to_srid)
	xr.to_pj, e = cproj.InitPlus(to_epsg.String())
	if e != nil {
		err = e
		xr.from_pj.Free()
		return
	}

	xfr = &xr
	return
}

func parseBody(req *http.Request) (xfr *xformReq, err error) {
	var xr xformReq
	var e error

	dec := json.NewDecoder(req.Body)
	jreq := make(map[string]interface{})
	e = dec.Decode(&jreq)
	if e != nil {
		err = e
		return
	}

	xr.coords = make([]float64, 0, 3)

	xi, ok := jreq["x"]
	if !ok {
		err = errors.New("missing x coordinate")
		return
	}
	x, ok := xi.(float64)
	if !ok {
		err = errors.New("invalid x coordinate")
		return
	}
	xr.coords = append(xr.coords, x)

	yi, ok := jreq["y"]
	if !ok {
		err = errors.New("missing y coordinate")
		return
	}
	y, ok := yi.(float64)
	if !ok {
		err = errors.New("invalid y coordinate")
		return
	}
	xr.coords = append(xr.coords, y)

	zi, ok := jreq["z"]
	if ok {
		z, ok := zi.(float64)
		if !ok {
			err = errors.New("invalid z coordinate")
			return
		}
		xr.coords = append(xr.coords, z)
	}

	from_sridi, ok := jreq["from_srid"]
	if !ok {
		err = errors.New("missing from_srid")
		return
	}
	from_srid, ok := from_sridi.(float64)
	if !ok {
		err = errors.New("invalid from_srid")
		return
	}
	from_epsg := fmt.Sprintf("+init=epsg:%d", int(from_srid))

	xr.from_pj, e = cproj.InitPlus(from_epsg)
	if e != nil {
		err = e
		return
	}

	to_sridi, ok := jreq["to_srid"]
	if !ok {
		err = errors.New("missing to_srid")
		xr.from_pj.Free()
		return
	}
	to_srid, ok := to_sridi.(float64)
	if !ok {
		err = errors.New("invalid to_srid")
		xr.from_pj.Free()
		return
	}
	to_epsg := fmt.Sprintf("+init=epsg:%d", int(to_srid))
	xr.to_pj, e = cproj.InitPlus(to_epsg)
	if e != nil {
		err = e
		xr.from_pj.Free()
		return
	}

	xfr = &xr
	return
}

func JsonProj(w http.ResponseWriter, req *http.Request) {
	var xr *xformReq
	var err error

	ctype := req.Header.Get("Content-Type")
	isJson := (ctype == "application/json" || ctype == "text/javascript")
	if req.Method == "POST" && isJson {
		xr, err = parseBody(req)
	} else {
		xr, err = parseForm(req)
	}
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer xr.Free()

	result, err := cproj.Transform(xr.from_pj, xr.to_pj, xr.coords)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	r := make(map[string]interface{})
	r["x"] = result[0]
	r["y"] = result[1]
	if len(result) == 3 {
		r["z"] = result[2]
	}

	webx.JsonResponse(w, r)
}
