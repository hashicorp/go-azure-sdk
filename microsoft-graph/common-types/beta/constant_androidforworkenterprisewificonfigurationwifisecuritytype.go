package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidForWorkEnterpriseWiFiConfigurationWiFiSecurityType string

const (
	AndroidForWorkEnterpriseWiFiConfigurationWiFiSecurityType_Open           AndroidForWorkEnterpriseWiFiConfigurationWiFiSecurityType = "open"
	AndroidForWorkEnterpriseWiFiConfigurationWiFiSecurityType_Wpa2Enterprise AndroidForWorkEnterpriseWiFiConfigurationWiFiSecurityType = "wpa2Enterprise"
	AndroidForWorkEnterpriseWiFiConfigurationWiFiSecurityType_WpaEnterprise  AndroidForWorkEnterpriseWiFiConfigurationWiFiSecurityType = "wpaEnterprise"
)

func PossibleValuesForAndroidForWorkEnterpriseWiFiConfigurationWiFiSecurityType() []string {
	return []string{
		string(AndroidForWorkEnterpriseWiFiConfigurationWiFiSecurityType_Open),
		string(AndroidForWorkEnterpriseWiFiConfigurationWiFiSecurityType_Wpa2Enterprise),
		string(AndroidForWorkEnterpriseWiFiConfigurationWiFiSecurityType_WpaEnterprise),
	}
}

func (s *AndroidForWorkEnterpriseWiFiConfigurationWiFiSecurityType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAndroidForWorkEnterpriseWiFiConfigurationWiFiSecurityType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAndroidForWorkEnterpriseWiFiConfigurationWiFiSecurityType(input string) (*AndroidForWorkEnterpriseWiFiConfigurationWiFiSecurityType, error) {
	vals := map[string]AndroidForWorkEnterpriseWiFiConfigurationWiFiSecurityType{
		"open":           AndroidForWorkEnterpriseWiFiConfigurationWiFiSecurityType_Open,
		"wpa2enterprise": AndroidForWorkEnterpriseWiFiConfigurationWiFiSecurityType_Wpa2Enterprise,
		"wpaenterprise":  AndroidForWorkEnterpriseWiFiConfigurationWiFiSecurityType_WpaEnterprise,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AndroidForWorkEnterpriseWiFiConfigurationWiFiSecurityType(input)
	return &out, nil
}
