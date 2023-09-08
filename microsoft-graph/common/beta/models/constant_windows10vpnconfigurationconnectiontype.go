package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Windows10VpnConfigurationConnectionType string

const (
	Windows10VpnConfigurationConnectionTypeautomatic                  Windows10VpnConfigurationConnectionType = "Automatic"
	Windows10VpnConfigurationConnectionTypecheckPointCapsuleVpn       Windows10VpnConfigurationConnectionType = "CheckPointCapsuleVpn"
	Windows10VpnConfigurationConnectionTypeciscoAnyConnect            Windows10VpnConfigurationConnectionType = "CiscoAnyConnect"
	Windows10VpnConfigurationConnectionTypecitrix                     Windows10VpnConfigurationConnectionType = "Citrix"
	Windows10VpnConfigurationConnectionTypedellSonicWallMobileConnect Windows10VpnConfigurationConnectionType = "DellSonicWallMobileConnect"
	Windows10VpnConfigurationConnectionTypef5EdgeClient               Windows10VpnConfigurationConnectionType = "F5EdgeClient"
	Windows10VpnConfigurationConnectionTypeikEv2                      Windows10VpnConfigurationConnectionType = "IkEv2"
	Windows10VpnConfigurationConnectionTypel2tp                       Windows10VpnConfigurationConnectionType = "L2tp"
	Windows10VpnConfigurationConnectionTypemicrosoftTunnel            Windows10VpnConfigurationConnectionType = "MicrosoftTunnel"
	Windows10VpnConfigurationConnectionTypepaloAltoGlobalProtect      Windows10VpnConfigurationConnectionType = "PaloAltoGlobalProtect"
	Windows10VpnConfigurationConnectionTypepptp                       Windows10VpnConfigurationConnectionType = "Pptp"
	Windows10VpnConfigurationConnectionTypepulseSecure                Windows10VpnConfigurationConnectionType = "PulseSecure"
)

func PossibleValuesForWindows10VpnConfigurationConnectionType() []string {
	return []string{
		string(Windows10VpnConfigurationConnectionTypeautomatic),
		string(Windows10VpnConfigurationConnectionTypecheckPointCapsuleVpn),
		string(Windows10VpnConfigurationConnectionTypeciscoAnyConnect),
		string(Windows10VpnConfigurationConnectionTypecitrix),
		string(Windows10VpnConfigurationConnectionTypedellSonicWallMobileConnect),
		string(Windows10VpnConfigurationConnectionTypef5EdgeClient),
		string(Windows10VpnConfigurationConnectionTypeikEv2),
		string(Windows10VpnConfigurationConnectionTypel2tp),
		string(Windows10VpnConfigurationConnectionTypemicrosoftTunnel),
		string(Windows10VpnConfigurationConnectionTypepaloAltoGlobalProtect),
		string(Windows10VpnConfigurationConnectionTypepptp),
		string(Windows10VpnConfigurationConnectionTypepulseSecure),
	}
}

func parseWindows10VpnConfigurationConnectionType(input string) (*Windows10VpnConfigurationConnectionType, error) {
	vals := map[string]Windows10VpnConfigurationConnectionType{
		"automatic":                  Windows10VpnConfigurationConnectionTypeautomatic,
		"checkpointcapsulevpn":       Windows10VpnConfigurationConnectionTypecheckPointCapsuleVpn,
		"ciscoanyconnect":            Windows10VpnConfigurationConnectionTypeciscoAnyConnect,
		"citrix":                     Windows10VpnConfigurationConnectionTypecitrix,
		"dellsonicwallmobileconnect": Windows10VpnConfigurationConnectionTypedellSonicWallMobileConnect,
		"f5edgeclient":               Windows10VpnConfigurationConnectionTypef5EdgeClient,
		"ikev2":                      Windows10VpnConfigurationConnectionTypeikEv2,
		"l2tp":                       Windows10VpnConfigurationConnectionTypel2tp,
		"microsofttunnel":            Windows10VpnConfigurationConnectionTypemicrosoftTunnel,
		"paloaltoglobalprotect":      Windows10VpnConfigurationConnectionTypepaloAltoGlobalProtect,
		"pptp":                       Windows10VpnConfigurationConnectionTypepptp,
		"pulsesecure":                Windows10VpnConfigurationConnectionTypepulseSecure,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Windows10VpnConfigurationConnectionType(input)
	return &out, nil
}
