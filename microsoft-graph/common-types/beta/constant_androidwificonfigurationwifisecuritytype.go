package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidWiFiConfigurationWiFiSecurityType string

const (
	AndroidWiFiConfigurationWiFiSecurityType_Open           AndroidWiFiConfigurationWiFiSecurityType = "open"
	AndroidWiFiConfigurationWiFiSecurityType_Wpa2Enterprise AndroidWiFiConfigurationWiFiSecurityType = "wpa2Enterprise"
	AndroidWiFiConfigurationWiFiSecurityType_WpaEnterprise  AndroidWiFiConfigurationWiFiSecurityType = "wpaEnterprise"
)

func PossibleValuesForAndroidWiFiConfigurationWiFiSecurityType() []string {
	return []string{
		string(AndroidWiFiConfigurationWiFiSecurityType_Open),
		string(AndroidWiFiConfigurationWiFiSecurityType_Wpa2Enterprise),
		string(AndroidWiFiConfigurationWiFiSecurityType_WpaEnterprise),
	}
}

func (s *AndroidWiFiConfigurationWiFiSecurityType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAndroidWiFiConfigurationWiFiSecurityType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAndroidWiFiConfigurationWiFiSecurityType(input string) (*AndroidWiFiConfigurationWiFiSecurityType, error) {
	vals := map[string]AndroidWiFiConfigurationWiFiSecurityType{
		"open":           AndroidWiFiConfigurationWiFiSecurityType_Open,
		"wpa2enterprise": AndroidWiFiConfigurationWiFiSecurityType_Wpa2Enterprise,
		"wpaenterprise":  AndroidWiFiConfigurationWiFiSecurityType_WpaEnterprise,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AndroidWiFiConfigurationWiFiSecurityType(input)
	return &out, nil
}
