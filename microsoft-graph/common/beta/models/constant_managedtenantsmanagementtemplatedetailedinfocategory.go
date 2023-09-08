package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ManagedTenantsManagementTemplateDetailedInfoCategory string

const (
	ManagedTenantsManagementTemplateDetailedInfoCategorycustom   ManagedTenantsManagementTemplateDetailedInfoCategory = "Custom"
	ManagedTenantsManagementTemplateDetailedInfoCategorydata     ManagedTenantsManagementTemplateDetailedInfoCategory = "Data"
	ManagedTenantsManagementTemplateDetailedInfoCategorydevices  ManagedTenantsManagementTemplateDetailedInfoCategory = "Devices"
	ManagedTenantsManagementTemplateDetailedInfoCategoryidentity ManagedTenantsManagementTemplateDetailedInfoCategory = "Identity"
)

func PossibleValuesForManagedTenantsManagementTemplateDetailedInfoCategory() []string {
	return []string{
		string(ManagedTenantsManagementTemplateDetailedInfoCategorycustom),
		string(ManagedTenantsManagementTemplateDetailedInfoCategorydata),
		string(ManagedTenantsManagementTemplateDetailedInfoCategorydevices),
		string(ManagedTenantsManagementTemplateDetailedInfoCategoryidentity),
	}
}

func parseManagedTenantsManagementTemplateDetailedInfoCategory(input string) (*ManagedTenantsManagementTemplateDetailedInfoCategory, error) {
	vals := map[string]ManagedTenantsManagementTemplateDetailedInfoCategory{
		"custom":   ManagedTenantsManagementTemplateDetailedInfoCategorycustom,
		"data":     ManagedTenantsManagementTemplateDetailedInfoCategorydata,
		"devices":  ManagedTenantsManagementTemplateDetailedInfoCategorydevices,
		"identity": ManagedTenantsManagementTemplateDetailedInfoCategoryidentity,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ManagedTenantsManagementTemplateDetailedInfoCategory(input)
	return &out, nil
}
