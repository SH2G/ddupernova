func solution(n_str string) string {
	var r string
	for _, s := range n_str {
		if s != '0' {
			r += string(s)
		}else if r != "" {
			r += string(s)
		}
	}
	return r
}