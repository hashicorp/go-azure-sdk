package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Windows10EndpointProtectionConfigurationDmaGuardDeviceEnumerationPolicy string

const (
	Windows10EndpointProtectionConfigurationDmaGuardDeviceEnumerationPolicy_AllowAll      Windows10EndpointProtectionConfigurationDmaGuardDeviceEnumerationPolicy = "allowAll"
	Windows10EndpointProtectionConfigurationDmaGuardDeviceEnumerationPolicy_BlockAll      Windows10EndpointProtectionConfigurationDmaGuardDeviceEnumerationPolicy = "blockAll"
	Windows10EndpointProtectionConfigurationDmaGuardDeviceEnumerationPolicy_DeviceDefault Windows10EndpointProtectionConfigurationDmaGuardDeviceEnumerationPolicy = "deviceDefault"
)

func PossibleValuesForWindows10EndpointProtectionConfigurationDmaGuardDeviceEnumerationPolicy() []string {
	return []string{
		string(Windows10EndpointProtectionConfigurationDmaGuardDeviceEnumerationPolicy_AllowAll),
		string(Windows10EndpointProtectionConfigurationDmaGuardDeviceEnumerationPolicy_BlockAll),
		string(Windows10EndpointProtectionConfigurationDmaGuardDeviceEnumerationPolicy_DeviceDefault),
	}
}

func (s *Windows10EndpointProtectionConfigurationDmaGuardDeviceEnumerationPolicy) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseWindows10EndpointProtectionConfigurationDmaGuardDeviceEnumerationPolicy(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseWindows10EndpointProtectionConfigurationDmaGuardDeviceEnumerationPolicy(input string) (*Windows10EndpointProtectionConfigurationDmaGuardDeviceEnumerationPolicy, error) {
	vals := map[string]Windows10EndpointProtectionConfigurationDmaGuardDeviceEnumerationPolicy{
		"allowall":      Windows10EndpointProtectionConfigurationDmaGuardDeviceEnumerationPolicy_AllowAll,
		"blockall":      Windows10EndpointProtectionConfigurationDmaGuardDeviceEnumerationPolicy_BlockAll,
		"devicedefault": Windows10EndpointProtectionConfigurationDmaGuardDeviceEnumerationPolicy_DeviceDefault,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Windows10EndpointProtectionConfigurationDmaGuardDeviceEnumerationPolicy(input)
	return &out, nil
}
