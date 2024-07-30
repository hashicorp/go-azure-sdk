package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AppleVpnConfigurationConnectionType string

const (
	AppleVpnConfigurationConnectionType_AlwaysOn                   AppleVpnConfigurationConnectionType = "alwaysOn"
	AppleVpnConfigurationConnectionType_CheckPointCapsuleVpn       AppleVpnConfigurationConnectionType = "checkPointCapsuleVpn"
	AppleVpnConfigurationConnectionType_CiscoAnyConnect            AppleVpnConfigurationConnectionType = "ciscoAnyConnect"
	AppleVpnConfigurationConnectionType_CiscoAnyConnectV2          AppleVpnConfigurationConnectionType = "ciscoAnyConnectV2"
	AppleVpnConfigurationConnectionType_CiscoIPSec                 AppleVpnConfigurationConnectionType = "ciscoIPSec"
	AppleVpnConfigurationConnectionType_Citrix                     AppleVpnConfigurationConnectionType = "citrix"
	AppleVpnConfigurationConnectionType_CitrixSso                  AppleVpnConfigurationConnectionType = "citrixSso"
	AppleVpnConfigurationConnectionType_CustomVpn                  AppleVpnConfigurationConnectionType = "customVpn"
	AppleVpnConfigurationConnectionType_DellSonicWallMobileConnect AppleVpnConfigurationConnectionType = "dellSonicWallMobileConnect"
	AppleVpnConfigurationConnectionType_F5Access2018               AppleVpnConfigurationConnectionType = "f5Access2018"
	AppleVpnConfigurationConnectionType_F5EdgeClient               AppleVpnConfigurationConnectionType = "f5EdgeClient"
	AppleVpnConfigurationConnectionType_IkEv2                      AppleVpnConfigurationConnectionType = "ikEv2"
	AppleVpnConfigurationConnectionType_MicrosoftProtect           AppleVpnConfigurationConnectionType = "microsoftProtect"
	AppleVpnConfigurationConnectionType_MicrosoftTunnel            AppleVpnConfigurationConnectionType = "microsoftTunnel"
	AppleVpnConfigurationConnectionType_NetMotionMobility          AppleVpnConfigurationConnectionType = "netMotionMobility"
	AppleVpnConfigurationConnectionType_PaloAltoGlobalProtect      AppleVpnConfigurationConnectionType = "paloAltoGlobalProtect"
	AppleVpnConfigurationConnectionType_PaloAltoGlobalProtectV2    AppleVpnConfigurationConnectionType = "paloAltoGlobalProtectV2"
	AppleVpnConfigurationConnectionType_PulseSecure                AppleVpnConfigurationConnectionType = "pulseSecure"
	AppleVpnConfigurationConnectionType_ZscalerPrivateAccess       AppleVpnConfigurationConnectionType = "zscalerPrivateAccess"
)

func PossibleValuesForAppleVpnConfigurationConnectionType() []string {
	return []string{
		string(AppleVpnConfigurationConnectionType_AlwaysOn),
		string(AppleVpnConfigurationConnectionType_CheckPointCapsuleVpn),
		string(AppleVpnConfigurationConnectionType_CiscoAnyConnect),
		string(AppleVpnConfigurationConnectionType_CiscoAnyConnectV2),
		string(AppleVpnConfigurationConnectionType_CiscoIPSec),
		string(AppleVpnConfigurationConnectionType_Citrix),
		string(AppleVpnConfigurationConnectionType_CitrixSso),
		string(AppleVpnConfigurationConnectionType_CustomVpn),
		string(AppleVpnConfigurationConnectionType_DellSonicWallMobileConnect),
		string(AppleVpnConfigurationConnectionType_F5Access2018),
		string(AppleVpnConfigurationConnectionType_F5EdgeClient),
		string(AppleVpnConfigurationConnectionType_IkEv2),
		string(AppleVpnConfigurationConnectionType_MicrosoftProtect),
		string(AppleVpnConfigurationConnectionType_MicrosoftTunnel),
		string(AppleVpnConfigurationConnectionType_NetMotionMobility),
		string(AppleVpnConfigurationConnectionType_PaloAltoGlobalProtect),
		string(AppleVpnConfigurationConnectionType_PaloAltoGlobalProtectV2),
		string(AppleVpnConfigurationConnectionType_PulseSecure),
		string(AppleVpnConfigurationConnectionType_ZscalerPrivateAccess),
	}
}

func (s *AppleVpnConfigurationConnectionType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAppleVpnConfigurationConnectionType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAppleVpnConfigurationConnectionType(input string) (*AppleVpnConfigurationConnectionType, error) {
	vals := map[string]AppleVpnConfigurationConnectionType{
		"alwayson":                   AppleVpnConfigurationConnectionType_AlwaysOn,
		"checkpointcapsulevpn":       AppleVpnConfigurationConnectionType_CheckPointCapsuleVpn,
		"ciscoanyconnect":            AppleVpnConfigurationConnectionType_CiscoAnyConnect,
		"ciscoanyconnectv2":          AppleVpnConfigurationConnectionType_CiscoAnyConnectV2,
		"ciscoipsec":                 AppleVpnConfigurationConnectionType_CiscoIPSec,
		"citrix":                     AppleVpnConfigurationConnectionType_Citrix,
		"citrixsso":                  AppleVpnConfigurationConnectionType_CitrixSso,
		"customvpn":                  AppleVpnConfigurationConnectionType_CustomVpn,
		"dellsonicwallmobileconnect": AppleVpnConfigurationConnectionType_DellSonicWallMobileConnect,
		"f5access2018":               AppleVpnConfigurationConnectionType_F5Access2018,
		"f5edgeclient":               AppleVpnConfigurationConnectionType_F5EdgeClient,
		"ikev2":                      AppleVpnConfigurationConnectionType_IkEv2,
		"microsoftprotect":           AppleVpnConfigurationConnectionType_MicrosoftProtect,
		"microsofttunnel":            AppleVpnConfigurationConnectionType_MicrosoftTunnel,
		"netmotionmobility":          AppleVpnConfigurationConnectionType_NetMotionMobility,
		"paloaltoglobalprotect":      AppleVpnConfigurationConnectionType_PaloAltoGlobalProtect,
		"paloaltoglobalprotectv2":    AppleVpnConfigurationConnectionType_PaloAltoGlobalProtectV2,
		"pulsesecure":                AppleVpnConfigurationConnectionType_PulseSecure,
		"zscalerprivateaccess":       AppleVpnConfigurationConnectionType_ZscalerPrivateAccess,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AppleVpnConfigurationConnectionType(input)
	return &out, nil
}
