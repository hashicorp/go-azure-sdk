package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IosVpnConfigurationConnectionType string

const (
	IosVpnConfigurationConnectionTypealwaysOn                   IosVpnConfigurationConnectionType = "AlwaysOn"
	IosVpnConfigurationConnectionTypecheckPointCapsuleVpn       IosVpnConfigurationConnectionType = "CheckPointCapsuleVpn"
	IosVpnConfigurationConnectionTypeciscoAnyConnect            IosVpnConfigurationConnectionType = "CiscoAnyConnect"
	IosVpnConfigurationConnectionTypeciscoAnyConnectV2          IosVpnConfigurationConnectionType = "CiscoAnyConnectV2"
	IosVpnConfigurationConnectionTypeciscoIPSec                 IosVpnConfigurationConnectionType = "CiscoIPSec"
	IosVpnConfigurationConnectionTypecitrix                     IosVpnConfigurationConnectionType = "Citrix"
	IosVpnConfigurationConnectionTypecitrixSso                  IosVpnConfigurationConnectionType = "CitrixSso"
	IosVpnConfigurationConnectionTypecustomVpn                  IosVpnConfigurationConnectionType = "CustomVpn"
	IosVpnConfigurationConnectionTypedellSonicWallMobileConnect IosVpnConfigurationConnectionType = "DellSonicWallMobileConnect"
	IosVpnConfigurationConnectionTypef5Access2018               IosVpnConfigurationConnectionType = "F5Access2018"
	IosVpnConfigurationConnectionTypef5EdgeClient               IosVpnConfigurationConnectionType = "F5EdgeClient"
	IosVpnConfigurationConnectionTypeikEv2                      IosVpnConfigurationConnectionType = "IkEv2"
	IosVpnConfigurationConnectionTypemicrosoftProtect           IosVpnConfigurationConnectionType = "MicrosoftProtect"
	IosVpnConfigurationConnectionTypemicrosoftTunnel            IosVpnConfigurationConnectionType = "MicrosoftTunnel"
	IosVpnConfigurationConnectionTypenetMotionMobility          IosVpnConfigurationConnectionType = "NetMotionMobility"
	IosVpnConfigurationConnectionTypepaloAltoGlobalProtect      IosVpnConfigurationConnectionType = "PaloAltoGlobalProtect"
	IosVpnConfigurationConnectionTypepaloAltoGlobalProtectV2    IosVpnConfigurationConnectionType = "PaloAltoGlobalProtectV2"
	IosVpnConfigurationConnectionTypepulseSecure                IosVpnConfigurationConnectionType = "PulseSecure"
	IosVpnConfigurationConnectionTypezscalerPrivateAccess       IosVpnConfigurationConnectionType = "ZscalerPrivateAccess"
)

func PossibleValuesForIosVpnConfigurationConnectionType() []string {
	return []string{
		string(IosVpnConfigurationConnectionTypealwaysOn),
		string(IosVpnConfigurationConnectionTypecheckPointCapsuleVpn),
		string(IosVpnConfigurationConnectionTypeciscoAnyConnect),
		string(IosVpnConfigurationConnectionTypeciscoAnyConnectV2),
		string(IosVpnConfigurationConnectionTypeciscoIPSec),
		string(IosVpnConfigurationConnectionTypecitrix),
		string(IosVpnConfigurationConnectionTypecitrixSso),
		string(IosVpnConfigurationConnectionTypecustomVpn),
		string(IosVpnConfigurationConnectionTypedellSonicWallMobileConnect),
		string(IosVpnConfigurationConnectionTypef5Access2018),
		string(IosVpnConfigurationConnectionTypef5EdgeClient),
		string(IosVpnConfigurationConnectionTypeikEv2),
		string(IosVpnConfigurationConnectionTypemicrosoftProtect),
		string(IosVpnConfigurationConnectionTypemicrosoftTunnel),
		string(IosVpnConfigurationConnectionTypenetMotionMobility),
		string(IosVpnConfigurationConnectionTypepaloAltoGlobalProtect),
		string(IosVpnConfigurationConnectionTypepaloAltoGlobalProtectV2),
		string(IosVpnConfigurationConnectionTypepulseSecure),
		string(IosVpnConfigurationConnectionTypezscalerPrivateAccess),
	}
}

func parseIosVpnConfigurationConnectionType(input string) (*IosVpnConfigurationConnectionType, error) {
	vals := map[string]IosVpnConfigurationConnectionType{
		"alwayson":                   IosVpnConfigurationConnectionTypealwaysOn,
		"checkpointcapsulevpn":       IosVpnConfigurationConnectionTypecheckPointCapsuleVpn,
		"ciscoanyconnect":            IosVpnConfigurationConnectionTypeciscoAnyConnect,
		"ciscoanyconnectv2":          IosVpnConfigurationConnectionTypeciscoAnyConnectV2,
		"ciscoipsec":                 IosVpnConfigurationConnectionTypeciscoIPSec,
		"citrix":                     IosVpnConfigurationConnectionTypecitrix,
		"citrixsso":                  IosVpnConfigurationConnectionTypecitrixSso,
		"customvpn":                  IosVpnConfigurationConnectionTypecustomVpn,
		"dellsonicwallmobileconnect": IosVpnConfigurationConnectionTypedellSonicWallMobileConnect,
		"f5access2018":               IosVpnConfigurationConnectionTypef5Access2018,
		"f5edgeclient":               IosVpnConfigurationConnectionTypef5EdgeClient,
		"ikev2":                      IosVpnConfigurationConnectionTypeikEv2,
		"microsoftprotect":           IosVpnConfigurationConnectionTypemicrosoftProtect,
		"microsofttunnel":            IosVpnConfigurationConnectionTypemicrosoftTunnel,
		"netmotionmobility":          IosVpnConfigurationConnectionTypenetMotionMobility,
		"paloaltoglobalprotect":      IosVpnConfigurationConnectionTypepaloAltoGlobalProtect,
		"paloaltoglobalprotectv2":    IosVpnConfigurationConnectionTypepaloAltoGlobalProtectV2,
		"pulsesecure":                IosVpnConfigurationConnectionTypepulseSecure,
		"zscalerprivateaccess":       IosVpnConfigurationConnectionTypezscalerPrivateAccess,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := IosVpnConfigurationConnectionType(input)
	return &out, nil
}
