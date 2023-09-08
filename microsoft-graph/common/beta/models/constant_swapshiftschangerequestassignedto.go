package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SwapShiftsChangeRequestAssignedTo string

const (
	SwapShiftsChangeRequestAssignedTomanager   SwapShiftsChangeRequestAssignedTo = "Manager"
	SwapShiftsChangeRequestAssignedTorecipient SwapShiftsChangeRequestAssignedTo = "Recipient"
	SwapShiftsChangeRequestAssignedTosender    SwapShiftsChangeRequestAssignedTo = "Sender"
	SwapShiftsChangeRequestAssignedTosystem    SwapShiftsChangeRequestAssignedTo = "System"
)

func PossibleValuesForSwapShiftsChangeRequestAssignedTo() []string {
	return []string{
		string(SwapShiftsChangeRequestAssignedTomanager),
		string(SwapShiftsChangeRequestAssignedTorecipient),
		string(SwapShiftsChangeRequestAssignedTosender),
		string(SwapShiftsChangeRequestAssignedTosystem),
	}
}

func parseSwapShiftsChangeRequestAssignedTo(input string) (*SwapShiftsChangeRequestAssignedTo, error) {
	vals := map[string]SwapShiftsChangeRequestAssignedTo{
		"manager":   SwapShiftsChangeRequestAssignedTomanager,
		"recipient": SwapShiftsChangeRequestAssignedTorecipient,
		"sender":    SwapShiftsChangeRequestAssignedTosender,
		"system":    SwapShiftsChangeRequestAssignedTosystem,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SwapShiftsChangeRequestAssignedTo(input)
	return &out, nil
}
