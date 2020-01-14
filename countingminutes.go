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

type time struct {
	Minutes  int
	Hours    int
	Meridien string
}

func CountingMinutes(str string) string {

	// code goes here
	return str
}

func parseTime(str string) time {
	isAm := strings.Contains(str, "am")
	meridien := "am"
	hours, _ := strconv.Atoi(str[0:2])
	minutes, _ := strconv.Atoi(str[3:5])

	if !isAm {
		meridien = "pm"
	}

	return time{
		Minutes:  minutes,
		Hours:    hours,
		Meridien: meridien,
	}
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func main() {
	str := "02:30am-02:40am"
	hours := strings.Split(str, "-")
	firstHour := parseTime(hours[0])
	secondHour := parseTime(hours[1])
	const HourInMinutes int = 60
	const DayInMinutes int = 24 * HourInMinutes
	var totalMinutes int

	if firstHour.Meridien == secondHour.Meridien {
		if firstHour.Hours < secondHour.Hours {
			totalMinutes += (secondHour.Hours - firstHour.Hours) * HourInMinutes

			if firstHour.Minutes < secondHour.Minutes {
				totalMinutes += abs(firstHour.Minutes - secondHour.Minutes)
			} else if firstHour.Minutes > secondHour.Minutes {
				totalMinutes += (secondHour.Minutes - firstHour.Minutes)
			}
		} else if firstHour.Hours == secondHour.Hours {
			if firstHour.Minutes < secondHour.Minutes {
				totalMinutes += abs(firstHour.Minutes - secondHour.Minutes)
			} else if firstHour.Minutes > secondHour.Minutes {
				fmt.Println((secondHour.Minutes - firstHour.Minutes))
				totalMinutes += (DayInMinutes + (secondHour.Minutes - firstHour.Minutes))
			}
		} else {
			totalMinutes += DayInMinutes + (firstHour.Hours - secondHour.Hours)
		}

	}

	fmt.Println(totalMinutes)
}
