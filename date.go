package date

import "time"

type Date struct {
	jd int // Julian day
}

func FromTime(t time.Time) (Date, error) {
	year, month, day := t.Date()
	return FromDate(year, month, day)
}

func FromJulianDay(jd int) Date {
	if jd < 0 {
		jd = 0
	}
	return Date{jd}
}

func FromDate(year int, month time.Month, day int) (Date, error) {
	if err := checkDate(year, month, day); err != nil {
		return Date{}, err
	}
	jd := julianDayFromDate(year, int(month), day)
	return Date{jd}, nil
}

func CurrentDate() Date {
	now := time.Now()
	year, month, day := now.Date()
	jd := julianDayFromDate(year, int(month), day)
	return Date{jd}
}

func (d Date) Date() (year int, month time.Month, day int) {
	var m int
	year, m, day = julianDayToDate(d.jd)
	month = time.Month(m)
	return
}

func (d Date) JulianDay() int {
	return d.jd
}

func (d Date) Weekday() time.Weekday {
	return time.Weekday((d.jd + 1) % 7)
}

func (a Date) DaysTo(b Date) int {
	return b.jd - a.jd
}

func (d Date) AddDays(ndays int) Date {
	return FromJulianDay(d.jd + ndays)
}

func (d Date) String() string {
	return d.Format("2006-01-02")
}

func (d Date) Format(layout string) string {
	year, month, day := d.Date()
	t := time.Date(year, month, day, 0, 0, 0, 0, time.UTC)
	return t.Format(layout)
}

func Parse(layout, value string) (Date, error) {
	t, err := time.Parse(layout, value)
	if err != nil {
		return Date{}, err
	}
	year, month, day := t.Date()
	return FromDate(year, month, day)
}
