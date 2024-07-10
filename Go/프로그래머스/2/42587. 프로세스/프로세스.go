func solution(priorities []int, location int) int {
	m := make([][2]int, len(priorities))
	for i, v := range priorities {
		m[i] = [2]int{i, v} 
	}

	r := 0
	for len(m) > 0 {
		one := m[0] 
		m = m[1:]
		var b int
		for _, v := range m {
			if v[1] > one[1] {
				b = 1
				break
			}
		}
		if b == 1 {
			m = append(m, one)
		} else {
			r++
			if one[0] == location {
				return r
			}
		}
	}

	return r
}