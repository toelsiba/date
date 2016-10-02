package date

import (
	"testing"
	"time"
)

func TestDate(t *testing.T) {

	for jd := 0; jd < 10000000; jd++ {

		d1 := FromJulianDay(jd)

		d2, err := FromDate(d1.Date())
		if err != nil {
			t.Fatal(err)
		}

		if d1.jd != d2.jd {
			t.Fatalf("jd(%d) != jd(%d)", d1.jd, d2.jd)
		}
	}
}

func TestZero(t *testing.T) {

	d1, err := FromDate(-1, time.December, 31)
	if err != nil {
		t.Fatal(err)
	}

	d2, err := FromDate(1, time.January, 1)
	if err != nil {
		t.Fatal(err)
	}

	d1_next := d1.AddDays(1)

	if d1_next.jd != d2.jd {
		t.Fatalf("jd(%d) != jd(%d)", d1_next.jd, d2.jd)
	}
}

func TestPassage(t *testing.T) {

	d1, err := FromDate(lastJulianDay.year, lastJulianDay.month, lastJulianDay.day)
	if err != nil {
		t.Fatal(err)
	}

	d2, err := FromDate(firstGregorianDay.year, firstGregorianDay.month, firstGregorianDay.day)
	if err != nil {
		t.Fatal(err)
	}

	d1_next := d1.AddDays(1)

	if d1_next.jd != d2.jd {
		t.Fatalf("jd(%d) != jd(%d)", d1_next.jd, d2.jd)
	}
}
