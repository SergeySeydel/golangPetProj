package triangleArea

import "math"

func Distance(x1, y1, x2, y2 float64) float64{
	return math.Sqrt(math.Pow(x2-x1,2)) + math.Pow(y2-y1,2)
}

func TriangleArea(x1, y1, x2, y2, x3, y3 float64) float64{

	a := Distance(x1, y1, x2, y2)
	b := Distance(x1, y1, x3, y3)
	c := Distance(x2, y2, x3, y3)
	s := (a + b + c) / 2
	return math.Sqrt(s * (s - a) * (s - b) * (s - c))
}

