package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Windows10EndpointProtectionConfigurationDeviceGuardSecureBootWithDMA string

const (
	Windows10EndpointProtectionConfigurationDeviceGuardSecureBootWithDMA_NotConfigured Windows10EndpointProtectionConfigurationDeviceGuardSecureBootWithDMA = "notConfigured"
	Windows10EndpointProtectionConfigurationDeviceGuardSecureBootWithDMA_WithDMA       Windows10EndpointProtectionConfigurationDeviceGuardSecureBootWithDMA = "withDMA"
	Windows10EndpointProtectionConfigurationDeviceGuardSecureBootWithDMA_WithoutDMA    Windows10EndpointProtectionConfigurationDeviceGuardSecureBootWithDMA = "withoutDMA"
)

func PossibleValuesForWindows10EndpointProtectionConfigurationDeviceGuardSecureBootWithDMA() []string {
	return []string{
		string(Windows10EndpointProtectionConfigurationDeviceGuardSecureBootWithDMA_NotConfigured),
		string(Windows10EndpointProtectionConfigurationDeviceGuardSecureBootWithDMA_WithDMA),
		string(Windows10EndpointProtectionConfigurationDeviceGuardSecureBootWithDMA_WithoutDMA),
	}
}

func (s *Windows10EndpointProtectionConfigurationDeviceGuardSecureBootWithDMA) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseWindows10EndpointProtectionConfigurationDeviceGuardSecureBootWithDMA(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseWindows10EndpointProtectionConfigurationDeviceGuardSecureBootWithDMA(input string) (*Windows10EndpointProtectionConfigurationDeviceGuardSecureBootWithDMA, error) {
	vals := map[string]Windows10EndpointProtectionConfigurationDeviceGuardSecureBootWithDMA{
		"notconfigured": Windows10EndpointProtectionConfigurationDeviceGuardSecureBootWithDMA_NotConfigured,
		"withdma":       Windows10EndpointProtectionConfigurationDeviceGuardSecureBootWithDMA_WithDMA,
		"withoutdma":    Windows10EndpointProtectionConfigurationDeviceGuardSecureBootWithDMA_WithoutDMA,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Windows10EndpointProtectionConfigurationDeviceGuardSecureBootWithDMA(input)
	return &out, nil
}
