package main

import (
	"fmt"
	"sort"
)

type Person struct {
	First string
	Age   int
}

type ByAge []Person

func (a ByAge) Len() int {
	return len(a)
}

func (a ByAge) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

func (a ByAge) Less(i, j int) bool {
	return a[i].Age < a[j].Age
}

func main() {
	me := Person{"martin", 29}
	brother := Person{"noah", 20}
	sisterOne := Person{"miranda", 26}
	sisterTwo := Person{"alexis", 23}

	family := []Person{me, brother, sisterOne, sisterTwo}

	fmt.Println(family)

	sort.Sort(ByAge(family))

	fmt.Println(family)
}
