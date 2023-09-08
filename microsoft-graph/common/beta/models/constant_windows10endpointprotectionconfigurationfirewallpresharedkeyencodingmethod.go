package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Windows10EndpointProtectionConfigurationFirewallPreSharedKeyEncodingMethod string

const (
	Windows10EndpointProtectionConfigurationFirewallPreSharedKeyEncodingMethoddeviceDefault Windows10EndpointProtectionConfigurationFirewallPreSharedKeyEncodingMethod = "DeviceDefault"
	Windows10EndpointProtectionConfigurationFirewallPreSharedKeyEncodingMethodnone          Windows10EndpointProtectionConfigurationFirewallPreSharedKeyEncodingMethod = "None"
	Windows10EndpointProtectionConfigurationFirewallPreSharedKeyEncodingMethodutF8          Windows10EndpointProtectionConfigurationFirewallPreSharedKeyEncodingMethod = "UtF8"
)

func PossibleValuesForWindows10EndpointProtectionConfigurationFirewallPreSharedKeyEncodingMethod() []string {
	return []string{
		string(Windows10EndpointProtectionConfigurationFirewallPreSharedKeyEncodingMethoddeviceDefault),
		string(Windows10EndpointProtectionConfigurationFirewallPreSharedKeyEncodingMethodnone),
		string(Windows10EndpointProtectionConfigurationFirewallPreSharedKeyEncodingMethodutF8),
	}
}

func parseWindows10EndpointProtectionConfigurationFirewallPreSharedKeyEncodingMethod(input string) (*Windows10EndpointProtectionConfigurationFirewallPreSharedKeyEncodingMethod, error) {
	vals := map[string]Windows10EndpointProtectionConfigurationFirewallPreSharedKeyEncodingMethod{
		"devicedefault": Windows10EndpointProtectionConfigurationFirewallPreSharedKeyEncodingMethoddeviceDefault,
		"none":          Windows10EndpointProtectionConfigurationFirewallPreSharedKeyEncodingMethodnone,
		"utf8":          Windows10EndpointProtectionConfigurationFirewallPreSharedKeyEncodingMethodutF8,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Windows10EndpointProtectionConfigurationFirewallPreSharedKeyEncodingMethod(input)
	return &out, nil
}
