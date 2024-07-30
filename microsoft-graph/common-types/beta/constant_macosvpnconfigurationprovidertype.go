package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MacOSVpnConfigurationProviderType string

const (
	MacOSVpnConfigurationProviderType_AppProxy      MacOSVpnConfigurationProviderType = "appProxy"
	MacOSVpnConfigurationProviderType_NotConfigured MacOSVpnConfigurationProviderType = "notConfigured"
	MacOSVpnConfigurationProviderType_PacketTunnel  MacOSVpnConfigurationProviderType = "packetTunnel"
)

func PossibleValuesForMacOSVpnConfigurationProviderType() []string {
	return []string{
		string(MacOSVpnConfigurationProviderType_AppProxy),
		string(MacOSVpnConfigurationProviderType_NotConfigured),
		string(MacOSVpnConfigurationProviderType_PacketTunnel),
	}
}

func (s *MacOSVpnConfigurationProviderType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseMacOSVpnConfigurationProviderType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseMacOSVpnConfigurationProviderType(input string) (*MacOSVpnConfigurationProviderType, error) {
	vals := map[string]MacOSVpnConfigurationProviderType{
		"appproxy":      MacOSVpnConfigurationProviderType_AppProxy,
		"notconfigured": MacOSVpnConfigurationProviderType_NotConfigured,
		"packettunnel":  MacOSVpnConfigurationProviderType_PacketTunnel,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := MacOSVpnConfigurationProviderType(input)
	return &out, nil
}
