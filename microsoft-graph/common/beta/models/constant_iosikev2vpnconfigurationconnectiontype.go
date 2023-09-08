package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IosikEv2VpnConfigurationConnectionType string

const (
	IosikEv2VpnConfigurationConnectionTypealwaysOn                   IosikEv2VpnConfigurationConnectionType = "AlwaysOn"
	IosikEv2VpnConfigurationConnectionTypecheckPointCapsuleVpn       IosikEv2VpnConfigurationConnectionType = "CheckPointCapsuleVpn"
	IosikEv2VpnConfigurationConnectionTypeciscoAnyConnect            IosikEv2VpnConfigurationConnectionType = "CiscoAnyConnect"
	IosikEv2VpnConfigurationConnectionTypeciscoAnyConnectV2          IosikEv2VpnConfigurationConnectionType = "CiscoAnyConnectV2"
	IosikEv2VpnConfigurationConnectionTypeciscoIPSec                 IosikEv2VpnConfigurationConnectionType = "CiscoIPSec"
	IosikEv2VpnConfigurationConnectionTypecitrix                     IosikEv2VpnConfigurationConnectionType = "Citrix"
	IosikEv2VpnConfigurationConnectionTypecitrixSso                  IosikEv2VpnConfigurationConnectionType = "CitrixSso"
	IosikEv2VpnConfigurationConnectionTypecustomVpn                  IosikEv2VpnConfigurationConnectionType = "CustomVpn"
	IosikEv2VpnConfigurationConnectionTypedellSonicWallMobileConnect IosikEv2VpnConfigurationConnectionType = "DellSonicWallMobileConnect"
	IosikEv2VpnConfigurationConnectionTypef5Access2018               IosikEv2VpnConfigurationConnectionType = "F5Access2018"
	IosikEv2VpnConfigurationConnectionTypef5EdgeClient               IosikEv2VpnConfigurationConnectionType = "F5EdgeClient"
	IosikEv2VpnConfigurationConnectionTypeikEv2                      IosikEv2VpnConfigurationConnectionType = "IkEv2"
	IosikEv2VpnConfigurationConnectionTypemicrosoftProtect           IosikEv2VpnConfigurationConnectionType = "MicrosoftProtect"
	IosikEv2VpnConfigurationConnectionTypemicrosoftTunnel            IosikEv2VpnConfigurationConnectionType = "MicrosoftTunnel"
	IosikEv2VpnConfigurationConnectionTypenetMotionMobility          IosikEv2VpnConfigurationConnectionType = "NetMotionMobility"
	IosikEv2VpnConfigurationConnectionTypepaloAltoGlobalProtect      IosikEv2VpnConfigurationConnectionType = "PaloAltoGlobalProtect"
	IosikEv2VpnConfigurationConnectionTypepaloAltoGlobalProtectV2    IosikEv2VpnConfigurationConnectionType = "PaloAltoGlobalProtectV2"
	IosikEv2VpnConfigurationConnectionTypepulseSecure                IosikEv2VpnConfigurationConnectionType = "PulseSecure"
	IosikEv2VpnConfigurationConnectionTypezscalerPrivateAccess       IosikEv2VpnConfigurationConnectionType = "ZscalerPrivateAccess"
)

func PossibleValuesForIosikEv2VpnConfigurationConnectionType() []string {
	return []string{
		string(IosikEv2VpnConfigurationConnectionTypealwaysOn),
		string(IosikEv2VpnConfigurationConnectionTypecheckPointCapsuleVpn),
		string(IosikEv2VpnConfigurationConnectionTypeciscoAnyConnect),
		string(IosikEv2VpnConfigurationConnectionTypeciscoAnyConnectV2),
		string(IosikEv2VpnConfigurationConnectionTypeciscoIPSec),
		string(IosikEv2VpnConfigurationConnectionTypecitrix),
		string(IosikEv2VpnConfigurationConnectionTypecitrixSso),
		string(IosikEv2VpnConfigurationConnectionTypecustomVpn),
		string(IosikEv2VpnConfigurationConnectionTypedellSonicWallMobileConnect),
		string(IosikEv2VpnConfigurationConnectionTypef5Access2018),
		string(IosikEv2VpnConfigurationConnectionTypef5EdgeClient),
		string(IosikEv2VpnConfigurationConnectionTypeikEv2),
		string(IosikEv2VpnConfigurationConnectionTypemicrosoftProtect),
		string(IosikEv2VpnConfigurationConnectionTypemicrosoftTunnel),
		string(IosikEv2VpnConfigurationConnectionTypenetMotionMobility),
		string(IosikEv2VpnConfigurationConnectionTypepaloAltoGlobalProtect),
		string(IosikEv2VpnConfigurationConnectionTypepaloAltoGlobalProtectV2),
		string(IosikEv2VpnConfigurationConnectionTypepulseSecure),
		string(IosikEv2VpnConfigurationConnectionTypezscalerPrivateAccess),
	}
}

func parseIosikEv2VpnConfigurationConnectionType(input string) (*IosikEv2VpnConfigurationConnectionType, error) {
	vals := map[string]IosikEv2VpnConfigurationConnectionType{
		"alwayson":                   IosikEv2VpnConfigurationConnectionTypealwaysOn,
		"checkpointcapsulevpn":       IosikEv2VpnConfigurationConnectionTypecheckPointCapsuleVpn,
		"ciscoanyconnect":            IosikEv2VpnConfigurationConnectionTypeciscoAnyConnect,
		"ciscoanyconnectv2":          IosikEv2VpnConfigurationConnectionTypeciscoAnyConnectV2,
		"ciscoipsec":                 IosikEv2VpnConfigurationConnectionTypeciscoIPSec,
		"citrix":                     IosikEv2VpnConfigurationConnectionTypecitrix,
		"citrixsso":                  IosikEv2VpnConfigurationConnectionTypecitrixSso,
		"customvpn":                  IosikEv2VpnConfigurationConnectionTypecustomVpn,
		"dellsonicwallmobileconnect": IosikEv2VpnConfigurationConnectionTypedellSonicWallMobileConnect,
		"f5access2018":               IosikEv2VpnConfigurationConnectionTypef5Access2018,
		"f5edgeclient":               IosikEv2VpnConfigurationConnectionTypef5EdgeClient,
		"ikev2":                      IosikEv2VpnConfigurationConnectionTypeikEv2,
		"microsoftprotect":           IosikEv2VpnConfigurationConnectionTypemicrosoftProtect,
		"microsofttunnel":            IosikEv2VpnConfigurationConnectionTypemicrosoftTunnel,
		"netmotionmobility":          IosikEv2VpnConfigurationConnectionTypenetMotionMobility,
		"paloaltoglobalprotect":      IosikEv2VpnConfigurationConnectionTypepaloAltoGlobalProtect,
		"paloaltoglobalprotectv2":    IosikEv2VpnConfigurationConnectionTypepaloAltoGlobalProtectV2,
		"pulsesecure":                IosikEv2VpnConfigurationConnectionTypepulseSecure,
		"zscalerprivateaccess":       IosikEv2VpnConfigurationConnectionTypezscalerPrivateAccess,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := IosikEv2VpnConfigurationConnectionType(input)
	return &out, nil
}
