package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AospDeviceOwnerWiFiConfigurationWiFiSecurityType string

const (
	AospDeviceOwnerWiFiConfigurationWiFiSecurityType_Open          AospDeviceOwnerWiFiConfigurationWiFiSecurityType = "open"
	AospDeviceOwnerWiFiConfigurationWiFiSecurityType_Wep           AospDeviceOwnerWiFiConfigurationWiFiSecurityType = "wep"
	AospDeviceOwnerWiFiConfigurationWiFiSecurityType_WpaEnterprise AospDeviceOwnerWiFiConfigurationWiFiSecurityType = "wpaEnterprise"
	AospDeviceOwnerWiFiConfigurationWiFiSecurityType_WpaPersonal   AospDeviceOwnerWiFiConfigurationWiFiSecurityType = "wpaPersonal"
)

func PossibleValuesForAospDeviceOwnerWiFiConfigurationWiFiSecurityType() []string {
	return []string{
		string(AospDeviceOwnerWiFiConfigurationWiFiSecurityType_Open),
		string(AospDeviceOwnerWiFiConfigurationWiFiSecurityType_Wep),
		string(AospDeviceOwnerWiFiConfigurationWiFiSecurityType_WpaEnterprise),
		string(AospDeviceOwnerWiFiConfigurationWiFiSecurityType_WpaPersonal),
	}
}

func (s *AospDeviceOwnerWiFiConfigurationWiFiSecurityType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAospDeviceOwnerWiFiConfigurationWiFiSecurityType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAospDeviceOwnerWiFiConfigurationWiFiSecurityType(input string) (*AospDeviceOwnerWiFiConfigurationWiFiSecurityType, error) {
	vals := map[string]AospDeviceOwnerWiFiConfigurationWiFiSecurityType{
		"open":          AospDeviceOwnerWiFiConfigurationWiFiSecurityType_Open,
		"wep":           AospDeviceOwnerWiFiConfigurationWiFiSecurityType_Wep,
		"wpaenterprise": AospDeviceOwnerWiFiConfigurationWiFiSecurityType_WpaEnterprise,
		"wpapersonal":   AospDeviceOwnerWiFiConfigurationWiFiSecurityType_WpaPersonal,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AospDeviceOwnerWiFiConfigurationWiFiSecurityType(input)
	return &out, nil
}
