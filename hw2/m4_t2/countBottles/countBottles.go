package countBottles


func CountBottles(a int) string{

	switch{
	case a == 1:
		return "butilka"
	case a > 1 && a < 5:
		return "butilki"
	case a > 4:
		return "butilok"
	default:
		return "butilok"
	}
}