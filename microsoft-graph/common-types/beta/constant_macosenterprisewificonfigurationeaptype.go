package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MacOSEnterpriseWiFiConfigurationEapType string

const (
	MacOSEnterpriseWiFiConfigurationEapType_EapFast MacOSEnterpriseWiFiConfigurationEapType = "eapFast"
	MacOSEnterpriseWiFiConfigurationEapType_EapSim  MacOSEnterpriseWiFiConfigurationEapType = "eapSim"
	MacOSEnterpriseWiFiConfigurationEapType_EapTls  MacOSEnterpriseWiFiConfigurationEapType = "eapTls"
	MacOSEnterpriseWiFiConfigurationEapType_EapTtls MacOSEnterpriseWiFiConfigurationEapType = "eapTtls"
	MacOSEnterpriseWiFiConfigurationEapType_Leap    MacOSEnterpriseWiFiConfigurationEapType = "leap"
	MacOSEnterpriseWiFiConfigurationEapType_Peap    MacOSEnterpriseWiFiConfigurationEapType = "peap"
	MacOSEnterpriseWiFiConfigurationEapType_Teap    MacOSEnterpriseWiFiConfigurationEapType = "teap"
)

func PossibleValuesForMacOSEnterpriseWiFiConfigurationEapType() []string {
	return []string{
		string(MacOSEnterpriseWiFiConfigurationEapType_EapFast),
		string(MacOSEnterpriseWiFiConfigurationEapType_EapSim),
		string(MacOSEnterpriseWiFiConfigurationEapType_EapTls),
		string(MacOSEnterpriseWiFiConfigurationEapType_EapTtls),
		string(MacOSEnterpriseWiFiConfigurationEapType_Leap),
		string(MacOSEnterpriseWiFiConfigurationEapType_Peap),
		string(MacOSEnterpriseWiFiConfigurationEapType_Teap),
	}
}

func (s *MacOSEnterpriseWiFiConfigurationEapType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseMacOSEnterpriseWiFiConfigurationEapType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseMacOSEnterpriseWiFiConfigurationEapType(input string) (*MacOSEnterpriseWiFiConfigurationEapType, error) {
	vals := map[string]MacOSEnterpriseWiFiConfigurationEapType{
		"eapfast": MacOSEnterpriseWiFiConfigurationEapType_EapFast,
		"eapsim":  MacOSEnterpriseWiFiConfigurationEapType_EapSim,
		"eaptls":  MacOSEnterpriseWiFiConfigurationEapType_EapTls,
		"eapttls": MacOSEnterpriseWiFiConfigurationEapType_EapTtls,
		"leap":    MacOSEnterpriseWiFiConfigurationEapType_Leap,
		"peap":    MacOSEnterpriseWiFiConfigurationEapType_Peap,
		"teap":    MacOSEnterpriseWiFiConfigurationEapType_Teap,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := MacOSEnterpriseWiFiConfigurationEapType(input)
	return &out, nil
}
