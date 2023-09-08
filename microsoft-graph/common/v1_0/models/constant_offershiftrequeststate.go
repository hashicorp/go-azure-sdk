package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type OfferShiftRequestState string

const (
	OfferShiftRequestStateapproved OfferShiftRequestState = "Approved"
	OfferShiftRequestStatedeclined OfferShiftRequestState = "Declined"
	OfferShiftRequestStatepending  OfferShiftRequestState = "Pending"
)

func PossibleValuesForOfferShiftRequestState() []string {
	return []string{
		string(OfferShiftRequestStateapproved),
		string(OfferShiftRequestStatedeclined),
		string(OfferShiftRequestStatepending),
	}
}

func parseOfferShiftRequestState(input string) (*OfferShiftRequestState, error) {
	vals := map[string]OfferShiftRequestState{
		"approved": OfferShiftRequestStateapproved,
		"declined": OfferShiftRequestStatedeclined,
		"pending":  OfferShiftRequestStatepending,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := OfferShiftRequestState(input)
	return &out, nil
}
