package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AppleVpnConfigurationConnectionType string

const (
	AppleVpnConfigurationConnectionTypealwaysOn                   AppleVpnConfigurationConnectionType = "AlwaysOn"
	AppleVpnConfigurationConnectionTypecheckPointCapsuleVpn       AppleVpnConfigurationConnectionType = "CheckPointCapsuleVpn"
	AppleVpnConfigurationConnectionTypeciscoAnyConnect            AppleVpnConfigurationConnectionType = "CiscoAnyConnect"
	AppleVpnConfigurationConnectionTypeciscoAnyConnectV2          AppleVpnConfigurationConnectionType = "CiscoAnyConnectV2"
	AppleVpnConfigurationConnectionTypeciscoIPSec                 AppleVpnConfigurationConnectionType = "CiscoIPSec"
	AppleVpnConfigurationConnectionTypecitrix                     AppleVpnConfigurationConnectionType = "Citrix"
	AppleVpnConfigurationConnectionTypecitrixSso                  AppleVpnConfigurationConnectionType = "CitrixSso"
	AppleVpnConfigurationConnectionTypecustomVpn                  AppleVpnConfigurationConnectionType = "CustomVpn"
	AppleVpnConfigurationConnectionTypedellSonicWallMobileConnect AppleVpnConfigurationConnectionType = "DellSonicWallMobileConnect"
	AppleVpnConfigurationConnectionTypef5Access2018               AppleVpnConfigurationConnectionType = "F5Access2018"
	AppleVpnConfigurationConnectionTypef5EdgeClient               AppleVpnConfigurationConnectionType = "F5EdgeClient"
	AppleVpnConfigurationConnectionTypeikEv2                      AppleVpnConfigurationConnectionType = "IkEv2"
	AppleVpnConfigurationConnectionTypemicrosoftProtect           AppleVpnConfigurationConnectionType = "MicrosoftProtect"
	AppleVpnConfigurationConnectionTypemicrosoftTunnel            AppleVpnConfigurationConnectionType = "MicrosoftTunnel"
	AppleVpnConfigurationConnectionTypenetMotionMobility          AppleVpnConfigurationConnectionType = "NetMotionMobility"
	AppleVpnConfigurationConnectionTypepaloAltoGlobalProtect      AppleVpnConfigurationConnectionType = "PaloAltoGlobalProtect"
	AppleVpnConfigurationConnectionTypepaloAltoGlobalProtectV2    AppleVpnConfigurationConnectionType = "PaloAltoGlobalProtectV2"
	AppleVpnConfigurationConnectionTypepulseSecure                AppleVpnConfigurationConnectionType = "PulseSecure"
	AppleVpnConfigurationConnectionTypezscalerPrivateAccess       AppleVpnConfigurationConnectionType = "ZscalerPrivateAccess"
)

func PossibleValuesForAppleVpnConfigurationConnectionType() []string {
	return []string{
		string(AppleVpnConfigurationConnectionTypealwaysOn),
		string(AppleVpnConfigurationConnectionTypecheckPointCapsuleVpn),
		string(AppleVpnConfigurationConnectionTypeciscoAnyConnect),
		string(AppleVpnConfigurationConnectionTypeciscoAnyConnectV2),
		string(AppleVpnConfigurationConnectionTypeciscoIPSec),
		string(AppleVpnConfigurationConnectionTypecitrix),
		string(AppleVpnConfigurationConnectionTypecitrixSso),
		string(AppleVpnConfigurationConnectionTypecustomVpn),
		string(AppleVpnConfigurationConnectionTypedellSonicWallMobileConnect),
		string(AppleVpnConfigurationConnectionTypef5Access2018),
		string(AppleVpnConfigurationConnectionTypef5EdgeClient),
		string(AppleVpnConfigurationConnectionTypeikEv2),
		string(AppleVpnConfigurationConnectionTypemicrosoftProtect),
		string(AppleVpnConfigurationConnectionTypemicrosoftTunnel),
		string(AppleVpnConfigurationConnectionTypenetMotionMobility),
		string(AppleVpnConfigurationConnectionTypepaloAltoGlobalProtect),
		string(AppleVpnConfigurationConnectionTypepaloAltoGlobalProtectV2),
		string(AppleVpnConfigurationConnectionTypepulseSecure),
		string(AppleVpnConfigurationConnectionTypezscalerPrivateAccess),
	}
}

func parseAppleVpnConfigurationConnectionType(input string) (*AppleVpnConfigurationConnectionType, error) {
	vals := map[string]AppleVpnConfigurationConnectionType{
		"alwayson":                   AppleVpnConfigurationConnectionTypealwaysOn,
		"checkpointcapsulevpn":       AppleVpnConfigurationConnectionTypecheckPointCapsuleVpn,
		"ciscoanyconnect":            AppleVpnConfigurationConnectionTypeciscoAnyConnect,
		"ciscoanyconnectv2":          AppleVpnConfigurationConnectionTypeciscoAnyConnectV2,
		"ciscoipsec":                 AppleVpnConfigurationConnectionTypeciscoIPSec,
		"citrix":                     AppleVpnConfigurationConnectionTypecitrix,
		"citrixsso":                  AppleVpnConfigurationConnectionTypecitrixSso,
		"customvpn":                  AppleVpnConfigurationConnectionTypecustomVpn,
		"dellsonicwallmobileconnect": AppleVpnConfigurationConnectionTypedellSonicWallMobileConnect,
		"f5access2018":               AppleVpnConfigurationConnectionTypef5Access2018,
		"f5edgeclient":               AppleVpnConfigurationConnectionTypef5EdgeClient,
		"ikev2":                      AppleVpnConfigurationConnectionTypeikEv2,
		"microsoftprotect":           AppleVpnConfigurationConnectionTypemicrosoftProtect,
		"microsofttunnel":            AppleVpnConfigurationConnectionTypemicrosoftTunnel,
		"netmotionmobility":          AppleVpnConfigurationConnectionTypenetMotionMobility,
		"paloaltoglobalprotect":      AppleVpnConfigurationConnectionTypepaloAltoGlobalProtect,
		"paloaltoglobalprotectv2":    AppleVpnConfigurationConnectionTypepaloAltoGlobalProtectV2,
		"pulsesecure":                AppleVpnConfigurationConnectionTypepulseSecure,
		"zscalerprivateaccess":       AppleVpnConfigurationConnectionTypezscalerPrivateAccess,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AppleVpnConfigurationConnectionType(input)
	return &out, nil
}
