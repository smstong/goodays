package main

import (
	"fmt"
)

func reverse(s string) string {
	var r string
	for _, v := range s {
		r = string(v) + r
	}
	return r
}

func leapYear(year int) bool {
	if year%400 == 0 {
		return true
	}
	if year%100 == 0 {
		return false
	}
	if year%4 == 0 {
		return true
	}
	return false
}

func maxDayOfMonth(year, month int) int {
	if leapYear(year) && month == 2 {
		return 29
	}
	if !leapYear(year) && month == 2 {
		return 28
	}
	switch month {
	case 1, 3, 5, 7, 8, 10, 12:
		return 31
	default:
		return 30
	}
}

type Date struct {
	Year, Month, Day int
}

func main() {
	year := 1
	month := 1
	day := 1
	var date string
	var maxDay int
	var goodDays []Date

	for year = 1; year <= 9999; year++ {
		for month = 1; month <= 12; month++ {
			maxDay = maxDayOfMonth(year, month)
			for day = 1; day <= maxDay; day++ {
				date = fmt.Sprintf("%04d%02d%02d", year, month, day)
				if date == reverse(date) {
					goodDays = append(goodDays, Date{year, month, day})
				}
			}
		}
	}
	for i := 0; i < len(goodDays)-1; i++ {
		diffYears := goodDays[i+1].Year - goodDays[i].Year
		fmt.Printf("%04d-%02d-%02d: next: %d years\n",
			goodDays[i].Year, goodDays[i].Month, goodDays[i].Day, diffYears)
	}
}
