package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IosManagedAppProtectionAppActionIfDeviceComplianceRequired string

const (
	IosManagedAppProtectionAppActionIfDeviceComplianceRequiredblock IosManagedAppProtectionAppActionIfDeviceComplianceRequired = "Block"
	IosManagedAppProtectionAppActionIfDeviceComplianceRequiredwarn  IosManagedAppProtectionAppActionIfDeviceComplianceRequired = "Warn"
	IosManagedAppProtectionAppActionIfDeviceComplianceRequiredwipe  IosManagedAppProtectionAppActionIfDeviceComplianceRequired = "Wipe"
)

func PossibleValuesForIosManagedAppProtectionAppActionIfDeviceComplianceRequired() []string {
	return []string{
		string(IosManagedAppProtectionAppActionIfDeviceComplianceRequiredblock),
		string(IosManagedAppProtectionAppActionIfDeviceComplianceRequiredwarn),
		string(IosManagedAppProtectionAppActionIfDeviceComplianceRequiredwipe),
	}
}

func parseIosManagedAppProtectionAppActionIfDeviceComplianceRequired(input string) (*IosManagedAppProtectionAppActionIfDeviceComplianceRequired, error) {
	vals := map[string]IosManagedAppProtectionAppActionIfDeviceComplianceRequired{
		"block": IosManagedAppProtectionAppActionIfDeviceComplianceRequiredblock,
		"warn":  IosManagedAppProtectionAppActionIfDeviceComplianceRequiredwarn,
		"wipe":  IosManagedAppProtectionAppActionIfDeviceComplianceRequiredwipe,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := IosManagedAppProtectionAppActionIfDeviceComplianceRequired(input)
	return &out, nil
}
