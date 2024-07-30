package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidDeviceOwnerVpnConfigurationConnectionType string

const (
	AndroidDeviceOwnerVpnConfigurationConnectionType_CheckPointCapsuleVpn       AndroidDeviceOwnerVpnConfigurationConnectionType = "checkPointCapsuleVpn"
	AndroidDeviceOwnerVpnConfigurationConnectionType_CiscoAnyConnect            AndroidDeviceOwnerVpnConfigurationConnectionType = "ciscoAnyConnect"
	AndroidDeviceOwnerVpnConfigurationConnectionType_Citrix                     AndroidDeviceOwnerVpnConfigurationConnectionType = "citrix"
	AndroidDeviceOwnerVpnConfigurationConnectionType_DellSonicWallMobileConnect AndroidDeviceOwnerVpnConfigurationConnectionType = "dellSonicWallMobileConnect"
	AndroidDeviceOwnerVpnConfigurationConnectionType_F5EdgeClient               AndroidDeviceOwnerVpnConfigurationConnectionType = "f5EdgeClient"
	AndroidDeviceOwnerVpnConfigurationConnectionType_MicrosoftProtect           AndroidDeviceOwnerVpnConfigurationConnectionType = "microsoftProtect"
	AndroidDeviceOwnerVpnConfigurationConnectionType_MicrosoftTunnel            AndroidDeviceOwnerVpnConfigurationConnectionType = "microsoftTunnel"
	AndroidDeviceOwnerVpnConfigurationConnectionType_NetMotionMobility          AndroidDeviceOwnerVpnConfigurationConnectionType = "netMotionMobility"
	AndroidDeviceOwnerVpnConfigurationConnectionType_PulseSecure                AndroidDeviceOwnerVpnConfigurationConnectionType = "pulseSecure"
)

func PossibleValuesForAndroidDeviceOwnerVpnConfigurationConnectionType() []string {
	return []string{
		string(AndroidDeviceOwnerVpnConfigurationConnectionType_CheckPointCapsuleVpn),
		string(AndroidDeviceOwnerVpnConfigurationConnectionType_CiscoAnyConnect),
		string(AndroidDeviceOwnerVpnConfigurationConnectionType_Citrix),
		string(AndroidDeviceOwnerVpnConfigurationConnectionType_DellSonicWallMobileConnect),
		string(AndroidDeviceOwnerVpnConfigurationConnectionType_F5EdgeClient),
		string(AndroidDeviceOwnerVpnConfigurationConnectionType_MicrosoftProtect),
		string(AndroidDeviceOwnerVpnConfigurationConnectionType_MicrosoftTunnel),
		string(AndroidDeviceOwnerVpnConfigurationConnectionType_NetMotionMobility),
		string(AndroidDeviceOwnerVpnConfigurationConnectionType_PulseSecure),
	}
}

func (s *AndroidDeviceOwnerVpnConfigurationConnectionType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAndroidDeviceOwnerVpnConfigurationConnectionType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAndroidDeviceOwnerVpnConfigurationConnectionType(input string) (*AndroidDeviceOwnerVpnConfigurationConnectionType, error) {
	vals := map[string]AndroidDeviceOwnerVpnConfigurationConnectionType{
		"checkpointcapsulevpn":       AndroidDeviceOwnerVpnConfigurationConnectionType_CheckPointCapsuleVpn,
		"ciscoanyconnect":            AndroidDeviceOwnerVpnConfigurationConnectionType_CiscoAnyConnect,
		"citrix":                     AndroidDeviceOwnerVpnConfigurationConnectionType_Citrix,
		"dellsonicwallmobileconnect": AndroidDeviceOwnerVpnConfigurationConnectionType_DellSonicWallMobileConnect,
		"f5edgeclient":               AndroidDeviceOwnerVpnConfigurationConnectionType_F5EdgeClient,
		"microsoftprotect":           AndroidDeviceOwnerVpnConfigurationConnectionType_MicrosoftProtect,
		"microsofttunnel":            AndroidDeviceOwnerVpnConfigurationConnectionType_MicrosoftTunnel,
		"netmotionmobility":          AndroidDeviceOwnerVpnConfigurationConnectionType_NetMotionMobility,
		"pulsesecure":                AndroidDeviceOwnerVpnConfigurationConnectionType_PulseSecure,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AndroidDeviceOwnerVpnConfigurationConnectionType(input)
	return &out, nil
}
