package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ManagedTenantsManagementActionCategory string

const (
	ManagedTenantsManagementActionCategorycustom   ManagedTenantsManagementActionCategory = "Custom"
	ManagedTenantsManagementActionCategorydata     ManagedTenantsManagementActionCategory = "Data"
	ManagedTenantsManagementActionCategorydevices  ManagedTenantsManagementActionCategory = "Devices"
	ManagedTenantsManagementActionCategoryidentity ManagedTenantsManagementActionCategory = "Identity"
)

func PossibleValuesForManagedTenantsManagementActionCategory() []string {
	return []string{
		string(ManagedTenantsManagementActionCategorycustom),
		string(ManagedTenantsManagementActionCategorydata),
		string(ManagedTenantsManagementActionCategorydevices),
		string(ManagedTenantsManagementActionCategoryidentity),
	}
}

func parseManagedTenantsManagementActionCategory(input string) (*ManagedTenantsManagementActionCategory, error) {
	vals := map[string]ManagedTenantsManagementActionCategory{
		"custom":   ManagedTenantsManagementActionCategorycustom,
		"data":     ManagedTenantsManagementActionCategorydata,
		"devices":  ManagedTenantsManagementActionCategorydevices,
		"identity": ManagedTenantsManagementActionCategoryidentity,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ManagedTenantsManagementActionCategory(input)
	return &out, nil
}
