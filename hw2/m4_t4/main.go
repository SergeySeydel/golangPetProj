package main

import "fmt"

func main(){

	var flag,cycleCounts int

	fmt.Print("Enter size of checkmate desk: ")
	fmt.Scan(&cycleCounts)

	if cycleCounts < 0 || cycleCounts > 20{
		fmt.Println("Error, please enter correct size")
		return
	}

	for i := 0; i < cycleCounts; i++{
		for j := 0; j < cycleCounts;j++{
			fmt.Printf("%d ", flag)
			if flag{
				flag = 0
			}
			else{
				flag = 1
			}
		}
		fmt.Print("\n")
	}