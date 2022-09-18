package main

func FindSum(nums []int, res int) bool {
	m := make(map[int]int)
	for _, num := range nums {
		_, ok := m[num]
		if ok {
			continue
		}
		m[res - num] = 1
	}
	return false
}
/*
func main() {
	fmt.Println("test")
}
*/
