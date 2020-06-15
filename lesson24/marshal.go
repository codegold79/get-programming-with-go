// Experiment: marshal.go
// Write a program that outputs coordinates in JSON format, expanding on work done for
// the preceding quick check. The JSON output should provide each coordinate in decimal
// degrees (DD) as well as the degrees, minutes, seconds format:
// {
// 	"decimal": 135.9,
// 	"dms": "135°54'0.0\" E",
// 	"degrees": 135,
// 	"minutes": 54,
// 	"seconds": 0,
// 	"hemisphere": "E"
// }
// This can be achieved without modifying the coordinate structure by satisfying the
// json.Marshaler interface to customize the JSON. The MarshalJSON method you write may
// make use of json.Marshal.
// NOTE To calculate decimal degrees, you’ll need the decimal method introduced in lesson 22.
package main

import (
	"encoding/json"
	"fmt"
)

type coordinate struct {
	d, m, s float64
	h       rune
}

type jsonCoordinate struct {
	Deg, Min, Sec, Dec float64
	DMS                string
	Hem                string
}

func main() {
	coords := coordinate{
		d: 135,
		m: 54,
		s: 0,
		h: 'E',
	}

	if b, err := coords.MarshalJSON(); err != nil {
		fmt.Println("marshalling json:", err)
	} else {
		fmt.Println(string(b))
	}
}

func (c coordinate) MarshalJSON() ([]byte, error) {
	jc := jsonCoordinate{
		Deg: c.d,
		Min: c.m,
		Sec: c.s,
		Hem: string(c.h),
		DMS: fmt.Sprintf("%1.0f°%1.0f'%1.2f\" %c", c.d, c.m, c.s, c.h),
		Dec: c.decimal(),
	}

	return json.Marshal(jc)
}

func (c coordinate) decimal() float64 {
	sign := 1.0
	switch c.h {
	case 'S', 'W', 's', 'w':
		sign = -1
	}
	return sign * (c.d + c.m/60 + c.s/3600)
}
