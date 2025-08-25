package booking

import (
    "time"
    "fmt"
)

// Schedule returns a time.Time from a string containing a date.
func Schedule(date string) time.Time {
    parsed, _ := time.Parse("1/02/2006 15:04:05", date)
	return parsed
}

// HasPassed returns whether a date has passed.
func HasPassed(date string) bool {
	appointmentTime, _ := time.Parse("January 2, 2006 15:04:05", date)
    return appointmentTime.Before(time.Now())
}

// IsAfternoonAppointment returns whether a time is in the afternoon.
func IsAfternoonAppointment(date string) bool {
	appointmentTime, _ := time.Parse("Monday, January 2, 2006 15:04:05", date)
    return 12 <= appointmentTime.Hour() && appointmentTime.Hour() < 18
}

// Description returns a formatted string of the appointment time.
func Description(date string) string {
	appointmentTime, _ := time.Parse("1/2/2006 15:04:05", date)
    properFormat := appointmentTime.Format("Monday, January 2, 2006, at 15:04")
    return fmt.Sprintf("You have an appointment on %s.", properFormat)
}

// AnniversaryDate returns a Time with this year's anniversary.
func AnniversaryDate() time.Time {
	return time.Date(time.Now().Year(), time.September, 15, 0, 0, 0, 0, time.UTC)
}
