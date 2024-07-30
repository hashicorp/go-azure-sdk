package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidWorkProfileEnterpriseWiFiConfigurationEapType string

const (
	AndroidWorkProfileEnterpriseWiFiConfigurationEapType_EapTls  AndroidWorkProfileEnterpriseWiFiConfigurationEapType = "eapTls"
	AndroidWorkProfileEnterpriseWiFiConfigurationEapType_EapTtls AndroidWorkProfileEnterpriseWiFiConfigurationEapType = "eapTtls"
	AndroidWorkProfileEnterpriseWiFiConfigurationEapType_Peap    AndroidWorkProfileEnterpriseWiFiConfigurationEapType = "peap"
)

func PossibleValuesForAndroidWorkProfileEnterpriseWiFiConfigurationEapType() []string {
	return []string{
		string(AndroidWorkProfileEnterpriseWiFiConfigurationEapType_EapTls),
		string(AndroidWorkProfileEnterpriseWiFiConfigurationEapType_EapTtls),
		string(AndroidWorkProfileEnterpriseWiFiConfigurationEapType_Peap),
	}
}

func (s *AndroidWorkProfileEnterpriseWiFiConfigurationEapType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAndroidWorkProfileEnterpriseWiFiConfigurationEapType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAndroidWorkProfileEnterpriseWiFiConfigurationEapType(input string) (*AndroidWorkProfileEnterpriseWiFiConfigurationEapType, error) {
	vals := map[string]AndroidWorkProfileEnterpriseWiFiConfigurationEapType{
		"eaptls":  AndroidWorkProfileEnterpriseWiFiConfigurationEapType_EapTls,
		"eapttls": AndroidWorkProfileEnterpriseWiFiConfigurationEapType_EapTtls,
		"peap":    AndroidWorkProfileEnterpriseWiFiConfigurationEapType_Peap,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AndroidWorkProfileEnterpriseWiFiConfigurationEapType(input)
	return &out, nil
}
