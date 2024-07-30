package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RemoteAssistancePartnerOnboardingStatus string

const (
	RemoteAssistancePartnerOnboardingStatus_NotOnboarded RemoteAssistancePartnerOnboardingStatus = "notOnboarded"
	RemoteAssistancePartnerOnboardingStatus_Onboarded    RemoteAssistancePartnerOnboardingStatus = "onboarded"
	RemoteAssistancePartnerOnboardingStatus_Onboarding   RemoteAssistancePartnerOnboardingStatus = "onboarding"
)

func PossibleValuesForRemoteAssistancePartnerOnboardingStatus() []string {
	return []string{
		string(RemoteAssistancePartnerOnboardingStatus_NotOnboarded),
		string(RemoteAssistancePartnerOnboardingStatus_Onboarded),
		string(RemoteAssistancePartnerOnboardingStatus_Onboarding),
	}
}

func (s *RemoteAssistancePartnerOnboardingStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseRemoteAssistancePartnerOnboardingStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseRemoteAssistancePartnerOnboardingStatus(input string) (*RemoteAssistancePartnerOnboardingStatus, error) {
	vals := map[string]RemoteAssistancePartnerOnboardingStatus{
		"notonboarded": RemoteAssistancePartnerOnboardingStatus_NotOnboarded,
		"onboarded":    RemoteAssistancePartnerOnboardingStatus_Onboarded,
		"onboarding":   RemoteAssistancePartnerOnboardingStatus_Onboarding,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := RemoteAssistancePartnerOnboardingStatus(input)
	return &out, nil
}
