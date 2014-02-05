package main

import (
    "errors"
    "flag"
    "fmt"
    "github.com/jquag/faker/util"
    "os"
    "strconv"
    "time"
)

type flagTimeStamp time.Time

func (t *flagTimeStamp) String() string {
    return time.Time(*t).String()
}

//TODO JQ: don't use these special Args because the default value is jacked up
func (t *flagTimeStamp) Set(value string) (err error) {
    var ts time.Time
    if len(value) == 8 {
        ts, err = time.Parse("20060102", value)
    } else {
        ts, err = time.Parse("20060102150405", value)
    }
    *t = flagTimeStamp(ts)
    return err
}

var flagSet *flag.FlagSet
var timeStamp flagTimeStamp = flagTimeStamp(time.Now())
var roller util.Roller
var filename string
var count int

func init() {
    flagSet = flag.NewFlagSet("main", flag.ExitOnError)
    flagSet.Usage = printUsage
    flagSet.Var(&roller, "roll", "the roller to use on subsequent files\n\t> format = <roll-type>[<sign><integer>]\n\t> where roll-type is year|month|week|day|hour|min|sec\n\t> e.g. day+1")
    flagSet.Var(&roller, "r", "shorthand for --roll")
    flagSet.Var(&timeStamp, "time", "initial timestamp to use for the files\n\t> defaults to the current time\n\t> format = YYYYMMDD[hhmmss]")
    flagSet.Var(&timeStamp, "t", "shorthand for --time")
}

func printUsage() {
    fmt.Println("usage: faker [options] <filename> [<count>]")
    fmt.Println("Options:")
    flagSet.VisitAll(func(f *flag.Flag) {
        dashes := "-"
        if len(f.Name) > 1 {
            dashes = "--"
        }
        fmt.Printf("%s%s:\t%s\n", dashes, f.Name, f.Usage)
    })
}

func printHelp() {
    fmt.Println("faker, the file maker - a utility for creating files")
    printUsage()
}

func parseArgs() error {
    flagSet.Parse(os.Args[1:])

    switch flagSet.NArg() {
    case 1:
        count = 1
    case 2:
        var err error
        if count, err = strconv.Atoi(flagSet.Arg(1)); err != nil || count <= 0 {
            return errors.New("invalid usage: count must be a positive integer")
        }
    default:
        return errors.New("invalid usage: wrong number of arguments")
    }
    filename = flagSet.Arg(0)
    return nil
}

func main() {
    if err := parseArgs(); err != nil {
        fmt.Printf("%s\n", err)
        printUsage()
    }

    filenames := util.GetFilenames(filename, count)
    for _, name := range filenames {
        fmt.Printf("created: %s\n", name)
    }
}
