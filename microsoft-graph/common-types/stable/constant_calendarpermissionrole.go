package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CalendarPermissionRole string

const (
	CalendarPermissionRole_Custom                            CalendarPermissionRole = "custom"
	CalendarPermissionRole_DelegateWithPrivateEventAccess    CalendarPermissionRole = "delegateWithPrivateEventAccess"
	CalendarPermissionRole_DelegateWithoutPrivateEventAccess CalendarPermissionRole = "delegateWithoutPrivateEventAccess"
	CalendarPermissionRole_FreeBusyRead                      CalendarPermissionRole = "freeBusyRead"
	CalendarPermissionRole_LimitedRead                       CalendarPermissionRole = "limitedRead"
	CalendarPermissionRole_None                              CalendarPermissionRole = "none"
	CalendarPermissionRole_Read                              CalendarPermissionRole = "read"
	CalendarPermissionRole_Write                             CalendarPermissionRole = "write"
)

func PossibleValuesForCalendarPermissionRole() []string {
	return []string{
		string(CalendarPermissionRole_Custom),
		string(CalendarPermissionRole_DelegateWithPrivateEventAccess),
		string(CalendarPermissionRole_DelegateWithoutPrivateEventAccess),
		string(CalendarPermissionRole_FreeBusyRead),
		string(CalendarPermissionRole_LimitedRead),
		string(CalendarPermissionRole_None),
		string(CalendarPermissionRole_Read),
		string(CalendarPermissionRole_Write),
	}
}

func (s *CalendarPermissionRole) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseCalendarPermissionRole(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseCalendarPermissionRole(input string) (*CalendarPermissionRole, error) {
	vals := map[string]CalendarPermissionRole{
		"custom":                            CalendarPermissionRole_Custom,
		"delegatewithprivateeventaccess":    CalendarPermissionRole_DelegateWithPrivateEventAccess,
		"delegatewithoutprivateeventaccess": CalendarPermissionRole_DelegateWithoutPrivateEventAccess,
		"freebusyread":                      CalendarPermissionRole_FreeBusyRead,
		"limitedread":                       CalendarPermissionRole_LimitedRead,
		"none":                              CalendarPermissionRole_None,
		"read":                              CalendarPermissionRole_Read,
		"write":                             CalendarPermissionRole_Write,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := CalendarPermissionRole(input)
	return &out, nil
}
