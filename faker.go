package main

import (
    "errors"
    "flag"
    "fmt"
    "github.com/jquag/faker/roller"
    "os"
    "strconv"
    "strings"
    "time"
)

const dateLayout, timeLayout = "20060102", "150405"

var flagSet *flag.FlagSet

//flags
var timeStamp time.Time
var timeStampString = time.Now().Format(dateLayout + timeLayout)
var rollerString string
var rollr *roller.Roller

//args
var filename string
var count int

func init() {
    flagSet = flag.NewFlagSet("main", flag.ExitOnError)
    flagSet.Usage = printHelp
    flagSet.StringVar(&rollerString, "roll", rollerString, "the roller to use on subsequent files, format=(year|month|week|day|hour|min|sec)[<sign><int>], e.g. day+3")
    flagSet.StringVar(&rollerString, "r", rollerString, "shorthand for -roll")
    flagSet.StringVar(&timeStampString, "time", timeStampString, "initial timestamp to use for the files, format = YYYYMMDD[hhmmss]")
    flagSet.StringVar(&timeStampString, "t", timeStampString, "shorthand for -time")
}

func printUsage() {
    fmt.Println("usage: faker [options] <filename> [<count>]")
    fmt.Println("Options:")
    flagSet.PrintDefaults()
}

func printHelp() {
    fmt.Println("faker, file maker - a utility for creating empty files with rolling timestamps")
    printUsage()
}

func parseArgs() error {
    flagSet.Parse(os.Args[1:])

    timeStampLayout := dateLayout
    if len(timeStampString) > 8 {
        timeStampLayout += timeLayout
    }
    var err error
    if timeStamp, err = time.Parse(timeStampLayout, timeStampString); err != nil {
        return errors.New("invalid usage: bad timestamp")
    }

    if rollerString != "" {
        if rollr, err = roller.New(rollerString); err != nil {
            return errors.New("invalid usage: bad roller")
        }
    }

    switch flagSet.NArg() {
    case 1:
        count = 1
    case 2:
        if count, err = strconv.Atoi(flagSet.Arg(1)); err != nil || count <= 0 {
            return errors.New("invalid usage: count must be a positive integer")
        }
    default:
        return errors.New("invalid usage: wrong number of arguments")
    }
    filename = flagSet.Arg(0)
    return nil
}

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

func createFile(name string, ts time.Time) {
    if _, err := os.Stat(name); err == nil {
        fmt.Printf("%s already exists, not touching\n", name)
    } else if _, err := os.Create(name); err != nil {
        fmt.Printf("failed to create: %s\n", name)
    } else {
        if err := os.Chtimes(name, timeStamp, timeStamp); err != nil {
            fmt.Printf("created: %s, but failed to set the correct timestamp\n", name)
        } else {
            fmt.Printf("created: %s @ %s\n", name, timeStamp.Format(dateLayout+timeLayout))
        }
    }
}

func main() {
    if err := parseArgs(); err != nil {
        fmt.Printf("%s\n", err)
        printUsage()
        return
    }

    filenames := GetFilenames(filename, count)

    for _, name := range filenames {
        createFile(name, timeStamp)
        if rollr != nil {
            timeStamp, _ = rollr.Roll(timeStamp)
        }
    }
}
