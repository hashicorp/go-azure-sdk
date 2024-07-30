package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MacOSEnterpriseWiFiConfigurationWiFiSecurityType string

const (
	MacOSEnterpriseWiFiConfigurationWiFiSecurityType_Open           MacOSEnterpriseWiFiConfigurationWiFiSecurityType = "open"
	MacOSEnterpriseWiFiConfigurationWiFiSecurityType_Wep            MacOSEnterpriseWiFiConfigurationWiFiSecurityType = "wep"
	MacOSEnterpriseWiFiConfigurationWiFiSecurityType_Wpa2Enterprise MacOSEnterpriseWiFiConfigurationWiFiSecurityType = "wpa2Enterprise"
	MacOSEnterpriseWiFiConfigurationWiFiSecurityType_Wpa2Personal   MacOSEnterpriseWiFiConfigurationWiFiSecurityType = "wpa2Personal"
	MacOSEnterpriseWiFiConfigurationWiFiSecurityType_WpaEnterprise  MacOSEnterpriseWiFiConfigurationWiFiSecurityType = "wpaEnterprise"
	MacOSEnterpriseWiFiConfigurationWiFiSecurityType_WpaPersonal    MacOSEnterpriseWiFiConfigurationWiFiSecurityType = "wpaPersonal"
)

func PossibleValuesForMacOSEnterpriseWiFiConfigurationWiFiSecurityType() []string {
	return []string{
		string(MacOSEnterpriseWiFiConfigurationWiFiSecurityType_Open),
		string(MacOSEnterpriseWiFiConfigurationWiFiSecurityType_Wep),
		string(MacOSEnterpriseWiFiConfigurationWiFiSecurityType_Wpa2Enterprise),
		string(MacOSEnterpriseWiFiConfigurationWiFiSecurityType_Wpa2Personal),
		string(MacOSEnterpriseWiFiConfigurationWiFiSecurityType_WpaEnterprise),
		string(MacOSEnterpriseWiFiConfigurationWiFiSecurityType_WpaPersonal),
	}
}

func (s *MacOSEnterpriseWiFiConfigurationWiFiSecurityType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseMacOSEnterpriseWiFiConfigurationWiFiSecurityType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseMacOSEnterpriseWiFiConfigurationWiFiSecurityType(input string) (*MacOSEnterpriseWiFiConfigurationWiFiSecurityType, error) {
	vals := map[string]MacOSEnterpriseWiFiConfigurationWiFiSecurityType{
		"open":           MacOSEnterpriseWiFiConfigurationWiFiSecurityType_Open,
		"wep":            MacOSEnterpriseWiFiConfigurationWiFiSecurityType_Wep,
		"wpa2enterprise": MacOSEnterpriseWiFiConfigurationWiFiSecurityType_Wpa2Enterprise,
		"wpa2personal":   MacOSEnterpriseWiFiConfigurationWiFiSecurityType_Wpa2Personal,
		"wpaenterprise":  MacOSEnterpriseWiFiConfigurationWiFiSecurityType_WpaEnterprise,
		"wpapersonal":    MacOSEnterpriseWiFiConfigurationWiFiSecurityType_WpaPersonal,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := MacOSEnterpriseWiFiConfigurationWiFiSecurityType(input)
	return &out, nil
}
