# Calendar Written In Go

Inspired by the UNIX command "cal", this is an implementation written
in pure Go. This comes in handy in restricted Git Bash sessions where
you don't have access to the entire UNIX set of commands.

## Installation

```sh
go install github.com/mojotx/cal/...@latest
```

You can also simply clone the repository and then install locally with:

```sh
go install -v ./...
```

This will install the cal binary into your `${GOBIN}` directory, e.g., `$HOME/go/bin`.

## Examples

### Current Date

```text
$ cal
     July 2025
Su Mo Tu We Th Fr Sa
       1  2  3  4  5
 6  7  8  9 10 11 12
13 14 15 16 17 18 19
20 21 22 23 24 25 26
27 28 29 30 31
```

### Specific month and date

```text
$ cal 2 2024
   February 2024
Su Mo Tu We Th Fr Sa
             1  2  3
 4  5  6  7  8  9 10
11 12 13 14 15 16 17
18 19 20 21 22 23 24
25 26 27 28 29
```

### Entire year

```text
$ cal 2025
    January 2025           February 2025             March 2025
Su Mo Tu We Th Fr Sa    Su Mo Tu We Th Fr Sa    Su Mo Tu We Th Fr Sa
          1  2  3  4                       1                       1
 5  6  7  8  9 10 11     2  3  4  5  6  7  8     2  3  4  5  6  7  8
12 13 14 15 16 17 18     9 10 11 12 13 14 15     9 10 11 12 13 14 15
19 20 21 22 23 24 25    16 17 18 19 20 21 22    16 17 18 19 20 21 22
26 27 28 29 30 31       23 24 25 26 27 28       23 24 25 26 27 28 29
                                                30 31

     April 2025               May 2025               June 2025
Su Mo Tu We Th Fr Sa    Su Mo Tu We Th Fr Sa    Su Mo Tu We Th Fr Sa
       1  2  3  4  5                 1  2  3     1  2  3  4  5  6  7
 6  7  8  9 10 11 12     4  5  6  7  8  9 10     8  9 10 11 12 13 14
13 14 15 16 17 18 19    11 12 13 14 15 16 17    15 16 17 18 19 20 21
20 21 22 23 24 25 26    18 19 20 21 22 23 24    22 23 24 25 26 27 28
27 28 29 30             25 26 27 28 29 30 31    29 30

     July 2025              August 2025            September 2025
Su Mo Tu We Th Fr Sa    Su Mo Tu We Th Fr Sa    Su Mo Tu We Th Fr Sa
       1  2  3  4  5                    1  2        1  2  3  4  5  6
 6  7  8  9 10 11 12     3  4  5  6  7  8  9     7  8  9 10 11 12 13
13 14 15 16 17 18 19    10 11 12 13 14 15 16    14 15 16 17 18 19 20
20 21 22 23 24 25 26    17 18 19 20 21 22 23    21 22 23 24 25 26 27
27 28 29 30 31          24 25 26 27 28 29 30    28 29 30
                        31

    October 2025           November 2025           December 2025
Su Mo Tu We Th Fr Sa    Su Mo Tu We Th Fr Sa    Su Mo Tu We Th Fr Sa
          1  2  3  4                       1        1  2  3  4  5  6
 5  6  7  8  9 10 11     2  3  4  5  6  7  8     7  8  9 10 11 12 13
12 13 14 15 16 17 18     9 10 11 12 13 14 15    14 15 16 17 18 19 20
19 20 21 22 23 24 25    16 17 18 19 20 21 22    21 22 23 24 25 26 27
26 27 28 29 30 31       23 24 25 26 27 28 29    28 29 30 31
                        30
```
