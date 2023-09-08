package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MultiTenantOrganizationState string

const (
	MultiTenantOrganizationStateactive   MultiTenantOrganizationState = "Active"
	MultiTenantOrganizationStateinactive MultiTenantOrganizationState = "Inactive"
)

func PossibleValuesForMultiTenantOrganizationState() []string {
	return []string{
		string(MultiTenantOrganizationStateactive),
		string(MultiTenantOrganizationStateinactive),
	}
}

func parseMultiTenantOrganizationState(input string) (*MultiTenantOrganizationState, error) {
	vals := map[string]MultiTenantOrganizationState{
		"active":   MultiTenantOrganizationStateactive,
		"inactive": MultiTenantOrganizationStateinactive,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := MultiTenantOrganizationState(input)
	return &out, nil
}
