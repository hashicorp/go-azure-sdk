package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SwapShiftsChangeRequestState string

const (
	SwapShiftsChangeRequestStateapproved SwapShiftsChangeRequestState = "Approved"
	SwapShiftsChangeRequestStatedeclined SwapShiftsChangeRequestState = "Declined"
	SwapShiftsChangeRequestStatepending  SwapShiftsChangeRequestState = "Pending"
)

func PossibleValuesForSwapShiftsChangeRequestState() []string {
	return []string{
		string(SwapShiftsChangeRequestStateapproved),
		string(SwapShiftsChangeRequestStatedeclined),
		string(SwapShiftsChangeRequestStatepending),
	}
}

func parseSwapShiftsChangeRequestState(input string) (*SwapShiftsChangeRequestState, error) {
	vals := map[string]SwapShiftsChangeRequestState{
		"approved": SwapShiftsChangeRequestStateapproved,
		"declined": SwapShiftsChangeRequestStatedeclined,
		"pending":  SwapShiftsChangeRequestStatepending,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SwapShiftsChangeRequestState(input)
	return &out, nil
}
