package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Windows10EndpointProtectionConfigurationDmaGuardDeviceEnumerationPolicy string

const (
	Windows10EndpointProtectionConfigurationDmaGuardDeviceEnumerationPolicyallowAll      Windows10EndpointProtectionConfigurationDmaGuardDeviceEnumerationPolicy = "AllowAll"
	Windows10EndpointProtectionConfigurationDmaGuardDeviceEnumerationPolicyblockAll      Windows10EndpointProtectionConfigurationDmaGuardDeviceEnumerationPolicy = "BlockAll"
	Windows10EndpointProtectionConfigurationDmaGuardDeviceEnumerationPolicydeviceDefault Windows10EndpointProtectionConfigurationDmaGuardDeviceEnumerationPolicy = "DeviceDefault"
)

func PossibleValuesForWindows10EndpointProtectionConfigurationDmaGuardDeviceEnumerationPolicy() []string {
	return []string{
		string(Windows10EndpointProtectionConfigurationDmaGuardDeviceEnumerationPolicyallowAll),
		string(Windows10EndpointProtectionConfigurationDmaGuardDeviceEnumerationPolicyblockAll),
		string(Windows10EndpointProtectionConfigurationDmaGuardDeviceEnumerationPolicydeviceDefault),
	}
}

func parseWindows10EndpointProtectionConfigurationDmaGuardDeviceEnumerationPolicy(input string) (*Windows10EndpointProtectionConfigurationDmaGuardDeviceEnumerationPolicy, error) {
	vals := map[string]Windows10EndpointProtectionConfigurationDmaGuardDeviceEnumerationPolicy{
		"allowall":      Windows10EndpointProtectionConfigurationDmaGuardDeviceEnumerationPolicyallowAll,
		"blockall":      Windows10EndpointProtectionConfigurationDmaGuardDeviceEnumerationPolicyblockAll,
		"devicedefault": Windows10EndpointProtectionConfigurationDmaGuardDeviceEnumerationPolicydeviceDefault,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Windows10EndpointProtectionConfigurationDmaGuardDeviceEnumerationPolicy(input)
	return &out, nil
}
