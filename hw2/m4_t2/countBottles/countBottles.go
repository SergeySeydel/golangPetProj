package countBottles


func CountBottles(a int) string{

	var b int
	b = a % 10

	switch{
	case b == 1:
		return "butilka"
	case b > 1 && a < 5:
		return "butilki"
	case b > 4:
		return "butilok"
	default:
		return "butilok"
	}
}