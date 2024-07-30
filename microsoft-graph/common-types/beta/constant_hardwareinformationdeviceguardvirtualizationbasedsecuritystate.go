package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type HardwareInformationDeviceGuardVirtualizationBasedSecurityState string

const (
	HardwareInformationDeviceGuardVirtualizationBasedSecurityState_DoesNotMeetHardwareRequirements HardwareInformationDeviceGuardVirtualizationBasedSecurityState = "doesNotMeetHardwareRequirements"
	HardwareInformationDeviceGuardVirtualizationBasedSecurityState_NotConfigured                   HardwareInformationDeviceGuardVirtualizationBasedSecurityState = "notConfigured"
	HardwareInformationDeviceGuardVirtualizationBasedSecurityState_NotLicensed                     HardwareInformationDeviceGuardVirtualizationBasedSecurityState = "notLicensed"
	HardwareInformationDeviceGuardVirtualizationBasedSecurityState_Other                           HardwareInformationDeviceGuardVirtualizationBasedSecurityState = "other"
	HardwareInformationDeviceGuardVirtualizationBasedSecurityState_RebootRequired                  HardwareInformationDeviceGuardVirtualizationBasedSecurityState = "rebootRequired"
	HardwareInformationDeviceGuardVirtualizationBasedSecurityState_Require64BitArchitecture        HardwareInformationDeviceGuardVirtualizationBasedSecurityState = "require64BitArchitecture"
	HardwareInformationDeviceGuardVirtualizationBasedSecurityState_Running                         HardwareInformationDeviceGuardVirtualizationBasedSecurityState = "running"
)

func PossibleValuesForHardwareInformationDeviceGuardVirtualizationBasedSecurityState() []string {
	return []string{
		string(HardwareInformationDeviceGuardVirtualizationBasedSecurityState_DoesNotMeetHardwareRequirements),
		string(HardwareInformationDeviceGuardVirtualizationBasedSecurityState_NotConfigured),
		string(HardwareInformationDeviceGuardVirtualizationBasedSecurityState_NotLicensed),
		string(HardwareInformationDeviceGuardVirtualizationBasedSecurityState_Other),
		string(HardwareInformationDeviceGuardVirtualizationBasedSecurityState_RebootRequired),
		string(HardwareInformationDeviceGuardVirtualizationBasedSecurityState_Require64BitArchitecture),
		string(HardwareInformationDeviceGuardVirtualizationBasedSecurityState_Running),
	}
}

func (s *HardwareInformationDeviceGuardVirtualizationBasedSecurityState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseHardwareInformationDeviceGuardVirtualizationBasedSecurityState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseHardwareInformationDeviceGuardVirtualizationBasedSecurityState(input string) (*HardwareInformationDeviceGuardVirtualizationBasedSecurityState, error) {
	vals := map[string]HardwareInformationDeviceGuardVirtualizationBasedSecurityState{
		"doesnotmeethardwarerequirements": HardwareInformationDeviceGuardVirtualizationBasedSecurityState_DoesNotMeetHardwareRequirements,
		"notconfigured":                   HardwareInformationDeviceGuardVirtualizationBasedSecurityState_NotConfigured,
		"notlicensed":                     HardwareInformationDeviceGuardVirtualizationBasedSecurityState_NotLicensed,
		"other":                           HardwareInformationDeviceGuardVirtualizationBasedSecurityState_Other,
		"rebootrequired":                  HardwareInformationDeviceGuardVirtualizationBasedSecurityState_RebootRequired,
		"require64bitarchitecture":        HardwareInformationDeviceGuardVirtualizationBasedSecurityState_Require64BitArchitecture,
		"running":                         HardwareInformationDeviceGuardVirtualizationBasedSecurityState_Running,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := HardwareInformationDeviceGuardVirtualizationBasedSecurityState(input)
	return &out, nil
}
