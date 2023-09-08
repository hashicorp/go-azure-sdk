package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type HardwareInformationDeviceGuardLocalSystemAuthorityCredentialGuardState string

const (
	HardwareInformationDeviceGuardLocalSystemAuthorityCredentialGuardStatenotConfigured                         HardwareInformationDeviceGuardLocalSystemAuthorityCredentialGuardState = "NotConfigured"
	HardwareInformationDeviceGuardLocalSystemAuthorityCredentialGuardStatenotLicensed                           HardwareInformationDeviceGuardLocalSystemAuthorityCredentialGuardState = "NotLicensed"
	HardwareInformationDeviceGuardLocalSystemAuthorityCredentialGuardStaterebootRequired                        HardwareInformationDeviceGuardLocalSystemAuthorityCredentialGuardState = "RebootRequired"
	HardwareInformationDeviceGuardLocalSystemAuthorityCredentialGuardStaterunning                               HardwareInformationDeviceGuardLocalSystemAuthorityCredentialGuardState = "Running"
	HardwareInformationDeviceGuardLocalSystemAuthorityCredentialGuardStatevirtualizationBasedSecurityNotRunning HardwareInformationDeviceGuardLocalSystemAuthorityCredentialGuardState = "VirtualizationBasedSecurityNotRunning"
)

func PossibleValuesForHardwareInformationDeviceGuardLocalSystemAuthorityCredentialGuardState() []string {
	return []string{
		string(HardwareInformationDeviceGuardLocalSystemAuthorityCredentialGuardStatenotConfigured),
		string(HardwareInformationDeviceGuardLocalSystemAuthorityCredentialGuardStatenotLicensed),
		string(HardwareInformationDeviceGuardLocalSystemAuthorityCredentialGuardStaterebootRequired),
		string(HardwareInformationDeviceGuardLocalSystemAuthorityCredentialGuardStaterunning),
		string(HardwareInformationDeviceGuardLocalSystemAuthorityCredentialGuardStatevirtualizationBasedSecurityNotRunning),
	}
}

func parseHardwareInformationDeviceGuardLocalSystemAuthorityCredentialGuardState(input string) (*HardwareInformationDeviceGuardLocalSystemAuthorityCredentialGuardState, error) {
	vals := map[string]HardwareInformationDeviceGuardLocalSystemAuthorityCredentialGuardState{
		"notconfigured":                         HardwareInformationDeviceGuardLocalSystemAuthorityCredentialGuardStatenotConfigured,
		"notlicensed":                           HardwareInformationDeviceGuardLocalSystemAuthorityCredentialGuardStatenotLicensed,
		"rebootrequired":                        HardwareInformationDeviceGuardLocalSystemAuthorityCredentialGuardStaterebootRequired,
		"running":                               HardwareInformationDeviceGuardLocalSystemAuthorityCredentialGuardStaterunning,
		"virtualizationbasedsecuritynotrunning": HardwareInformationDeviceGuardLocalSystemAuthorityCredentialGuardStatevirtualizationBasedSecurityNotRunning,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := HardwareInformationDeviceGuardLocalSystemAuthorityCredentialGuardState(input)
	return &out, nil
}
