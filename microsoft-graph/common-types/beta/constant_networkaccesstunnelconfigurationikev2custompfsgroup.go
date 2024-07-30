package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type NetworkaccessTunnelConfigurationIKEv2CustomPfsGroup string

const (
	NetworkaccessTunnelConfigurationIKEv2CustomPfsGroup_Ecp256  NetworkaccessTunnelConfigurationIKEv2CustomPfsGroup = "ecp256"
	NetworkaccessTunnelConfigurationIKEv2CustomPfsGroup_Ecp384  NetworkaccessTunnelConfigurationIKEv2CustomPfsGroup = "ecp384"
	NetworkaccessTunnelConfigurationIKEv2CustomPfsGroup_None    NetworkaccessTunnelConfigurationIKEv2CustomPfsGroup = "none"
	NetworkaccessTunnelConfigurationIKEv2CustomPfsGroup_Pfs1    NetworkaccessTunnelConfigurationIKEv2CustomPfsGroup = "pfs1"
	NetworkaccessTunnelConfigurationIKEv2CustomPfsGroup_Pfs14   NetworkaccessTunnelConfigurationIKEv2CustomPfsGroup = "pfs14"
	NetworkaccessTunnelConfigurationIKEv2CustomPfsGroup_Pfs2    NetworkaccessTunnelConfigurationIKEv2CustomPfsGroup = "pfs2"
	NetworkaccessTunnelConfigurationIKEv2CustomPfsGroup_Pfs2048 NetworkaccessTunnelConfigurationIKEv2CustomPfsGroup = "pfs2048"
	NetworkaccessTunnelConfigurationIKEv2CustomPfsGroup_Pfs24   NetworkaccessTunnelConfigurationIKEv2CustomPfsGroup = "pfs24"
	NetworkaccessTunnelConfigurationIKEv2CustomPfsGroup_Pfsmm   NetworkaccessTunnelConfigurationIKEv2CustomPfsGroup = "pfsmm"
)

func PossibleValuesForNetworkaccessTunnelConfigurationIKEv2CustomPfsGroup() []string {
	return []string{
		string(NetworkaccessTunnelConfigurationIKEv2CustomPfsGroup_Ecp256),
		string(NetworkaccessTunnelConfigurationIKEv2CustomPfsGroup_Ecp384),
		string(NetworkaccessTunnelConfigurationIKEv2CustomPfsGroup_None),
		string(NetworkaccessTunnelConfigurationIKEv2CustomPfsGroup_Pfs1),
		string(NetworkaccessTunnelConfigurationIKEv2CustomPfsGroup_Pfs14),
		string(NetworkaccessTunnelConfigurationIKEv2CustomPfsGroup_Pfs2),
		string(NetworkaccessTunnelConfigurationIKEv2CustomPfsGroup_Pfs2048),
		string(NetworkaccessTunnelConfigurationIKEv2CustomPfsGroup_Pfs24),
		string(NetworkaccessTunnelConfigurationIKEv2CustomPfsGroup_Pfsmm),
	}
}

func (s *NetworkaccessTunnelConfigurationIKEv2CustomPfsGroup) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseNetworkaccessTunnelConfigurationIKEv2CustomPfsGroup(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseNetworkaccessTunnelConfigurationIKEv2CustomPfsGroup(input string) (*NetworkaccessTunnelConfigurationIKEv2CustomPfsGroup, error) {
	vals := map[string]NetworkaccessTunnelConfigurationIKEv2CustomPfsGroup{
		"ecp256":  NetworkaccessTunnelConfigurationIKEv2CustomPfsGroup_Ecp256,
		"ecp384":  NetworkaccessTunnelConfigurationIKEv2CustomPfsGroup_Ecp384,
		"none":    NetworkaccessTunnelConfigurationIKEv2CustomPfsGroup_None,
		"pfs1":    NetworkaccessTunnelConfigurationIKEv2CustomPfsGroup_Pfs1,
		"pfs14":   NetworkaccessTunnelConfigurationIKEv2CustomPfsGroup_Pfs14,
		"pfs2":    NetworkaccessTunnelConfigurationIKEv2CustomPfsGroup_Pfs2,
		"pfs2048": NetworkaccessTunnelConfigurationIKEv2CustomPfsGroup_Pfs2048,
		"pfs24":   NetworkaccessTunnelConfigurationIKEv2CustomPfsGroup_Pfs24,
		"pfsmm":   NetworkaccessTunnelConfigurationIKEv2CustomPfsGroup_Pfsmm,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := NetworkaccessTunnelConfigurationIKEv2CustomPfsGroup(input)
	return &out, nil
}
