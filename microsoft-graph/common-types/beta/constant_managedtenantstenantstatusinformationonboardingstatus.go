package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ManagedTenantsTenantStatusInformationOnboardingStatus string

const (
	ManagedTenantsTenantStatusInformationOnboardingStatus_Active     ManagedTenantsTenantStatusInformationOnboardingStatus = "active"
	ManagedTenantsTenantStatusInformationOnboardingStatus_InProcess  ManagedTenantsTenantStatusInformationOnboardingStatus = "inProcess"
	ManagedTenantsTenantStatusInformationOnboardingStatus_Inactive   ManagedTenantsTenantStatusInformationOnboardingStatus = "inactive"
	ManagedTenantsTenantStatusInformationOnboardingStatus_Ineligible ManagedTenantsTenantStatusInformationOnboardingStatus = "ineligible"
)

func PossibleValuesForManagedTenantsTenantStatusInformationOnboardingStatus() []string {
	return []string{
		string(ManagedTenantsTenantStatusInformationOnboardingStatus_Active),
		string(ManagedTenantsTenantStatusInformationOnboardingStatus_InProcess),
		string(ManagedTenantsTenantStatusInformationOnboardingStatus_Inactive),
		string(ManagedTenantsTenantStatusInformationOnboardingStatus_Ineligible),
	}
}

func (s *ManagedTenantsTenantStatusInformationOnboardingStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseManagedTenantsTenantStatusInformationOnboardingStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseManagedTenantsTenantStatusInformationOnboardingStatus(input string) (*ManagedTenantsTenantStatusInformationOnboardingStatus, error) {
	vals := map[string]ManagedTenantsTenantStatusInformationOnboardingStatus{
		"active":     ManagedTenantsTenantStatusInformationOnboardingStatus_Active,
		"inprocess":  ManagedTenantsTenantStatusInformationOnboardingStatus_InProcess,
		"inactive":   ManagedTenantsTenantStatusInformationOnboardingStatus_Inactive,
		"ineligible": ManagedTenantsTenantStatusInformationOnboardingStatus_Ineligible,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ManagedTenantsTenantStatusInformationOnboardingStatus(input)
	return &out, nil
}
