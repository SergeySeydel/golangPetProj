package isRightTriangle

import "math"

func Distance(x1, y1, x2, y2 float64) float64{
	return math.Sqrt(math.Pow(x2-x1,2)) + math.Pow(y2-y1,2)
}

func IsRightTriangle(x1, y1, x2, y2, x3, y3 float64) bool{

	a := math.Pow(Distance(x1, y1, x2, y2), 2)
	b := math.Pow(Distance(x1, y1, x3, y3), 2)
	c := math.Pow(Distance(x2, y2, x3, y3), 2)
	return math.Abs(a+b-c) < 0.000001 || math.Abs(a+c-b) < 0.000001 || math.Abs(b+c-a) < 0.000001
}

