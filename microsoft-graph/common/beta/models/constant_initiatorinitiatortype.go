package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type InitiatorInitiatorType string

const (
	InitiatorInitiatorTypeapplication InitiatorInitiatorType = "Application"
	InitiatorInitiatorTypesystem      InitiatorInitiatorType = "System"
	InitiatorInitiatorTypeuser        InitiatorInitiatorType = "User"
)

func PossibleValuesForInitiatorInitiatorType() []string {
	return []string{
		string(InitiatorInitiatorTypeapplication),
		string(InitiatorInitiatorTypesystem),
		string(InitiatorInitiatorTypeuser),
	}
}

func parseInitiatorInitiatorType(input string) (*InitiatorInitiatorType, error) {
	vals := map[string]InitiatorInitiatorType{
		"application": InitiatorInitiatorTypeapplication,
		"system":      InitiatorInitiatorTypesystem,
		"user":        InitiatorInitiatorTypeuser,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := InitiatorInitiatorType(input)
	return &out, nil
}
