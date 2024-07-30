package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CalendarPermissionAllowedRoles string

const (
	CalendarPermissionAllowedRoles_Custom                            CalendarPermissionAllowedRoles = "custom"
	CalendarPermissionAllowedRoles_DelegateWithPrivateEventAccess    CalendarPermissionAllowedRoles = "delegateWithPrivateEventAccess"
	CalendarPermissionAllowedRoles_DelegateWithoutPrivateEventAccess CalendarPermissionAllowedRoles = "delegateWithoutPrivateEventAccess"
	CalendarPermissionAllowedRoles_FreeBusyRead                      CalendarPermissionAllowedRoles = "freeBusyRead"
	CalendarPermissionAllowedRoles_LimitedRead                       CalendarPermissionAllowedRoles = "limitedRead"
	CalendarPermissionAllowedRoles_None                              CalendarPermissionAllowedRoles = "none"
	CalendarPermissionAllowedRoles_Read                              CalendarPermissionAllowedRoles = "read"
	CalendarPermissionAllowedRoles_Write                             CalendarPermissionAllowedRoles = "write"
)

func PossibleValuesForCalendarPermissionAllowedRoles() []string {
	return []string{
		string(CalendarPermissionAllowedRoles_Custom),
		string(CalendarPermissionAllowedRoles_DelegateWithPrivateEventAccess),
		string(CalendarPermissionAllowedRoles_DelegateWithoutPrivateEventAccess),
		string(CalendarPermissionAllowedRoles_FreeBusyRead),
		string(CalendarPermissionAllowedRoles_LimitedRead),
		string(CalendarPermissionAllowedRoles_None),
		string(CalendarPermissionAllowedRoles_Read),
		string(CalendarPermissionAllowedRoles_Write),
	}
}

func (s *CalendarPermissionAllowedRoles) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseCalendarPermissionAllowedRoles(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseCalendarPermissionAllowedRoles(input string) (*CalendarPermissionAllowedRoles, error) {
	vals := map[string]CalendarPermissionAllowedRoles{
		"custom":                            CalendarPermissionAllowedRoles_Custom,
		"delegatewithprivateeventaccess":    CalendarPermissionAllowedRoles_DelegateWithPrivateEventAccess,
		"delegatewithoutprivateeventaccess": CalendarPermissionAllowedRoles_DelegateWithoutPrivateEventAccess,
		"freebusyread":                      CalendarPermissionAllowedRoles_FreeBusyRead,
		"limitedread":                       CalendarPermissionAllowedRoles_LimitedRead,
		"none":                              CalendarPermissionAllowedRoles_None,
		"read":                              CalendarPermissionAllowedRoles_Read,
		"write":                             CalendarPermissionAllowedRoles_Write,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := CalendarPermissionAllowedRoles(input)
	return &out, nil
}
