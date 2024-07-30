package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MacOSVpnConfigurationConnectionType string

const (
	MacOSVpnConfigurationConnectionType_AlwaysOn                   MacOSVpnConfigurationConnectionType = "alwaysOn"
	MacOSVpnConfigurationConnectionType_CheckPointCapsuleVpn       MacOSVpnConfigurationConnectionType = "checkPointCapsuleVpn"
	MacOSVpnConfigurationConnectionType_CiscoAnyConnect            MacOSVpnConfigurationConnectionType = "ciscoAnyConnect"
	MacOSVpnConfigurationConnectionType_CiscoAnyConnectV2          MacOSVpnConfigurationConnectionType = "ciscoAnyConnectV2"
	MacOSVpnConfigurationConnectionType_CiscoIPSec                 MacOSVpnConfigurationConnectionType = "ciscoIPSec"
	MacOSVpnConfigurationConnectionType_Citrix                     MacOSVpnConfigurationConnectionType = "citrix"
	MacOSVpnConfigurationConnectionType_CitrixSso                  MacOSVpnConfigurationConnectionType = "citrixSso"
	MacOSVpnConfigurationConnectionType_CustomVpn                  MacOSVpnConfigurationConnectionType = "customVpn"
	MacOSVpnConfigurationConnectionType_DellSonicWallMobileConnect MacOSVpnConfigurationConnectionType = "dellSonicWallMobileConnect"
	MacOSVpnConfigurationConnectionType_F5Access2018               MacOSVpnConfigurationConnectionType = "f5Access2018"
	MacOSVpnConfigurationConnectionType_F5EdgeClient               MacOSVpnConfigurationConnectionType = "f5EdgeClient"
	MacOSVpnConfigurationConnectionType_IkEv2                      MacOSVpnConfigurationConnectionType = "ikEv2"
	MacOSVpnConfigurationConnectionType_MicrosoftProtect           MacOSVpnConfigurationConnectionType = "microsoftProtect"
	MacOSVpnConfigurationConnectionType_MicrosoftTunnel            MacOSVpnConfigurationConnectionType = "microsoftTunnel"
	MacOSVpnConfigurationConnectionType_NetMotionMobility          MacOSVpnConfigurationConnectionType = "netMotionMobility"
	MacOSVpnConfigurationConnectionType_PaloAltoGlobalProtect      MacOSVpnConfigurationConnectionType = "paloAltoGlobalProtect"
	MacOSVpnConfigurationConnectionType_PaloAltoGlobalProtectV2    MacOSVpnConfigurationConnectionType = "paloAltoGlobalProtectV2"
	MacOSVpnConfigurationConnectionType_PulseSecure                MacOSVpnConfigurationConnectionType = "pulseSecure"
	MacOSVpnConfigurationConnectionType_ZscalerPrivateAccess       MacOSVpnConfigurationConnectionType = "zscalerPrivateAccess"
)

func PossibleValuesForMacOSVpnConfigurationConnectionType() []string {
	return []string{
		string(MacOSVpnConfigurationConnectionType_AlwaysOn),
		string(MacOSVpnConfigurationConnectionType_CheckPointCapsuleVpn),
		string(MacOSVpnConfigurationConnectionType_CiscoAnyConnect),
		string(MacOSVpnConfigurationConnectionType_CiscoAnyConnectV2),
		string(MacOSVpnConfigurationConnectionType_CiscoIPSec),
		string(MacOSVpnConfigurationConnectionType_Citrix),
		string(MacOSVpnConfigurationConnectionType_CitrixSso),
		string(MacOSVpnConfigurationConnectionType_CustomVpn),
		string(MacOSVpnConfigurationConnectionType_DellSonicWallMobileConnect),
		string(MacOSVpnConfigurationConnectionType_F5Access2018),
		string(MacOSVpnConfigurationConnectionType_F5EdgeClient),
		string(MacOSVpnConfigurationConnectionType_IkEv2),
		string(MacOSVpnConfigurationConnectionType_MicrosoftProtect),
		string(MacOSVpnConfigurationConnectionType_MicrosoftTunnel),
		string(MacOSVpnConfigurationConnectionType_NetMotionMobility),
		string(MacOSVpnConfigurationConnectionType_PaloAltoGlobalProtect),
		string(MacOSVpnConfigurationConnectionType_PaloAltoGlobalProtectV2),
		string(MacOSVpnConfigurationConnectionType_PulseSecure),
		string(MacOSVpnConfigurationConnectionType_ZscalerPrivateAccess),
	}
}

func (s *MacOSVpnConfigurationConnectionType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseMacOSVpnConfigurationConnectionType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseMacOSVpnConfigurationConnectionType(input string) (*MacOSVpnConfigurationConnectionType, error) {
	vals := map[string]MacOSVpnConfigurationConnectionType{
		"alwayson":                   MacOSVpnConfigurationConnectionType_AlwaysOn,
		"checkpointcapsulevpn":       MacOSVpnConfigurationConnectionType_CheckPointCapsuleVpn,
		"ciscoanyconnect":            MacOSVpnConfigurationConnectionType_CiscoAnyConnect,
		"ciscoanyconnectv2":          MacOSVpnConfigurationConnectionType_CiscoAnyConnectV2,
		"ciscoipsec":                 MacOSVpnConfigurationConnectionType_CiscoIPSec,
		"citrix":                     MacOSVpnConfigurationConnectionType_Citrix,
		"citrixsso":                  MacOSVpnConfigurationConnectionType_CitrixSso,
		"customvpn":                  MacOSVpnConfigurationConnectionType_CustomVpn,
		"dellsonicwallmobileconnect": MacOSVpnConfigurationConnectionType_DellSonicWallMobileConnect,
		"f5access2018":               MacOSVpnConfigurationConnectionType_F5Access2018,
		"f5edgeclient":               MacOSVpnConfigurationConnectionType_F5EdgeClient,
		"ikev2":                      MacOSVpnConfigurationConnectionType_IkEv2,
		"microsoftprotect":           MacOSVpnConfigurationConnectionType_MicrosoftProtect,
		"microsofttunnel":            MacOSVpnConfigurationConnectionType_MicrosoftTunnel,
		"netmotionmobility":          MacOSVpnConfigurationConnectionType_NetMotionMobility,
		"paloaltoglobalprotect":      MacOSVpnConfigurationConnectionType_PaloAltoGlobalProtect,
		"paloaltoglobalprotectv2":    MacOSVpnConfigurationConnectionType_PaloAltoGlobalProtectV2,
		"pulsesecure":                MacOSVpnConfigurationConnectionType_PulseSecure,
		"zscalerprivateaccess":       MacOSVpnConfigurationConnectionType_ZscalerPrivateAccess,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := MacOSVpnConfigurationConnectionType(input)
	return &out, nil
}
