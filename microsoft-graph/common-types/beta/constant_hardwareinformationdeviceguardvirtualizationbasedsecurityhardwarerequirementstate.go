package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type HardwareInformationDeviceGuardVirtualizationBasedSecurityHardwareRequirementState string

const (
	HardwareInformationDeviceGuardVirtualizationBasedSecurityHardwareRequirementState_DmaProtectionRequired        HardwareInformationDeviceGuardVirtualizationBasedSecurityHardwareRequirementState = "dmaProtectionRequired"
	HardwareInformationDeviceGuardVirtualizationBasedSecurityHardwareRequirementState_HyperVNotAvailable           HardwareInformationDeviceGuardVirtualizationBasedSecurityHardwareRequirementState = "hyperVNotAvailable"
	HardwareInformationDeviceGuardVirtualizationBasedSecurityHardwareRequirementState_HyperVNotSupportedForGuestVM HardwareInformationDeviceGuardVirtualizationBasedSecurityHardwareRequirementState = "hyperVNotSupportedForGuestVM"
	HardwareInformationDeviceGuardVirtualizationBasedSecurityHardwareRequirementState_MeetHardwareRequirements     HardwareInformationDeviceGuardVirtualizationBasedSecurityHardwareRequirementState = "meetHardwareRequirements"
	HardwareInformationDeviceGuardVirtualizationBasedSecurityHardwareRequirementState_SecureBootRequired           HardwareInformationDeviceGuardVirtualizationBasedSecurityHardwareRequirementState = "secureBootRequired"
)

func PossibleValuesForHardwareInformationDeviceGuardVirtualizationBasedSecurityHardwareRequirementState() []string {
	return []string{
		string(HardwareInformationDeviceGuardVirtualizationBasedSecurityHardwareRequirementState_DmaProtectionRequired),
		string(HardwareInformationDeviceGuardVirtualizationBasedSecurityHardwareRequirementState_HyperVNotAvailable),
		string(HardwareInformationDeviceGuardVirtualizationBasedSecurityHardwareRequirementState_HyperVNotSupportedForGuestVM),
		string(HardwareInformationDeviceGuardVirtualizationBasedSecurityHardwareRequirementState_MeetHardwareRequirements),
		string(HardwareInformationDeviceGuardVirtualizationBasedSecurityHardwareRequirementState_SecureBootRequired),
	}
}

func (s *HardwareInformationDeviceGuardVirtualizationBasedSecurityHardwareRequirementState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseHardwareInformationDeviceGuardVirtualizationBasedSecurityHardwareRequirementState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseHardwareInformationDeviceGuardVirtualizationBasedSecurityHardwareRequirementState(input string) (*HardwareInformationDeviceGuardVirtualizationBasedSecurityHardwareRequirementState, error) {
	vals := map[string]HardwareInformationDeviceGuardVirtualizationBasedSecurityHardwareRequirementState{
		"dmaprotectionrequired":        HardwareInformationDeviceGuardVirtualizationBasedSecurityHardwareRequirementState_DmaProtectionRequired,
		"hypervnotavailable":           HardwareInformationDeviceGuardVirtualizationBasedSecurityHardwareRequirementState_HyperVNotAvailable,
		"hypervnotsupportedforguestvm": HardwareInformationDeviceGuardVirtualizationBasedSecurityHardwareRequirementState_HyperVNotSupportedForGuestVM,
		"meethardwarerequirements":     HardwareInformationDeviceGuardVirtualizationBasedSecurityHardwareRequirementState_MeetHardwareRequirements,
		"securebootrequired":           HardwareInformationDeviceGuardVirtualizationBasedSecurityHardwareRequirementState_SecureBootRequired,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := HardwareInformationDeviceGuardVirtualizationBasedSecurityHardwareRequirementState(input)
	return &out, nil
}
