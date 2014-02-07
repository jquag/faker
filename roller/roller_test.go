package roller

import (
    "errors"
    "github.com/stretchr/testify/assert"
    "testing"
    "time"
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
        assert.Equal(t, tcase.out, tcase.in.String())
    }
}

var casesForNew = []struct {
    in     string
    out    *Roller
    outErr error
}{
    {"year", &Roller{"year", 1}, nil},
    {"month+1", &Roller{"month", 1}, nil},
    {"week-1", &Roller{"week", -1}, nil},
    {"day+0", &Roller{"day", 0}, nil},
    {"hour+123", &Roller{"hour", 123}, nil},
    {"minny+5", nil, errors.New("invalid roller: minny+5")},
}

func TestNew(t *testing.T) {
    for _, tcase := range casesForNew {
        r, e := New(tcase.in)
        assert.Equal(t, tcase.out, r)
        assert.Equal(t, tcase.outErr, e)
    }
}

const tl = "20060102150405"

type rollCase struct {
    inc      int
    expected string
}

var startTime, _ = time.Parse(tl, "20140206194019")

var casesForRollYear = []rollCase{
    {1, "20150206194019"},
    {4, "20180206194019"},
    {-2, "20120206194019"},
}

var casesForRollWeek = []rollCase{
    {1, "20140213194019"},
    {4, "20140306194019"},
    {-2, "20140123194019"},
}

var casesForRollHour = []rollCase{
    {1, "20140206204019"},
    {5, "20140207004019"},
    {-3, "20140206164019"},
}

func TestRollYear(t *testing.T) {
    testRoll(casesForRollYear, "year", t)
}

func TestRollWeek(t *testing.T) {
    testRoll(casesForRollWeek, "week", t)
}

func TestRollHour(t *testing.T) {
    testRoll(casesForRollHour, "hour", t)
}

func testRoll(cases []rollCase, rollString string, t *testing.T) {
    for _, tcase := range cases {
        r := &Roller{rollString, tcase.inc}
        actual, _ := r.Roll(startTime)
        expected, _ := time.Parse(tl, tcase.expected)
        assert.Equal(t, expected, actual, "invalid result rolling %s by %d", rollString, tcase.inc)
    }
}
