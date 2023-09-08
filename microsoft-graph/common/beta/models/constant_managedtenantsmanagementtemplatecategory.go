package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ManagedTenantsManagementTemplateCategory string

const (
	ManagedTenantsManagementTemplateCategorycustom   ManagedTenantsManagementTemplateCategory = "Custom"
	ManagedTenantsManagementTemplateCategorydata     ManagedTenantsManagementTemplateCategory = "Data"
	ManagedTenantsManagementTemplateCategorydevices  ManagedTenantsManagementTemplateCategory = "Devices"
	ManagedTenantsManagementTemplateCategoryidentity ManagedTenantsManagementTemplateCategory = "Identity"
)

func PossibleValuesForManagedTenantsManagementTemplateCategory() []string {
	return []string{
		string(ManagedTenantsManagementTemplateCategorycustom),
		string(ManagedTenantsManagementTemplateCategorydata),
		string(ManagedTenantsManagementTemplateCategorydevices),
		string(ManagedTenantsManagementTemplateCategoryidentity),
	}
}

func parseManagedTenantsManagementTemplateCategory(input string) (*ManagedTenantsManagementTemplateCategory, error) {
	vals := map[string]ManagedTenantsManagementTemplateCategory{
		"custom":   ManagedTenantsManagementTemplateCategorycustom,
		"data":     ManagedTenantsManagementTemplateCategorydata,
		"devices":  ManagedTenantsManagementTemplateCategorydevices,
		"identity": ManagedTenantsManagementTemplateCategoryidentity,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ManagedTenantsManagementTemplateCategory(input)
	return &out, nil
}
