package common

import "strconv"

func Int642String(param int64) string {
	return strconv.FormatInt(param, 10)
}

func Int2String(param int) string {
	return strconv.Itoa(param)
}
