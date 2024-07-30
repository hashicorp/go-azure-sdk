package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IosVpnConfigurationConnectionType string

const (
	IosVpnConfigurationConnectionType_AlwaysOn                   IosVpnConfigurationConnectionType = "alwaysOn"
	IosVpnConfigurationConnectionType_CheckPointCapsuleVpn       IosVpnConfigurationConnectionType = "checkPointCapsuleVpn"
	IosVpnConfigurationConnectionType_CiscoAnyConnect            IosVpnConfigurationConnectionType = "ciscoAnyConnect"
	IosVpnConfigurationConnectionType_CiscoAnyConnectV2          IosVpnConfigurationConnectionType = "ciscoAnyConnectV2"
	IosVpnConfigurationConnectionType_CiscoIPSec                 IosVpnConfigurationConnectionType = "ciscoIPSec"
	IosVpnConfigurationConnectionType_Citrix                     IosVpnConfigurationConnectionType = "citrix"
	IosVpnConfigurationConnectionType_CitrixSso                  IosVpnConfigurationConnectionType = "citrixSso"
	IosVpnConfigurationConnectionType_CustomVpn                  IosVpnConfigurationConnectionType = "customVpn"
	IosVpnConfigurationConnectionType_DellSonicWallMobileConnect IosVpnConfigurationConnectionType = "dellSonicWallMobileConnect"
	IosVpnConfigurationConnectionType_F5Access2018               IosVpnConfigurationConnectionType = "f5Access2018"
	IosVpnConfigurationConnectionType_F5EdgeClient               IosVpnConfigurationConnectionType = "f5EdgeClient"
	IosVpnConfigurationConnectionType_IkEv2                      IosVpnConfigurationConnectionType = "ikEv2"
	IosVpnConfigurationConnectionType_MicrosoftProtect           IosVpnConfigurationConnectionType = "microsoftProtect"
	IosVpnConfigurationConnectionType_MicrosoftTunnel            IosVpnConfigurationConnectionType = "microsoftTunnel"
	IosVpnConfigurationConnectionType_NetMotionMobility          IosVpnConfigurationConnectionType = "netMotionMobility"
	IosVpnConfigurationConnectionType_PaloAltoGlobalProtect      IosVpnConfigurationConnectionType = "paloAltoGlobalProtect"
	IosVpnConfigurationConnectionType_PaloAltoGlobalProtectV2    IosVpnConfigurationConnectionType = "paloAltoGlobalProtectV2"
	IosVpnConfigurationConnectionType_PulseSecure                IosVpnConfigurationConnectionType = "pulseSecure"
	IosVpnConfigurationConnectionType_ZscalerPrivateAccess       IosVpnConfigurationConnectionType = "zscalerPrivateAccess"
)

func PossibleValuesForIosVpnConfigurationConnectionType() []string {
	return []string{
		string(IosVpnConfigurationConnectionType_AlwaysOn),
		string(IosVpnConfigurationConnectionType_CheckPointCapsuleVpn),
		string(IosVpnConfigurationConnectionType_CiscoAnyConnect),
		string(IosVpnConfigurationConnectionType_CiscoAnyConnectV2),
		string(IosVpnConfigurationConnectionType_CiscoIPSec),
		string(IosVpnConfigurationConnectionType_Citrix),
		string(IosVpnConfigurationConnectionType_CitrixSso),
		string(IosVpnConfigurationConnectionType_CustomVpn),
		string(IosVpnConfigurationConnectionType_DellSonicWallMobileConnect),
		string(IosVpnConfigurationConnectionType_F5Access2018),
		string(IosVpnConfigurationConnectionType_F5EdgeClient),
		string(IosVpnConfigurationConnectionType_IkEv2),
		string(IosVpnConfigurationConnectionType_MicrosoftProtect),
		string(IosVpnConfigurationConnectionType_MicrosoftTunnel),
		string(IosVpnConfigurationConnectionType_NetMotionMobility),
		string(IosVpnConfigurationConnectionType_PaloAltoGlobalProtect),
		string(IosVpnConfigurationConnectionType_PaloAltoGlobalProtectV2),
		string(IosVpnConfigurationConnectionType_PulseSecure),
		string(IosVpnConfigurationConnectionType_ZscalerPrivateAccess),
	}
}

func (s *IosVpnConfigurationConnectionType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseIosVpnConfigurationConnectionType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseIosVpnConfigurationConnectionType(input string) (*IosVpnConfigurationConnectionType, error) {
	vals := map[string]IosVpnConfigurationConnectionType{
		"alwayson":                   IosVpnConfigurationConnectionType_AlwaysOn,
		"checkpointcapsulevpn":       IosVpnConfigurationConnectionType_CheckPointCapsuleVpn,
		"ciscoanyconnect":            IosVpnConfigurationConnectionType_CiscoAnyConnect,
		"ciscoanyconnectv2":          IosVpnConfigurationConnectionType_CiscoAnyConnectV2,
		"ciscoipsec":                 IosVpnConfigurationConnectionType_CiscoIPSec,
		"citrix":                     IosVpnConfigurationConnectionType_Citrix,
		"citrixsso":                  IosVpnConfigurationConnectionType_CitrixSso,
		"customvpn":                  IosVpnConfigurationConnectionType_CustomVpn,
		"dellsonicwallmobileconnect": IosVpnConfigurationConnectionType_DellSonicWallMobileConnect,
		"f5access2018":               IosVpnConfigurationConnectionType_F5Access2018,
		"f5edgeclient":               IosVpnConfigurationConnectionType_F5EdgeClient,
		"ikev2":                      IosVpnConfigurationConnectionType_IkEv2,
		"microsoftprotect":           IosVpnConfigurationConnectionType_MicrosoftProtect,
		"microsofttunnel":            IosVpnConfigurationConnectionType_MicrosoftTunnel,
		"netmotionmobility":          IosVpnConfigurationConnectionType_NetMotionMobility,
		"paloaltoglobalprotect":      IosVpnConfigurationConnectionType_PaloAltoGlobalProtect,
		"paloaltoglobalprotectv2":    IosVpnConfigurationConnectionType_PaloAltoGlobalProtectV2,
		"pulsesecure":                IosVpnConfigurationConnectionType_PulseSecure,
		"zscalerprivateaccess":       IosVpnConfigurationConnectionType_ZscalerPrivateAccess,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := IosVpnConfigurationConnectionType(input)
	return &out, nil
}
