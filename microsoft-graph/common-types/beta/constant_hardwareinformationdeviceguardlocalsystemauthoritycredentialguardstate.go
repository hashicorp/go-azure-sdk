package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type HardwareInformationDeviceGuardLocalSystemAuthorityCredentialGuardState string

const (
	HardwareInformationDeviceGuardLocalSystemAuthorityCredentialGuardState_NotConfigured                         HardwareInformationDeviceGuardLocalSystemAuthorityCredentialGuardState = "notConfigured"
	HardwareInformationDeviceGuardLocalSystemAuthorityCredentialGuardState_NotLicensed                           HardwareInformationDeviceGuardLocalSystemAuthorityCredentialGuardState = "notLicensed"
	HardwareInformationDeviceGuardLocalSystemAuthorityCredentialGuardState_RebootRequired                        HardwareInformationDeviceGuardLocalSystemAuthorityCredentialGuardState = "rebootRequired"
	HardwareInformationDeviceGuardLocalSystemAuthorityCredentialGuardState_Running                               HardwareInformationDeviceGuardLocalSystemAuthorityCredentialGuardState = "running"
	HardwareInformationDeviceGuardLocalSystemAuthorityCredentialGuardState_VirtualizationBasedSecurityNotRunning HardwareInformationDeviceGuardLocalSystemAuthorityCredentialGuardState = "virtualizationBasedSecurityNotRunning"
)

func PossibleValuesForHardwareInformationDeviceGuardLocalSystemAuthorityCredentialGuardState() []string {
	return []string{
		string(HardwareInformationDeviceGuardLocalSystemAuthorityCredentialGuardState_NotConfigured),
		string(HardwareInformationDeviceGuardLocalSystemAuthorityCredentialGuardState_NotLicensed),
		string(HardwareInformationDeviceGuardLocalSystemAuthorityCredentialGuardState_RebootRequired),
		string(HardwareInformationDeviceGuardLocalSystemAuthorityCredentialGuardState_Running),
		string(HardwareInformationDeviceGuardLocalSystemAuthorityCredentialGuardState_VirtualizationBasedSecurityNotRunning),
	}
}

func (s *HardwareInformationDeviceGuardLocalSystemAuthorityCredentialGuardState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseHardwareInformationDeviceGuardLocalSystemAuthorityCredentialGuardState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseHardwareInformationDeviceGuardLocalSystemAuthorityCredentialGuardState(input string) (*HardwareInformationDeviceGuardLocalSystemAuthorityCredentialGuardState, error) {
	vals := map[string]HardwareInformationDeviceGuardLocalSystemAuthorityCredentialGuardState{
		"notconfigured":                         HardwareInformationDeviceGuardLocalSystemAuthorityCredentialGuardState_NotConfigured,
		"notlicensed":                           HardwareInformationDeviceGuardLocalSystemAuthorityCredentialGuardState_NotLicensed,
		"rebootrequired":                        HardwareInformationDeviceGuardLocalSystemAuthorityCredentialGuardState_RebootRequired,
		"running":                               HardwareInformationDeviceGuardLocalSystemAuthorityCredentialGuardState_Running,
		"virtualizationbasedsecuritynotrunning": HardwareInformationDeviceGuardLocalSystemAuthorityCredentialGuardState_VirtualizationBasedSecurityNotRunning,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := HardwareInformationDeviceGuardLocalSystemAuthorityCredentialGuardState(input)
	return &out, nil
}
