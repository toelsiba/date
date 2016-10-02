# date
Package for work with date in golang

Partly is an analog Qt QDate Class.

## Installation
```bash
go get github.com/toelsiba/date
```
## Usage
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
