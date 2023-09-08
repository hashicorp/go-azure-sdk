package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type OpenShiftChangeRequestState string

const (
	OpenShiftChangeRequestStateapproved OpenShiftChangeRequestState = "Approved"
	OpenShiftChangeRequestStatedeclined OpenShiftChangeRequestState = "Declined"
	OpenShiftChangeRequestStatepending  OpenShiftChangeRequestState = "Pending"
)

func PossibleValuesForOpenShiftChangeRequestState() []string {
	return []string{
		string(OpenShiftChangeRequestStateapproved),
		string(OpenShiftChangeRequestStatedeclined),
		string(OpenShiftChangeRequestStatepending),
	}
}

func parseOpenShiftChangeRequestState(input string) (*OpenShiftChangeRequestState, error) {
	vals := map[string]OpenShiftChangeRequestState{
		"approved": OpenShiftChangeRequestStateapproved,
		"declined": OpenShiftChangeRequestStatedeclined,
		"pending":  OpenShiftChangeRequestStatepending,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := OpenShiftChangeRequestState(input)
	return &out, nil
}
