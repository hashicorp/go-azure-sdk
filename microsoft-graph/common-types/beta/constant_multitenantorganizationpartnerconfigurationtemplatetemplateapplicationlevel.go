package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MultiTenantOrganizationPartnerConfigurationTemplateTemplateApplicationLevel string

const (
	MultiTenantOrganizationPartnerConfigurationTemplateTemplateApplicationLevel_ExistingPartners MultiTenantOrganizationPartnerConfigurationTemplateTemplateApplicationLevel = "existingPartners"
	MultiTenantOrganizationPartnerConfigurationTemplateTemplateApplicationLevel_NewPartners      MultiTenantOrganizationPartnerConfigurationTemplateTemplateApplicationLevel = "newPartners"
	MultiTenantOrganizationPartnerConfigurationTemplateTemplateApplicationLevel_None             MultiTenantOrganizationPartnerConfigurationTemplateTemplateApplicationLevel = "none"
)

func PossibleValuesForMultiTenantOrganizationPartnerConfigurationTemplateTemplateApplicationLevel() []string {
	return []string{
		string(MultiTenantOrganizationPartnerConfigurationTemplateTemplateApplicationLevel_ExistingPartners),
		string(MultiTenantOrganizationPartnerConfigurationTemplateTemplateApplicationLevel_NewPartners),
		string(MultiTenantOrganizationPartnerConfigurationTemplateTemplateApplicationLevel_None),
	}
}

func (s *MultiTenantOrganizationPartnerConfigurationTemplateTemplateApplicationLevel) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseMultiTenantOrganizationPartnerConfigurationTemplateTemplateApplicationLevel(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseMultiTenantOrganizationPartnerConfigurationTemplateTemplateApplicationLevel(input string) (*MultiTenantOrganizationPartnerConfigurationTemplateTemplateApplicationLevel, error) {
	vals := map[string]MultiTenantOrganizationPartnerConfigurationTemplateTemplateApplicationLevel{
		"existingpartners": MultiTenantOrganizationPartnerConfigurationTemplateTemplateApplicationLevel_ExistingPartners,
		"newpartners":      MultiTenantOrganizationPartnerConfigurationTemplateTemplateApplicationLevel_NewPartners,
		"none":             MultiTenantOrganizationPartnerConfigurationTemplateTemplateApplicationLevel_None,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := MultiTenantOrganizationPartnerConfigurationTemplateTemplateApplicationLevel(input)
	return &out, nil
}
