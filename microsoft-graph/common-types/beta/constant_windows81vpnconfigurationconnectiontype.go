package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Windows81VpnConfigurationConnectionType string

const (
	Windows81VpnConfigurationConnectionType_CheckPointCapsuleVpn       Windows81VpnConfigurationConnectionType = "checkPointCapsuleVpn"
	Windows81VpnConfigurationConnectionType_DellSonicWallMobileConnect Windows81VpnConfigurationConnectionType = "dellSonicWallMobileConnect"
	Windows81VpnConfigurationConnectionType_F5EdgeClient               Windows81VpnConfigurationConnectionType = "f5EdgeClient"
	Windows81VpnConfigurationConnectionType_PulseSecure                Windows81VpnConfigurationConnectionType = "pulseSecure"
)

func PossibleValuesForWindows81VpnConfigurationConnectionType() []string {
	return []string{
		string(Windows81VpnConfigurationConnectionType_CheckPointCapsuleVpn),
		string(Windows81VpnConfigurationConnectionType_DellSonicWallMobileConnect),
		string(Windows81VpnConfigurationConnectionType_F5EdgeClient),
		string(Windows81VpnConfigurationConnectionType_PulseSecure),
	}
}

func (s *Windows81VpnConfigurationConnectionType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseWindows81VpnConfigurationConnectionType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseWindows81VpnConfigurationConnectionType(input string) (*Windows81VpnConfigurationConnectionType, error) {
	vals := map[string]Windows81VpnConfigurationConnectionType{
		"checkpointcapsulevpn":       Windows81VpnConfigurationConnectionType_CheckPointCapsuleVpn,
		"dellsonicwallmobileconnect": Windows81VpnConfigurationConnectionType_DellSonicWallMobileConnect,
		"f5edgeclient":               Windows81VpnConfigurationConnectionType_F5EdgeClient,
		"pulsesecure":                Windows81VpnConfigurationConnectionType_PulseSecure,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Windows81VpnConfigurationConnectionType(input)
	return &out, nil
}
