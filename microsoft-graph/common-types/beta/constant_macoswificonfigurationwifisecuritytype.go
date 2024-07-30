package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MacOSWiFiConfigurationWiFiSecurityType string

const (
	MacOSWiFiConfigurationWiFiSecurityType_Open           MacOSWiFiConfigurationWiFiSecurityType = "open"
	MacOSWiFiConfigurationWiFiSecurityType_Wep            MacOSWiFiConfigurationWiFiSecurityType = "wep"
	MacOSWiFiConfigurationWiFiSecurityType_Wpa2Enterprise MacOSWiFiConfigurationWiFiSecurityType = "wpa2Enterprise"
	MacOSWiFiConfigurationWiFiSecurityType_Wpa2Personal   MacOSWiFiConfigurationWiFiSecurityType = "wpa2Personal"
	MacOSWiFiConfigurationWiFiSecurityType_WpaEnterprise  MacOSWiFiConfigurationWiFiSecurityType = "wpaEnterprise"
	MacOSWiFiConfigurationWiFiSecurityType_WpaPersonal    MacOSWiFiConfigurationWiFiSecurityType = "wpaPersonal"
)

func PossibleValuesForMacOSWiFiConfigurationWiFiSecurityType() []string {
	return []string{
		string(MacOSWiFiConfigurationWiFiSecurityType_Open),
		string(MacOSWiFiConfigurationWiFiSecurityType_Wep),
		string(MacOSWiFiConfigurationWiFiSecurityType_Wpa2Enterprise),
		string(MacOSWiFiConfigurationWiFiSecurityType_Wpa2Personal),
		string(MacOSWiFiConfigurationWiFiSecurityType_WpaEnterprise),
		string(MacOSWiFiConfigurationWiFiSecurityType_WpaPersonal),
	}
}

func (s *MacOSWiFiConfigurationWiFiSecurityType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseMacOSWiFiConfigurationWiFiSecurityType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseMacOSWiFiConfigurationWiFiSecurityType(input string) (*MacOSWiFiConfigurationWiFiSecurityType, error) {
	vals := map[string]MacOSWiFiConfigurationWiFiSecurityType{
		"open":           MacOSWiFiConfigurationWiFiSecurityType_Open,
		"wep":            MacOSWiFiConfigurationWiFiSecurityType_Wep,
		"wpa2enterprise": MacOSWiFiConfigurationWiFiSecurityType_Wpa2Enterprise,
		"wpa2personal":   MacOSWiFiConfigurationWiFiSecurityType_Wpa2Personal,
		"wpaenterprise":  MacOSWiFiConfigurationWiFiSecurityType_WpaEnterprise,
		"wpapersonal":    MacOSWiFiConfigurationWiFiSecurityType_WpaPersonal,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := MacOSWiFiConfigurationWiFiSecurityType(input)
	return &out, nil
}
