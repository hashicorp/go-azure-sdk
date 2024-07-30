package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IosikEv2VpnConfigurationConnectionType string

const (
	IosikEv2VpnConfigurationConnectionType_AlwaysOn                   IosikEv2VpnConfigurationConnectionType = "alwaysOn"
	IosikEv2VpnConfigurationConnectionType_CheckPointCapsuleVpn       IosikEv2VpnConfigurationConnectionType = "checkPointCapsuleVpn"
	IosikEv2VpnConfigurationConnectionType_CiscoAnyConnect            IosikEv2VpnConfigurationConnectionType = "ciscoAnyConnect"
	IosikEv2VpnConfigurationConnectionType_CiscoAnyConnectV2          IosikEv2VpnConfigurationConnectionType = "ciscoAnyConnectV2"
	IosikEv2VpnConfigurationConnectionType_CiscoIPSec                 IosikEv2VpnConfigurationConnectionType = "ciscoIPSec"
	IosikEv2VpnConfigurationConnectionType_Citrix                     IosikEv2VpnConfigurationConnectionType = "citrix"
	IosikEv2VpnConfigurationConnectionType_CitrixSso                  IosikEv2VpnConfigurationConnectionType = "citrixSso"
	IosikEv2VpnConfigurationConnectionType_CustomVpn                  IosikEv2VpnConfigurationConnectionType = "customVpn"
	IosikEv2VpnConfigurationConnectionType_DellSonicWallMobileConnect IosikEv2VpnConfigurationConnectionType = "dellSonicWallMobileConnect"
	IosikEv2VpnConfigurationConnectionType_F5Access2018               IosikEv2VpnConfigurationConnectionType = "f5Access2018"
	IosikEv2VpnConfigurationConnectionType_F5EdgeClient               IosikEv2VpnConfigurationConnectionType = "f5EdgeClient"
	IosikEv2VpnConfigurationConnectionType_IkEv2                      IosikEv2VpnConfigurationConnectionType = "ikEv2"
	IosikEv2VpnConfigurationConnectionType_MicrosoftProtect           IosikEv2VpnConfigurationConnectionType = "microsoftProtect"
	IosikEv2VpnConfigurationConnectionType_MicrosoftTunnel            IosikEv2VpnConfigurationConnectionType = "microsoftTunnel"
	IosikEv2VpnConfigurationConnectionType_NetMotionMobility          IosikEv2VpnConfigurationConnectionType = "netMotionMobility"
	IosikEv2VpnConfigurationConnectionType_PaloAltoGlobalProtect      IosikEv2VpnConfigurationConnectionType = "paloAltoGlobalProtect"
	IosikEv2VpnConfigurationConnectionType_PaloAltoGlobalProtectV2    IosikEv2VpnConfigurationConnectionType = "paloAltoGlobalProtectV2"
	IosikEv2VpnConfigurationConnectionType_PulseSecure                IosikEv2VpnConfigurationConnectionType = "pulseSecure"
	IosikEv2VpnConfigurationConnectionType_ZscalerPrivateAccess       IosikEv2VpnConfigurationConnectionType = "zscalerPrivateAccess"
)

func PossibleValuesForIosikEv2VpnConfigurationConnectionType() []string {
	return []string{
		string(IosikEv2VpnConfigurationConnectionType_AlwaysOn),
		string(IosikEv2VpnConfigurationConnectionType_CheckPointCapsuleVpn),
		string(IosikEv2VpnConfigurationConnectionType_CiscoAnyConnect),
		string(IosikEv2VpnConfigurationConnectionType_CiscoAnyConnectV2),
		string(IosikEv2VpnConfigurationConnectionType_CiscoIPSec),
		string(IosikEv2VpnConfigurationConnectionType_Citrix),
		string(IosikEv2VpnConfigurationConnectionType_CitrixSso),
		string(IosikEv2VpnConfigurationConnectionType_CustomVpn),
		string(IosikEv2VpnConfigurationConnectionType_DellSonicWallMobileConnect),
		string(IosikEv2VpnConfigurationConnectionType_F5Access2018),
		string(IosikEv2VpnConfigurationConnectionType_F5EdgeClient),
		string(IosikEv2VpnConfigurationConnectionType_IkEv2),
		string(IosikEv2VpnConfigurationConnectionType_MicrosoftProtect),
		string(IosikEv2VpnConfigurationConnectionType_MicrosoftTunnel),
		string(IosikEv2VpnConfigurationConnectionType_NetMotionMobility),
		string(IosikEv2VpnConfigurationConnectionType_PaloAltoGlobalProtect),
		string(IosikEv2VpnConfigurationConnectionType_PaloAltoGlobalProtectV2),
		string(IosikEv2VpnConfigurationConnectionType_PulseSecure),
		string(IosikEv2VpnConfigurationConnectionType_ZscalerPrivateAccess),
	}
}

func (s *IosikEv2VpnConfigurationConnectionType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseIosikEv2VpnConfigurationConnectionType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseIosikEv2VpnConfigurationConnectionType(input string) (*IosikEv2VpnConfigurationConnectionType, error) {
	vals := map[string]IosikEv2VpnConfigurationConnectionType{
		"alwayson":                   IosikEv2VpnConfigurationConnectionType_AlwaysOn,
		"checkpointcapsulevpn":       IosikEv2VpnConfigurationConnectionType_CheckPointCapsuleVpn,
		"ciscoanyconnect":            IosikEv2VpnConfigurationConnectionType_CiscoAnyConnect,
		"ciscoanyconnectv2":          IosikEv2VpnConfigurationConnectionType_CiscoAnyConnectV2,
		"ciscoipsec":                 IosikEv2VpnConfigurationConnectionType_CiscoIPSec,
		"citrix":                     IosikEv2VpnConfigurationConnectionType_Citrix,
		"citrixsso":                  IosikEv2VpnConfigurationConnectionType_CitrixSso,
		"customvpn":                  IosikEv2VpnConfigurationConnectionType_CustomVpn,
		"dellsonicwallmobileconnect": IosikEv2VpnConfigurationConnectionType_DellSonicWallMobileConnect,
		"f5access2018":               IosikEv2VpnConfigurationConnectionType_F5Access2018,
		"f5edgeclient":               IosikEv2VpnConfigurationConnectionType_F5EdgeClient,
		"ikev2":                      IosikEv2VpnConfigurationConnectionType_IkEv2,
		"microsoftprotect":           IosikEv2VpnConfigurationConnectionType_MicrosoftProtect,
		"microsofttunnel":            IosikEv2VpnConfigurationConnectionType_MicrosoftTunnel,
		"netmotionmobility":          IosikEv2VpnConfigurationConnectionType_NetMotionMobility,
		"paloaltoglobalprotect":      IosikEv2VpnConfigurationConnectionType_PaloAltoGlobalProtect,
		"paloaltoglobalprotectv2":    IosikEv2VpnConfigurationConnectionType_PaloAltoGlobalProtectV2,
		"pulsesecure":                IosikEv2VpnConfigurationConnectionType_PulseSecure,
		"zscalerprivateaccess":       IosikEv2VpnConfigurationConnectionType_ZscalerPrivateAccess,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := IosikEv2VpnConfigurationConnectionType(input)
	return &out, nil
}
