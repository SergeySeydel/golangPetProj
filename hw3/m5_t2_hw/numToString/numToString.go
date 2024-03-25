package numToString

import (
	"fmt"
	"strconv"
)

func NumToString(num string) string{

	result := ""
	for i := 0; i < len(num); i += 2{
		n, err := strconv.Atoi(num[i:i+2])
		if err != nil{
			fmt.Println("Error with Atoi", err)
			return ""
		}
		if n == 26{
			result += " "
		} else {
			result += string('a' + n)
		}

	}

	return result

}