package main

import "time"

func StrToTime(timeStr string) time.Time {
	t, _ := time.Parse(time.RFC3339, timeStr)
	return t
}

func StrToTimePtr(timeStr string) *time.Time {
	t := StrToTime(timeStr)
	return &t
}
