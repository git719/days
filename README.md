# days
Command line calendar days calculator.

## Why
Almost all of these days calculations can of course be done with most Unix `date` commands, or with many other tools. So why a dedicated `days` utility? Because these simple calculations seem to occur very, and having a simple dedicated utility for them is quite handy. Using the `date` command, and.or most other ways, seems to always get to elaborate. For instance, given someone's birth date, you can quickly pop up their age by doing: 

```
days 1995-07-14
-9921 (27 years + 66 days)
```

Or maybe there's a need to quickly calculate how many days and years passed between 2 historical dates or years: 
```
days 1492-07-01 1776-12-01
103882 (284 years + 222 days)
```

Or maybe one simply needs what the date was X days ago, or what it will be X days into the future: 
```
days -342
2021-10-04

days +90
2022-12-10
```


## Getting started
To compile, you obviously need to have GoLang installed and properly setup in your system, with `$GOPATH` set up correctly (typically at `$HOME/go`). Also setup `$GOPATH/bin/` in your `$PATH`, since that's where the executable binary will be placed.

From a `bash` shell: 

```
go mod init days
go mod tidy
./build
```

To build in Windows you have to have a BASH shell such as [GitBASH](https://www.git-scm.com/download/win). However, to build from a regular Windows Command Prompt, just run the corresponding line in the `build` file (`go build ...`).

This utility has been successfully built and tested on macOS, Linux, and Windows.

## Known Issues
- All calculations are based on UTC timezone.

## Usage
```
days Calendar days calculator v1
     -DAYS                    Print the date from number of DAYS ago, e.g. -11
     +DAYS                    Print what the date will be in number of DAYS into the future, e.g. +6
     YYYY-MM-DD               Print number of days since given date, or into the future
     YYYY-MM-DD YYYY-MM-DD    Print number of days between the 2 dates
```
