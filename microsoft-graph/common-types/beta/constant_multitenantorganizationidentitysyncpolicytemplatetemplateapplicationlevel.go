package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MultiTenantOrganizationIdentitySyncPolicyTemplateTemplateApplicationLevel string

const (
	MultiTenantOrganizationIdentitySyncPolicyTemplateTemplateApplicationLevel_ExistingPartners MultiTenantOrganizationIdentitySyncPolicyTemplateTemplateApplicationLevel = "existingPartners"
	MultiTenantOrganizationIdentitySyncPolicyTemplateTemplateApplicationLevel_NewPartners      MultiTenantOrganizationIdentitySyncPolicyTemplateTemplateApplicationLevel = "newPartners"
	MultiTenantOrganizationIdentitySyncPolicyTemplateTemplateApplicationLevel_None             MultiTenantOrganizationIdentitySyncPolicyTemplateTemplateApplicationLevel = "none"
)

func PossibleValuesForMultiTenantOrganizationIdentitySyncPolicyTemplateTemplateApplicationLevel() []string {
	return []string{
		string(MultiTenantOrganizationIdentitySyncPolicyTemplateTemplateApplicationLevel_ExistingPartners),
		string(MultiTenantOrganizationIdentitySyncPolicyTemplateTemplateApplicationLevel_NewPartners),
		string(MultiTenantOrganizationIdentitySyncPolicyTemplateTemplateApplicationLevel_None),
	}
}

func (s *MultiTenantOrganizationIdentitySyncPolicyTemplateTemplateApplicationLevel) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseMultiTenantOrganizationIdentitySyncPolicyTemplateTemplateApplicationLevel(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseMultiTenantOrganizationIdentitySyncPolicyTemplateTemplateApplicationLevel(input string) (*MultiTenantOrganizationIdentitySyncPolicyTemplateTemplateApplicationLevel, error) {
	vals := map[string]MultiTenantOrganizationIdentitySyncPolicyTemplateTemplateApplicationLevel{
		"existingpartners": MultiTenantOrganizationIdentitySyncPolicyTemplateTemplateApplicationLevel_ExistingPartners,
		"newpartners":      MultiTenantOrganizationIdentitySyncPolicyTemplateTemplateApplicationLevel_NewPartners,
		"none":             MultiTenantOrganizationIdentitySyncPolicyTemplateTemplateApplicationLevel_None,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := MultiTenantOrganizationIdentitySyncPolicyTemplateTemplateApplicationLevel(input)
	return &out, nil
}
