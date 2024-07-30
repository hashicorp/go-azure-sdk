package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IosWiFiConfigurationWiFiSecurityType string

const (
	IosWiFiConfigurationWiFiSecurityType_Open           IosWiFiConfigurationWiFiSecurityType = "open"
	IosWiFiConfigurationWiFiSecurityType_Wep            IosWiFiConfigurationWiFiSecurityType = "wep"
	IosWiFiConfigurationWiFiSecurityType_Wpa2Enterprise IosWiFiConfigurationWiFiSecurityType = "wpa2Enterprise"
	IosWiFiConfigurationWiFiSecurityType_Wpa2Personal   IosWiFiConfigurationWiFiSecurityType = "wpa2Personal"
	IosWiFiConfigurationWiFiSecurityType_WpaEnterprise  IosWiFiConfigurationWiFiSecurityType = "wpaEnterprise"
	IosWiFiConfigurationWiFiSecurityType_WpaPersonal    IosWiFiConfigurationWiFiSecurityType = "wpaPersonal"
)

func PossibleValuesForIosWiFiConfigurationWiFiSecurityType() []string {
	return []string{
		string(IosWiFiConfigurationWiFiSecurityType_Open),
		string(IosWiFiConfigurationWiFiSecurityType_Wep),
		string(IosWiFiConfigurationWiFiSecurityType_Wpa2Enterprise),
		string(IosWiFiConfigurationWiFiSecurityType_Wpa2Personal),
		string(IosWiFiConfigurationWiFiSecurityType_WpaEnterprise),
		string(IosWiFiConfigurationWiFiSecurityType_WpaPersonal),
	}
}

func (s *IosWiFiConfigurationWiFiSecurityType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseIosWiFiConfigurationWiFiSecurityType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseIosWiFiConfigurationWiFiSecurityType(input string) (*IosWiFiConfigurationWiFiSecurityType, error) {
	vals := map[string]IosWiFiConfigurationWiFiSecurityType{
		"open":           IosWiFiConfigurationWiFiSecurityType_Open,
		"wep":            IosWiFiConfigurationWiFiSecurityType_Wep,
		"wpa2enterprise": IosWiFiConfigurationWiFiSecurityType_Wpa2Enterprise,
		"wpa2personal":   IosWiFiConfigurationWiFiSecurityType_Wpa2Personal,
		"wpaenterprise":  IosWiFiConfigurationWiFiSecurityType_WpaEnterprise,
		"wpapersonal":    IosWiFiConfigurationWiFiSecurityType_WpaPersonal,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := IosWiFiConfigurationWiFiSecurityType(input)
	return &out, nil
}
