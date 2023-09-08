package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidWorkProfileVpnConfigurationConnectionType string

const (
	AndroidWorkProfileVpnConfigurationConnectionTypecheckPointCapsuleVpn       AndroidWorkProfileVpnConfigurationConnectionType = "CheckPointCapsuleVpn"
	AndroidWorkProfileVpnConfigurationConnectionTypeciscoAnyConnect            AndroidWorkProfileVpnConfigurationConnectionType = "CiscoAnyConnect"
	AndroidWorkProfileVpnConfigurationConnectionTypecitrix                     AndroidWorkProfileVpnConfigurationConnectionType = "Citrix"
	AndroidWorkProfileVpnConfigurationConnectionTypedellSonicWallMobileConnect AndroidWorkProfileVpnConfigurationConnectionType = "DellSonicWallMobileConnect"
	AndroidWorkProfileVpnConfigurationConnectionTypef5EdgeClient               AndroidWorkProfileVpnConfigurationConnectionType = "F5EdgeClient"
	AndroidWorkProfileVpnConfigurationConnectionTypemicrosoftProtect           AndroidWorkProfileVpnConfigurationConnectionType = "MicrosoftProtect"
	AndroidWorkProfileVpnConfigurationConnectionTypemicrosoftTunnel            AndroidWorkProfileVpnConfigurationConnectionType = "MicrosoftTunnel"
	AndroidWorkProfileVpnConfigurationConnectionTypenetMotionMobility          AndroidWorkProfileVpnConfigurationConnectionType = "NetMotionMobility"
	AndroidWorkProfileVpnConfigurationConnectionTypepaloAltoGlobalProtect      AndroidWorkProfileVpnConfigurationConnectionType = "PaloAltoGlobalProtect"
	AndroidWorkProfileVpnConfigurationConnectionTypepulseSecure                AndroidWorkProfileVpnConfigurationConnectionType = "PulseSecure"
)

func PossibleValuesForAndroidWorkProfileVpnConfigurationConnectionType() []string {
	return []string{
		string(AndroidWorkProfileVpnConfigurationConnectionTypecheckPointCapsuleVpn),
		string(AndroidWorkProfileVpnConfigurationConnectionTypeciscoAnyConnect),
		string(AndroidWorkProfileVpnConfigurationConnectionTypecitrix),
		string(AndroidWorkProfileVpnConfigurationConnectionTypedellSonicWallMobileConnect),
		string(AndroidWorkProfileVpnConfigurationConnectionTypef5EdgeClient),
		string(AndroidWorkProfileVpnConfigurationConnectionTypemicrosoftProtect),
		string(AndroidWorkProfileVpnConfigurationConnectionTypemicrosoftTunnel),
		string(AndroidWorkProfileVpnConfigurationConnectionTypenetMotionMobility),
		string(AndroidWorkProfileVpnConfigurationConnectionTypepaloAltoGlobalProtect),
		string(AndroidWorkProfileVpnConfigurationConnectionTypepulseSecure),
	}
}

func parseAndroidWorkProfileVpnConfigurationConnectionType(input string) (*AndroidWorkProfileVpnConfigurationConnectionType, error) {
	vals := map[string]AndroidWorkProfileVpnConfigurationConnectionType{
		"checkpointcapsulevpn":       AndroidWorkProfileVpnConfigurationConnectionTypecheckPointCapsuleVpn,
		"ciscoanyconnect":            AndroidWorkProfileVpnConfigurationConnectionTypeciscoAnyConnect,
		"citrix":                     AndroidWorkProfileVpnConfigurationConnectionTypecitrix,
		"dellsonicwallmobileconnect": AndroidWorkProfileVpnConfigurationConnectionTypedellSonicWallMobileConnect,
		"f5edgeclient":               AndroidWorkProfileVpnConfigurationConnectionTypef5EdgeClient,
		"microsoftprotect":           AndroidWorkProfileVpnConfigurationConnectionTypemicrosoftProtect,
		"microsofttunnel":            AndroidWorkProfileVpnConfigurationConnectionTypemicrosoftTunnel,
		"netmotionmobility":          AndroidWorkProfileVpnConfigurationConnectionTypenetMotionMobility,
		"paloaltoglobalprotect":      AndroidWorkProfileVpnConfigurationConnectionTypepaloAltoGlobalProtect,
		"pulsesecure":                AndroidWorkProfileVpnConfigurationConnectionTypepulseSecure,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AndroidWorkProfileVpnConfigurationConnectionType(input)
	return &out, nil
}
