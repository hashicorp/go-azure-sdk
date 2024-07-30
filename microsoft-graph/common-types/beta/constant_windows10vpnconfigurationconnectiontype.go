package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Windows10VpnConfigurationConnectionType string

const (
	Windows10VpnConfigurationConnectionType_Automatic                  Windows10VpnConfigurationConnectionType = "automatic"
	Windows10VpnConfigurationConnectionType_CheckPointCapsuleVpn       Windows10VpnConfigurationConnectionType = "checkPointCapsuleVpn"
	Windows10VpnConfigurationConnectionType_CiscoAnyConnect            Windows10VpnConfigurationConnectionType = "ciscoAnyConnect"
	Windows10VpnConfigurationConnectionType_Citrix                     Windows10VpnConfigurationConnectionType = "citrix"
	Windows10VpnConfigurationConnectionType_DellSonicWallMobileConnect Windows10VpnConfigurationConnectionType = "dellSonicWallMobileConnect"
	Windows10VpnConfigurationConnectionType_F5EdgeClient               Windows10VpnConfigurationConnectionType = "f5EdgeClient"
	Windows10VpnConfigurationConnectionType_IkEv2                      Windows10VpnConfigurationConnectionType = "ikEv2"
	Windows10VpnConfigurationConnectionType_L2tp                       Windows10VpnConfigurationConnectionType = "l2tp"
	Windows10VpnConfigurationConnectionType_MicrosoftTunnel            Windows10VpnConfigurationConnectionType = "microsoftTunnel"
	Windows10VpnConfigurationConnectionType_PaloAltoGlobalProtect      Windows10VpnConfigurationConnectionType = "paloAltoGlobalProtect"
	Windows10VpnConfigurationConnectionType_Pptp                       Windows10VpnConfigurationConnectionType = "pptp"
	Windows10VpnConfigurationConnectionType_PulseSecure                Windows10VpnConfigurationConnectionType = "pulseSecure"
)

func PossibleValuesForWindows10VpnConfigurationConnectionType() []string {
	return []string{
		string(Windows10VpnConfigurationConnectionType_Automatic),
		string(Windows10VpnConfigurationConnectionType_CheckPointCapsuleVpn),
		string(Windows10VpnConfigurationConnectionType_CiscoAnyConnect),
		string(Windows10VpnConfigurationConnectionType_Citrix),
		string(Windows10VpnConfigurationConnectionType_DellSonicWallMobileConnect),
		string(Windows10VpnConfigurationConnectionType_F5EdgeClient),
		string(Windows10VpnConfigurationConnectionType_IkEv2),
		string(Windows10VpnConfigurationConnectionType_L2tp),
		string(Windows10VpnConfigurationConnectionType_MicrosoftTunnel),
		string(Windows10VpnConfigurationConnectionType_PaloAltoGlobalProtect),
		string(Windows10VpnConfigurationConnectionType_Pptp),
		string(Windows10VpnConfigurationConnectionType_PulseSecure),
	}
}

func (s *Windows10VpnConfigurationConnectionType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseWindows10VpnConfigurationConnectionType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseWindows10VpnConfigurationConnectionType(input string) (*Windows10VpnConfigurationConnectionType, error) {
	vals := map[string]Windows10VpnConfigurationConnectionType{
		"automatic":                  Windows10VpnConfigurationConnectionType_Automatic,
		"checkpointcapsulevpn":       Windows10VpnConfigurationConnectionType_CheckPointCapsuleVpn,
		"ciscoanyconnect":            Windows10VpnConfigurationConnectionType_CiscoAnyConnect,
		"citrix":                     Windows10VpnConfigurationConnectionType_Citrix,
		"dellsonicwallmobileconnect": Windows10VpnConfigurationConnectionType_DellSonicWallMobileConnect,
		"f5edgeclient":               Windows10VpnConfigurationConnectionType_F5EdgeClient,
		"ikev2":                      Windows10VpnConfigurationConnectionType_IkEv2,
		"l2tp":                       Windows10VpnConfigurationConnectionType_L2tp,
		"microsofttunnel":            Windows10VpnConfigurationConnectionType_MicrosoftTunnel,
		"paloaltoglobalprotect":      Windows10VpnConfigurationConnectionType_PaloAltoGlobalProtect,
		"pptp":                       Windows10VpnConfigurationConnectionType_Pptp,
		"pulsesecure":                Windows10VpnConfigurationConnectionType_PulseSecure,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Windows10VpnConfigurationConnectionType(input)
	return &out, nil
}
