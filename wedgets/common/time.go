package common

import (
	"fmt"
	"strings"
	"time"
)

func GetNowSecond() int64 {
	return time.Now().Unix()
}

func GetNowMillisecond() int64 {
	return time.Now().UnixNano() / 1e6
}

func GetZeroOfHourTime() int64 {
	nowTime := GetNowSecond()
	return nowTime - nowTime%3600
}

func GetZeroTime() int64 {
	nowTime := GetNowSecond()
	return nowTime - nowTime%86400
}

func GetZeroTimeByTimeZone(timeZone int64) int64 {
	nowTime := GetNowSecond()
	return nowTime - (nowTime+timeZone*3600)%86400
}

// "2023-07-18 14:30:15.000+0800"
func FormatTimeWithTimeZoneOffset(t time.Time) string {
	// 使用time.Format获取不包含时区偏移量的时间字符串
	const layout = "2006-01-02 15:04:05.000"
	baseTimeStr := t.Format(layout)

	// 获取时区偏移量
	_, offset := t.Zone()
	sign := "+"
	if offset < 0 {
		sign = "-"
		offset = -offset // 转换为正数以方便处理
	}

	// 转换时区偏移量为小时和分钟
	hours := offset / 3600
	minutes := (offset % 3600) / 60

	// 格式化小时和分钟为两位数字，并添加到时间字符串的末尾
	timeZoneOffset := sign + fmt.Sprintf("%02d%02d", hours, minutes)

	// 返回完整的时间字符串，包括时区偏移量
	return baseTimeStr + timeZoneOffset
}

// Fri, 12 Jul 2013 09:13:05 GMT
func GetGMTTime() string {
	now := time.Now()
	utcTime := now.UTC()
	//1
	//return utcTime.Format("Mon, 02 Jan 2006 15:04:05 GMT")

	//2.
	formattedTime := utcTime.Format(time.RFC1123Z)
	formattedTimeGMT := strings.ReplaceAll(formattedTime, "+0000", "GMT")
	return formattedTimeGMT
}