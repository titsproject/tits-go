package main

import "fmt"

func main() {

	tits := [30][74]string{}
	n := ""

	for i := 0; i <= 29; i++ {
		if i < 10 {
			n = " "
		} else {
			n = ""
		}

		fmt.Print(n, i, " ")

		for j := 0; j <= 73; j++ {
			fmt.Print(tits[i][j], "0")
		}

		fmt.Println()
	}

}
