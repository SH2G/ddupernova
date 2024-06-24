import "sort"

func solution(array []int, commands [][]int) []int {
	var r []int
	for _, c := range commands {
		f := make([]int, 0)
		f = append(f, array[c[0]-1 : c[1]]...)
		sort.Ints(f)
		r = append(r, f[c[2]-1])
	}
	return r
}
