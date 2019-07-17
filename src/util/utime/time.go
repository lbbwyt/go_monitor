package utime

import (
	log "github.com/astaxie/beego/logs"
	"fmt"
	"math"
	"time"
)

var WeekDayMap = map[string]int{
	"Monday":    1,
	"Tuesday":   2,
	"Wednesday": 3,
	"Thursday":  4,
	"Friday":    5,
	"Saturday":  6,
	"Sunday":    0,
}

var MonthMap = map[string]int{
	"January":   1,
	"February":  2,
	"March":     3,
	"April":     4,
	"May":       5,
	"June":      6,
	"July":      7,
	"August":    8,
	"September": 9,
	"October":   10,
	"November":  11,
	"December":  12,
}

var OffsetTime int64 = 0

//设置补偿时间
func SetOffsetTime(t int64) {
	OffsetTime = t
}

//标准时间转时间戳
func Date2Unix(Y int, M int, D int, H int, I int, S int) int64 {
	str := fmt.Sprintf("%d-%02d-%02d %02d:%02d:%02d", Y, M, D, H, I, S)
	return DateStr2Unix(str)
}

//错误返回-1
func DateStr2Unix(str string) int64 {
	timeLayout := "2006-01-02 15:04:05"
	_, offset := GetNow().Zone()
	t, err := time.Parse(timeLayout, str)
	if err != nil {
		log.Error("转换标准时间出错")
		return -1
	}
	return t.Unix() - int64(offset)
}

//获取当前时间戳
func GetNowUnix() int64 {
	return GetNow().Unix()
}

//
func GetNowUnixNano() int64 {
	return GetNow().UnixNano()
}

//获取当前时间
func GetNow() time.Time {
	now := time.Now().Add(time.Duration(OffsetTime) * time.Second)
	return now
}

//获取当天零点
func GetDayDot() int64 {
	_, offset := GetNow().Zone()
	t := (GetNow().Unix()/86400)*86400 - int64(offset)
	return t
}

// 获取本周第一天 (星期一) format:2006-01-02 15:04:05
func GetWeekFirstTime() time.Time {
	now := GetNow()
	year, month, day := now.Date()
	todayBegin := time.Date(year, month, day, 0, 0, 0, 0, now.Location())
	weekday := int(todayBegin.Weekday()) // 0 ~ 6

	if weekday < 1 {
		weekday = weekday + 6
	} else {
		weekday = weekday - 1
	}
	return todayBegin.AddDate(0, 0, -weekday)
}

// 获取本月第一天 format:2006-01-02 15:04:05
func GetMonthFirstTime() time.Time {
	now := GetNow()
	year, month, _ := now.Date()
	return time.Date(year, month, 1, 0, 0, 0, 0, now.Location())
}

//获取月开始
func GetMonthStartTime(i int) time.Time {
	now := GetNow()
	year, month, _ := now.Date()
	return time.Date(year, month+time.Month(i), 1, 0, 0, 0, 0, now.Location())
}

// 获取上月第一天 format:2006-01-02 15:04:05
func GetLastMonthFirstTime() time.Time {
	return GetMonthFirstTime().AddDate(0, -1, 0)
}

func WeekDay() int {
	wd := GetNow().Weekday().String()
	return WeekDayMap[wd]
}

func WeekOfYear() (Year int, Week int) {
	w := WeekDay()
	m := 0
	if w == 0 {
		m = 6
	} else {
		m = w - 1
	}
	d1 := GetNowUnix() - int64(m*86400)
	tm := time.Unix(d1, 0)
	return tm.ISOWeek()
}

//是否同一天
func IsSameDay(time1 int64, time2 int64) bool {
	tm1 := time.Unix(time1, 0)
	tm2 := time.Unix(time2, 0)
	y1, m1, d1 := tm1.Date()
	y2, m2, d2 := tm2.Date()
	fmt.Println(y1, m1, d1, y2, m2, d2)
	return y1 == y2 && m1 == m2 && d1 == d2
}

//下一个时间点，小时
func NextHour(h int64, m int64, s int64) int {
	n2 := int(h*3600 + m*60 + s)
	return NextHourSec(n2)
}
func NextHourSec(n2 int) int {
	th := GetNow().Hour()
	tm := GetNow().Minute()
	ts := GetNow().Second()
	n1 := th*3600 + tm*60 + ts
	if n1 >= n2 {
		return n2 + 86400 - n1
	} else {
		return n2 - n1
	}
}

//下一个时间点，分钟
func NextMin(m int64, s int64) int {
	n2 := int(m*60 + s)
	return NextMinSec(n2)
}

//下一个时间点，分钟
func NextMinSec(n2 int) int {
	tm := GetNow().Minute()
	ts := GetNow().Second()
	n1 := tm*60 + ts
	if n1 >= n2 {
		return n2 + 3600 - n1
	} else {
		return n2 - n1
	}
}

//间隔天数
func IntervalDay(t1, t2 int64) int64 {
	if t1 > t2 {
		t3 := t1
		t1 = t2
		t2 = t3
	}
	tm1 := time.Unix(t1, 0)
	tm2 := time.Unix(t2, 0)
	y1 := tm1.Year()
	m1 := tm1.Month().String()
	d1 := tm1.Day()
	y2 := tm2.Year()
	m2 := tm2.Month().String()
	d2 := tm2.Day()
	u1 := Date2Unix(y1, MonthMap[m1], d1, 0, 0, 0)
	u2 := Date2Unix(y2, MonthMap[m2], d2, 0, 0, 0)
	f := math.Floor((float64(u2) - float64(u1)) / 86400)
	return int64(f)
}

func GetDateTime() (y int, m int, d int, h int, i int, s int) {
	now := GetNow()
	y, M, d := now.Date()
	m = MonthMap[M.String()]
	h = now.Hour()
	i = now.Minute()
	s = now.Second()
	return
}

//获取时间零点
func GetDayZore(i int64) int64 {
	_, offset := GetNow().Zone()
	t := (GetNow().Unix()/86400)*86400 - int64(offset)
	return t
}

//标准时间转时间戳
func Unix2YearMonth(timestamp int64) int32 {
	tm := time.Unix(timestamp, 0)
	y, M, _ := tm.Date()
	m := MonthMap[M.String()]
	return int32(y)*100 + int32(m)
}

func Unix2YearMonthList(st, et int64) []int32 {
	if st > et {
		et, st = st, et
	}
	endMonth := Unix2YearMonth(et)
	startMonth := Unix2YearMonth(st)
	list := make([]int32, 0)
	list = append(list, startMonth)
	if endMonth == startMonth {
		return list
	}
	for i := 1; i <= 100; i++ {
		cur := startMonth % 100
		if cur == 12 {
			startMonth = startMonth + 100 - 11
			list = append(list, startMonth)
		} else {
			startMonth = startMonth + 1
			list = append(list, startMonth)
		}
		if startMonth == endMonth {
			return list
		}
	}
	return list
}
