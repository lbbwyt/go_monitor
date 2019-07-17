package common

import (
	"fmt"
	"time"
)

// 给定一个时间，获取这个时间当天的开始时间和结束时间
func GetDayBeginEnd(date time.Time) (zore, end string) {
	year, month, day := date.Date()
	zore = time.Date(year, month, day, 0, 0, 0, 0, time.Local).Format("2006-01-02 15:04:05")
	end = time.Date(year, month, day, 23, 59, 59, 999, time.Local).Format("2006-01-02 15:04:05")
	return
}

// 给定一个时间，获取这个时间所在月份的开始日期和结束日期
func GetMonthBeginEnd(date time.Time) (begin, end string) {
	year, month, _ := date.Date()
	return fmt.Sprintf("%d-%d-1", year, month), fmt.Sprintf("%d-%d-%d", year, month, getMonthLastDay(year, int(month)))
}

// 给定一个时间，获取这个时间所在月份的前一个月的开始日期和结束日期
func GetLastMonthBeginEnd(date time.Time) (begin, end string) {
	_, _, day := date.Date()
	year, month, _ := date.AddDate(0, -1, 1 - day).Date()
	return fmt.Sprintf("%d-%d-1", year, month), fmt.Sprintf("%d-%d-%d", year, month, getMonthLastDay(year, int(month)))
}

// 给定一个时间，获取这个时间所在周的第一天（星期天）
func GetWeekBegin(date time.Time) string {
	year, month, day := date.Date()
	week := date.Weekday()
	if day > 7 || (day < 7 && day >= int(week+1)) { // 星期第一天不跨月
		return fmt.Sprintf("%d-%d-%d", year, month, day-int(week))
	} else { // 星期第一天跨月
		if month == 1 { // 跨年
			month = 12
			year--
		} else { // 不跨年
			month--
		}
		lastDay := getMonthLastDay(year, int(month))
		weekBegin := lastDay - (int(week) - day)
		return fmt.Sprintf("%d-%d-%d", year, month, weekBegin)
	}
}

// 获取月份最后一天
func getMonthLastDay(year, month int) int {
	if month == 4 || month == 6 || month == 9 || month == 11 {
		return 30
	} else if month == 2 {
		if (year%4 == 0 && year%100 != 0) || year%400 == 0 { // 润年
			return 29
		} else {
			return 28
		}
	} else {
		return 31
	}
}
