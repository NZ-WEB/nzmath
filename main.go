package nzmath

import "math"

func Cosd(angle float64) float64 {
	return math.Cos(angle * math.Pi / 180.0)
}

func Rad2deg(angle float64) float64 {
	return angle * 180.0 / math.Pi
}

func Sind(angle float64) float64 {
	return math.Sin(angle * math.Pi / 180.0)
}

func Tand(angle float64) float64 {
	return math.Tan(angle * math.Pi / 180.0)
}

func Atand(angle float64) float64 {
	return math.Atan(angle) * math.Pi / 180.0
}
