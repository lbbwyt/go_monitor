package utime

import (
	"testing"
)

func TestGetDayDot(t *testing.T) {
	d := GetDayDot()
	if d == 0 {
		t.Error("测试失败", d)
		return
	}
	t.Log("测试成功", d)
}

func TestGetNowUnix(t *testing.T) {
	d := GetNowUnix()
	if d == 0 {
		t.Error("测试失败", d)
		return
	}
	t.Log("测试成功", d)
}

func TestGetMonthFirstTime(t *testing.T) {
	d := GetMonthFirstTime()
	t.Log(d)
	t.Log(GetLastMonthFirstTime())
}

func TestDate2Unix(t *testing.T) {
	d := Date2Unix(2018, 8, 12, 9, 12, 50)
	t.Log("结果", d)
}

func TestWeekOfYear(t *testing.T) {
	y, w := WeekOfYear()
	t.Log("结果", y, w)
}

func TestWeekDay(t *testing.T) {
	w := WeekDay()
	t.Log("结果", w)
}

func TestIsSameDay(t *testing.T) {
	t.Log(IsSameDay(1540870970, 1538502170))
}

func TestNextHour(t *testing.T) {
	t.Log(NextHour(11, 0, 0))
}
func TestNextMin(t *testing.T) {
	t.Log(NextMin(10, 0))
}
func TestIntervalDay(t *testing.T) {
	t.Log(IntervalDay(1540964927, 1540914527))
}

func TestUnix2YearMonthList() {
	t.Log(Unix2YearMonthList(1540964927, 1540914527))
}
