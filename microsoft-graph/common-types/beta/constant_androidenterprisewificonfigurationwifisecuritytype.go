package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidEnterpriseWiFiConfigurationWiFiSecurityType string

const (
	AndroidEnterpriseWiFiConfigurationWiFiSecurityType_Open           AndroidEnterpriseWiFiConfigurationWiFiSecurityType = "open"
	AndroidEnterpriseWiFiConfigurationWiFiSecurityType_Wpa2Enterprise AndroidEnterpriseWiFiConfigurationWiFiSecurityType = "wpa2Enterprise"
	AndroidEnterpriseWiFiConfigurationWiFiSecurityType_WpaEnterprise  AndroidEnterpriseWiFiConfigurationWiFiSecurityType = "wpaEnterprise"
)

func PossibleValuesForAndroidEnterpriseWiFiConfigurationWiFiSecurityType() []string {
	return []string{
		string(AndroidEnterpriseWiFiConfigurationWiFiSecurityType_Open),
		string(AndroidEnterpriseWiFiConfigurationWiFiSecurityType_Wpa2Enterprise),
		string(AndroidEnterpriseWiFiConfigurationWiFiSecurityType_WpaEnterprise),
	}
}

func (s *AndroidEnterpriseWiFiConfigurationWiFiSecurityType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAndroidEnterpriseWiFiConfigurationWiFiSecurityType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAndroidEnterpriseWiFiConfigurationWiFiSecurityType(input string) (*AndroidEnterpriseWiFiConfigurationWiFiSecurityType, error) {
	vals := map[string]AndroidEnterpriseWiFiConfigurationWiFiSecurityType{
		"open":           AndroidEnterpriseWiFiConfigurationWiFiSecurityType_Open,
		"wpa2enterprise": AndroidEnterpriseWiFiConfigurationWiFiSecurityType_Wpa2Enterprise,
		"wpaenterprise":  AndroidEnterpriseWiFiConfigurationWiFiSecurityType_WpaEnterprise,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AndroidEnterpriseWiFiConfigurationWiFiSecurityType(input)
	return &out, nil
}
