package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidVpnConfigurationConnectionType string

const (
	AndroidVpnConfigurationConnectionTypecheckPointCapsuleVpn       AndroidVpnConfigurationConnectionType = "CheckPointCapsuleVpn"
	AndroidVpnConfigurationConnectionTypeciscoAnyConnect            AndroidVpnConfigurationConnectionType = "CiscoAnyConnect"
	AndroidVpnConfigurationConnectionTypecitrix                     AndroidVpnConfigurationConnectionType = "Citrix"
	AndroidVpnConfigurationConnectionTypedellSonicWallMobileConnect AndroidVpnConfigurationConnectionType = "DellSonicWallMobileConnect"
	AndroidVpnConfigurationConnectionTypef5EdgeClient               AndroidVpnConfigurationConnectionType = "F5EdgeClient"
	AndroidVpnConfigurationConnectionTypemicrosoftProtect           AndroidVpnConfigurationConnectionType = "MicrosoftProtect"
	AndroidVpnConfigurationConnectionTypemicrosoftTunnel            AndroidVpnConfigurationConnectionType = "MicrosoftTunnel"
	AndroidVpnConfigurationConnectionTypenetMotionMobility          AndroidVpnConfigurationConnectionType = "NetMotionMobility"
	AndroidVpnConfigurationConnectionTypepulseSecure                AndroidVpnConfigurationConnectionType = "PulseSecure"
)

func PossibleValuesForAndroidVpnConfigurationConnectionType() []string {
	return []string{
		string(AndroidVpnConfigurationConnectionTypecheckPointCapsuleVpn),
		string(AndroidVpnConfigurationConnectionTypeciscoAnyConnect),
		string(AndroidVpnConfigurationConnectionTypecitrix),
		string(AndroidVpnConfigurationConnectionTypedellSonicWallMobileConnect),
		string(AndroidVpnConfigurationConnectionTypef5EdgeClient),
		string(AndroidVpnConfigurationConnectionTypemicrosoftProtect),
		string(AndroidVpnConfigurationConnectionTypemicrosoftTunnel),
		string(AndroidVpnConfigurationConnectionTypenetMotionMobility),
		string(AndroidVpnConfigurationConnectionTypepulseSecure),
	}
}

func parseAndroidVpnConfigurationConnectionType(input string) (*AndroidVpnConfigurationConnectionType, error) {
	vals := map[string]AndroidVpnConfigurationConnectionType{
		"checkpointcapsulevpn":       AndroidVpnConfigurationConnectionTypecheckPointCapsuleVpn,
		"ciscoanyconnect":            AndroidVpnConfigurationConnectionTypeciscoAnyConnect,
		"citrix":                     AndroidVpnConfigurationConnectionTypecitrix,
		"dellsonicwallmobileconnect": AndroidVpnConfigurationConnectionTypedellSonicWallMobileConnect,
		"f5edgeclient":               AndroidVpnConfigurationConnectionTypef5EdgeClient,
		"microsoftprotect":           AndroidVpnConfigurationConnectionTypemicrosoftProtect,
		"microsofttunnel":            AndroidVpnConfigurationConnectionTypemicrosoftTunnel,
		"netmotionmobility":          AndroidVpnConfigurationConnectionTypenetMotionMobility,
		"pulsesecure":                AndroidVpnConfigurationConnectionTypepulseSecure,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AndroidVpnConfigurationConnectionType(input)
	return &out, nil
}
