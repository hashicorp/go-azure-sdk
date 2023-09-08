package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DefaultManagedAppProtectionAppActionIfDeviceComplianceRequired string

const (
	DefaultManagedAppProtectionAppActionIfDeviceComplianceRequiredblock DefaultManagedAppProtectionAppActionIfDeviceComplianceRequired = "Block"
	DefaultManagedAppProtectionAppActionIfDeviceComplianceRequiredwarn  DefaultManagedAppProtectionAppActionIfDeviceComplianceRequired = "Warn"
	DefaultManagedAppProtectionAppActionIfDeviceComplianceRequiredwipe  DefaultManagedAppProtectionAppActionIfDeviceComplianceRequired = "Wipe"
)

func PossibleValuesForDefaultManagedAppProtectionAppActionIfDeviceComplianceRequired() []string {
	return []string{
		string(DefaultManagedAppProtectionAppActionIfDeviceComplianceRequiredblock),
		string(DefaultManagedAppProtectionAppActionIfDeviceComplianceRequiredwarn),
		string(DefaultManagedAppProtectionAppActionIfDeviceComplianceRequiredwipe),
	}
}

func parseDefaultManagedAppProtectionAppActionIfDeviceComplianceRequired(input string) (*DefaultManagedAppProtectionAppActionIfDeviceComplianceRequired, error) {
	vals := map[string]DefaultManagedAppProtectionAppActionIfDeviceComplianceRequired{
		"block": DefaultManagedAppProtectionAppActionIfDeviceComplianceRequiredblock,
		"warn":  DefaultManagedAppProtectionAppActionIfDeviceComplianceRequiredwarn,
		"wipe":  DefaultManagedAppProtectionAppActionIfDeviceComplianceRequiredwipe,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DefaultManagedAppProtectionAppActionIfDeviceComplianceRequired(input)
	return &out, nil
}
