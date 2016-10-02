package date

// Julian calendar until October 4, 1582
// Algorithm from Frequently Asked Questions about Calendars by Claus Toendering
func julianDayFromJulianDate(year, month, day int) (jd int) {
	a := (14 - month) / 12
	e := (153*(month+(12*a)-3) + 2) / 5
	e += (1461 * (year + 4800 - a)) / 4
	e += day - 32083
	return e
}

// Gregorian calendar starting from October 15, 1582
// Algorithm from Henry F. Fliegel and Thomas C. Van Flandern
func julianDayFromGregorianDate(year, month, day int) (jd int) {
	e := (1461 * (year + 4800 + (month-14)/12)) / 4
	e += (367 * (month - 2 - 12*((month-14)/12))) / 12
	e += -(3 * ((year + 4900 + (month-14)/12) / 100)) / 4
	e += day - 32075
	return e
}

// Julian calendar until October 4, 1582
// Algorithm from Frequently Asked Questions about Calendars by Claus Toendering
func julianDayToJulianDate(jd int) (year, month, day int) {
	jd += 32082
	dd := (4*jd + 3) / 1461
	ee := jd - (1461*dd)/4
	mm := ((5 * ee) + 2) / 153
	day = ee - (153*mm+2)/5 + 1
	month = mm + 3 - 12*(mm/10)
	year = dd - 4800 + (mm / 10)
	if year <= 0 {
		year--
	}
	return
}

// Gregorian calendar starting from October 15, 1582
// This algorithm is from Henry F. Fliegel and Thomas C. Van Flandern
func julianDayToGregorianDate(jd int) (year, month, day int) {
	ell := uint64(jd) + 68569
	n := (4 * ell) / 146097
	ell = ell - (146097*n+3)/4
	i := (4000 * (ell + 1)) / 1461001
	ell = ell - (1461*i)/4 + 31
	j := (80 * ell) / 2447
	day = int(ell - (2447*j)/80)
	ell = j / 11
	month = int(j + 2 - (12 * ell))
	year = int(100*(n-49) + i + ell)
	return
}

func julianDayFromDate(year, month, day int) (jd int) {
	if year < 0 {
		year++
	}
	// the day following October 4, 1582 is October 15, 1582
	if year > 1582 || (year == 1582 && (month > 10 || (month == 10 && day >= 15))) {
		// First Gregorian Day 15/10/1582
		jd = julianDayFromGregorianDate(year, month, day)
		return
	}
	if year < 1582 || (year == 1582 && (month < 10 || (month == 10 && day <= 4))) {
		// Last Julian Day 04/10/1582
		jd = julianDayFromJulianDate(year, month, day)
		return
	}
	jd = 0
	return
}

func julianDayToDate(jd int) (year, month, day int) {
	if jd >= 0 {
		if jd >= 2299161 {
			year, month, day = julianDayToGregorianDate(jd)
		} else {
			year, month, day = julianDayToJulianDate(jd)
		}
	}
	return
}
