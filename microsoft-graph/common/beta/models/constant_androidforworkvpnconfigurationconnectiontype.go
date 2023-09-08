package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidForWorkVpnConfigurationConnectionType string

const (
	AndroidForWorkVpnConfigurationConnectionTypecheckPointCapsuleVpn       AndroidForWorkVpnConfigurationConnectionType = "CheckPointCapsuleVpn"
	AndroidForWorkVpnConfigurationConnectionTypeciscoAnyConnect            AndroidForWorkVpnConfigurationConnectionType = "CiscoAnyConnect"
	AndroidForWorkVpnConfigurationConnectionTypecitrix                     AndroidForWorkVpnConfigurationConnectionType = "Citrix"
	AndroidForWorkVpnConfigurationConnectionTypedellSonicWallMobileConnect AndroidForWorkVpnConfigurationConnectionType = "DellSonicWallMobileConnect"
	AndroidForWorkVpnConfigurationConnectionTypef5EdgeClient               AndroidForWorkVpnConfigurationConnectionType = "F5EdgeClient"
	AndroidForWorkVpnConfigurationConnectionTypepulseSecure                AndroidForWorkVpnConfigurationConnectionType = "PulseSecure"
)

func PossibleValuesForAndroidForWorkVpnConfigurationConnectionType() []string {
	return []string{
		string(AndroidForWorkVpnConfigurationConnectionTypecheckPointCapsuleVpn),
		string(AndroidForWorkVpnConfigurationConnectionTypeciscoAnyConnect),
		string(AndroidForWorkVpnConfigurationConnectionTypecitrix),
		string(AndroidForWorkVpnConfigurationConnectionTypedellSonicWallMobileConnect),
		string(AndroidForWorkVpnConfigurationConnectionTypef5EdgeClient),
		string(AndroidForWorkVpnConfigurationConnectionTypepulseSecure),
	}
}

func parseAndroidForWorkVpnConfigurationConnectionType(input string) (*AndroidForWorkVpnConfigurationConnectionType, error) {
	vals := map[string]AndroidForWorkVpnConfigurationConnectionType{
		"checkpointcapsulevpn":       AndroidForWorkVpnConfigurationConnectionTypecheckPointCapsuleVpn,
		"ciscoanyconnect":            AndroidForWorkVpnConfigurationConnectionTypeciscoAnyConnect,
		"citrix":                     AndroidForWorkVpnConfigurationConnectionTypecitrix,
		"dellsonicwallmobileconnect": AndroidForWorkVpnConfigurationConnectionTypedellSonicWallMobileConnect,
		"f5edgeclient":               AndroidForWorkVpnConfigurationConnectionTypef5EdgeClient,
		"pulsesecure":                AndroidForWorkVpnConfigurationConnectionTypepulseSecure,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AndroidForWorkVpnConfigurationConnectionType(input)
	return &out, nil
}
