package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SubjectRightsRequestStatus string

const (
	SubjectRightsRequestStatusactive SubjectRightsRequestStatus = "Active"
	SubjectRightsRequestStatusclosed SubjectRightsRequestStatus = "Closed"
)

func PossibleValuesForSubjectRightsRequestStatus() []string {
	return []string{
		string(SubjectRightsRequestStatusactive),
		string(SubjectRightsRequestStatusclosed),
	}
}

func parseSubjectRightsRequestStatus(input string) (*SubjectRightsRequestStatus, error) {
	vals := map[string]SubjectRightsRequestStatus{
		"active": SubjectRightsRequestStatusactive,
		"closed": SubjectRightsRequestStatusclosed,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SubjectRightsRequestStatus(input)
	return &out, nil
}
