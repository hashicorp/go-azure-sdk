package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MultiTenantOrganizationIdentitySyncPolicyTemplateTemplateApplicationLevel string

const (
	MultiTenantOrganizationIdentitySyncPolicyTemplateTemplateApplicationLevelexistingPartners MultiTenantOrganizationIdentitySyncPolicyTemplateTemplateApplicationLevel = "ExistingPartners"
	MultiTenantOrganizationIdentitySyncPolicyTemplateTemplateApplicationLevelnewPartners      MultiTenantOrganizationIdentitySyncPolicyTemplateTemplateApplicationLevel = "NewPartners"
	MultiTenantOrganizationIdentitySyncPolicyTemplateTemplateApplicationLevelnone             MultiTenantOrganizationIdentitySyncPolicyTemplateTemplateApplicationLevel = "None"
)

func PossibleValuesForMultiTenantOrganizationIdentitySyncPolicyTemplateTemplateApplicationLevel() []string {
	return []string{
		string(MultiTenantOrganizationIdentitySyncPolicyTemplateTemplateApplicationLevelexistingPartners),
		string(MultiTenantOrganizationIdentitySyncPolicyTemplateTemplateApplicationLevelnewPartners),
		string(MultiTenantOrganizationIdentitySyncPolicyTemplateTemplateApplicationLevelnone),
	}
}

func parseMultiTenantOrganizationIdentitySyncPolicyTemplateTemplateApplicationLevel(input string) (*MultiTenantOrganizationIdentitySyncPolicyTemplateTemplateApplicationLevel, error) {
	vals := map[string]MultiTenantOrganizationIdentitySyncPolicyTemplateTemplateApplicationLevel{
		"existingpartners": MultiTenantOrganizationIdentitySyncPolicyTemplateTemplateApplicationLevelexistingPartners,
		"newpartners":      MultiTenantOrganizationIdentitySyncPolicyTemplateTemplateApplicationLevelnewPartners,
		"none":             MultiTenantOrganizationIdentitySyncPolicyTemplateTemplateApplicationLevelnone,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := MultiTenantOrganizationIdentitySyncPolicyTemplateTemplateApplicationLevel(input)
	return &out, nil
}
