func solution(numbers []int) float64 {
   var w float64 = 0
    for _, v := range numbers {
        w += float64(v)
    }
    return w/float64(len(numbers))
}