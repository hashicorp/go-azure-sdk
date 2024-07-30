package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidDeviceOwnerEnterpriseWiFiConfigurationWiFiSecurityType string

const (
	AndroidDeviceOwnerEnterpriseWiFiConfigurationWiFiSecurityType_Open          AndroidDeviceOwnerEnterpriseWiFiConfigurationWiFiSecurityType = "open"
	AndroidDeviceOwnerEnterpriseWiFiConfigurationWiFiSecurityType_Wep           AndroidDeviceOwnerEnterpriseWiFiConfigurationWiFiSecurityType = "wep"
	AndroidDeviceOwnerEnterpriseWiFiConfigurationWiFiSecurityType_WpaEnterprise AndroidDeviceOwnerEnterpriseWiFiConfigurationWiFiSecurityType = "wpaEnterprise"
	AndroidDeviceOwnerEnterpriseWiFiConfigurationWiFiSecurityType_WpaPersonal   AndroidDeviceOwnerEnterpriseWiFiConfigurationWiFiSecurityType = "wpaPersonal"
)

func PossibleValuesForAndroidDeviceOwnerEnterpriseWiFiConfigurationWiFiSecurityType() []string {
	return []string{
		string(AndroidDeviceOwnerEnterpriseWiFiConfigurationWiFiSecurityType_Open),
		string(AndroidDeviceOwnerEnterpriseWiFiConfigurationWiFiSecurityType_Wep),
		string(AndroidDeviceOwnerEnterpriseWiFiConfigurationWiFiSecurityType_WpaEnterprise),
		string(AndroidDeviceOwnerEnterpriseWiFiConfigurationWiFiSecurityType_WpaPersonal),
	}
}

func (s *AndroidDeviceOwnerEnterpriseWiFiConfigurationWiFiSecurityType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAndroidDeviceOwnerEnterpriseWiFiConfigurationWiFiSecurityType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAndroidDeviceOwnerEnterpriseWiFiConfigurationWiFiSecurityType(input string) (*AndroidDeviceOwnerEnterpriseWiFiConfigurationWiFiSecurityType, error) {
	vals := map[string]AndroidDeviceOwnerEnterpriseWiFiConfigurationWiFiSecurityType{
		"open":          AndroidDeviceOwnerEnterpriseWiFiConfigurationWiFiSecurityType_Open,
		"wep":           AndroidDeviceOwnerEnterpriseWiFiConfigurationWiFiSecurityType_Wep,
		"wpaenterprise": AndroidDeviceOwnerEnterpriseWiFiConfigurationWiFiSecurityType_WpaEnterprise,
		"wpapersonal":   AndroidDeviceOwnerEnterpriseWiFiConfigurationWiFiSecurityType_WpaPersonal,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AndroidDeviceOwnerEnterpriseWiFiConfigurationWiFiSecurityType(input)
	return &out, nil
}
