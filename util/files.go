package util

import (
    "fmt"
    "strings"
)

func GetFilenames(name string, count int) (names []string) {
    if count <= 1 {
        return []string{name}
    }

    suffixStartIndex := strings.LastIndex(name, ".")
    filenamePrefix, filenameSuffix := name, ""
    if suffixStartIndex != -1 {
        filenamePrefix = name[0:suffixStartIndex]
        filenameSuffix = name[suffixStartIndex:]
    }
    names = make([]string, count)
    for i := 0; i < count; i++ {
        names[i] = fmt.Sprintf("%s%d%s", filenamePrefix, i+1, filenameSuffix)
    }
    return
}
