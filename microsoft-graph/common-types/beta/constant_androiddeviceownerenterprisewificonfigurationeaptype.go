package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidDeviceOwnerEnterpriseWiFiConfigurationEapType string

const (
	AndroidDeviceOwnerEnterpriseWiFiConfigurationEapType_EapTls  AndroidDeviceOwnerEnterpriseWiFiConfigurationEapType = "eapTls"
	AndroidDeviceOwnerEnterpriseWiFiConfigurationEapType_EapTtls AndroidDeviceOwnerEnterpriseWiFiConfigurationEapType = "eapTtls"
	AndroidDeviceOwnerEnterpriseWiFiConfigurationEapType_Peap    AndroidDeviceOwnerEnterpriseWiFiConfigurationEapType = "peap"
)

func PossibleValuesForAndroidDeviceOwnerEnterpriseWiFiConfigurationEapType() []string {
	return []string{
		string(AndroidDeviceOwnerEnterpriseWiFiConfigurationEapType_EapTls),
		string(AndroidDeviceOwnerEnterpriseWiFiConfigurationEapType_EapTtls),
		string(AndroidDeviceOwnerEnterpriseWiFiConfigurationEapType_Peap),
	}
}

func (s *AndroidDeviceOwnerEnterpriseWiFiConfigurationEapType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAndroidDeviceOwnerEnterpriseWiFiConfigurationEapType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAndroidDeviceOwnerEnterpriseWiFiConfigurationEapType(input string) (*AndroidDeviceOwnerEnterpriseWiFiConfigurationEapType, error) {
	vals := map[string]AndroidDeviceOwnerEnterpriseWiFiConfigurationEapType{
		"eaptls":  AndroidDeviceOwnerEnterpriseWiFiConfigurationEapType_EapTls,
		"eapttls": AndroidDeviceOwnerEnterpriseWiFiConfigurationEapType_EapTtls,
		"peap":    AndroidDeviceOwnerEnterpriseWiFiConfigurationEapType_Peap,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AndroidDeviceOwnerEnterpriseWiFiConfigurationEapType(input)
	return &out, nil
}
