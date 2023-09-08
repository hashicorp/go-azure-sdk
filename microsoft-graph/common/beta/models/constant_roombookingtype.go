package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RoomBookingType string

const (
	RoomBookingTypereserved RoomBookingType = "Reserved"
	RoomBookingTypestandard RoomBookingType = "Standard"
	RoomBookingTypeunknown  RoomBookingType = "Unknown"
)

func PossibleValuesForRoomBookingType() []string {
	return []string{
		string(RoomBookingTypereserved),
		string(RoomBookingTypestandard),
		string(RoomBookingTypeunknown),
	}
}

func parseRoomBookingType(input string) (*RoomBookingType, error) {
	vals := map[string]RoomBookingType{
		"reserved": RoomBookingTypereserved,
		"standard": RoomBookingTypestandard,
		"unknown":  RoomBookingTypeunknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := RoomBookingType(input)
	return &out, nil
}
