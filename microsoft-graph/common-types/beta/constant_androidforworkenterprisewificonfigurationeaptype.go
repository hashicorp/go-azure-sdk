package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidForWorkEnterpriseWiFiConfigurationEapType string

const (
	AndroidForWorkEnterpriseWiFiConfigurationEapType_EapTls  AndroidForWorkEnterpriseWiFiConfigurationEapType = "eapTls"
	AndroidForWorkEnterpriseWiFiConfigurationEapType_EapTtls AndroidForWorkEnterpriseWiFiConfigurationEapType = "eapTtls"
	AndroidForWorkEnterpriseWiFiConfigurationEapType_Peap    AndroidForWorkEnterpriseWiFiConfigurationEapType = "peap"
)

func PossibleValuesForAndroidForWorkEnterpriseWiFiConfigurationEapType() []string {
	return []string{
		string(AndroidForWorkEnterpriseWiFiConfigurationEapType_EapTls),
		string(AndroidForWorkEnterpriseWiFiConfigurationEapType_EapTtls),
		string(AndroidForWorkEnterpriseWiFiConfigurationEapType_Peap),
	}
}

func (s *AndroidForWorkEnterpriseWiFiConfigurationEapType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAndroidForWorkEnterpriseWiFiConfigurationEapType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAndroidForWorkEnterpriseWiFiConfigurationEapType(input string) (*AndroidForWorkEnterpriseWiFiConfigurationEapType, error) {
	vals := map[string]AndroidForWorkEnterpriseWiFiConfigurationEapType{
		"eaptls":  AndroidForWorkEnterpriseWiFiConfigurationEapType_EapTls,
		"eapttls": AndroidForWorkEnterpriseWiFiConfigurationEapType_EapTtls,
		"peap":    AndroidForWorkEnterpriseWiFiConfigurationEapType_Peap,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AndroidForWorkEnterpriseWiFiConfigurationEapType(input)
	return &out, nil
}
