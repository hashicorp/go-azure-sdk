package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidForWorkWiFiConfigurationWiFiSecurityType string

const (
	AndroidForWorkWiFiConfigurationWiFiSecurityType_Open           AndroidForWorkWiFiConfigurationWiFiSecurityType = "open"
	AndroidForWorkWiFiConfigurationWiFiSecurityType_Wpa2Enterprise AndroidForWorkWiFiConfigurationWiFiSecurityType = "wpa2Enterprise"
	AndroidForWorkWiFiConfigurationWiFiSecurityType_WpaEnterprise  AndroidForWorkWiFiConfigurationWiFiSecurityType = "wpaEnterprise"
)

func PossibleValuesForAndroidForWorkWiFiConfigurationWiFiSecurityType() []string {
	return []string{
		string(AndroidForWorkWiFiConfigurationWiFiSecurityType_Open),
		string(AndroidForWorkWiFiConfigurationWiFiSecurityType_Wpa2Enterprise),
		string(AndroidForWorkWiFiConfigurationWiFiSecurityType_WpaEnterprise),
	}
}

func (s *AndroidForWorkWiFiConfigurationWiFiSecurityType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAndroidForWorkWiFiConfigurationWiFiSecurityType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAndroidForWorkWiFiConfigurationWiFiSecurityType(input string) (*AndroidForWorkWiFiConfigurationWiFiSecurityType, error) {
	vals := map[string]AndroidForWorkWiFiConfigurationWiFiSecurityType{
		"open":           AndroidForWorkWiFiConfigurationWiFiSecurityType_Open,
		"wpa2enterprise": AndroidForWorkWiFiConfigurationWiFiSecurityType_Wpa2Enterprise,
		"wpaenterprise":  AndroidForWorkWiFiConfigurationWiFiSecurityType_WpaEnterprise,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AndroidForWorkWiFiConfigurationWiFiSecurityType(input)
	return &out, nil
}
