package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AospDeviceOwnerEnterpriseWiFiConfigurationEapType string

const (
	AospDeviceOwnerEnterpriseWiFiConfigurationEapType_EapTls  AospDeviceOwnerEnterpriseWiFiConfigurationEapType = "eapTls"
	AospDeviceOwnerEnterpriseWiFiConfigurationEapType_EapTtls AospDeviceOwnerEnterpriseWiFiConfigurationEapType = "eapTtls"
	AospDeviceOwnerEnterpriseWiFiConfigurationEapType_Peap    AospDeviceOwnerEnterpriseWiFiConfigurationEapType = "peap"
)

func PossibleValuesForAospDeviceOwnerEnterpriseWiFiConfigurationEapType() []string {
	return []string{
		string(AospDeviceOwnerEnterpriseWiFiConfigurationEapType_EapTls),
		string(AospDeviceOwnerEnterpriseWiFiConfigurationEapType_EapTtls),
		string(AospDeviceOwnerEnterpriseWiFiConfigurationEapType_Peap),
	}
}

func (s *AospDeviceOwnerEnterpriseWiFiConfigurationEapType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAospDeviceOwnerEnterpriseWiFiConfigurationEapType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAospDeviceOwnerEnterpriseWiFiConfigurationEapType(input string) (*AospDeviceOwnerEnterpriseWiFiConfigurationEapType, error) {
	vals := map[string]AospDeviceOwnerEnterpriseWiFiConfigurationEapType{
		"eaptls":  AospDeviceOwnerEnterpriseWiFiConfigurationEapType_EapTls,
		"eapttls": AospDeviceOwnerEnterpriseWiFiConfigurationEapType_EapTtls,
		"peap":    AospDeviceOwnerEnterpriseWiFiConfigurationEapType_Peap,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AospDeviceOwnerEnterpriseWiFiConfigurationEapType(input)
	return &out, nil
}
