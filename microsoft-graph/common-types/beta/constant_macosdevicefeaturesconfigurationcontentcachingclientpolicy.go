package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MacOSDeviceFeaturesConfigurationContentCachingClientPolicy string

const (
	MacOSDeviceFeaturesConfigurationContentCachingClientPolicy_ClientsInCustomLocalNetworks             MacOSDeviceFeaturesConfigurationContentCachingClientPolicy = "clientsInCustomLocalNetworks"
	MacOSDeviceFeaturesConfigurationContentCachingClientPolicy_ClientsInCustomLocalNetworksWithFallback MacOSDeviceFeaturesConfigurationContentCachingClientPolicy = "clientsInCustomLocalNetworksWithFallback"
	MacOSDeviceFeaturesConfigurationContentCachingClientPolicy_ClientsInLocalNetwork                    MacOSDeviceFeaturesConfigurationContentCachingClientPolicy = "clientsInLocalNetwork"
	MacOSDeviceFeaturesConfigurationContentCachingClientPolicy_ClientsWithSamePublicIpAddress           MacOSDeviceFeaturesConfigurationContentCachingClientPolicy = "clientsWithSamePublicIpAddress"
	MacOSDeviceFeaturesConfigurationContentCachingClientPolicy_NotConfigured                            MacOSDeviceFeaturesConfigurationContentCachingClientPolicy = "notConfigured"
)

func PossibleValuesForMacOSDeviceFeaturesConfigurationContentCachingClientPolicy() []string {
	return []string{
		string(MacOSDeviceFeaturesConfigurationContentCachingClientPolicy_ClientsInCustomLocalNetworks),
		string(MacOSDeviceFeaturesConfigurationContentCachingClientPolicy_ClientsInCustomLocalNetworksWithFallback),
		string(MacOSDeviceFeaturesConfigurationContentCachingClientPolicy_ClientsInLocalNetwork),
		string(MacOSDeviceFeaturesConfigurationContentCachingClientPolicy_ClientsWithSamePublicIpAddress),
		string(MacOSDeviceFeaturesConfigurationContentCachingClientPolicy_NotConfigured),
	}
}

func (s *MacOSDeviceFeaturesConfigurationContentCachingClientPolicy) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseMacOSDeviceFeaturesConfigurationContentCachingClientPolicy(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseMacOSDeviceFeaturesConfigurationContentCachingClientPolicy(input string) (*MacOSDeviceFeaturesConfigurationContentCachingClientPolicy, error) {
	vals := map[string]MacOSDeviceFeaturesConfigurationContentCachingClientPolicy{
		"clientsincustomlocalnetworks":             MacOSDeviceFeaturesConfigurationContentCachingClientPolicy_ClientsInCustomLocalNetworks,
		"clientsincustomlocalnetworkswithfallback": MacOSDeviceFeaturesConfigurationContentCachingClientPolicy_ClientsInCustomLocalNetworksWithFallback,
		"clientsinlocalnetwork":                    MacOSDeviceFeaturesConfigurationContentCachingClientPolicy_ClientsInLocalNetwork,
		"clientswithsamepublicipaddress":           MacOSDeviceFeaturesConfigurationContentCachingClientPolicy_ClientsWithSamePublicIpAddress,
		"notconfigured":                            MacOSDeviceFeaturesConfigurationContentCachingClientPolicy_NotConfigured,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := MacOSDeviceFeaturesConfigurationContentCachingClientPolicy(input)
	return &out, nil
}
