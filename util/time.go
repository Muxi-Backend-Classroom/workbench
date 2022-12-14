package util

import (
	"strconv"
	"time"
)

func String2Time(value string) (time.Time, error) {
	return time.Parse(time.RFC3339, value)
}

func GetCurrentTime() string {
	timestamp := time.Now().Unix()
	tm := time.Unix(timestamp, 0)
	return tm.Format(time.RFC3339)
}

func GetTimeFromTimestamp(stamp int64) string {
	t := time.Unix(stamp, 0)
	return t.Format(time.RFC3339)
}

func Time2String(t time.Time) string {
	lo, _ := time.LoadLocation("Local")
	return t.In(lo).Format(time.RFC3339)
}

// GetTimeMonth returns the time after the specified duration(month).
func GetTimeMonth(month int) string {
	currentTime := time.Now()
	res := currentTime.AddDate(0, month, 0)
	return res.Format(time.RFC3339)
}

// GetTimeDay returns the time after the specified duration(day).
func GetTimeDay(day int) string {
	currentTime := time.Now()
	res := currentTime.AddDate(0, 0, day)
	return res.Format(time.RFC3339)
}

// GetTimeMinute returns the time after the specified duration(minute).
func GetTimeMinute(minute int) string {
	currentTime := time.Now()
	m, _ := time.ParseDuration(strconv.Itoa(minute) + "m")
	res := currentTime.Add(m)
	return res.Format(time.RFC3339)
}

// GetTimeHour returns the time after the specified duration(hour).
func GetTimeHour(hour int) string {
	currentTime := time.Now()
	h, _ := time.ParseDuration(strconv.Itoa(hour) + "h")
	res := currentTime.Add(h)
	return res.Format(time.RFC3339)
}

// GetTimeYear returns the time after the specified duration(year).
func GetTimeYear(year int) string {
	currentTime := time.Now()
	res := currentTime.AddDate(year, 0, 0)
	return res.Format(time.RFC3339)
}
