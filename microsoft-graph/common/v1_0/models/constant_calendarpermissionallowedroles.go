package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CalendarPermissionAllowedRoles string

const (
	CalendarPermissionAllowedRolescustom                            CalendarPermissionAllowedRoles = "Custom"
	CalendarPermissionAllowedRolesdelegateWithPrivateEventAccess    CalendarPermissionAllowedRoles = "DelegateWithPrivateEventAccess"
	CalendarPermissionAllowedRolesdelegateWithoutPrivateEventAccess CalendarPermissionAllowedRoles = "DelegateWithoutPrivateEventAccess"
	CalendarPermissionAllowedRolesfreeBusyRead                      CalendarPermissionAllowedRoles = "FreeBusyRead"
	CalendarPermissionAllowedRoleslimitedRead                       CalendarPermissionAllowedRoles = "LimitedRead"
	CalendarPermissionAllowedRolesnone                              CalendarPermissionAllowedRoles = "None"
	CalendarPermissionAllowedRolesread                              CalendarPermissionAllowedRoles = "Read"
	CalendarPermissionAllowedRoleswrite                             CalendarPermissionAllowedRoles = "Write"
)

func PossibleValuesForCalendarPermissionAllowedRoles() []string {
	return []string{
		string(CalendarPermissionAllowedRolescustom),
		string(CalendarPermissionAllowedRolesdelegateWithPrivateEventAccess),
		string(CalendarPermissionAllowedRolesdelegateWithoutPrivateEventAccess),
		string(CalendarPermissionAllowedRolesfreeBusyRead),
		string(CalendarPermissionAllowedRoleslimitedRead),
		string(CalendarPermissionAllowedRolesnone),
		string(CalendarPermissionAllowedRolesread),
		string(CalendarPermissionAllowedRoleswrite),
	}
}

func parseCalendarPermissionAllowedRoles(input string) (*CalendarPermissionAllowedRoles, error) {
	vals := map[string]CalendarPermissionAllowedRoles{
		"custom":                            CalendarPermissionAllowedRolescustom,
		"delegatewithprivateeventaccess":    CalendarPermissionAllowedRolesdelegateWithPrivateEventAccess,
		"delegatewithoutprivateeventaccess": CalendarPermissionAllowedRolesdelegateWithoutPrivateEventAccess,
		"freebusyread":                      CalendarPermissionAllowedRolesfreeBusyRead,
		"limitedread":                       CalendarPermissionAllowedRoleslimitedRead,
		"none":                              CalendarPermissionAllowedRolesnone,
		"read":                              CalendarPermissionAllowedRolesread,
		"write":                             CalendarPermissionAllowedRoleswrite,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := CalendarPermissionAllowedRoles(input)
	return &out, nil
}
