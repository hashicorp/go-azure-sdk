package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RemoteAssistancePartnerOnboardingStatus string

const (
	RemoteAssistancePartnerOnboardingStatusnotOnboarded RemoteAssistancePartnerOnboardingStatus = "NotOnboarded"
	RemoteAssistancePartnerOnboardingStatusonboarded    RemoteAssistancePartnerOnboardingStatus = "Onboarded"
	RemoteAssistancePartnerOnboardingStatusonboarding   RemoteAssistancePartnerOnboardingStatus = "Onboarding"
)

func PossibleValuesForRemoteAssistancePartnerOnboardingStatus() []string {
	return []string{
		string(RemoteAssistancePartnerOnboardingStatusnotOnboarded),
		string(RemoteAssistancePartnerOnboardingStatusonboarded),
		string(RemoteAssistancePartnerOnboardingStatusonboarding),
	}
}

func parseRemoteAssistancePartnerOnboardingStatus(input string) (*RemoteAssistancePartnerOnboardingStatus, error) {
	vals := map[string]RemoteAssistancePartnerOnboardingStatus{
		"notonboarded": RemoteAssistancePartnerOnboardingStatusnotOnboarded,
		"onboarded":    RemoteAssistancePartnerOnboardingStatusonboarded,
		"onboarding":   RemoteAssistancePartnerOnboardingStatusonboarding,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := RemoteAssistancePartnerOnboardingStatus(input)
	return &out, nil
}
