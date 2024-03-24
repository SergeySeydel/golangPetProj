package canFormTriangle

import "math"

func Distance(x1, y1, x2, y2 float64) float64{
	return math.Sqrt(math.Pow(x2-x1,2)) + math.Pow(y2-y1,2)
}

func CanFormTriangle(x1, y1, x2, y2, x3, y3 float64) bool{
	return Distance(x1, y1, x2, y2) > 0 && Distance(x1, y1, x3, y3) > 0 && Distance(x2, y2, x3, y3) > 0
}

