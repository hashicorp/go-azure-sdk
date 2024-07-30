package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IosVpnConfigurationProviderType string

const (
	IosVpnConfigurationProviderType_AppProxy      IosVpnConfigurationProviderType = "appProxy"
	IosVpnConfigurationProviderType_NotConfigured IosVpnConfigurationProviderType = "notConfigured"
	IosVpnConfigurationProviderType_PacketTunnel  IosVpnConfigurationProviderType = "packetTunnel"
)

func PossibleValuesForIosVpnConfigurationProviderType() []string {
	return []string{
		string(IosVpnConfigurationProviderType_AppProxy),
		string(IosVpnConfigurationProviderType_NotConfigured),
		string(IosVpnConfigurationProviderType_PacketTunnel),
	}
}

func (s *IosVpnConfigurationProviderType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseIosVpnConfigurationProviderType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseIosVpnConfigurationProviderType(input string) (*IosVpnConfigurationProviderType, error) {
	vals := map[string]IosVpnConfigurationProviderType{
		"appproxy":      IosVpnConfigurationProviderType_AppProxy,
		"notconfigured": IosVpnConfigurationProviderType_NotConfigured,
		"packettunnel":  IosVpnConfigurationProviderType_PacketTunnel,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := IosVpnConfigurationProviderType(input)
	return &out, nil
}
