package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityDeviceEvidenceOnboardingStatus string

const (
	SecurityDeviceEvidenceOnboardingStatuscanBeOnboarded   SecurityDeviceEvidenceOnboardingStatus = "CanBeOnboarded"
	SecurityDeviceEvidenceOnboardingStatusinsufficientInfo SecurityDeviceEvidenceOnboardingStatus = "InsufficientInfo"
	SecurityDeviceEvidenceOnboardingStatusonboarded        SecurityDeviceEvidenceOnboardingStatus = "Onboarded"
	SecurityDeviceEvidenceOnboardingStatusunsupported      SecurityDeviceEvidenceOnboardingStatus = "Unsupported"
)

func PossibleValuesForSecurityDeviceEvidenceOnboardingStatus() []string {
	return []string{
		string(SecurityDeviceEvidenceOnboardingStatuscanBeOnboarded),
		string(SecurityDeviceEvidenceOnboardingStatusinsufficientInfo),
		string(SecurityDeviceEvidenceOnboardingStatusonboarded),
		string(SecurityDeviceEvidenceOnboardingStatusunsupported),
	}
}

func parseSecurityDeviceEvidenceOnboardingStatus(input string) (*SecurityDeviceEvidenceOnboardingStatus, error) {
	vals := map[string]SecurityDeviceEvidenceOnboardingStatus{
		"canbeonboarded":   SecurityDeviceEvidenceOnboardingStatuscanBeOnboarded,
		"insufficientinfo": SecurityDeviceEvidenceOnboardingStatusinsufficientInfo,
		"onboarded":        SecurityDeviceEvidenceOnboardingStatusonboarded,
		"unsupported":      SecurityDeviceEvidenceOnboardingStatusunsupported,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityDeviceEvidenceOnboardingStatus(input)
	return &out, nil
}
