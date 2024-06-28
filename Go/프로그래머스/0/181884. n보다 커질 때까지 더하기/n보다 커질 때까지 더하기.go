func solution(numbers []int, n int) int {
	var s int
	for _, b := range numbers {
		if s <= n {
			s += b
			continue
		}
		break
	}
	return s
}
