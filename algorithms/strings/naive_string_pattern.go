package main

import "log"

// Naive string matching Algorithm
// Worst Case O(n*m)
func main() {
	text := `Sed ut perspiciatis accusantium unde omnis iste natus 
		error sit voluptatem accusantium doloremque laudantium`
	pattern := "accusantium"
	if exist, indexList := NaiveString(text, pattern); exist {
		for _, index := range indexList {
			log.Printf("Pattern was found at index %d", index)
			log.Println(text[index : index+len(pattern)])
		}
	} else {
		log.Println("Pattern was not found")
	}
	// 2009/11/10 23:00:00 Pattern was found at index 20
	// 2009/11/10 23:00:00 accusantium
	// 2009/11/10 23:00:00 Pattern was found at index 78
	// 2009/11/10 23:00:00 accusantium
}

func NaiveString(text, pattern string) (bool, []int) {
	lenText := len(text)
	lenPattern := len(pattern)
	limit := lenText - lenPattern
	var indexList []int

	for i := 0; i <= limit; i++ {
		j := 0
		for ; j < lenPattern; j++ {
			if text[i+j] != pattern[j] {
				break
			}
		}
		if j == lenPattern {
			indexList = append(indexList, i)
		}
	}

	if len(indexList) > 0 {
		return true, indexList
	}

	return false, indexList
}
