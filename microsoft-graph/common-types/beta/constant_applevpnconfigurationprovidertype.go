package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AppleVpnConfigurationProviderType string

const (
	AppleVpnConfigurationProviderType_AppProxy      AppleVpnConfigurationProviderType = "appProxy"
	AppleVpnConfigurationProviderType_NotConfigured AppleVpnConfigurationProviderType = "notConfigured"
	AppleVpnConfigurationProviderType_PacketTunnel  AppleVpnConfigurationProviderType = "packetTunnel"
)

func PossibleValuesForAppleVpnConfigurationProviderType() []string {
	return []string{
		string(AppleVpnConfigurationProviderType_AppProxy),
		string(AppleVpnConfigurationProviderType_NotConfigured),
		string(AppleVpnConfigurationProviderType_PacketTunnel),
	}
}

func (s *AppleVpnConfigurationProviderType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAppleVpnConfigurationProviderType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAppleVpnConfigurationProviderType(input string) (*AppleVpnConfigurationProviderType, error) {
	vals := map[string]AppleVpnConfigurationProviderType{
		"appproxy":      AppleVpnConfigurationProviderType_AppProxy,
		"notconfigured": AppleVpnConfigurationProviderType_NotConfigured,
		"packettunnel":  AppleVpnConfigurationProviderType_PacketTunnel,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AppleVpnConfigurationProviderType(input)
	return &out, nil
}
