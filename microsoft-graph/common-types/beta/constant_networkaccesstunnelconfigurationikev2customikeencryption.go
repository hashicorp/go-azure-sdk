package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type NetworkaccessTunnelConfigurationIKEv2CustomIkeEncryption string

const (
	NetworkaccessTunnelConfigurationIKEv2CustomIkeEncryption_Aes128    NetworkaccessTunnelConfigurationIKEv2CustomIkeEncryption = "aes128"
	NetworkaccessTunnelConfigurationIKEv2CustomIkeEncryption_Aes192    NetworkaccessTunnelConfigurationIKEv2CustomIkeEncryption = "aes192"
	NetworkaccessTunnelConfigurationIKEv2CustomIkeEncryption_Aes256    NetworkaccessTunnelConfigurationIKEv2CustomIkeEncryption = "aes256"
	NetworkaccessTunnelConfigurationIKEv2CustomIkeEncryption_GcmAes128 NetworkaccessTunnelConfigurationIKEv2CustomIkeEncryption = "gcmAes128"
	NetworkaccessTunnelConfigurationIKEv2CustomIkeEncryption_GcmAes256 NetworkaccessTunnelConfigurationIKEv2CustomIkeEncryption = "gcmAes256"
)

func PossibleValuesForNetworkaccessTunnelConfigurationIKEv2CustomIkeEncryption() []string {
	return []string{
		string(NetworkaccessTunnelConfigurationIKEv2CustomIkeEncryption_Aes128),
		string(NetworkaccessTunnelConfigurationIKEv2CustomIkeEncryption_Aes192),
		string(NetworkaccessTunnelConfigurationIKEv2CustomIkeEncryption_Aes256),
		string(NetworkaccessTunnelConfigurationIKEv2CustomIkeEncryption_GcmAes128),
		string(NetworkaccessTunnelConfigurationIKEv2CustomIkeEncryption_GcmAes256),
	}
}

func (s *NetworkaccessTunnelConfigurationIKEv2CustomIkeEncryption) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseNetworkaccessTunnelConfigurationIKEv2CustomIkeEncryption(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseNetworkaccessTunnelConfigurationIKEv2CustomIkeEncryption(input string) (*NetworkaccessTunnelConfigurationIKEv2CustomIkeEncryption, error) {
	vals := map[string]NetworkaccessTunnelConfigurationIKEv2CustomIkeEncryption{
		"aes128":    NetworkaccessTunnelConfigurationIKEv2CustomIkeEncryption_Aes128,
		"aes192":    NetworkaccessTunnelConfigurationIKEv2CustomIkeEncryption_Aes192,
		"aes256":    NetworkaccessTunnelConfigurationIKEv2CustomIkeEncryption_Aes256,
		"gcmaes128": NetworkaccessTunnelConfigurationIKEv2CustomIkeEncryption_GcmAes128,
		"gcmaes256": NetworkaccessTunnelConfigurationIKEv2CustomIkeEncryption_GcmAes256,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := NetworkaccessTunnelConfigurationIKEv2CustomIkeEncryption(input)
	return &out, nil
}
