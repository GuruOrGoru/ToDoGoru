package main

import (
	"strconv"
	"time"
)

func ConvertStringToBool(str string) (bool, error) {
	return strconv.ParseBool(str)
}

func ConvertStringToInt(str string) (int, error) {
	return strconv.Atoi(str)
}

func ConvertStringToTime(str string) (time.Time, error) {
	layout := "2006-01-02 15:04:05"
	return time.Parse(layout, str)
}
