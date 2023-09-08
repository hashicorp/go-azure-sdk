package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ManagedTenantsTenantStatusInformationTenantOnboardingEligibilityReason string

const (
	ManagedTenantsTenantStatusInformationTenantOnboardingEligibilityReasoncontractType             ManagedTenantsTenantStatusInformationTenantOnboardingEligibilityReason = "ContractType"
	ManagedTenantsTenantStatusInformationTenantOnboardingEligibilityReasondelegatedAdminPrivileges ManagedTenantsTenantStatusInformationTenantOnboardingEligibilityReason = "DelegatedAdminPrivileges"
	ManagedTenantsTenantStatusInformationTenantOnboardingEligibilityReasonlicense                  ManagedTenantsTenantStatusInformationTenantOnboardingEligibilityReason = "License"
	ManagedTenantsTenantStatusInformationTenantOnboardingEligibilityReasonnone                     ManagedTenantsTenantStatusInformationTenantOnboardingEligibilityReason = "None"
	ManagedTenantsTenantStatusInformationTenantOnboardingEligibilityReasonusersCount               ManagedTenantsTenantStatusInformationTenantOnboardingEligibilityReason = "UsersCount"
)

func PossibleValuesForManagedTenantsTenantStatusInformationTenantOnboardingEligibilityReason() []string {
	return []string{
		string(ManagedTenantsTenantStatusInformationTenantOnboardingEligibilityReasoncontractType),
		string(ManagedTenantsTenantStatusInformationTenantOnboardingEligibilityReasondelegatedAdminPrivileges),
		string(ManagedTenantsTenantStatusInformationTenantOnboardingEligibilityReasonlicense),
		string(ManagedTenantsTenantStatusInformationTenantOnboardingEligibilityReasonnone),
		string(ManagedTenantsTenantStatusInformationTenantOnboardingEligibilityReasonusersCount),
	}
}

func parseManagedTenantsTenantStatusInformationTenantOnboardingEligibilityReason(input string) (*ManagedTenantsTenantStatusInformationTenantOnboardingEligibilityReason, error) {
	vals := map[string]ManagedTenantsTenantStatusInformationTenantOnboardingEligibilityReason{
		"contracttype":             ManagedTenantsTenantStatusInformationTenantOnboardingEligibilityReasoncontractType,
		"delegatedadminprivileges": ManagedTenantsTenantStatusInformationTenantOnboardingEligibilityReasondelegatedAdminPrivileges,
		"license":                  ManagedTenantsTenantStatusInformationTenantOnboardingEligibilityReasonlicense,
		"none":                     ManagedTenantsTenantStatusInformationTenantOnboardingEligibilityReasonnone,
		"userscount":               ManagedTenantsTenantStatusInformationTenantOnboardingEligibilityReasonusersCount,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ManagedTenantsTenantStatusInformationTenantOnboardingEligibilityReason(input)
	return &out, nil
}
