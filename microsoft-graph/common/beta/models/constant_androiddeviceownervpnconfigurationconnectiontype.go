package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidDeviceOwnerVpnConfigurationConnectionType string

const (
	AndroidDeviceOwnerVpnConfigurationConnectionTypecheckPointCapsuleVpn       AndroidDeviceOwnerVpnConfigurationConnectionType = "CheckPointCapsuleVpn"
	AndroidDeviceOwnerVpnConfigurationConnectionTypeciscoAnyConnect            AndroidDeviceOwnerVpnConfigurationConnectionType = "CiscoAnyConnect"
	AndroidDeviceOwnerVpnConfigurationConnectionTypecitrix                     AndroidDeviceOwnerVpnConfigurationConnectionType = "Citrix"
	AndroidDeviceOwnerVpnConfigurationConnectionTypedellSonicWallMobileConnect AndroidDeviceOwnerVpnConfigurationConnectionType = "DellSonicWallMobileConnect"
	AndroidDeviceOwnerVpnConfigurationConnectionTypef5EdgeClient               AndroidDeviceOwnerVpnConfigurationConnectionType = "F5EdgeClient"
	AndroidDeviceOwnerVpnConfigurationConnectionTypemicrosoftProtect           AndroidDeviceOwnerVpnConfigurationConnectionType = "MicrosoftProtect"
	AndroidDeviceOwnerVpnConfigurationConnectionTypemicrosoftTunnel            AndroidDeviceOwnerVpnConfigurationConnectionType = "MicrosoftTunnel"
	AndroidDeviceOwnerVpnConfigurationConnectionTypenetMotionMobility          AndroidDeviceOwnerVpnConfigurationConnectionType = "NetMotionMobility"
	AndroidDeviceOwnerVpnConfigurationConnectionTypepulseSecure                AndroidDeviceOwnerVpnConfigurationConnectionType = "PulseSecure"
)

func PossibleValuesForAndroidDeviceOwnerVpnConfigurationConnectionType() []string {
	return []string{
		string(AndroidDeviceOwnerVpnConfigurationConnectionTypecheckPointCapsuleVpn),
		string(AndroidDeviceOwnerVpnConfigurationConnectionTypeciscoAnyConnect),
		string(AndroidDeviceOwnerVpnConfigurationConnectionTypecitrix),
		string(AndroidDeviceOwnerVpnConfigurationConnectionTypedellSonicWallMobileConnect),
		string(AndroidDeviceOwnerVpnConfigurationConnectionTypef5EdgeClient),
		string(AndroidDeviceOwnerVpnConfigurationConnectionTypemicrosoftProtect),
		string(AndroidDeviceOwnerVpnConfigurationConnectionTypemicrosoftTunnel),
		string(AndroidDeviceOwnerVpnConfigurationConnectionTypenetMotionMobility),
		string(AndroidDeviceOwnerVpnConfigurationConnectionTypepulseSecure),
	}
}

func parseAndroidDeviceOwnerVpnConfigurationConnectionType(input string) (*AndroidDeviceOwnerVpnConfigurationConnectionType, error) {
	vals := map[string]AndroidDeviceOwnerVpnConfigurationConnectionType{
		"checkpointcapsulevpn":       AndroidDeviceOwnerVpnConfigurationConnectionTypecheckPointCapsuleVpn,
		"ciscoanyconnect":            AndroidDeviceOwnerVpnConfigurationConnectionTypeciscoAnyConnect,
		"citrix":                     AndroidDeviceOwnerVpnConfigurationConnectionTypecitrix,
		"dellsonicwallmobileconnect": AndroidDeviceOwnerVpnConfigurationConnectionTypedellSonicWallMobileConnect,
		"f5edgeclient":               AndroidDeviceOwnerVpnConfigurationConnectionTypef5EdgeClient,
		"microsoftprotect":           AndroidDeviceOwnerVpnConfigurationConnectionTypemicrosoftProtect,
		"microsofttunnel":            AndroidDeviceOwnerVpnConfigurationConnectionTypemicrosoftTunnel,
		"netmotionmobility":          AndroidDeviceOwnerVpnConfigurationConnectionTypenetMotionMobility,
		"pulsesecure":                AndroidDeviceOwnerVpnConfigurationConnectionTypepulseSecure,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AndroidDeviceOwnerVpnConfigurationConnectionType(input)
	return &out, nil
}
