package BaseConversion

import "testing"


type testParam struct {
    in uint
    out string
}


var testParams = []testParam {
    {1, "1"},
    {4, "4"},
    {39,"D"},
    {62, "10"},
    {63, "11"},
    {64, "12"},
    {465, "7v"},
}

func TestSolution(t *testing.T) {
    for _,tp := range testParams {
        got := ConvertToBase62(tp.in)
        if got != tp.out {
            t.Errorf("on input: %d got: %v, wanted: %v", tp.in, got, tp.out)
        }
    }
}
