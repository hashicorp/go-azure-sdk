package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type NetworkaccessTenantStatusOnboardingStatus string

const (
	NetworkaccessTenantStatusOnboardingStatus_Offboarded               NetworkaccessTenantStatusOnboardingStatus = "offboarded"
	NetworkaccessTenantStatusOnboardingStatus_OffboardingErrorOccurred NetworkaccessTenantStatusOnboardingStatus = "offboardingErrorOccurred"
	NetworkaccessTenantStatusOnboardingStatus_OffboardingInProgress    NetworkaccessTenantStatusOnboardingStatus = "offboardingInProgress"
	NetworkaccessTenantStatusOnboardingStatus_Onboarded                NetworkaccessTenantStatusOnboardingStatus = "onboarded"
	NetworkaccessTenantStatusOnboardingStatus_OnboardingErrorOccurred  NetworkaccessTenantStatusOnboardingStatus = "onboardingErrorOccurred"
	NetworkaccessTenantStatusOnboardingStatus_OnboardingInProgress     NetworkaccessTenantStatusOnboardingStatus = "onboardingInProgress"
)

func PossibleValuesForNetworkaccessTenantStatusOnboardingStatus() []string {
	return []string{
		string(NetworkaccessTenantStatusOnboardingStatus_Offboarded),
		string(NetworkaccessTenantStatusOnboardingStatus_OffboardingErrorOccurred),
		string(NetworkaccessTenantStatusOnboardingStatus_OffboardingInProgress),
		string(NetworkaccessTenantStatusOnboardingStatus_Onboarded),
		string(NetworkaccessTenantStatusOnboardingStatus_OnboardingErrorOccurred),
		string(NetworkaccessTenantStatusOnboardingStatus_OnboardingInProgress),
	}
}

func (s *NetworkaccessTenantStatusOnboardingStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseNetworkaccessTenantStatusOnboardingStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseNetworkaccessTenantStatusOnboardingStatus(input string) (*NetworkaccessTenantStatusOnboardingStatus, error) {
	vals := map[string]NetworkaccessTenantStatusOnboardingStatus{
		"offboarded":               NetworkaccessTenantStatusOnboardingStatus_Offboarded,
		"offboardingerroroccurred": NetworkaccessTenantStatusOnboardingStatus_OffboardingErrorOccurred,
		"offboardinginprogress":    NetworkaccessTenantStatusOnboardingStatus_OffboardingInProgress,
		"onboarded":                NetworkaccessTenantStatusOnboardingStatus_Onboarded,
		"onboardingerroroccurred":  NetworkaccessTenantStatusOnboardingStatus_OnboardingErrorOccurred,
		"onboardinginprogress":     NetworkaccessTenantStatusOnboardingStatus_OnboardingInProgress,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := NetworkaccessTenantStatusOnboardingStatus(input)
	return &out, nil
}
