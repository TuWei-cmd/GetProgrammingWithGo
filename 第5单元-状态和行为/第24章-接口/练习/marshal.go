/*扩展并编写一个用JSON格式输出坐标的程序
JSON输出分别应该用十进制(DD)和度分秒(DMS)两种格式提供每个坐标
{
	"decimal": 135.9,
	"dms" : "135°54'0.0\" E",
	"degree": 135,
	"minutes": 54,
	"seconds": 0,
	"hemisphere": "E"
}
*/
package main

import (
	"encoding/json"
	"fmt"
)

type coordinate struct {
	d, m, s float64
	h       rune
}

func (c coordinate) String() string {
	return fmt.Sprintf("%v°%v'%.1f\"%c", c.d, c.m, c.s, c.h)
}

type jsonLocation struct {
	Decimal    float64 `json:"decimal"`
	Dms        string  `json:"dms"`
	Degree     float64 `json:"degree"`
	Minutes    float64 `json:"minutes"`
	Seconds    float64 `json:"seconds"`
	Hemisphere rune    `json:"hemisphere"`
}

func (c coordinate) decimal() float64 {
	sign := 1.0
	switch c.h {
	case 'S', 'W', 's', 'w':
		sign = -1
	}
	return sign * (c.d + c.m/60 + c.s/3600)
}

func (c coordinate) MarshalJSON() ([]byte, error) {
	var jsonC = jsonLocation{
		Decimal:    c.decimal(),
		Dms:        c.String(),
		Degree:     c.d,
		Minutes:    c.m,
		Seconds:    c.s,
		Hemisphere: c.h,
	}
	return json.Marshal(jsonC)
}

type location struct {
	lat, long coordinate
}

func (l location) String() string {
	LAT, _ := json.Marshal(l.lat)
	LONG, _ := json.Marshal(l.long)
	return fmt.Sprintf("lat:\n%v\nlong:\n%v", string(LAT), string(LONG))
}

func main() {
	elysium := location{
		lat:  coordinate{4, 30, 0.0, 'N'},
		long: coordinate{135, 54, 0.0, 'E'},
	}
	fmt.Println(elysium)
}
