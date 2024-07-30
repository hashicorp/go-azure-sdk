package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidWorkProfileVpnConfigurationConnectionType string

const (
	AndroidWorkProfileVpnConfigurationConnectionType_CheckPointCapsuleVpn       AndroidWorkProfileVpnConfigurationConnectionType = "checkPointCapsuleVpn"
	AndroidWorkProfileVpnConfigurationConnectionType_CiscoAnyConnect            AndroidWorkProfileVpnConfigurationConnectionType = "ciscoAnyConnect"
	AndroidWorkProfileVpnConfigurationConnectionType_Citrix                     AndroidWorkProfileVpnConfigurationConnectionType = "citrix"
	AndroidWorkProfileVpnConfigurationConnectionType_DellSonicWallMobileConnect AndroidWorkProfileVpnConfigurationConnectionType = "dellSonicWallMobileConnect"
	AndroidWorkProfileVpnConfigurationConnectionType_F5EdgeClient               AndroidWorkProfileVpnConfigurationConnectionType = "f5EdgeClient"
	AndroidWorkProfileVpnConfigurationConnectionType_MicrosoftProtect           AndroidWorkProfileVpnConfigurationConnectionType = "microsoftProtect"
	AndroidWorkProfileVpnConfigurationConnectionType_MicrosoftTunnel            AndroidWorkProfileVpnConfigurationConnectionType = "microsoftTunnel"
	AndroidWorkProfileVpnConfigurationConnectionType_NetMotionMobility          AndroidWorkProfileVpnConfigurationConnectionType = "netMotionMobility"
	AndroidWorkProfileVpnConfigurationConnectionType_PaloAltoGlobalProtect      AndroidWorkProfileVpnConfigurationConnectionType = "paloAltoGlobalProtect"
	AndroidWorkProfileVpnConfigurationConnectionType_PulseSecure                AndroidWorkProfileVpnConfigurationConnectionType = "pulseSecure"
)

func PossibleValuesForAndroidWorkProfileVpnConfigurationConnectionType() []string {
	return []string{
		string(AndroidWorkProfileVpnConfigurationConnectionType_CheckPointCapsuleVpn),
		string(AndroidWorkProfileVpnConfigurationConnectionType_CiscoAnyConnect),
		string(AndroidWorkProfileVpnConfigurationConnectionType_Citrix),
		string(AndroidWorkProfileVpnConfigurationConnectionType_DellSonicWallMobileConnect),
		string(AndroidWorkProfileVpnConfigurationConnectionType_F5EdgeClient),
		string(AndroidWorkProfileVpnConfigurationConnectionType_MicrosoftProtect),
		string(AndroidWorkProfileVpnConfigurationConnectionType_MicrosoftTunnel),
		string(AndroidWorkProfileVpnConfigurationConnectionType_NetMotionMobility),
		string(AndroidWorkProfileVpnConfigurationConnectionType_PaloAltoGlobalProtect),
		string(AndroidWorkProfileVpnConfigurationConnectionType_PulseSecure),
	}
}

func (s *AndroidWorkProfileVpnConfigurationConnectionType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAndroidWorkProfileVpnConfigurationConnectionType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAndroidWorkProfileVpnConfigurationConnectionType(input string) (*AndroidWorkProfileVpnConfigurationConnectionType, error) {
	vals := map[string]AndroidWorkProfileVpnConfigurationConnectionType{
		"checkpointcapsulevpn":       AndroidWorkProfileVpnConfigurationConnectionType_CheckPointCapsuleVpn,
		"ciscoanyconnect":            AndroidWorkProfileVpnConfigurationConnectionType_CiscoAnyConnect,
		"citrix":                     AndroidWorkProfileVpnConfigurationConnectionType_Citrix,
		"dellsonicwallmobileconnect": AndroidWorkProfileVpnConfigurationConnectionType_DellSonicWallMobileConnect,
		"f5edgeclient":               AndroidWorkProfileVpnConfigurationConnectionType_F5EdgeClient,
		"microsoftprotect":           AndroidWorkProfileVpnConfigurationConnectionType_MicrosoftProtect,
		"microsofttunnel":            AndroidWorkProfileVpnConfigurationConnectionType_MicrosoftTunnel,
		"netmotionmobility":          AndroidWorkProfileVpnConfigurationConnectionType_NetMotionMobility,
		"paloaltoglobalprotect":      AndroidWorkProfileVpnConfigurationConnectionType_PaloAltoGlobalProtect,
		"pulsesecure":                AndroidWorkProfileVpnConfigurationConnectionType_PulseSecure,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AndroidWorkProfileVpnConfigurationConnectionType(input)
	return &out, nil
}
