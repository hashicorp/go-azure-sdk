package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PayloadCompatibleAssignmentFilterPayloadType string

const (
	PayloadCompatibleAssignmentFilterPayloadTypeenrollmentRestrictions PayloadCompatibleAssignmentFilterPayloadType = "EnrollmentRestrictions"
	PayloadCompatibleAssignmentFilterPayloadTypenotSet                 PayloadCompatibleAssignmentFilterPayloadType = "NotSet"
)

func PossibleValuesForPayloadCompatibleAssignmentFilterPayloadType() []string {
	return []string{
		string(PayloadCompatibleAssignmentFilterPayloadTypeenrollmentRestrictions),
		string(PayloadCompatibleAssignmentFilterPayloadTypenotSet),
	}
}

func parsePayloadCompatibleAssignmentFilterPayloadType(input string) (*PayloadCompatibleAssignmentFilterPayloadType, error) {
	vals := map[string]PayloadCompatibleAssignmentFilterPayloadType{
		"enrollmentrestrictions": PayloadCompatibleAssignmentFilterPayloadTypeenrollmentRestrictions,
		"notset":                 PayloadCompatibleAssignmentFilterPayloadTypenotSet,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := PayloadCompatibleAssignmentFilterPayloadType(input)
	return &out, nil
}
