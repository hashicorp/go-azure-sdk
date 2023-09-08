package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TeamsAppResourceSpecificPermissionPermissionType string

const (
	TeamsAppResourceSpecificPermissionPermissionTypeapplication TeamsAppResourceSpecificPermissionPermissionType = "Application"
	TeamsAppResourceSpecificPermissionPermissionTypedelegated   TeamsAppResourceSpecificPermissionPermissionType = "Delegated"
)

func PossibleValuesForTeamsAppResourceSpecificPermissionPermissionType() []string {
	return []string{
		string(TeamsAppResourceSpecificPermissionPermissionTypeapplication),
		string(TeamsAppResourceSpecificPermissionPermissionTypedelegated),
	}
}

func parseTeamsAppResourceSpecificPermissionPermissionType(input string) (*TeamsAppResourceSpecificPermissionPermissionType, error) {
	vals := map[string]TeamsAppResourceSpecificPermissionPermissionType{
		"application": TeamsAppResourceSpecificPermissionPermissionTypeapplication,
		"delegated":   TeamsAppResourceSpecificPermissionPermissionTypedelegated,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := TeamsAppResourceSpecificPermissionPermissionType(input)
	return &out, nil
}
