package arrplay

import "fmt"
import "strconv"
import "sort"

func stats(a []int) {
    fmt.Printf("Slice a fixed: %d\n", a)
    fmt.Printf("Slice a has cap of: %d\n", cap(a))
    fmt.Printf("Slice a has len: %d\n", len(a))
    fmt.Printf("-----------------\n")
}

func loopy(a []int) map[string]int {
    var r = make(map[string]int)
    for i, v := range(a) {
        l := len(strconv.Itoa(i))
        if l == 1 {
            r["string 0" + strconv.Itoa(i)] = v
        } else {
            r["string " + strconv.Itoa(i)] = v
        }
    }
    return r
}

func mapformat(m map[string]int) {
    // display in order by first making a slice index
    var keys []string
    for k, _ := range m {
        keys = append(keys, k)
    }
    sort.Strings(keys)

    for _, v := range keys {
        fmt.Printf("\t%s\t%d\n", v, m[v])
    }
}

func Dothing() {
    one := [...]int{1,2,3,4,5,6,7,8,9,10,11}
    // stats(one[:])
    // stats(one[:5])
    var two [len(one)+10]int
    copy(two[len(one):],one[:])
    // stats(two[:])
    mapformat(loopy(one[:]))

}