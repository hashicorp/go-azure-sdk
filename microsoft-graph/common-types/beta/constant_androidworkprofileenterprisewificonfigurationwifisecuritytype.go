package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidWorkProfileEnterpriseWiFiConfigurationWiFiSecurityType string

const (
	AndroidWorkProfileEnterpriseWiFiConfigurationWiFiSecurityType_Open           AndroidWorkProfileEnterpriseWiFiConfigurationWiFiSecurityType = "open"
	AndroidWorkProfileEnterpriseWiFiConfigurationWiFiSecurityType_Wpa2Enterprise AndroidWorkProfileEnterpriseWiFiConfigurationWiFiSecurityType = "wpa2Enterprise"
	AndroidWorkProfileEnterpriseWiFiConfigurationWiFiSecurityType_WpaEnterprise  AndroidWorkProfileEnterpriseWiFiConfigurationWiFiSecurityType = "wpaEnterprise"
)

func PossibleValuesForAndroidWorkProfileEnterpriseWiFiConfigurationWiFiSecurityType() []string {
	return []string{
		string(AndroidWorkProfileEnterpriseWiFiConfigurationWiFiSecurityType_Open),
		string(AndroidWorkProfileEnterpriseWiFiConfigurationWiFiSecurityType_Wpa2Enterprise),
		string(AndroidWorkProfileEnterpriseWiFiConfigurationWiFiSecurityType_WpaEnterprise),
	}
}

func (s *AndroidWorkProfileEnterpriseWiFiConfigurationWiFiSecurityType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAndroidWorkProfileEnterpriseWiFiConfigurationWiFiSecurityType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAndroidWorkProfileEnterpriseWiFiConfigurationWiFiSecurityType(input string) (*AndroidWorkProfileEnterpriseWiFiConfigurationWiFiSecurityType, error) {
	vals := map[string]AndroidWorkProfileEnterpriseWiFiConfigurationWiFiSecurityType{
		"open":           AndroidWorkProfileEnterpriseWiFiConfigurationWiFiSecurityType_Open,
		"wpa2enterprise": AndroidWorkProfileEnterpriseWiFiConfigurationWiFiSecurityType_Wpa2Enterprise,
		"wpaenterprise":  AndroidWorkProfileEnterpriseWiFiConfigurationWiFiSecurityType_WpaEnterprise,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AndroidWorkProfileEnterpriseWiFiConfigurationWiFiSecurityType(input)
	return &out, nil
}
