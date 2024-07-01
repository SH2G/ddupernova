func solution(answers []int) []int {
	one := []int{1, 2, 3, 4, 5}
	two := []int{2, 1, 2, 3, 2, 4, 2, 5}
	three := []int{3, 3, 1, 1, 2, 2, 4, 4, 5, 5}

	var o int
	var w int
	var h int

	for i, v := range answers {
		if v == one[i%len(one)] {
			o += 1
		}
		if v == two[i%len(two)] {
			w += 1
		}
		if v == three[i%len(three)] {
			h += 1
		}
	}

	r := []int{}
	t := []int{o, w, h}
	selectionSort(t)

	if o == t[len(t)-1] {
		r = append(r, 1)
	}

	if w == t[len(t)-1] {
		r = append(r, 2)
	}

	if h == t[len(t)-1] {
		r = append(r, 3)
	}

	return r
}

func selectionSort(s []int) []int {
	for i := 0; i < len(s)-1; i++ {
		for j := i + 1; j < len(s); j++ {
			if s[j] < s[i] {
				s[i], s[j] = s[j], s[i]
			}
		}
	}
	return s
}