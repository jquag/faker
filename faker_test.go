package main

import (
    "github.com/stretchr/testify/assert"
    "testing"
)

var casesForGetFilenams = []struct {
    name  string
    count int
    out   []string
}{
    {"webapp.log", 2, []string{"webapp1.log", "webapp2.log"}},
    {"noextension", 2, []string{"noextension1", "noextension2"}},
    {"singleFile.txt", 1, []string{"singleFile.txt"}},
    {"dumb", -1, []string{"dumb"}},
    {"webapp.log.19790924", 2, []string{"webapp.log1.19790924", "webapp.log2.19790924"}},
}

func TestGetFilenames(t *testing.T) {
    for _, tcase := range casesForGetFilenams {
        result := GetFilenames(tcase.name, tcase.count)
        assert.Equal(t, tcase.out, result)
    }
}
