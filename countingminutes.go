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
	Value    int
}

const minutesInHour int = 60
const halfDayInHours int = 12
const halfDayInMinutes int = 12 * minutesInHour
const dayInMinutes int = 24 * minutesInHour

func CountingMinutes(str string) string {
	hours := strings.Split(str, "-")
	firstHour := parseTime(hours[0])
	secondHour := parseTime(hours[1])

	fmt.Println(firstHour)
	fmt.Println(secondHour)

	var totalMinutes int

	if (firstHour.Value > halfDayInMinutes && secondHour.Value > halfDayInMinutes) ||
		(firstHour.Value < halfDayInMinutes && secondHour.Value < halfDayInMinutes) {
		if firstHour.Value > secondHour.Value {
			totalMinutes += (dayInMinutes) - (firstHour.Value - secondHour.Value)
		} else {
			totalMinutes += secondHour.Value - firstHour.Value
		}
	} else {
		totalMinutes += halfDayInMinutes
		totalMinutes += (secondHour.Value + halfDayInMinutes) - firstHour.Value
	}
	text := strconv.Itoa(totalMinutes)

	return text
}

func parseTime(str string) time {
	isAm := strings.Contains(str, "am")
	index := strings.Index(str, ":")
	meridien := "am"
	hours, _ := strconv.Atoi(str[0:index])
	minutes, _ := strconv.Atoi(str[index+1 : index+3])
	var value int

	if !isAm {
		meridien = "pm"
	}

	if meridien == "am" {
		value += minutesInHour * hours
		value += minutes
	} else {
		value += minutesInHour*(hours+halfDayInHours) + minutes
	}

	return time{
		Minutes:  minutes,
		Hours:    hours,
		Meridien: meridien,
		Value:    value,
	}
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func main() {
	str := "11:00am-2:10pm"

	fmt.Println(CountingMinutes(str))
}

// 1. For input "12:30pm-12:00am" the output was incorrect. The correct output is 690

// 2. For input "1:00pm-11:00am" the output was incorrect. The correct output is 1320

// 3. For input "2:03pm-1:39pm" the output was incorrect. The correct output is 1416

// 4. For input "1:23am-1:08am" the output was incorrect. The correct output is 1425

// 5. For input "2:08pm-2:00am" the output was incorrect. The correct output is 712

// 6. For input "2:00pm-3:00pm" the output was incorrect. The correct output is 60
//
// 7. For input "11:00am-2:10pm" the output was incorrect. The correct output is 190

// 8. For input "12:31pm-12:34pm" the output was incorrect. The correct output is 3

// 9. For input "3:00pm-4:45am" the output was incorrect. The correct output is 825

// 10. For input "5:00pm-5:11pm" the output was incorrect. The correct output is 11
