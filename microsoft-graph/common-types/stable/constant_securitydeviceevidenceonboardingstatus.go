package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityDeviceEvidenceOnboardingStatus string

const (
	SecurityDeviceEvidenceOnboardingStatus_CanBeOnboarded   SecurityDeviceEvidenceOnboardingStatus = "canBeOnboarded"
	SecurityDeviceEvidenceOnboardingStatus_InsufficientInfo SecurityDeviceEvidenceOnboardingStatus = "insufficientInfo"
	SecurityDeviceEvidenceOnboardingStatus_Onboarded        SecurityDeviceEvidenceOnboardingStatus = "onboarded"
	SecurityDeviceEvidenceOnboardingStatus_Unsupported      SecurityDeviceEvidenceOnboardingStatus = "unsupported"
)

func PossibleValuesForSecurityDeviceEvidenceOnboardingStatus() []string {
	return []string{
		string(SecurityDeviceEvidenceOnboardingStatus_CanBeOnboarded),
		string(SecurityDeviceEvidenceOnboardingStatus_InsufficientInfo),
		string(SecurityDeviceEvidenceOnboardingStatus_Onboarded),
		string(SecurityDeviceEvidenceOnboardingStatus_Unsupported),
	}
}

func (s *SecurityDeviceEvidenceOnboardingStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSecurityDeviceEvidenceOnboardingStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSecurityDeviceEvidenceOnboardingStatus(input string) (*SecurityDeviceEvidenceOnboardingStatus, error) {
	vals := map[string]SecurityDeviceEvidenceOnboardingStatus{
		"canbeonboarded":   SecurityDeviceEvidenceOnboardingStatus_CanBeOnboarded,
		"insufficientinfo": SecurityDeviceEvidenceOnboardingStatus_InsufficientInfo,
		"onboarded":        SecurityDeviceEvidenceOnboardingStatus_Onboarded,
		"unsupported":      SecurityDeviceEvidenceOnboardingStatus_Unsupported,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityDeviceEvidenceOnboardingStatus(input)
	return &out, nil
}
