// Have the function CountingMinutes(str) take the str parameter being passed which will be
// two times (each properly formatted with a colon and am or pm) separated by a hyphen and
// return the total number of minutes between the two times.

// The time will be in a 12 hour clock format. For example: if str is 9:00am-10:00am then the output should be 60.
// If str is 1:00pm-11:00am the output should be 1320.

// Input: "12:30pm-12:00am"
// Output: 690
// Input: "1:23am-1:08am"
// Output: 1425

package main

import (
	"fmt"
	"strconv"
	"strings"
)

type Time struct {
	Minutes  int
	Hours    int
	Meridien string
}

func CountingMinutes(str string) string {

	// code goes here
	return str
}

func parseTime(str string) Time {
	isAm := strings.Contains(str, "am")
	meridien := "am"
	hours, _ := strconv.Atoi(str[0:2])
	minutes, _ := strconv.Atoi(str[3:5])

	if !isAm {
		meridien = "pm"
	}

	fmt.Println(hours)
	fmt.Println(minutes)

	return Time{
		Minutes:  minutes,
		Hours:    hours,
		Meridien: meridien,
	}
}

func main() {
	str := "12:30pm-12:00am"
	hours := strings.Split(str, "-")
	firstHour := parseTime(hours[0])
	secondHour := parseTime(hours[1])
	var totalMinutes int

	if firstHour.Meridien == secondHour.Meridien {

	}
	fmt.Println(totalMinutes)
}
