package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidDeviceOwnerWiFiConfigurationWiFiSecurityType string

const (
	AndroidDeviceOwnerWiFiConfigurationWiFiSecurityType_Open          AndroidDeviceOwnerWiFiConfigurationWiFiSecurityType = "open"
	AndroidDeviceOwnerWiFiConfigurationWiFiSecurityType_Wep           AndroidDeviceOwnerWiFiConfigurationWiFiSecurityType = "wep"
	AndroidDeviceOwnerWiFiConfigurationWiFiSecurityType_WpaEnterprise AndroidDeviceOwnerWiFiConfigurationWiFiSecurityType = "wpaEnterprise"
	AndroidDeviceOwnerWiFiConfigurationWiFiSecurityType_WpaPersonal   AndroidDeviceOwnerWiFiConfigurationWiFiSecurityType = "wpaPersonal"
)

func PossibleValuesForAndroidDeviceOwnerWiFiConfigurationWiFiSecurityType() []string {
	return []string{
		string(AndroidDeviceOwnerWiFiConfigurationWiFiSecurityType_Open),
		string(AndroidDeviceOwnerWiFiConfigurationWiFiSecurityType_Wep),
		string(AndroidDeviceOwnerWiFiConfigurationWiFiSecurityType_WpaEnterprise),
		string(AndroidDeviceOwnerWiFiConfigurationWiFiSecurityType_WpaPersonal),
	}
}

func (s *AndroidDeviceOwnerWiFiConfigurationWiFiSecurityType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAndroidDeviceOwnerWiFiConfigurationWiFiSecurityType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAndroidDeviceOwnerWiFiConfigurationWiFiSecurityType(input string) (*AndroidDeviceOwnerWiFiConfigurationWiFiSecurityType, error) {
	vals := map[string]AndroidDeviceOwnerWiFiConfigurationWiFiSecurityType{
		"open":          AndroidDeviceOwnerWiFiConfigurationWiFiSecurityType_Open,
		"wep":           AndroidDeviceOwnerWiFiConfigurationWiFiSecurityType_Wep,
		"wpaenterprise": AndroidDeviceOwnerWiFiConfigurationWiFiSecurityType_WpaEnterprise,
		"wpapersonal":   AndroidDeviceOwnerWiFiConfigurationWiFiSecurityType_WpaPersonal,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AndroidDeviceOwnerWiFiConfigurationWiFiSecurityType(input)
	return &out, nil
}
