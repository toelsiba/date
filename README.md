# date
Package for work with date in golang

Partly is an analog Qt QDate Class.

## Features:
  * dates before or equal 4 October 1582 considered by Julian calendar.
  * dates equal or after 15 October 1582 considered by Gregorian calendar.
  * dates more than 4 October and less 15 October 1582 year is invalid (error: date.ErrPassage).
  * year zero is invalid (if try call function FromDate(0, month, day) then return error ErrZeroYear).
  * dates represented as the Julian day (first Julian day (jd = 0) is 1 January -4713 year), dates less first JD is invalid.

## Installation
```bash
go get github.com/toelsiba/date
```
## Usage

### Example usage JulianDay & Weekday:
```go
package main

import (
	"fmt"
	"log"
	"time"

	"github.com/toelsiba/date"
)

func main() {
	// Nikolay Ivanovich Pirogov was born on 25 November 1810
	d, err := date.FromDate(1810, time.November, 25)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("For date %s:\n", d)
	fmt.Println("\t- JulianDay =", d.JulianDay())
	fmt.Println("\t- Weekday =", d.Weekday())
}
```
result:
```
For date 1810-11-25:
	- JulianDay = 2382477
	- Weekday = Sunday
```

### Example usage AddDays:
```go
package main

import (
	"fmt"
	"log"
	"time"

	"github.com/toelsiba/date"
)

func main() {
	d, err := date.FromDate(-1, time.December, 31)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(d)
	d = d.AddDays(1)
	fmt.Println(d)
}
```
result:
```
-0001-12-31
0001-01-01
```
