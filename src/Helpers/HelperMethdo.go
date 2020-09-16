package Helpers

import "strings"

func CleanDate(CreatedAt string) string {
	month := strings.Split(CreatedAt, " ")[1]
	day := strings.Split(CreatedAt, " ")[2]
	year := strings.Split(CreatedAt, " ")[5]
	hour := strings.Split(CreatedAt, " ")[3]

	if month == "Jan" {
		month = "01"
	} else if month == "Feb" {
		month = "02"
	} else if month == "Mar" {
		month = "03"
	} else if month == "Apr" {
		month = "04"
	} else if month == "May" {
		month = "05"
	} else if month == "Jun" {
		month = "06"
	} else if month == "Jul" {
		month = "07"
	} else if month == "Aug" {
		month = "08"
	} else if month == "Sep" {
		month = "09"
	} else if month == "Oct" {
		month = "10"
	} else if month == "Nov" {
		month = "11"
	} else if month == "Dec" {
		month = "12"
	}
	full_date := year + "-" + month + "-" + day + " " + hour
	return full_date
}

func SubstrDate(CreatedAt string) string {
	//20200913231435
	a := []rune(CreatedAt)
	year := string(a[0:4])
	month := string(a[4:6])
	day := string(a[6:8])
	hour := string(a[8:10])
	min := string(a[10:12])

	full_date := year + "-" + month + "-" + day + " " + hour + ":" + min
	return full_date
}

func SubstrDate2(CreatedAt string) string {
	//20200913231435
	a := []rune(CreatedAt)
	//year := string(a[0:4])
	month := string(a[4:6])
	day := string(a[6:8])
	full_date := month + "월" + day + "일"
	return full_date
}
