faker
=====

a utility for creating a bunch of files with various timestamps

usage: faker [options] <filename> [<count>]
Options:
  -r="": shorthand for -roll
  -roll="": the roller to use on subsequent files, format=(year|month|week|day|hour|min|sec)[<sign><int>], e.g. day+3
  -t="20140206223453": shorthand for -time
  -time="20140206223453": initial timestamp to use for the files, format = YYYYMMDD[hhmmss]
