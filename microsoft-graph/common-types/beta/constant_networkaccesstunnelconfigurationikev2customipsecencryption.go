package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type NetworkaccessTunnelConfigurationIKEv2CustomIpSecEncryption string

const (
	NetworkaccessTunnelConfigurationIKEv2CustomIpSecEncryption_GcmAes128 NetworkaccessTunnelConfigurationIKEv2CustomIpSecEncryption = "gcmAes128"
	NetworkaccessTunnelConfigurationIKEv2CustomIpSecEncryption_GcmAes192 NetworkaccessTunnelConfigurationIKEv2CustomIpSecEncryption = "gcmAes192"
	NetworkaccessTunnelConfigurationIKEv2CustomIpSecEncryption_GcmAes256 NetworkaccessTunnelConfigurationIKEv2CustomIpSecEncryption = "gcmAes256"
	NetworkaccessTunnelConfigurationIKEv2CustomIpSecEncryption_None      NetworkaccessTunnelConfigurationIKEv2CustomIpSecEncryption = "none"
)

func PossibleValuesForNetworkaccessTunnelConfigurationIKEv2CustomIpSecEncryption() []string {
	return []string{
		string(NetworkaccessTunnelConfigurationIKEv2CustomIpSecEncryption_GcmAes128),
		string(NetworkaccessTunnelConfigurationIKEv2CustomIpSecEncryption_GcmAes192),
		string(NetworkaccessTunnelConfigurationIKEv2CustomIpSecEncryption_GcmAes256),
		string(NetworkaccessTunnelConfigurationIKEv2CustomIpSecEncryption_None),
	}
}

func (s *NetworkaccessTunnelConfigurationIKEv2CustomIpSecEncryption) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseNetworkaccessTunnelConfigurationIKEv2CustomIpSecEncryption(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseNetworkaccessTunnelConfigurationIKEv2CustomIpSecEncryption(input string) (*NetworkaccessTunnelConfigurationIKEv2CustomIpSecEncryption, error) {
	vals := map[string]NetworkaccessTunnelConfigurationIKEv2CustomIpSecEncryption{
		"gcmaes128": NetworkaccessTunnelConfigurationIKEv2CustomIpSecEncryption_GcmAes128,
		"gcmaes192": NetworkaccessTunnelConfigurationIKEv2CustomIpSecEncryption_GcmAes192,
		"gcmaes256": NetworkaccessTunnelConfigurationIKEv2CustomIpSecEncryption_GcmAes256,
		"none":      NetworkaccessTunnelConfigurationIKEv2CustomIpSecEncryption_None,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := NetworkaccessTunnelConfigurationIKEv2CustomIpSecEncryption(input)
	return &out, nil
}
