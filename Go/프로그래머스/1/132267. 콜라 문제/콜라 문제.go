func solution(a int, b int, n int) int {
	var z int

	for i := 1; i <= 1000000; i++ {
		if n >= a {
			x := n % a
			y := (n / a) * b
			z += y
			n = x + y
		}
	}

	return z
}