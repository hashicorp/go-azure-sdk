package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type BookingWorkHoursDay string

const (
	BookingWorkHoursDay_Friday    BookingWorkHoursDay = "friday"
	BookingWorkHoursDay_Monday    BookingWorkHoursDay = "monday"
	BookingWorkHoursDay_Saturday  BookingWorkHoursDay = "saturday"
	BookingWorkHoursDay_Sunday    BookingWorkHoursDay = "sunday"
	BookingWorkHoursDay_Thursday  BookingWorkHoursDay = "thursday"
	BookingWorkHoursDay_Tuesday   BookingWorkHoursDay = "tuesday"
	BookingWorkHoursDay_Wednesday BookingWorkHoursDay = "wednesday"
)

func PossibleValuesForBookingWorkHoursDay() []string {
	return []string{
		string(BookingWorkHoursDay_Friday),
		string(BookingWorkHoursDay_Monday),
		string(BookingWorkHoursDay_Saturday),
		string(BookingWorkHoursDay_Sunday),
		string(BookingWorkHoursDay_Thursday),
		string(BookingWorkHoursDay_Tuesday),
		string(BookingWorkHoursDay_Wednesday),
	}
}

func (s *BookingWorkHoursDay) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseBookingWorkHoursDay(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseBookingWorkHoursDay(input string) (*BookingWorkHoursDay, error) {
	vals := map[string]BookingWorkHoursDay{
		"friday":    BookingWorkHoursDay_Friday,
		"monday":    BookingWorkHoursDay_Monday,
		"saturday":  BookingWorkHoursDay_Saturday,
		"sunday":    BookingWorkHoursDay_Sunday,
		"thursday":  BookingWorkHoursDay_Thursday,
		"tuesday":   BookingWorkHoursDay_Tuesday,
		"wednesday": BookingWorkHoursDay_Wednesday,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := BookingWorkHoursDay(input)
	return &out, nil
}
