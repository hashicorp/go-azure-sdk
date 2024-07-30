package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type NetworkaccessTunnelConfigurationIKEv2CustomDhGroup string

const (
	NetworkaccessTunnelConfigurationIKEv2CustomDhGroup_DhGroup14   NetworkaccessTunnelConfigurationIKEv2CustomDhGroup = "dhGroup14"
	NetworkaccessTunnelConfigurationIKEv2CustomDhGroup_DhGroup2048 NetworkaccessTunnelConfigurationIKEv2CustomDhGroup = "dhGroup2048"
	NetworkaccessTunnelConfigurationIKEv2CustomDhGroup_DhGroup24   NetworkaccessTunnelConfigurationIKEv2CustomDhGroup = "dhGroup24"
	NetworkaccessTunnelConfigurationIKEv2CustomDhGroup_Ecp256      NetworkaccessTunnelConfigurationIKEv2CustomDhGroup = "ecp256"
	NetworkaccessTunnelConfigurationIKEv2CustomDhGroup_Ecp384      NetworkaccessTunnelConfigurationIKEv2CustomDhGroup = "ecp384"
)

func PossibleValuesForNetworkaccessTunnelConfigurationIKEv2CustomDhGroup() []string {
	return []string{
		string(NetworkaccessTunnelConfigurationIKEv2CustomDhGroup_DhGroup14),
		string(NetworkaccessTunnelConfigurationIKEv2CustomDhGroup_DhGroup2048),
		string(NetworkaccessTunnelConfigurationIKEv2CustomDhGroup_DhGroup24),
		string(NetworkaccessTunnelConfigurationIKEv2CustomDhGroup_Ecp256),
		string(NetworkaccessTunnelConfigurationIKEv2CustomDhGroup_Ecp384),
	}
}

func (s *NetworkaccessTunnelConfigurationIKEv2CustomDhGroup) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseNetworkaccessTunnelConfigurationIKEv2CustomDhGroup(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseNetworkaccessTunnelConfigurationIKEv2CustomDhGroup(input string) (*NetworkaccessTunnelConfigurationIKEv2CustomDhGroup, error) {
	vals := map[string]NetworkaccessTunnelConfigurationIKEv2CustomDhGroup{
		"dhgroup14":   NetworkaccessTunnelConfigurationIKEv2CustomDhGroup_DhGroup14,
		"dhgroup2048": NetworkaccessTunnelConfigurationIKEv2CustomDhGroup_DhGroup2048,
		"dhgroup24":   NetworkaccessTunnelConfigurationIKEv2CustomDhGroup_DhGroup24,
		"ecp256":      NetworkaccessTunnelConfigurationIKEv2CustomDhGroup_Ecp256,
		"ecp384":      NetworkaccessTunnelConfigurationIKEv2CustomDhGroup_Ecp384,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := NetworkaccessTunnelConfigurationIKEv2CustomDhGroup(input)
	return &out, nil
}
