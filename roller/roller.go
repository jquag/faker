package roller

import (
    "errors"
    "fmt"
    "regexp"
    "strconv"
    "time"
)

var validRollerString = regexp.MustCompile(`(year|month|week|day|hour|min|sec)([+-]\d+)?$`)

type Roller struct {
    RollType  string
    Increment int
}

func (r *Roller) String() string {
    if r.Increment >= 0 {
        return fmt.Sprintf("%s+%d", r.RollType, r.Increment)
    } else {
        return fmt.Sprintf("%s%d", r.RollType, r.Increment)
    }
}

func New(value string) (r *Roller, err error) {
    matchGroups := validRollerString.FindStringSubmatch(value)
    if matchGroups == nil {
        err = fmt.Errorf("invalid roller: %s", value)
        return
    }

    r = new(Roller)
    r.RollType = matchGroups[1]
    if matchGroups[2] == "" {
        r.Increment = 1
    } else {
        r.Increment, _ = strconv.Atoi(matchGroups[2])
    }

    return
}

func (r *Roller) Roll(t time.Time) (time.Time, error) {
    switch r.RollType {
    case "year":
        return t.AddDate(r.Increment, 0, 0), nil
    case "month":
        return t.AddDate(0, r.Increment, 0), nil
    case "week":
        return t.AddDate(0, 0, r.Increment*7), nil
    case "day":
        return t.AddDate(0, 0, r.Increment), nil
    case "hour", "min", "sec":
        dur, err := time.ParseDuration(fmt.Sprintf("%d%s", r.Increment, r.RollType[0:1]))
        if err != nil {
            return t, err
        }
        return t.Add(dur), nil
    default:
        return t, errors.New("do not know how to roll the time with this roller")
    }
}
