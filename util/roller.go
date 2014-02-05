package util

import (
    "fmt"
    "regexp"
    "strconv"
)

var validRollerString = regexp.MustCompile(`(year|month|week|day|hour|min|sec)([+-]\d+)?$`)

type Roller struct {
    rollType  string
    increment int
}

func (r *Roller) String() string {
    if r.increment >= 0 {
        return fmt.Sprintf("%s+%d", r.rollType, r.increment)
    } else {
        return fmt.Sprintf("%s%d", r.rollType, r.increment)
    }
}

func (r *Roller) Set(value string) error {
    matchGroups := validRollerString.FindStringSubmatch(value)
    if matchGroups == nil {
        return fmt.Errorf("invalid roller: %s", value)
    }

    r.rollType = matchGroups[1]
    if matchGroups[2] == "" {
        r.increment = 1
    } else {
        r.increment, _ = strconv.Atoi(matchGroups[2])
    }

    return nil
}
