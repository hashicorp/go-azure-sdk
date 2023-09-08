package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MacOSGeneralDeviceConfigurationUpdateDelayPolicy string

const (
	MacOSGeneralDeviceConfigurationUpdateDelayPolicydelayAppUpdateVisibility     MacOSGeneralDeviceConfigurationUpdateDelayPolicy = "DelayAppUpdateVisibility"
	MacOSGeneralDeviceConfigurationUpdateDelayPolicydelayMajorOsUpdateVisibility MacOSGeneralDeviceConfigurationUpdateDelayPolicy = "DelayMajorOsUpdateVisibility"
	MacOSGeneralDeviceConfigurationUpdateDelayPolicydelayOSUpdateVisibility      MacOSGeneralDeviceConfigurationUpdateDelayPolicy = "DelayOSUpdateVisibility"
	MacOSGeneralDeviceConfigurationUpdateDelayPolicynone                         MacOSGeneralDeviceConfigurationUpdateDelayPolicy = "None"
)

func PossibleValuesForMacOSGeneralDeviceConfigurationUpdateDelayPolicy() []string {
	return []string{
		string(MacOSGeneralDeviceConfigurationUpdateDelayPolicydelayAppUpdateVisibility),
		string(MacOSGeneralDeviceConfigurationUpdateDelayPolicydelayMajorOsUpdateVisibility),
		string(MacOSGeneralDeviceConfigurationUpdateDelayPolicydelayOSUpdateVisibility),
		string(MacOSGeneralDeviceConfigurationUpdateDelayPolicynone),
	}
}

func parseMacOSGeneralDeviceConfigurationUpdateDelayPolicy(input string) (*MacOSGeneralDeviceConfigurationUpdateDelayPolicy, error) {
	vals := map[string]MacOSGeneralDeviceConfigurationUpdateDelayPolicy{
		"delayappupdatevisibility":     MacOSGeneralDeviceConfigurationUpdateDelayPolicydelayAppUpdateVisibility,
		"delaymajorosupdatevisibility": MacOSGeneralDeviceConfigurationUpdateDelayPolicydelayMajorOsUpdateVisibility,
		"delayosupdatevisibility":      MacOSGeneralDeviceConfigurationUpdateDelayPolicydelayOSUpdateVisibility,
		"none":                         MacOSGeneralDeviceConfigurationUpdateDelayPolicynone,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := MacOSGeneralDeviceConfigurationUpdateDelayPolicy(input)
	return &out, nil
}
