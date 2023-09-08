package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MacOSVpnConfigurationConnectionType string

const (
	MacOSVpnConfigurationConnectionTypealwaysOn                   MacOSVpnConfigurationConnectionType = "AlwaysOn"
	MacOSVpnConfigurationConnectionTypecheckPointCapsuleVpn       MacOSVpnConfigurationConnectionType = "CheckPointCapsuleVpn"
	MacOSVpnConfigurationConnectionTypeciscoAnyConnect            MacOSVpnConfigurationConnectionType = "CiscoAnyConnect"
	MacOSVpnConfigurationConnectionTypeciscoAnyConnectV2          MacOSVpnConfigurationConnectionType = "CiscoAnyConnectV2"
	MacOSVpnConfigurationConnectionTypeciscoIPSec                 MacOSVpnConfigurationConnectionType = "CiscoIPSec"
	MacOSVpnConfigurationConnectionTypecitrix                     MacOSVpnConfigurationConnectionType = "Citrix"
	MacOSVpnConfigurationConnectionTypecitrixSso                  MacOSVpnConfigurationConnectionType = "CitrixSso"
	MacOSVpnConfigurationConnectionTypecustomVpn                  MacOSVpnConfigurationConnectionType = "CustomVpn"
	MacOSVpnConfigurationConnectionTypedellSonicWallMobileConnect MacOSVpnConfigurationConnectionType = "DellSonicWallMobileConnect"
	MacOSVpnConfigurationConnectionTypef5Access2018               MacOSVpnConfigurationConnectionType = "F5Access2018"
	MacOSVpnConfigurationConnectionTypef5EdgeClient               MacOSVpnConfigurationConnectionType = "F5EdgeClient"
	MacOSVpnConfigurationConnectionTypeikEv2                      MacOSVpnConfigurationConnectionType = "IkEv2"
	MacOSVpnConfigurationConnectionTypemicrosoftProtect           MacOSVpnConfigurationConnectionType = "MicrosoftProtect"
	MacOSVpnConfigurationConnectionTypemicrosoftTunnel            MacOSVpnConfigurationConnectionType = "MicrosoftTunnel"
	MacOSVpnConfigurationConnectionTypenetMotionMobility          MacOSVpnConfigurationConnectionType = "NetMotionMobility"
	MacOSVpnConfigurationConnectionTypepaloAltoGlobalProtect      MacOSVpnConfigurationConnectionType = "PaloAltoGlobalProtect"
	MacOSVpnConfigurationConnectionTypepaloAltoGlobalProtectV2    MacOSVpnConfigurationConnectionType = "PaloAltoGlobalProtectV2"
	MacOSVpnConfigurationConnectionTypepulseSecure                MacOSVpnConfigurationConnectionType = "PulseSecure"
	MacOSVpnConfigurationConnectionTypezscalerPrivateAccess       MacOSVpnConfigurationConnectionType = "ZscalerPrivateAccess"
)

func PossibleValuesForMacOSVpnConfigurationConnectionType() []string {
	return []string{
		string(MacOSVpnConfigurationConnectionTypealwaysOn),
		string(MacOSVpnConfigurationConnectionTypecheckPointCapsuleVpn),
		string(MacOSVpnConfigurationConnectionTypeciscoAnyConnect),
		string(MacOSVpnConfigurationConnectionTypeciscoAnyConnectV2),
		string(MacOSVpnConfigurationConnectionTypeciscoIPSec),
		string(MacOSVpnConfigurationConnectionTypecitrix),
		string(MacOSVpnConfigurationConnectionTypecitrixSso),
		string(MacOSVpnConfigurationConnectionTypecustomVpn),
		string(MacOSVpnConfigurationConnectionTypedellSonicWallMobileConnect),
		string(MacOSVpnConfigurationConnectionTypef5Access2018),
		string(MacOSVpnConfigurationConnectionTypef5EdgeClient),
		string(MacOSVpnConfigurationConnectionTypeikEv2),
		string(MacOSVpnConfigurationConnectionTypemicrosoftProtect),
		string(MacOSVpnConfigurationConnectionTypemicrosoftTunnel),
		string(MacOSVpnConfigurationConnectionTypenetMotionMobility),
		string(MacOSVpnConfigurationConnectionTypepaloAltoGlobalProtect),
		string(MacOSVpnConfigurationConnectionTypepaloAltoGlobalProtectV2),
		string(MacOSVpnConfigurationConnectionTypepulseSecure),
		string(MacOSVpnConfigurationConnectionTypezscalerPrivateAccess),
	}
}

func parseMacOSVpnConfigurationConnectionType(input string) (*MacOSVpnConfigurationConnectionType, error) {
	vals := map[string]MacOSVpnConfigurationConnectionType{
		"alwayson":                   MacOSVpnConfigurationConnectionTypealwaysOn,
		"checkpointcapsulevpn":       MacOSVpnConfigurationConnectionTypecheckPointCapsuleVpn,
		"ciscoanyconnect":            MacOSVpnConfigurationConnectionTypeciscoAnyConnect,
		"ciscoanyconnectv2":          MacOSVpnConfigurationConnectionTypeciscoAnyConnectV2,
		"ciscoipsec":                 MacOSVpnConfigurationConnectionTypeciscoIPSec,
		"citrix":                     MacOSVpnConfigurationConnectionTypecitrix,
		"citrixsso":                  MacOSVpnConfigurationConnectionTypecitrixSso,
		"customvpn":                  MacOSVpnConfigurationConnectionTypecustomVpn,
		"dellsonicwallmobileconnect": MacOSVpnConfigurationConnectionTypedellSonicWallMobileConnect,
		"f5access2018":               MacOSVpnConfigurationConnectionTypef5Access2018,
		"f5edgeclient":               MacOSVpnConfigurationConnectionTypef5EdgeClient,
		"ikev2":                      MacOSVpnConfigurationConnectionTypeikEv2,
		"microsoftprotect":           MacOSVpnConfigurationConnectionTypemicrosoftProtect,
		"microsofttunnel":            MacOSVpnConfigurationConnectionTypemicrosoftTunnel,
		"netmotionmobility":          MacOSVpnConfigurationConnectionTypenetMotionMobility,
		"paloaltoglobalprotect":      MacOSVpnConfigurationConnectionTypepaloAltoGlobalProtect,
		"paloaltoglobalprotectv2":    MacOSVpnConfigurationConnectionTypepaloAltoGlobalProtectV2,
		"pulsesecure":                MacOSVpnConfigurationConnectionTypepulseSecure,
		"zscalerprivateaccess":       MacOSVpnConfigurationConnectionTypezscalerPrivateAccess,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := MacOSVpnConfigurationConnectionType(input)
	return &out, nil
}
