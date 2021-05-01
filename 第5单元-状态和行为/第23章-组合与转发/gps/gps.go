/*
使用gps结构表示全球定位系统，
有两个location结构和一个world结构组成，
前者用于表示当前位置和坐标位置，后者用于表示位置所在的世界
*/
package main

import (
	"fmt"
	"math"
)

// 位置
type location struct {
	lat, long float64
	name      string
}

// // location()返回名称，维度和经度的字符串
func (l location) description() string {
	return fmt.Sprintf("%v(lat:%v and long:%v)", l.name, l.lat, l.long)
}

// 世界
type world struct {
	radius float64
	name   string
}

// distance使用余弦球面定理计算两个位置的距离
func (w world) distance(p1, p2 location) float64 {
	s1, c1 := math.Sincos(rad(p1.lat))
	s2, c2 := math.Sincos(rad(p2.lat))
	clong := math.Cos(rad(p1.long - p2.long))
	return w.radius * math.Acos(s1*s2+c1*c2*clong)
}

type gps struct {
	position, target location
	planet           world
}

// // distance方法计算当前位置和目标位置的距离
func (gps gps) distance() float64 {
	return gps.planet.distance(gps.position, gps.target)
}
func (gps gps) message() string {
	return fmt.Sprintf("The distance form %v to %v: %v km", gps.position.description(), gps.target.description(), gps.distance())
}

type rover struct {
	roverName string
	gps
}

func main() {
	marsGps := gps{
		position: location{lat: -4.5895, long: 137.4417, name: "布莱德柏利着陆点"},
		target:   location{lat: 4.5, long: 135.9, name: "埃律西昂平原"},
		planet:   world{radius: 3389.5, name: "Mars"},
	}
	fmt.Printf("%+v\n", marsGps)
	curiocity := rover{
		roverName: "curiocity",
		gps:       marsGps,
	}
	fmt.Println(curiocity.message())
}

// rad()将角度转换为弧度，方便三角函数的计算
func rad(deg float64) float64 {
	return deg * math.Pi / 180.0
}
