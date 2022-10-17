package main

func MinItemProduce(n int, m int, day []int) int {
	t := 0

	if n > len(day) {
		return t
	}

	for {
		items := 0

		for i := 0; i < n; i++ {
			items += (t / day[i])

			if items >= m {
				return t
			}

			t++ // Increment time
		}
	}

	return t
}
