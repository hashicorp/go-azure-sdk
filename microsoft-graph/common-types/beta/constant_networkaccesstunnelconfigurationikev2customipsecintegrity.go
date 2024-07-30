package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type NetworkaccessTunnelConfigurationIKEv2CustomIpSecIntegrity string

const (
	NetworkaccessTunnelConfigurationIKEv2CustomIpSecIntegrity_GcmAes128 NetworkaccessTunnelConfigurationIKEv2CustomIpSecIntegrity = "gcmAes128"
	NetworkaccessTunnelConfigurationIKEv2CustomIpSecIntegrity_GcmAes192 NetworkaccessTunnelConfigurationIKEv2CustomIpSecIntegrity = "gcmAes192"
	NetworkaccessTunnelConfigurationIKEv2CustomIpSecIntegrity_GcmAes256 NetworkaccessTunnelConfigurationIKEv2CustomIpSecIntegrity = "gcmAes256"
	NetworkaccessTunnelConfigurationIKEv2CustomIpSecIntegrity_Sha256    NetworkaccessTunnelConfigurationIKEv2CustomIpSecIntegrity = "sha256"
)

func PossibleValuesForNetworkaccessTunnelConfigurationIKEv2CustomIpSecIntegrity() []string {
	return []string{
		string(NetworkaccessTunnelConfigurationIKEv2CustomIpSecIntegrity_GcmAes128),
		string(NetworkaccessTunnelConfigurationIKEv2CustomIpSecIntegrity_GcmAes192),
		string(NetworkaccessTunnelConfigurationIKEv2CustomIpSecIntegrity_GcmAes256),
		string(NetworkaccessTunnelConfigurationIKEv2CustomIpSecIntegrity_Sha256),
	}
}

func (s *NetworkaccessTunnelConfigurationIKEv2CustomIpSecIntegrity) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseNetworkaccessTunnelConfigurationIKEv2CustomIpSecIntegrity(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseNetworkaccessTunnelConfigurationIKEv2CustomIpSecIntegrity(input string) (*NetworkaccessTunnelConfigurationIKEv2CustomIpSecIntegrity, error) {
	vals := map[string]NetworkaccessTunnelConfigurationIKEv2CustomIpSecIntegrity{
		"gcmaes128": NetworkaccessTunnelConfigurationIKEv2CustomIpSecIntegrity_GcmAes128,
		"gcmaes192": NetworkaccessTunnelConfigurationIKEv2CustomIpSecIntegrity_GcmAes192,
		"gcmaes256": NetworkaccessTunnelConfigurationIKEv2CustomIpSecIntegrity_GcmAes256,
		"sha256":    NetworkaccessTunnelConfigurationIKEv2CustomIpSecIntegrity_Sha256,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := NetworkaccessTunnelConfigurationIKEv2CustomIpSecIntegrity(input)
	return &out, nil
}
