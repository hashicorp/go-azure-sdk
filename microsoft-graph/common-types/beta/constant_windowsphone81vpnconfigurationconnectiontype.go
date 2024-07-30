package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WindowsPhone81VpnConfigurationConnectionType string

const (
	WindowsPhone81VpnConfigurationConnectionType_CheckPointCapsuleVpn       WindowsPhone81VpnConfigurationConnectionType = "checkPointCapsuleVpn"
	WindowsPhone81VpnConfigurationConnectionType_DellSonicWallMobileConnect WindowsPhone81VpnConfigurationConnectionType = "dellSonicWallMobileConnect"
	WindowsPhone81VpnConfigurationConnectionType_F5EdgeClient               WindowsPhone81VpnConfigurationConnectionType = "f5EdgeClient"
	WindowsPhone81VpnConfigurationConnectionType_PulseSecure                WindowsPhone81VpnConfigurationConnectionType = "pulseSecure"
)

func PossibleValuesForWindowsPhone81VpnConfigurationConnectionType() []string {
	return []string{
		string(WindowsPhone81VpnConfigurationConnectionType_CheckPointCapsuleVpn),
		string(WindowsPhone81VpnConfigurationConnectionType_DellSonicWallMobileConnect),
		string(WindowsPhone81VpnConfigurationConnectionType_F5EdgeClient),
		string(WindowsPhone81VpnConfigurationConnectionType_PulseSecure),
	}
}

func (s *WindowsPhone81VpnConfigurationConnectionType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseWindowsPhone81VpnConfigurationConnectionType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseWindowsPhone81VpnConfigurationConnectionType(input string) (*WindowsPhone81VpnConfigurationConnectionType, error) {
	vals := map[string]WindowsPhone81VpnConfigurationConnectionType{
		"checkpointcapsulevpn":       WindowsPhone81VpnConfigurationConnectionType_CheckPointCapsuleVpn,
		"dellsonicwallmobileconnect": WindowsPhone81VpnConfigurationConnectionType_DellSonicWallMobileConnect,
		"f5edgeclient":               WindowsPhone81VpnConfigurationConnectionType_F5EdgeClient,
		"pulsesecure":                WindowsPhone81VpnConfigurationConnectionType_PulseSecure,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := WindowsPhone81VpnConfigurationConnectionType(input)
	return &out, nil
}
