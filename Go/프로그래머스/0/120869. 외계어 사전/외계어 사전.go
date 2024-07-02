func solution(spell []string, dic []string) int {
	var r int
	for _, d := range dic {
		var c int
		for _, s := range spell {
			for _, v := range d {
				if s == string(v) {
					c += 1
					break
				}
			
			}
		}
		if len(spell) <= c {
			r = 1
			break
		} else { 
			r = 2
		}
	}
	return r
}