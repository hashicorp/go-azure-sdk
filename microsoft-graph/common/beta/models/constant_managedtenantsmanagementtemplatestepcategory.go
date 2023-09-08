package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ManagedTenantsManagementTemplateStepCategory string

const (
	ManagedTenantsManagementTemplateStepCategorycustom   ManagedTenantsManagementTemplateStepCategory = "Custom"
	ManagedTenantsManagementTemplateStepCategorydata     ManagedTenantsManagementTemplateStepCategory = "Data"
	ManagedTenantsManagementTemplateStepCategorydevices  ManagedTenantsManagementTemplateStepCategory = "Devices"
	ManagedTenantsManagementTemplateStepCategoryidentity ManagedTenantsManagementTemplateStepCategory = "Identity"
)

func PossibleValuesForManagedTenantsManagementTemplateStepCategory() []string {
	return []string{
		string(ManagedTenantsManagementTemplateStepCategorycustom),
		string(ManagedTenantsManagementTemplateStepCategorydata),
		string(ManagedTenantsManagementTemplateStepCategorydevices),
		string(ManagedTenantsManagementTemplateStepCategoryidentity),
	}
}

func parseManagedTenantsManagementTemplateStepCategory(input string) (*ManagedTenantsManagementTemplateStepCategory, error) {
	vals := map[string]ManagedTenantsManagementTemplateStepCategory{
		"custom":   ManagedTenantsManagementTemplateStepCategorycustom,
		"data":     ManagedTenantsManagementTemplateStepCategorydata,
		"devices":  ManagedTenantsManagementTemplateStepCategorydevices,
		"identity": ManagedTenantsManagementTemplateStepCategoryidentity,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ManagedTenantsManagementTemplateStepCategory(input)
	return &out, nil
}
