package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ManagedTenantsTenantStatusInformationTenantOnboardingEligibilityReason string

const (
	ManagedTenantsTenantStatusInformationTenantOnboardingEligibilityReason_ContractType             ManagedTenantsTenantStatusInformationTenantOnboardingEligibilityReason = "contractType"
	ManagedTenantsTenantStatusInformationTenantOnboardingEligibilityReason_DelegatedAdminPrivileges ManagedTenantsTenantStatusInformationTenantOnboardingEligibilityReason = "delegatedAdminPrivileges"
	ManagedTenantsTenantStatusInformationTenantOnboardingEligibilityReason_License                  ManagedTenantsTenantStatusInformationTenantOnboardingEligibilityReason = "license"
	ManagedTenantsTenantStatusInformationTenantOnboardingEligibilityReason_None                     ManagedTenantsTenantStatusInformationTenantOnboardingEligibilityReason = "none"
	ManagedTenantsTenantStatusInformationTenantOnboardingEligibilityReason_UsersCount               ManagedTenantsTenantStatusInformationTenantOnboardingEligibilityReason = "usersCount"
)

func PossibleValuesForManagedTenantsTenantStatusInformationTenantOnboardingEligibilityReason() []string {
	return []string{
		string(ManagedTenantsTenantStatusInformationTenantOnboardingEligibilityReason_ContractType),
		string(ManagedTenantsTenantStatusInformationTenantOnboardingEligibilityReason_DelegatedAdminPrivileges),
		string(ManagedTenantsTenantStatusInformationTenantOnboardingEligibilityReason_License),
		string(ManagedTenantsTenantStatusInformationTenantOnboardingEligibilityReason_None),
		string(ManagedTenantsTenantStatusInformationTenantOnboardingEligibilityReason_UsersCount),
	}
}

func (s *ManagedTenantsTenantStatusInformationTenantOnboardingEligibilityReason) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseManagedTenantsTenantStatusInformationTenantOnboardingEligibilityReason(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseManagedTenantsTenantStatusInformationTenantOnboardingEligibilityReason(input string) (*ManagedTenantsTenantStatusInformationTenantOnboardingEligibilityReason, error) {
	vals := map[string]ManagedTenantsTenantStatusInformationTenantOnboardingEligibilityReason{
		"contracttype":             ManagedTenantsTenantStatusInformationTenantOnboardingEligibilityReason_ContractType,
		"delegatedadminprivileges": ManagedTenantsTenantStatusInformationTenantOnboardingEligibilityReason_DelegatedAdminPrivileges,
		"license":                  ManagedTenantsTenantStatusInformationTenantOnboardingEligibilityReason_License,
		"none":                     ManagedTenantsTenantStatusInformationTenantOnboardingEligibilityReason_None,
		"userscount":               ManagedTenantsTenantStatusInformationTenantOnboardingEligibilityReason_UsersCount,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ManagedTenantsTenantStatusInformationTenantOnboardingEligibilityReason(input)
	return &out, nil
}
