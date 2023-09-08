package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type BookingWorkHoursDay string

const (
	BookingWorkHoursDayfriday    BookingWorkHoursDay = "Friday"
	BookingWorkHoursDaymonday    BookingWorkHoursDay = "Monday"
	BookingWorkHoursDaysaturday  BookingWorkHoursDay = "Saturday"
	BookingWorkHoursDaysunday    BookingWorkHoursDay = "Sunday"
	BookingWorkHoursDaythursday  BookingWorkHoursDay = "Thursday"
	BookingWorkHoursDaytuesday   BookingWorkHoursDay = "Tuesday"
	BookingWorkHoursDaywednesday BookingWorkHoursDay = "Wednesday"
)

func PossibleValuesForBookingWorkHoursDay() []string {
	return []string{
		string(BookingWorkHoursDayfriday),
		string(BookingWorkHoursDaymonday),
		string(BookingWorkHoursDaysaturday),
		string(BookingWorkHoursDaysunday),
		string(BookingWorkHoursDaythursday),
		string(BookingWorkHoursDaytuesday),
		string(BookingWorkHoursDaywednesday),
	}
}

func parseBookingWorkHoursDay(input string) (*BookingWorkHoursDay, error) {
	vals := map[string]BookingWorkHoursDay{
		"friday":    BookingWorkHoursDayfriday,
		"monday":    BookingWorkHoursDaymonday,
		"saturday":  BookingWorkHoursDaysaturday,
		"sunday":    BookingWorkHoursDaysunday,
		"thursday":  BookingWorkHoursDaythursday,
		"tuesday":   BookingWorkHoursDaytuesday,
		"wednesday": BookingWorkHoursDaywednesday,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := BookingWorkHoursDay(input)
	return &out, nil
}
