package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PermissionGrantConditionSetPermissionType string

const (
	PermissionGrantConditionSetPermissionTypeapplication              PermissionGrantConditionSetPermissionType = "Application"
	PermissionGrantConditionSetPermissionTypedelegated                PermissionGrantConditionSetPermissionType = "Delegated"
	PermissionGrantConditionSetPermissionTypedelegatedUserConsentable PermissionGrantConditionSetPermissionType = "DelegatedUserConsentable"
)

func PossibleValuesForPermissionGrantConditionSetPermissionType() []string {
	return []string{
		string(PermissionGrantConditionSetPermissionTypeapplication),
		string(PermissionGrantConditionSetPermissionTypedelegated),
		string(PermissionGrantConditionSetPermissionTypedelegatedUserConsentable),
	}
}

func parsePermissionGrantConditionSetPermissionType(input string) (*PermissionGrantConditionSetPermissionType, error) {
	vals := map[string]PermissionGrantConditionSetPermissionType{
		"application":              PermissionGrantConditionSetPermissionTypeapplication,
		"delegated":                PermissionGrantConditionSetPermissionTypedelegated,
		"delegateduserconsentable": PermissionGrantConditionSetPermissionTypedelegatedUserConsentable,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := PermissionGrantConditionSetPermissionType(input)
	return &out, nil
}
