package tenants

import "strings"

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ResourceNameStatus string

const (
	ResourceNameStatusAllowed  ResourceNameStatus = "Allowed"
	ResourceNameStatusReserved ResourceNameStatus = "Reserved"
)

func PossibleValuesForResourceNameStatus() []string {
	return []string{
		string(ResourceNameStatusAllowed),
		string(ResourceNameStatusReserved),
	}
}

func parseResourceNameStatus(input string) (*ResourceNameStatus, error) {
	vals := map[string]ResourceNameStatus{
		"allowed":  ResourceNameStatusAllowed,
		"reserved": ResourceNameStatusReserved,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ResourceNameStatus(input)
	return &out, nil
}

type TenantCategory string

const (
	TenantCategoryHome        TenantCategory = "Home"
	TenantCategoryManagedBy   TenantCategory = "ManagedBy"
	TenantCategoryProjectedBy TenantCategory = "ProjectedBy"
)

func PossibleValuesForTenantCategory() []string {
	return []string{
		string(TenantCategoryHome),
		string(TenantCategoryManagedBy),
		string(TenantCategoryProjectedBy),
	}
}

func parseTenantCategory(input string) (*TenantCategory, error) {
	vals := map[string]TenantCategory{
		"home":        TenantCategoryHome,
		"managedby":   TenantCategoryManagedBy,
		"projectedby": TenantCategoryProjectedBy,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := TenantCategory(input)
	return &out, nil
}
