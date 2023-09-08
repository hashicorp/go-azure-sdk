package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ManagedAppProtectionAppActionIfDeviceComplianceRequired string

const (
	ManagedAppProtectionAppActionIfDeviceComplianceRequiredblock ManagedAppProtectionAppActionIfDeviceComplianceRequired = "Block"
	ManagedAppProtectionAppActionIfDeviceComplianceRequiredwarn  ManagedAppProtectionAppActionIfDeviceComplianceRequired = "Warn"
	ManagedAppProtectionAppActionIfDeviceComplianceRequiredwipe  ManagedAppProtectionAppActionIfDeviceComplianceRequired = "Wipe"
)

func PossibleValuesForManagedAppProtectionAppActionIfDeviceComplianceRequired() []string {
	return []string{
		string(ManagedAppProtectionAppActionIfDeviceComplianceRequiredblock),
		string(ManagedAppProtectionAppActionIfDeviceComplianceRequiredwarn),
		string(ManagedAppProtectionAppActionIfDeviceComplianceRequiredwipe),
	}
}

func parseManagedAppProtectionAppActionIfDeviceComplianceRequired(input string) (*ManagedAppProtectionAppActionIfDeviceComplianceRequired, error) {
	vals := map[string]ManagedAppProtectionAppActionIfDeviceComplianceRequired{
		"block": ManagedAppProtectionAppActionIfDeviceComplianceRequiredblock,
		"warn":  ManagedAppProtectionAppActionIfDeviceComplianceRequiredwarn,
		"wipe":  ManagedAppProtectionAppActionIfDeviceComplianceRequiredwipe,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ManagedAppProtectionAppActionIfDeviceComplianceRequired(input)
	return &out, nil
}
