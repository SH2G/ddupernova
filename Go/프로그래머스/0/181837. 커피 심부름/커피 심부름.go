func solution(order []string) int {
    var a int
    var c int
    for _, o := range order {
        if o == "anything" || o == "americanoice" || o == "iceamericano" || o == "americano" || o == "hotamericano" || o == "americanohot" {
          a += 4500
        }else {
            c += 5000 
        }
    }
    return a + c
}