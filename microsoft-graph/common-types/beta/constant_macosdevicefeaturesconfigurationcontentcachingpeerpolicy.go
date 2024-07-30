package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MacOSDeviceFeaturesConfigurationContentCachingPeerPolicy string

const (
	MacOSDeviceFeaturesConfigurationContentCachingPeerPolicy_NotConfigured                MacOSDeviceFeaturesConfigurationContentCachingPeerPolicy = "notConfigured"
	MacOSDeviceFeaturesConfigurationContentCachingPeerPolicy_PeersInCustomLocalNetworks   MacOSDeviceFeaturesConfigurationContentCachingPeerPolicy = "peersInCustomLocalNetworks"
	MacOSDeviceFeaturesConfigurationContentCachingPeerPolicy_PeersInLocalNetwork          MacOSDeviceFeaturesConfigurationContentCachingPeerPolicy = "peersInLocalNetwork"
	MacOSDeviceFeaturesConfigurationContentCachingPeerPolicy_PeersWithSamePublicIpAddress MacOSDeviceFeaturesConfigurationContentCachingPeerPolicy = "peersWithSamePublicIpAddress"
)

func PossibleValuesForMacOSDeviceFeaturesConfigurationContentCachingPeerPolicy() []string {
	return []string{
		string(MacOSDeviceFeaturesConfigurationContentCachingPeerPolicy_NotConfigured),
		string(MacOSDeviceFeaturesConfigurationContentCachingPeerPolicy_PeersInCustomLocalNetworks),
		string(MacOSDeviceFeaturesConfigurationContentCachingPeerPolicy_PeersInLocalNetwork),
		string(MacOSDeviceFeaturesConfigurationContentCachingPeerPolicy_PeersWithSamePublicIpAddress),
	}
}

func (s *MacOSDeviceFeaturesConfigurationContentCachingPeerPolicy) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseMacOSDeviceFeaturesConfigurationContentCachingPeerPolicy(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseMacOSDeviceFeaturesConfigurationContentCachingPeerPolicy(input string) (*MacOSDeviceFeaturesConfigurationContentCachingPeerPolicy, error) {
	vals := map[string]MacOSDeviceFeaturesConfigurationContentCachingPeerPolicy{
		"notconfigured":                MacOSDeviceFeaturesConfigurationContentCachingPeerPolicy_NotConfigured,
		"peersincustomlocalnetworks":   MacOSDeviceFeaturesConfigurationContentCachingPeerPolicy_PeersInCustomLocalNetworks,
		"peersinlocalnetwork":          MacOSDeviceFeaturesConfigurationContentCachingPeerPolicy_PeersInLocalNetwork,
		"peerswithsamepublicipaddress": MacOSDeviceFeaturesConfigurationContentCachingPeerPolicy_PeersWithSamePublicIpAddress,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := MacOSDeviceFeaturesConfigurationContentCachingPeerPolicy(input)
	return &out, nil
}
