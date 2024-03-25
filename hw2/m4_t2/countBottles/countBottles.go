package countBottles


func CountBottles(a int) string{

	
	switch{
	case a % 100 < 11 || a % 100 > 14:
		switch a % 10{
		case 1:
			return "butilka"
		case 2,3,4:
			return "butilki"
		default:
			return "butilok"
		}
	default:
		return "butilok"
		
	}
	
}