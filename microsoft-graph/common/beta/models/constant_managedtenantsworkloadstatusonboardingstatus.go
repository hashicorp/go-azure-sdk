package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ManagedTenantsWorkloadStatusOnboardingStatus string

const (
	ManagedTenantsWorkloadStatusOnboardingStatusnotOnboarded ManagedTenantsWorkloadStatusOnboardingStatus = "NotOnboarded"
	ManagedTenantsWorkloadStatusOnboardingStatusonboarded    ManagedTenantsWorkloadStatusOnboardingStatus = "Onboarded"
)

func PossibleValuesForManagedTenantsWorkloadStatusOnboardingStatus() []string {
	return []string{
		string(ManagedTenantsWorkloadStatusOnboardingStatusnotOnboarded),
		string(ManagedTenantsWorkloadStatusOnboardingStatusonboarded),
	}
}

func parseManagedTenantsWorkloadStatusOnboardingStatus(input string) (*ManagedTenantsWorkloadStatusOnboardingStatus, error) {
	vals := map[string]ManagedTenantsWorkloadStatusOnboardingStatus{
		"notonboarded": ManagedTenantsWorkloadStatusOnboardingStatusnotOnboarded,
		"onboarded":    ManagedTenantsWorkloadStatusOnboardingStatusonboarded,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ManagedTenantsWorkloadStatusOnboardingStatus(input)
	return &out, nil
}
