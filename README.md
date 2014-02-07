faker
=====

a utility for creating a bunch of files with various timestamps

usage:<code>faker [options] FILENAME [COUNT]</code>

Options:
* -r: shorthand for -roll
* -roll: the roller to use on subsequent files, format=(year|month|week|day|hour|min|sec)[(+|-)INT], e.g. day+3
* -t=CURRENT_TIME: shorthand for -time
* -time=CURRENT_TIME: initial timestamp to use for the files, format = YYYYMMDD[hhmmss]
