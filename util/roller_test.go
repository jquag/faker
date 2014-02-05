package util

import (
    "testing"
)

var stringTests = []struct {
    in  Roller
    out string
}{
    {Roller{"day", 0}, "day+0"},
    {Roller{"day", -0}, "day+0"},
    {Roller{"month", -1}, "month-1"},
    {Roller{"year", +1}, "year+1"},
    {Roller{"week", +123}, "week+123"},
}

func TestString(t *testing.T) {
    for _, tcase := range stringTests {
        if result := tcase.in.String(); result != tcase.out {
            t.Errorf("{%s, %d}.String() = %s, want %s", tcase.in.rollType, tcase.in.increment, result, tcase.out)
        }
    }
}

var setTests = []struct {
    in           string
    outType      string
    outIncrement int
    outErr       string
}{
    {"year", "year", 1, ""},
    {"month+1", "month", 1, ""},
    {"week-1", "week", -1, ""},
    {"day+0", "day", 0, ""},
    {"hour+123", "hour", 123, ""},
    {"minny+5", "", 0, "invalid roller: minny+5"},
}

func TestSet(t *testing.T) {
    for _, tcase := range setTests {
        r := new(Roller)
        r.Set(tcase.in)
        err := r.Set(tcase.in)
        var actualErrorString string
        if err != nil {
            actualErrorString = err.Error()
        }

        if actualErrorString != tcase.outErr {
            t.Errorf("r.Set(%s) = \"%s\", want \"%s\"", tcase.in, err, tcase.outErr)
        }
        if r.rollType != tcase.outType || r.increment != tcase.outIncrement {
            t.Errorf("After r.Set(\"%s\")... r={%s, %d}, want r={%s, %d}", tcase.in, r.rollType, r.increment, tcase.outType, tcase.outIncrement)
        }
    }
}
