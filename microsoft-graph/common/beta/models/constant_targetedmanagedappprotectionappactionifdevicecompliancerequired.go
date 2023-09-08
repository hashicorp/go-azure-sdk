package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TargetedManagedAppProtectionAppActionIfDeviceComplianceRequired string

const (
	TargetedManagedAppProtectionAppActionIfDeviceComplianceRequiredblock TargetedManagedAppProtectionAppActionIfDeviceComplianceRequired = "Block"
	TargetedManagedAppProtectionAppActionIfDeviceComplianceRequiredwarn  TargetedManagedAppProtectionAppActionIfDeviceComplianceRequired = "Warn"
	TargetedManagedAppProtectionAppActionIfDeviceComplianceRequiredwipe  TargetedManagedAppProtectionAppActionIfDeviceComplianceRequired = "Wipe"
)

func PossibleValuesForTargetedManagedAppProtectionAppActionIfDeviceComplianceRequired() []string {
	return []string{
		string(TargetedManagedAppProtectionAppActionIfDeviceComplianceRequiredblock),
		string(TargetedManagedAppProtectionAppActionIfDeviceComplianceRequiredwarn),
		string(TargetedManagedAppProtectionAppActionIfDeviceComplianceRequiredwipe),
	}
}

func parseTargetedManagedAppProtectionAppActionIfDeviceComplianceRequired(input string) (*TargetedManagedAppProtectionAppActionIfDeviceComplianceRequired, error) {
	vals := map[string]TargetedManagedAppProtectionAppActionIfDeviceComplianceRequired{
		"block": TargetedManagedAppProtectionAppActionIfDeviceComplianceRequiredblock,
		"warn":  TargetedManagedAppProtectionAppActionIfDeviceComplianceRequiredwarn,
		"wipe":  TargetedManagedAppProtectionAppActionIfDeviceComplianceRequiredwipe,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := TargetedManagedAppProtectionAppActionIfDeviceComplianceRequired(input)
	return &out, nil
}
