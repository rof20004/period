package period

import "time"

const (
	daysPerCycle   = 146097
	days0000To1970 = (daysPerCycle * 5) - (30*365 + 7)
)

type Period struct {
	Years  int
	Months int
	Days   int
}

func newPeriod(years, months, days int) Period {
	return Period{years, months, days}
}

func Between(start, end time.Time) Period {
	startYear, startMonth, startDay := start.Date()
	endYear, endMonth, endDay := end.Date()

	sMonth := int(startMonth)
	eMonth := int(endMonth)

	var totalMonths = getProlepticMonth(endYear, eMonth) - getProlepticMonth(startYear, sMonth)

	var days = endDay - startDay

	if totalMonths > 0 && days < 0 {
		totalMonths--
		var calcDate = start.AddDate(0, totalMonths, 0)
		calcYear, calcMont, calcDay := calcDate.Date()
		cMonth := int(calcMont)
		days = toEpochDay(endYear, eMonth, endDay) - toEpochDay(calcYear, cMonth, calcDay)
	} else if totalMonths < 0 && days > 0 {
		totalMonths++
		days -= lengthOfMonth(endYear, eMonth)
	}

	var years = totalMonths / 12
	var months = totalMonths % 12

	return newPeriod(years, months, days)
}

func getProlepticMonth(year, month int) int {
	return year*12 + month - 1
}

func toEpochDay(year, month, day int) int {
	var y = year
	var m = month
	var total = 365 * y
	if y >= 0 {
		total += (y+3)/4 - (y+99)/100 + (y+399)/400
	} else {
		total -= y/-4 - y/-100 + y/-400
	}
	total += (367*m - 362) / 12
	total += day - 1
	if m > 2 {
		total--
		if !isLeapYear(year) {
			total--
		}
	}
	return total - days0000To1970
}

func isLeapYear(year int) bool {
	return year%4 == 0 && year%100 != 0 || year%400 == 0
}

func lengthOfMonth(year, month int) int {
	switch month {
	case 2:
		if isLeapYear(year) {
			return 29
		}
		return 28
	case 4, 6, 9, 11:
		return 30
	default:
		return 31
	}
}
