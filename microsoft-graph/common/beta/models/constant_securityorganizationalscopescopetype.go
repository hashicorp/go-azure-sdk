package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityOrganizationalScopeScopeType string

const (
	SecurityOrganizationalScopeScopeTypedeviceGroup SecurityOrganizationalScopeScopeType = "DeviceGroup"
)

func PossibleValuesForSecurityOrganizationalScopeScopeType() []string {
	return []string{
		string(SecurityOrganizationalScopeScopeTypedeviceGroup),
	}
}

func parseSecurityOrganizationalScopeScopeType(input string) (*SecurityOrganizationalScopeScopeType, error) {
	vals := map[string]SecurityOrganizationalScopeScopeType{
		"devicegroup": SecurityOrganizationalScopeScopeTypedeviceGroup,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityOrganizationalScopeScopeType(input)
	return &out, nil
}
