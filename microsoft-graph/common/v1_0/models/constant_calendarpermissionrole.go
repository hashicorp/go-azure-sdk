package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CalendarPermissionRole string

const (
	CalendarPermissionRolecustom                            CalendarPermissionRole = "Custom"
	CalendarPermissionRoledelegateWithPrivateEventAccess    CalendarPermissionRole = "DelegateWithPrivateEventAccess"
	CalendarPermissionRoledelegateWithoutPrivateEventAccess CalendarPermissionRole = "DelegateWithoutPrivateEventAccess"
	CalendarPermissionRolefreeBusyRead                      CalendarPermissionRole = "FreeBusyRead"
	CalendarPermissionRolelimitedRead                       CalendarPermissionRole = "LimitedRead"
	CalendarPermissionRolenone                              CalendarPermissionRole = "None"
	CalendarPermissionRoleread                              CalendarPermissionRole = "Read"
	CalendarPermissionRolewrite                             CalendarPermissionRole = "Write"
)

func PossibleValuesForCalendarPermissionRole() []string {
	return []string{
		string(CalendarPermissionRolecustom),
		string(CalendarPermissionRoledelegateWithPrivateEventAccess),
		string(CalendarPermissionRoledelegateWithoutPrivateEventAccess),
		string(CalendarPermissionRolefreeBusyRead),
		string(CalendarPermissionRolelimitedRead),
		string(CalendarPermissionRolenone),
		string(CalendarPermissionRoleread),
		string(CalendarPermissionRolewrite),
	}
}

func parseCalendarPermissionRole(input string) (*CalendarPermissionRole, error) {
	vals := map[string]CalendarPermissionRole{
		"custom":                            CalendarPermissionRolecustom,
		"delegatewithprivateeventaccess":    CalendarPermissionRoledelegateWithPrivateEventAccess,
		"delegatewithoutprivateeventaccess": CalendarPermissionRoledelegateWithoutPrivateEventAccess,
		"freebusyread":                      CalendarPermissionRolefreeBusyRead,
		"limitedread":                       CalendarPermissionRolelimitedRead,
		"none":                              CalendarPermissionRolenone,
		"read":                              CalendarPermissionRoleread,
		"write":                             CalendarPermissionRolewrite,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := CalendarPermissionRole(input)
	return &out, nil
}
