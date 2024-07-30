package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidVpnConfigurationConnectionType string

const (
	AndroidVpnConfigurationConnectionType_CheckPointCapsuleVpn       AndroidVpnConfigurationConnectionType = "checkPointCapsuleVpn"
	AndroidVpnConfigurationConnectionType_CiscoAnyConnect            AndroidVpnConfigurationConnectionType = "ciscoAnyConnect"
	AndroidVpnConfigurationConnectionType_Citrix                     AndroidVpnConfigurationConnectionType = "citrix"
	AndroidVpnConfigurationConnectionType_DellSonicWallMobileConnect AndroidVpnConfigurationConnectionType = "dellSonicWallMobileConnect"
	AndroidVpnConfigurationConnectionType_F5EdgeClient               AndroidVpnConfigurationConnectionType = "f5EdgeClient"
	AndroidVpnConfigurationConnectionType_MicrosoftProtect           AndroidVpnConfigurationConnectionType = "microsoftProtect"
	AndroidVpnConfigurationConnectionType_MicrosoftTunnel            AndroidVpnConfigurationConnectionType = "microsoftTunnel"
	AndroidVpnConfigurationConnectionType_NetMotionMobility          AndroidVpnConfigurationConnectionType = "netMotionMobility"
	AndroidVpnConfigurationConnectionType_PulseSecure                AndroidVpnConfigurationConnectionType = "pulseSecure"
)

func PossibleValuesForAndroidVpnConfigurationConnectionType() []string {
	return []string{
		string(AndroidVpnConfigurationConnectionType_CheckPointCapsuleVpn),
		string(AndroidVpnConfigurationConnectionType_CiscoAnyConnect),
		string(AndroidVpnConfigurationConnectionType_Citrix),
		string(AndroidVpnConfigurationConnectionType_DellSonicWallMobileConnect),
		string(AndroidVpnConfigurationConnectionType_F5EdgeClient),
		string(AndroidVpnConfigurationConnectionType_MicrosoftProtect),
		string(AndroidVpnConfigurationConnectionType_MicrosoftTunnel),
		string(AndroidVpnConfigurationConnectionType_NetMotionMobility),
		string(AndroidVpnConfigurationConnectionType_PulseSecure),
	}
}

func (s *AndroidVpnConfigurationConnectionType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAndroidVpnConfigurationConnectionType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAndroidVpnConfigurationConnectionType(input string) (*AndroidVpnConfigurationConnectionType, error) {
	vals := map[string]AndroidVpnConfigurationConnectionType{
		"checkpointcapsulevpn":       AndroidVpnConfigurationConnectionType_CheckPointCapsuleVpn,
		"ciscoanyconnect":            AndroidVpnConfigurationConnectionType_CiscoAnyConnect,
		"citrix":                     AndroidVpnConfigurationConnectionType_Citrix,
		"dellsonicwallmobileconnect": AndroidVpnConfigurationConnectionType_DellSonicWallMobileConnect,
		"f5edgeclient":               AndroidVpnConfigurationConnectionType_F5EdgeClient,
		"microsoftprotect":           AndroidVpnConfigurationConnectionType_MicrosoftProtect,
		"microsofttunnel":            AndroidVpnConfigurationConnectionType_MicrosoftTunnel,
		"netmotionmobility":          AndroidVpnConfigurationConnectionType_NetMotionMobility,
		"pulsesecure":                AndroidVpnConfigurationConnectionType_PulseSecure,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AndroidVpnConfigurationConnectionType(input)
	return &out, nil
}
