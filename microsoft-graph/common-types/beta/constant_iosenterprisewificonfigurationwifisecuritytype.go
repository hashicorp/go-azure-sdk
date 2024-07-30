package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IosEnterpriseWiFiConfigurationWiFiSecurityType string

const (
	IosEnterpriseWiFiConfigurationWiFiSecurityType_Open           IosEnterpriseWiFiConfigurationWiFiSecurityType = "open"
	IosEnterpriseWiFiConfigurationWiFiSecurityType_Wep            IosEnterpriseWiFiConfigurationWiFiSecurityType = "wep"
	IosEnterpriseWiFiConfigurationWiFiSecurityType_Wpa2Enterprise IosEnterpriseWiFiConfigurationWiFiSecurityType = "wpa2Enterprise"
	IosEnterpriseWiFiConfigurationWiFiSecurityType_Wpa2Personal   IosEnterpriseWiFiConfigurationWiFiSecurityType = "wpa2Personal"
	IosEnterpriseWiFiConfigurationWiFiSecurityType_WpaEnterprise  IosEnterpriseWiFiConfigurationWiFiSecurityType = "wpaEnterprise"
	IosEnterpriseWiFiConfigurationWiFiSecurityType_WpaPersonal    IosEnterpriseWiFiConfigurationWiFiSecurityType = "wpaPersonal"
)

func PossibleValuesForIosEnterpriseWiFiConfigurationWiFiSecurityType() []string {
	return []string{
		string(IosEnterpriseWiFiConfigurationWiFiSecurityType_Open),
		string(IosEnterpriseWiFiConfigurationWiFiSecurityType_Wep),
		string(IosEnterpriseWiFiConfigurationWiFiSecurityType_Wpa2Enterprise),
		string(IosEnterpriseWiFiConfigurationWiFiSecurityType_Wpa2Personal),
		string(IosEnterpriseWiFiConfigurationWiFiSecurityType_WpaEnterprise),
		string(IosEnterpriseWiFiConfigurationWiFiSecurityType_WpaPersonal),
	}
}

func (s *IosEnterpriseWiFiConfigurationWiFiSecurityType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseIosEnterpriseWiFiConfigurationWiFiSecurityType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseIosEnterpriseWiFiConfigurationWiFiSecurityType(input string) (*IosEnterpriseWiFiConfigurationWiFiSecurityType, error) {
	vals := map[string]IosEnterpriseWiFiConfigurationWiFiSecurityType{
		"open":           IosEnterpriseWiFiConfigurationWiFiSecurityType_Open,
		"wep":            IosEnterpriseWiFiConfigurationWiFiSecurityType_Wep,
		"wpa2enterprise": IosEnterpriseWiFiConfigurationWiFiSecurityType_Wpa2Enterprise,
		"wpa2personal":   IosEnterpriseWiFiConfigurationWiFiSecurityType_Wpa2Personal,
		"wpaenterprise":  IosEnterpriseWiFiConfigurationWiFiSecurityType_WpaEnterprise,
		"wpapersonal":    IosEnterpriseWiFiConfigurationWiFiSecurityType_WpaPersonal,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := IosEnterpriseWiFiConfigurationWiFiSecurityType(input)
	return &out, nil
}
