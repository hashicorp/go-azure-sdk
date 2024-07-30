package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidEnterpriseWiFiConfigurationEapType string

const (
	AndroidEnterpriseWiFiConfigurationEapType_EapTls  AndroidEnterpriseWiFiConfigurationEapType = "eapTls"
	AndroidEnterpriseWiFiConfigurationEapType_EapTtls AndroidEnterpriseWiFiConfigurationEapType = "eapTtls"
	AndroidEnterpriseWiFiConfigurationEapType_Peap    AndroidEnterpriseWiFiConfigurationEapType = "peap"
)

func PossibleValuesForAndroidEnterpriseWiFiConfigurationEapType() []string {
	return []string{
		string(AndroidEnterpriseWiFiConfigurationEapType_EapTls),
		string(AndroidEnterpriseWiFiConfigurationEapType_EapTtls),
		string(AndroidEnterpriseWiFiConfigurationEapType_Peap),
	}
}

func (s *AndroidEnterpriseWiFiConfigurationEapType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAndroidEnterpriseWiFiConfigurationEapType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAndroidEnterpriseWiFiConfigurationEapType(input string) (*AndroidEnterpriseWiFiConfigurationEapType, error) {
	vals := map[string]AndroidEnterpriseWiFiConfigurationEapType{
		"eaptls":  AndroidEnterpriseWiFiConfigurationEapType_EapTls,
		"eapttls": AndroidEnterpriseWiFiConfigurationEapType_EapTtls,
		"peap":    AndroidEnterpriseWiFiConfigurationEapType_Peap,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AndroidEnterpriseWiFiConfigurationEapType(input)
	return &out, nil
}
