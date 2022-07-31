package main

//domino in a middle of R and L it stays still
//L and R pushes at the same time
//initial array don't change size
//two indexes for the start and end of an array

func pushDominoes(dominoes string) string {
	runes := []rune(dominoes)

	rightIndex := len(runes) - 1
	for i, item := range runes {

		if item == 'R' && i+1 < len(runes) && runes[i+1] == '.' {
			runes[i+1] = 'R'
		}

		if runes[rightIndex] == 'L' {
			runes[rightIndex-1] = 'L'
		}

		if i-1 == rightIndex {
			break
		}

		if rightIndex > 0 {
			rightIndex = rightIndex - 1
		}
	}

	dominoes = string(runes)

	return dominoes
}

/*
func main() {
	println(pushDominoes("RR.L") == "RR.L")
	println(pushDominoes(".L.R...LR..L..") == "LL.RR.LLRRLL..")
	println(pushDominoes("R....") == "RRRRR")
	println(pushDominoes("R....L") == "RRRLLL")
	println(pushDominoes("R.....L") == "RRR.LLL")
	println(pushDominoes("R.L...L") == "R.LLLLL")
	println(pushDominoes("L.L.R.R") == "LLL.RRR")
}
/*

*/
