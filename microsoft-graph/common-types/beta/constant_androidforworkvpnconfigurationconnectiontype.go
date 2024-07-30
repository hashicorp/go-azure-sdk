package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidForWorkVpnConfigurationConnectionType string

const (
	AndroidForWorkVpnConfigurationConnectionType_CheckPointCapsuleVpn       AndroidForWorkVpnConfigurationConnectionType = "checkPointCapsuleVpn"
	AndroidForWorkVpnConfigurationConnectionType_CiscoAnyConnect            AndroidForWorkVpnConfigurationConnectionType = "ciscoAnyConnect"
	AndroidForWorkVpnConfigurationConnectionType_Citrix                     AndroidForWorkVpnConfigurationConnectionType = "citrix"
	AndroidForWorkVpnConfigurationConnectionType_DellSonicWallMobileConnect AndroidForWorkVpnConfigurationConnectionType = "dellSonicWallMobileConnect"
	AndroidForWorkVpnConfigurationConnectionType_F5EdgeClient               AndroidForWorkVpnConfigurationConnectionType = "f5EdgeClient"
	AndroidForWorkVpnConfigurationConnectionType_PulseSecure                AndroidForWorkVpnConfigurationConnectionType = "pulseSecure"
)

func PossibleValuesForAndroidForWorkVpnConfigurationConnectionType() []string {
	return []string{
		string(AndroidForWorkVpnConfigurationConnectionType_CheckPointCapsuleVpn),
		string(AndroidForWorkVpnConfigurationConnectionType_CiscoAnyConnect),
		string(AndroidForWorkVpnConfigurationConnectionType_Citrix),
		string(AndroidForWorkVpnConfigurationConnectionType_DellSonicWallMobileConnect),
		string(AndroidForWorkVpnConfigurationConnectionType_F5EdgeClient),
		string(AndroidForWorkVpnConfigurationConnectionType_PulseSecure),
	}
}

func (s *AndroidForWorkVpnConfigurationConnectionType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAndroidForWorkVpnConfigurationConnectionType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAndroidForWorkVpnConfigurationConnectionType(input string) (*AndroidForWorkVpnConfigurationConnectionType, error) {
	vals := map[string]AndroidForWorkVpnConfigurationConnectionType{
		"checkpointcapsulevpn":       AndroidForWorkVpnConfigurationConnectionType_CheckPointCapsuleVpn,
		"ciscoanyconnect":            AndroidForWorkVpnConfigurationConnectionType_CiscoAnyConnect,
		"citrix":                     AndroidForWorkVpnConfigurationConnectionType_Citrix,
		"dellsonicwallmobileconnect": AndroidForWorkVpnConfigurationConnectionType_DellSonicWallMobileConnect,
		"f5edgeclient":               AndroidForWorkVpnConfigurationConnectionType_F5EdgeClient,
		"pulsesecure":                AndroidForWorkVpnConfigurationConnectionType_PulseSecure,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AndroidForWorkVpnConfigurationConnectionType(input)
	return &out, nil
}
