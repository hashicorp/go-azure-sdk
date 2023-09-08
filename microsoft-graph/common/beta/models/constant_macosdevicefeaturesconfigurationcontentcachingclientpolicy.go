package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MacOSDeviceFeaturesConfigurationContentCachingClientPolicy string

const (
	MacOSDeviceFeaturesConfigurationContentCachingClientPolicyclientsInCustomLocalNetworks             MacOSDeviceFeaturesConfigurationContentCachingClientPolicy = "ClientsInCustomLocalNetworks"
	MacOSDeviceFeaturesConfigurationContentCachingClientPolicyclientsInCustomLocalNetworksWithFallback MacOSDeviceFeaturesConfigurationContentCachingClientPolicy = "ClientsInCustomLocalNetworksWithFallback"
	MacOSDeviceFeaturesConfigurationContentCachingClientPolicyclientsInLocalNetwork                    MacOSDeviceFeaturesConfigurationContentCachingClientPolicy = "ClientsInLocalNetwork"
	MacOSDeviceFeaturesConfigurationContentCachingClientPolicyclientsWithSamePublicIpAddress           MacOSDeviceFeaturesConfigurationContentCachingClientPolicy = "ClientsWithSamePublicIpAddress"
	MacOSDeviceFeaturesConfigurationContentCachingClientPolicynotConfigured                            MacOSDeviceFeaturesConfigurationContentCachingClientPolicy = "NotConfigured"
)

func PossibleValuesForMacOSDeviceFeaturesConfigurationContentCachingClientPolicy() []string {
	return []string{
		string(MacOSDeviceFeaturesConfigurationContentCachingClientPolicyclientsInCustomLocalNetworks),
		string(MacOSDeviceFeaturesConfigurationContentCachingClientPolicyclientsInCustomLocalNetworksWithFallback),
		string(MacOSDeviceFeaturesConfigurationContentCachingClientPolicyclientsInLocalNetwork),
		string(MacOSDeviceFeaturesConfigurationContentCachingClientPolicyclientsWithSamePublicIpAddress),
		string(MacOSDeviceFeaturesConfigurationContentCachingClientPolicynotConfigured),
	}
}

func parseMacOSDeviceFeaturesConfigurationContentCachingClientPolicy(input string) (*MacOSDeviceFeaturesConfigurationContentCachingClientPolicy, error) {
	vals := map[string]MacOSDeviceFeaturesConfigurationContentCachingClientPolicy{
		"clientsincustomlocalnetworks":             MacOSDeviceFeaturesConfigurationContentCachingClientPolicyclientsInCustomLocalNetworks,
		"clientsincustomlocalnetworkswithfallback": MacOSDeviceFeaturesConfigurationContentCachingClientPolicyclientsInCustomLocalNetworksWithFallback,
		"clientsinlocalnetwork":                    MacOSDeviceFeaturesConfigurationContentCachingClientPolicyclientsInLocalNetwork,
		"clientswithsamepublicipaddress":           MacOSDeviceFeaturesConfigurationContentCachingClientPolicyclientsWithSamePublicIpAddress,
		"notconfigured":                            MacOSDeviceFeaturesConfigurationContentCachingClientPolicynotConfigured,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := MacOSDeviceFeaturesConfigurationContentCachingClientPolicy(input)
	return &out, nil
}
