package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ManagedTenantsTenantStatusInformationOnboardingStatus string

const (
	ManagedTenantsTenantStatusInformationOnboardingStatusactive     ManagedTenantsTenantStatusInformationOnboardingStatus = "Active"
	ManagedTenantsTenantStatusInformationOnboardingStatusinProcess  ManagedTenantsTenantStatusInformationOnboardingStatus = "InProcess"
	ManagedTenantsTenantStatusInformationOnboardingStatusinactive   ManagedTenantsTenantStatusInformationOnboardingStatus = "Inactive"
	ManagedTenantsTenantStatusInformationOnboardingStatusineligible ManagedTenantsTenantStatusInformationOnboardingStatus = "Ineligible"
)

func PossibleValuesForManagedTenantsTenantStatusInformationOnboardingStatus() []string {
	return []string{
		string(ManagedTenantsTenantStatusInformationOnboardingStatusactive),
		string(ManagedTenantsTenantStatusInformationOnboardingStatusinProcess),
		string(ManagedTenantsTenantStatusInformationOnboardingStatusinactive),
		string(ManagedTenantsTenantStatusInformationOnboardingStatusineligible),
	}
}

func parseManagedTenantsTenantStatusInformationOnboardingStatus(input string) (*ManagedTenantsTenantStatusInformationOnboardingStatus, error) {
	vals := map[string]ManagedTenantsTenantStatusInformationOnboardingStatus{
		"active":     ManagedTenantsTenantStatusInformationOnboardingStatusactive,
		"inprocess":  ManagedTenantsTenantStatusInformationOnboardingStatusinProcess,
		"inactive":   ManagedTenantsTenantStatusInformationOnboardingStatusinactive,
		"ineligible": ManagedTenantsTenantStatusInformationOnboardingStatusineligible,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ManagedTenantsTenantStatusInformationOnboardingStatus(input)
	return &out, nil
}
