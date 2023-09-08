package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type HardwareInformationDeviceGuardVirtualizationBasedSecurityState string

const (
	HardwareInformationDeviceGuardVirtualizationBasedSecurityStatedoesNotMeetHardwareRequirements HardwareInformationDeviceGuardVirtualizationBasedSecurityState = "DoesNotMeetHardwareRequirements"
	HardwareInformationDeviceGuardVirtualizationBasedSecurityStatenotConfigured                   HardwareInformationDeviceGuardVirtualizationBasedSecurityState = "NotConfigured"
	HardwareInformationDeviceGuardVirtualizationBasedSecurityStatenotLicensed                     HardwareInformationDeviceGuardVirtualizationBasedSecurityState = "NotLicensed"
	HardwareInformationDeviceGuardVirtualizationBasedSecurityStateother                           HardwareInformationDeviceGuardVirtualizationBasedSecurityState = "Other"
	HardwareInformationDeviceGuardVirtualizationBasedSecurityStaterebootRequired                  HardwareInformationDeviceGuardVirtualizationBasedSecurityState = "RebootRequired"
	HardwareInformationDeviceGuardVirtualizationBasedSecurityStaterequire64BitArchitecture        HardwareInformationDeviceGuardVirtualizationBasedSecurityState = "Require64BitArchitecture"
	HardwareInformationDeviceGuardVirtualizationBasedSecurityStaterunning                         HardwareInformationDeviceGuardVirtualizationBasedSecurityState = "Running"
)

func PossibleValuesForHardwareInformationDeviceGuardVirtualizationBasedSecurityState() []string {
	return []string{
		string(HardwareInformationDeviceGuardVirtualizationBasedSecurityStatedoesNotMeetHardwareRequirements),
		string(HardwareInformationDeviceGuardVirtualizationBasedSecurityStatenotConfigured),
		string(HardwareInformationDeviceGuardVirtualizationBasedSecurityStatenotLicensed),
		string(HardwareInformationDeviceGuardVirtualizationBasedSecurityStateother),
		string(HardwareInformationDeviceGuardVirtualizationBasedSecurityStaterebootRequired),
		string(HardwareInformationDeviceGuardVirtualizationBasedSecurityStaterequire64BitArchitecture),
		string(HardwareInformationDeviceGuardVirtualizationBasedSecurityStaterunning),
	}
}

func parseHardwareInformationDeviceGuardVirtualizationBasedSecurityState(input string) (*HardwareInformationDeviceGuardVirtualizationBasedSecurityState, error) {
	vals := map[string]HardwareInformationDeviceGuardVirtualizationBasedSecurityState{
		"doesnotmeethardwarerequirements": HardwareInformationDeviceGuardVirtualizationBasedSecurityStatedoesNotMeetHardwareRequirements,
		"notconfigured":                   HardwareInformationDeviceGuardVirtualizationBasedSecurityStatenotConfigured,
		"notlicensed":                     HardwareInformationDeviceGuardVirtualizationBasedSecurityStatenotLicensed,
		"other":                           HardwareInformationDeviceGuardVirtualizationBasedSecurityStateother,
		"rebootrequired":                  HardwareInformationDeviceGuardVirtualizationBasedSecurityStaterebootRequired,
		"require64bitarchitecture":        HardwareInformationDeviceGuardVirtualizationBasedSecurityStaterequire64BitArchitecture,
		"running":                         HardwareInformationDeviceGuardVirtualizationBasedSecurityStaterunning,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := HardwareInformationDeviceGuardVirtualizationBasedSecurityState(input)
	return &out, nil
}
