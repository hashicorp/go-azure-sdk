package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type NetworkaccessTunnelConfigurationIKEv2CustomIkeIntegrity string

const (
	NetworkaccessTunnelConfigurationIKEv2CustomIkeIntegrity_GcmAes128 NetworkaccessTunnelConfigurationIKEv2CustomIkeIntegrity = "gcmAes128"
	NetworkaccessTunnelConfigurationIKEv2CustomIkeIntegrity_GcmAes256 NetworkaccessTunnelConfigurationIKEv2CustomIkeIntegrity = "gcmAes256"
	NetworkaccessTunnelConfigurationIKEv2CustomIkeIntegrity_Sha256    NetworkaccessTunnelConfigurationIKEv2CustomIkeIntegrity = "sha256"
	NetworkaccessTunnelConfigurationIKEv2CustomIkeIntegrity_Sha384    NetworkaccessTunnelConfigurationIKEv2CustomIkeIntegrity = "sha384"
)

func PossibleValuesForNetworkaccessTunnelConfigurationIKEv2CustomIkeIntegrity() []string {
	return []string{
		string(NetworkaccessTunnelConfigurationIKEv2CustomIkeIntegrity_GcmAes128),
		string(NetworkaccessTunnelConfigurationIKEv2CustomIkeIntegrity_GcmAes256),
		string(NetworkaccessTunnelConfigurationIKEv2CustomIkeIntegrity_Sha256),
		string(NetworkaccessTunnelConfigurationIKEv2CustomIkeIntegrity_Sha384),
	}
}

func (s *NetworkaccessTunnelConfigurationIKEv2CustomIkeIntegrity) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseNetworkaccessTunnelConfigurationIKEv2CustomIkeIntegrity(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseNetworkaccessTunnelConfigurationIKEv2CustomIkeIntegrity(input string) (*NetworkaccessTunnelConfigurationIKEv2CustomIkeIntegrity, error) {
	vals := map[string]NetworkaccessTunnelConfigurationIKEv2CustomIkeIntegrity{
		"gcmaes128": NetworkaccessTunnelConfigurationIKEv2CustomIkeIntegrity_GcmAes128,
		"gcmaes256": NetworkaccessTunnelConfigurationIKEv2CustomIkeIntegrity_GcmAes256,
		"sha256":    NetworkaccessTunnelConfigurationIKEv2CustomIkeIntegrity_Sha256,
		"sha384":    NetworkaccessTunnelConfigurationIKEv2CustomIkeIntegrity_Sha384,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := NetworkaccessTunnelConfigurationIKEv2CustomIkeIntegrity(input)
	return &out, nil
}
