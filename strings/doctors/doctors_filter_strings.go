package main

import (
	"fmt"
	"sort"
	"strings"
)

func getBoardCertifiedOncologists(data [][]string) []string {
	var oncologists []string

	for _, row := range data[1:] {
		name := row[0]
		labels := row[1]

		if isBoardCertified(labels) && isOncologist(labels) {
			oncologists = append(oncologists, name)
		}
	}

	sort.Strings(oncologists)

	return oncologists
}

func isBoardCertified(labels string) bool {
	return strings.Contains(labels, "board certified")
}

func isOncologist(labels string) bool {
	return strings.Contains(labels, "oncology")
}

func solution(data [][]string, searchParams []string) []string {
	var onc []string

	for _, d := range data[1:] {
		name := d[0]
		labels := d[1]

		if filterSearch(labels, searchParams) {
			onc = append(onc, name)
		}
	}

	sort.Strings(onc)

	return onc
}

func filterSearch(labels string, search []string) bool {
	for _, p := range search {
		if p == "not sanctioned" {
			if strings.Contains(labels, "sanctioned") {
				return false
			}
		} else if p == "male" {
			if !strings.Contains(labels, "male") || strings.Contains(labels, "female") {
				return false
			}

		} else if strings.HasPrefix(p, "not") {
			ep := strings.TrimPrefix(p, "not")
			if strings.Contains(labels, ep) {
				return false
			}
		} else {
			if !strings.Contains(labels, p) {
				return false
			}
		}
	}
	return true
}

func main() {
	data := [][]string{
		{"name", "labels"},
		{"Matt", "board certified,primary care,male,takes_new_patients"},
		{"Belda", "board certified,internal medicine,female"},
		{"Wyatt", "primary care,male,takes_new_patients"},
		{"Emma", "board certified,oncology"},
		{"Aaron", "sanctioned,primary care"},
		{"Josh", "board certified, internal medicine, takes_new_patients"},
		{"Adrien", "oncology,board certified, takes_new_patients"},
		{"Andy", "internal medicine,male,sanctioned"},
	}

	searchParams := []string{"male"}

	filteredDoctors := solution(data, searchParams)

	fmt.Println(filteredDoctors)
}
