package utils

import "strconv"

// todo: move to config in yaml
// todo: singleton pattern

type TimeInfo struct {
	StartYear  int
	StartMonth int
}

func (this *TimeInfo) CalendarMth(t int) (res int) {
	num := (this.StartMonth + t) % 12
	if num == 0 {
		res = 12
	} else {
		res = num
	}
	return
}

func (this *TimeInfo) CalendarYr(t int) (res int) {
	if this.StartMonth == 12 {
		res = this.StartYear + (t+11)/12
	} else {
		res = this.StartYear + (this.StartMonth+t-1)/12
	}
	return
}

func (this *TimeInfo) T(Year, Month int) int {
	res := (Year-this.StartYear)*12 + Month - this.StartMonth
	return res
}

func (this *TimeInfo) YYYYMM(Year, Month int) string {
	res := strconv.Itoa(Year*100 + Month)
	return res
}
