package timekit

import (
	"strconv"
	"time"

	"github.com/dannav/hhmmss"
	"github.com/relvacode/iso8601"
)

// ParseJavaScriptTime will convert the number of milliseconds since the Unix Epoch parameter into Golang `time` format. As a result, the output of the JavaScript `getTime()` function can be used as the parameter in this function.
func ParseJavaScriptTime(i int64) time.Time {
	// Special thanks JavaScript timestamp to golang time.Time
	// https://gist.github.com/alextanhongpin/3b6b2ee47665ac9c1c32c805b86380a6
	return time.Unix(i/1000, (i%1000)*1000*1000)
}

// ParseJavaScriptTimeString will convert the string of milliseconds integers since the Unix Epoch parameter into Golang `time` format. As a result, the output of the JavaScript `getTime()` function can be used as the parameter in this function.
func ParseJavaScriptTimeString(s string) (time.Time, error) {
	i, valErr := strconv.ParseInt(s, 10, 64)
	if valErr != nil {
		return time.Now(), valErr
	}

	// Special thanks JavaScript timestamp to golang time.Time
	// https://gist.github.com/alextanhongpin/3b6b2ee47665ac9c1c32c805b86380a6
	return time.Unix(i/1000, (i%1000)*1000*1000), nil
}

// ToJavaScriptTime will return a Unix Epoch time value that your JavaScript code can read into JavaScript `Date` format. Example JavaScript code snippet of using the results of this function: `var date = new Date(UNIX_Timestamp * 1000);` as an example.
func ToJavaScriptTime(t time.Time) int64 {
	return t.Unix()
}

// ToISO8601String will convert the Golang `Date` format into an ISO 8601 formatted date/time string.
func ToISO8601String(t time.Time) string {
	return t.Format(time.RFC3339) // "How to convert ISO 8601 time in golang?" via https://stackoverflow.com/a/42217963
}

// ParseISO8601String converts ISO8601 compliant date-time string into a Golang `time.Time` object.
func ParseISO8601String(s string) (time.Time, error) {
	// Note: https://stackoverflow.com/q/38596079
	return iso8601.ParseString(s)
}

// ParseBubbleTime will convert the date/time string (ex: "Nov 11, 2011 11:00 am") used "https://bubble.io" into Golang `time`. You will find need of this function if the Bubble.io app you built will be making an API call to your Golang backend server.
func ParseBubbleTime(s string) (time.Time, error) {
	// Note: https://www.geeksforgeeks.org/time-formatting-in-golang/
	return time.Parse("Jan _2, 2006 15:04 am", s)
}

// ParseHourMinuteSecondDurationString will convert a HH:MM:SS string (example: "08:30:00") into duration.
func ParseHourMinuteSecondDurationString(s string) (time.Duration, error) {
	// Note: https://github.com/dannav/hhmmss
	return hhmmss.Parse(s)
}

// ToAmericanDateTimeString will convert the Golang Date/Time format into the American style of notation string as mentioned via https://en.wikipedia.org/wiki/Date_and_time_notation_in_the_United_States.
func ToAmericanDateTimeString(t time.Time) string {
	return t.Format("January 2, 2006 3:04:05 PM")
}

// ToAmericanDateString will convert the Golang Date/Time format into the American style of notation string as mentioned via https://en.wikipedia.org/wiki/Date_and_time_notation_in_the_United_States.
func ToAmericanDateString(t time.Time) string {
	return t.Format("January 2, 2006")
}

// To1AM will take entered date/time and return same date but time starts at 1 AM (01:00:00).
func To1AM(t time.Time) time.Time {
	// Create a new time.Time with the same date but set the time to 1 AM (01:00:00).
	oneAM := time.Date(t.Year(), t.Month(), t.Day(), 1, 0, 0, 0, t.Location())
	return oneAM
}

// GetMonthAbbreviation returns the 3-character abbreviation for the provided month.
func GetMonthAbbreviation(month time.Month) string {
	abbreviation, found := monthAbbreviations[month]
	if !found {
		return ""
	}
	return abbreviation
}

// GetMonthAbbreviationByInt returns the 3-character abbreviation for the provided month number.
func GetMonthAbbreviationByInt(month int) string {
	abbreviation, found := monthNumberAbbreviations[month]
	if !found {
		return ""
	}
	return abbreviation
}
