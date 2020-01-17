package main

import "time"

func NewDate(dateStr string) time.Time {
	ans, _ := time.Parse("2006-01-02", dateStr)
	return ans
}
