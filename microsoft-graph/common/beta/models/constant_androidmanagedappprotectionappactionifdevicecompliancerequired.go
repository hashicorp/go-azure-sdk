package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidManagedAppProtectionAppActionIfDeviceComplianceRequired string

const (
	AndroidManagedAppProtectionAppActionIfDeviceComplianceRequiredblock AndroidManagedAppProtectionAppActionIfDeviceComplianceRequired = "Block"
	AndroidManagedAppProtectionAppActionIfDeviceComplianceRequiredwarn  AndroidManagedAppProtectionAppActionIfDeviceComplianceRequired = "Warn"
	AndroidManagedAppProtectionAppActionIfDeviceComplianceRequiredwipe  AndroidManagedAppProtectionAppActionIfDeviceComplianceRequired = "Wipe"
)

func PossibleValuesForAndroidManagedAppProtectionAppActionIfDeviceComplianceRequired() []string {
	return []string{
		string(AndroidManagedAppProtectionAppActionIfDeviceComplianceRequiredblock),
		string(AndroidManagedAppProtectionAppActionIfDeviceComplianceRequiredwarn),
		string(AndroidManagedAppProtectionAppActionIfDeviceComplianceRequiredwipe),
	}
}

func parseAndroidManagedAppProtectionAppActionIfDeviceComplianceRequired(input string) (*AndroidManagedAppProtectionAppActionIfDeviceComplianceRequired, error) {
	vals := map[string]AndroidManagedAppProtectionAppActionIfDeviceComplianceRequired{
		"block": AndroidManagedAppProtectionAppActionIfDeviceComplianceRequiredblock,
		"warn":  AndroidManagedAppProtectionAppActionIfDeviceComplianceRequiredwarn,
		"wipe":  AndroidManagedAppProtectionAppActionIfDeviceComplianceRequiredwipe,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AndroidManagedAppProtectionAppActionIfDeviceComplianceRequired(input)
	return &out, nil
}
