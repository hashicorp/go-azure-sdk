package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AvailabilityItemStatus string

const (
	AvailabilityItemStatusavailable      AvailabilityItemStatus = "Available"
	AvailabilityItemStatusbusy           AvailabilityItemStatus = "Busy"
	AvailabilityItemStatusoutOfOffice    AvailabilityItemStatus = "OutOfOffice"
	AvailabilityItemStatusslotsAvailable AvailabilityItemStatus = "SlotsAvailable"
)

func PossibleValuesForAvailabilityItemStatus() []string {
	return []string{
		string(AvailabilityItemStatusavailable),
		string(AvailabilityItemStatusbusy),
		string(AvailabilityItemStatusoutOfOffice),
		string(AvailabilityItemStatusslotsAvailable),
	}
}

func parseAvailabilityItemStatus(input string) (*AvailabilityItemStatus, error) {
	vals := map[string]AvailabilityItemStatus{
		"available":      AvailabilityItemStatusavailable,
		"busy":           AvailabilityItemStatusbusy,
		"outofoffice":    AvailabilityItemStatusoutOfOffice,
		"slotsavailable": AvailabilityItemStatusslotsAvailable,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AvailabilityItemStatus(input)
	return &out, nil
}
