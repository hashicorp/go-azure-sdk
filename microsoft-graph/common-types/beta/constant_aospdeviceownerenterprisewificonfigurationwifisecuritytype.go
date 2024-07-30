package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AospDeviceOwnerEnterpriseWiFiConfigurationWiFiSecurityType string

const (
	AospDeviceOwnerEnterpriseWiFiConfigurationWiFiSecurityType_Open          AospDeviceOwnerEnterpriseWiFiConfigurationWiFiSecurityType = "open"
	AospDeviceOwnerEnterpriseWiFiConfigurationWiFiSecurityType_Wep           AospDeviceOwnerEnterpriseWiFiConfigurationWiFiSecurityType = "wep"
	AospDeviceOwnerEnterpriseWiFiConfigurationWiFiSecurityType_WpaEnterprise AospDeviceOwnerEnterpriseWiFiConfigurationWiFiSecurityType = "wpaEnterprise"
	AospDeviceOwnerEnterpriseWiFiConfigurationWiFiSecurityType_WpaPersonal   AospDeviceOwnerEnterpriseWiFiConfigurationWiFiSecurityType = "wpaPersonal"
)

func PossibleValuesForAospDeviceOwnerEnterpriseWiFiConfigurationWiFiSecurityType() []string {
	return []string{
		string(AospDeviceOwnerEnterpriseWiFiConfigurationWiFiSecurityType_Open),
		string(AospDeviceOwnerEnterpriseWiFiConfigurationWiFiSecurityType_Wep),
		string(AospDeviceOwnerEnterpriseWiFiConfigurationWiFiSecurityType_WpaEnterprise),
		string(AospDeviceOwnerEnterpriseWiFiConfigurationWiFiSecurityType_WpaPersonal),
	}
}

func (s *AospDeviceOwnerEnterpriseWiFiConfigurationWiFiSecurityType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAospDeviceOwnerEnterpriseWiFiConfigurationWiFiSecurityType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAospDeviceOwnerEnterpriseWiFiConfigurationWiFiSecurityType(input string) (*AospDeviceOwnerEnterpriseWiFiConfigurationWiFiSecurityType, error) {
	vals := map[string]AospDeviceOwnerEnterpriseWiFiConfigurationWiFiSecurityType{
		"open":          AospDeviceOwnerEnterpriseWiFiConfigurationWiFiSecurityType_Open,
		"wep":           AospDeviceOwnerEnterpriseWiFiConfigurationWiFiSecurityType_Wep,
		"wpaenterprise": AospDeviceOwnerEnterpriseWiFiConfigurationWiFiSecurityType_WpaEnterprise,
		"wpapersonal":   AospDeviceOwnerEnterpriseWiFiConfigurationWiFiSecurityType_WpaPersonal,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AospDeviceOwnerEnterpriseWiFiConfigurationWiFiSecurityType(input)
	return &out, nil
}
