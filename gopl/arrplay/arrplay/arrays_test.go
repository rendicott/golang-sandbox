package arrplay

import "testing"

func TestMapformat(t *testing.T) {
    desired :=
`        string 00       1
        string 01       2
        string 02       3
        string 03       4
        string 04       5
        string 05       6
        string 06       7
        string 07       8
        string 08       9
        string 09       10
        string 10       11`
    one := [...]int{1,2,3,4,5,6,7,8,9,10,11}
    output := mapformat(loopy(one[:]))
    if output != desired {
        t.Error("mapformat(loopy(one[:])) != desired")
    }
}