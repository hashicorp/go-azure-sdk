package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type OpenShiftChangeRequestAssignedTo string

const (
	OpenShiftChangeRequestAssignedTomanager   OpenShiftChangeRequestAssignedTo = "Manager"
	OpenShiftChangeRequestAssignedTorecipient OpenShiftChangeRequestAssignedTo = "Recipient"
	OpenShiftChangeRequestAssignedTosender    OpenShiftChangeRequestAssignedTo = "Sender"
	OpenShiftChangeRequestAssignedTosystem    OpenShiftChangeRequestAssignedTo = "System"
)

func PossibleValuesForOpenShiftChangeRequestAssignedTo() []string {
	return []string{
		string(OpenShiftChangeRequestAssignedTomanager),
		string(OpenShiftChangeRequestAssignedTorecipient),
		string(OpenShiftChangeRequestAssignedTosender),
		string(OpenShiftChangeRequestAssignedTosystem),
	}
}

func parseOpenShiftChangeRequestAssignedTo(input string) (*OpenShiftChangeRequestAssignedTo, error) {
	vals := map[string]OpenShiftChangeRequestAssignedTo{
		"manager":   OpenShiftChangeRequestAssignedTomanager,
		"recipient": OpenShiftChangeRequestAssignedTorecipient,
		"sender":    OpenShiftChangeRequestAssignedTosender,
		"system":    OpenShiftChangeRequestAssignedTosystem,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := OpenShiftChangeRequestAssignedTo(input)
	return &out, nil
}
