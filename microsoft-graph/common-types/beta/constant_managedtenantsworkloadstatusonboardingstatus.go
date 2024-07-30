package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ManagedTenantsWorkloadStatusOnboardingStatus string

const (
	ManagedTenantsWorkloadStatusOnboardingStatus_NotOnboarded ManagedTenantsWorkloadStatusOnboardingStatus = "notOnboarded"
	ManagedTenantsWorkloadStatusOnboardingStatus_Onboarded    ManagedTenantsWorkloadStatusOnboardingStatus = "onboarded"
)

func PossibleValuesForManagedTenantsWorkloadStatusOnboardingStatus() []string {
	return []string{
		string(ManagedTenantsWorkloadStatusOnboardingStatus_NotOnboarded),
		string(ManagedTenantsWorkloadStatusOnboardingStatus_Onboarded),
	}
}

func (s *ManagedTenantsWorkloadStatusOnboardingStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseManagedTenantsWorkloadStatusOnboardingStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseManagedTenantsWorkloadStatusOnboardingStatus(input string) (*ManagedTenantsWorkloadStatusOnboardingStatus, error) {
	vals := map[string]ManagedTenantsWorkloadStatusOnboardingStatus{
		"notonboarded": ManagedTenantsWorkloadStatusOnboardingStatus_NotOnboarded,
		"onboarded":    ManagedTenantsWorkloadStatusOnboardingStatus_Onboarded,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ManagedTenantsWorkloadStatusOnboardingStatus(input)
	return &out, nil
}
