package common

import (
	log "github.com/astaxie/beego/logs"
	"fmt"
	"testing"
	"time"
)

func TestGetTodayBeginEnd(t *testing.T) {
	begin, end := GetDayBeginEnd(time.Date(2018, 1, 2, 2, 6, 8, 0, time.Local))
	log.Info(fmt.Sprintf("begin: %s, end: %s", begin, end))
}

func TestGetMonthBeginEnd(t *testing.T) {
	begin, end := GetMonthBeginEnd(time.Date(2018, 2, 2, 2, 6, 8, 0, time.Local))
	log.Info(fmt.Sprintf("begin: %s, end: %s", begin, end))
}

func TestGetLastMonthBeginEnd(t *testing.T) {
	begin, end := GetLastMonthBeginEnd(time.Date(2018, 12, 2, 2, 6, 8, 0, time.Local))
	log.Info(fmt.Sprintf("begin: %s, end: %s", begin, end))
}

func TestGetWeekBeginEnd(t *testing.T) {
	str := GetWeekBegin(time.Date(2018, 11, 1, 2, 6, 8, 0, time.Local))
	log.Info(fmt.Sprintf("week: %s", str))
}
