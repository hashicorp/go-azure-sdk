package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TimeOffRequestState string

const (
	TimeOffRequestStateapproved TimeOffRequestState = "Approved"
	TimeOffRequestStatedeclined TimeOffRequestState = "Declined"
	TimeOffRequestStatepending  TimeOffRequestState = "Pending"
)

func PossibleValuesForTimeOffRequestState() []string {
	return []string{
		string(TimeOffRequestStateapproved),
		string(TimeOffRequestStatedeclined),
		string(TimeOffRequestStatepending),
	}
}

func parseTimeOffRequestState(input string) (*TimeOffRequestState, error) {
	vals := map[string]TimeOffRequestState{
		"approved": TimeOffRequestStateapproved,
		"declined": TimeOffRequestStatedeclined,
		"pending":  TimeOffRequestStatepending,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := TimeOffRequestState(input)
	return &out, nil
}
