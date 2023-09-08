package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Windows10EndpointProtectionConfigurationDeviceGuardSecureBootWithDMA string

const (
	Windows10EndpointProtectionConfigurationDeviceGuardSecureBootWithDMAnotConfigured Windows10EndpointProtectionConfigurationDeviceGuardSecureBootWithDMA = "NotConfigured"
	Windows10EndpointProtectionConfigurationDeviceGuardSecureBootWithDMAwithDMA       Windows10EndpointProtectionConfigurationDeviceGuardSecureBootWithDMA = "WithDMA"
	Windows10EndpointProtectionConfigurationDeviceGuardSecureBootWithDMAwithoutDMA    Windows10EndpointProtectionConfigurationDeviceGuardSecureBootWithDMA = "WithoutDMA"
)

func PossibleValuesForWindows10EndpointProtectionConfigurationDeviceGuardSecureBootWithDMA() []string {
	return []string{
		string(Windows10EndpointProtectionConfigurationDeviceGuardSecureBootWithDMAnotConfigured),
		string(Windows10EndpointProtectionConfigurationDeviceGuardSecureBootWithDMAwithDMA),
		string(Windows10EndpointProtectionConfigurationDeviceGuardSecureBootWithDMAwithoutDMA),
	}
}

func parseWindows10EndpointProtectionConfigurationDeviceGuardSecureBootWithDMA(input string) (*Windows10EndpointProtectionConfigurationDeviceGuardSecureBootWithDMA, error) {
	vals := map[string]Windows10EndpointProtectionConfigurationDeviceGuardSecureBootWithDMA{
		"notconfigured": Windows10EndpointProtectionConfigurationDeviceGuardSecureBootWithDMAnotConfigured,
		"withdma":       Windows10EndpointProtectionConfigurationDeviceGuardSecureBootWithDMAwithDMA,
		"withoutdma":    Windows10EndpointProtectionConfigurationDeviceGuardSecureBootWithDMAwithoutDMA,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Windows10EndpointProtectionConfigurationDeviceGuardSecureBootWithDMA(input)
	return &out, nil
}
