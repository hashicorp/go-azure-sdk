package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidWorkProfileWiFiConfigurationWiFiSecurityType string

const (
	AndroidWorkProfileWiFiConfigurationWiFiSecurityType_Open           AndroidWorkProfileWiFiConfigurationWiFiSecurityType = "open"
	AndroidWorkProfileWiFiConfigurationWiFiSecurityType_Wpa2Enterprise AndroidWorkProfileWiFiConfigurationWiFiSecurityType = "wpa2Enterprise"
	AndroidWorkProfileWiFiConfigurationWiFiSecurityType_WpaEnterprise  AndroidWorkProfileWiFiConfigurationWiFiSecurityType = "wpaEnterprise"
)

func PossibleValuesForAndroidWorkProfileWiFiConfigurationWiFiSecurityType() []string {
	return []string{
		string(AndroidWorkProfileWiFiConfigurationWiFiSecurityType_Open),
		string(AndroidWorkProfileWiFiConfigurationWiFiSecurityType_Wpa2Enterprise),
		string(AndroidWorkProfileWiFiConfigurationWiFiSecurityType_WpaEnterprise),
	}
}

func (s *AndroidWorkProfileWiFiConfigurationWiFiSecurityType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAndroidWorkProfileWiFiConfigurationWiFiSecurityType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAndroidWorkProfileWiFiConfigurationWiFiSecurityType(input string) (*AndroidWorkProfileWiFiConfigurationWiFiSecurityType, error) {
	vals := map[string]AndroidWorkProfileWiFiConfigurationWiFiSecurityType{
		"open":           AndroidWorkProfileWiFiConfigurationWiFiSecurityType_Open,
		"wpa2enterprise": AndroidWorkProfileWiFiConfigurationWiFiSecurityType_Wpa2Enterprise,
		"wpaenterprise":  AndroidWorkProfileWiFiConfigurationWiFiSecurityType_WpaEnterprise,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AndroidWorkProfileWiFiConfigurationWiFiSecurityType(input)
	return &out, nil
}
