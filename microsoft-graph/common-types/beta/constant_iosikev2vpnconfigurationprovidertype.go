package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IosikEv2VpnConfigurationProviderType string

const (
	IosikEv2VpnConfigurationProviderType_AppProxy      IosikEv2VpnConfigurationProviderType = "appProxy"
	IosikEv2VpnConfigurationProviderType_NotConfigured IosikEv2VpnConfigurationProviderType = "notConfigured"
	IosikEv2VpnConfigurationProviderType_PacketTunnel  IosikEv2VpnConfigurationProviderType = "packetTunnel"
)

func PossibleValuesForIosikEv2VpnConfigurationProviderType() []string {
	return []string{
		string(IosikEv2VpnConfigurationProviderType_AppProxy),
		string(IosikEv2VpnConfigurationProviderType_NotConfigured),
		string(IosikEv2VpnConfigurationProviderType_PacketTunnel),
	}
}

func (s *IosikEv2VpnConfigurationProviderType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseIosikEv2VpnConfigurationProviderType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseIosikEv2VpnConfigurationProviderType(input string) (*IosikEv2VpnConfigurationProviderType, error) {
	vals := map[string]IosikEv2VpnConfigurationProviderType{
		"appproxy":      IosikEv2VpnConfigurationProviderType_AppProxy,
		"notconfigured": IosikEv2VpnConfigurationProviderType_NotConfigured,
		"packettunnel":  IosikEv2VpnConfigurationProviderType_PacketTunnel,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := IosikEv2VpnConfigurationProviderType(input)
	return &out, nil
}
