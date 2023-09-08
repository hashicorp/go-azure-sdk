package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MultiTenantOrganizationPartnerConfigurationTemplateTemplateApplicationLevel string

const (
	MultiTenantOrganizationPartnerConfigurationTemplateTemplateApplicationLevelexistingPartners MultiTenantOrganizationPartnerConfigurationTemplateTemplateApplicationLevel = "ExistingPartners"
	MultiTenantOrganizationPartnerConfigurationTemplateTemplateApplicationLevelnewPartners      MultiTenantOrganizationPartnerConfigurationTemplateTemplateApplicationLevel = "NewPartners"
	MultiTenantOrganizationPartnerConfigurationTemplateTemplateApplicationLevelnone             MultiTenantOrganizationPartnerConfigurationTemplateTemplateApplicationLevel = "None"
)

func PossibleValuesForMultiTenantOrganizationPartnerConfigurationTemplateTemplateApplicationLevel() []string {
	return []string{
		string(MultiTenantOrganizationPartnerConfigurationTemplateTemplateApplicationLevelexistingPartners),
		string(MultiTenantOrganizationPartnerConfigurationTemplateTemplateApplicationLevelnewPartners),
		string(MultiTenantOrganizationPartnerConfigurationTemplateTemplateApplicationLevelnone),
	}
}

func parseMultiTenantOrganizationPartnerConfigurationTemplateTemplateApplicationLevel(input string) (*MultiTenantOrganizationPartnerConfigurationTemplateTemplateApplicationLevel, error) {
	vals := map[string]MultiTenantOrganizationPartnerConfigurationTemplateTemplateApplicationLevel{
		"existingpartners": MultiTenantOrganizationPartnerConfigurationTemplateTemplateApplicationLevelexistingPartners,
		"newpartners":      MultiTenantOrganizationPartnerConfigurationTemplateTemplateApplicationLevelnewPartners,
		"none":             MultiTenantOrganizationPartnerConfigurationTemplateTemplateApplicationLevelnone,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := MultiTenantOrganizationPartnerConfigurationTemplateTemplateApplicationLevel(input)
	return &out, nil
}
