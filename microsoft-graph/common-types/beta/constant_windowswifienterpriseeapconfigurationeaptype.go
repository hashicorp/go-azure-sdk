package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WindowsWifiEnterpriseEAPConfigurationEapType string

const (
	WindowsWifiEnterpriseEAPConfigurationEapType_EapFast WindowsWifiEnterpriseEAPConfigurationEapType = "eapFast"
	WindowsWifiEnterpriseEAPConfigurationEapType_EapSim  WindowsWifiEnterpriseEAPConfigurationEapType = "eapSim"
	WindowsWifiEnterpriseEAPConfigurationEapType_EapTls  WindowsWifiEnterpriseEAPConfigurationEapType = "eapTls"
	WindowsWifiEnterpriseEAPConfigurationEapType_EapTtls WindowsWifiEnterpriseEAPConfigurationEapType = "eapTtls"
	WindowsWifiEnterpriseEAPConfigurationEapType_Leap    WindowsWifiEnterpriseEAPConfigurationEapType = "leap"
	WindowsWifiEnterpriseEAPConfigurationEapType_Peap    WindowsWifiEnterpriseEAPConfigurationEapType = "peap"
	WindowsWifiEnterpriseEAPConfigurationEapType_Teap    WindowsWifiEnterpriseEAPConfigurationEapType = "teap"
)

func PossibleValuesForWindowsWifiEnterpriseEAPConfigurationEapType() []string {
	return []string{
		string(WindowsWifiEnterpriseEAPConfigurationEapType_EapFast),
		string(WindowsWifiEnterpriseEAPConfigurationEapType_EapSim),
		string(WindowsWifiEnterpriseEAPConfigurationEapType_EapTls),
		string(WindowsWifiEnterpriseEAPConfigurationEapType_EapTtls),
		string(WindowsWifiEnterpriseEAPConfigurationEapType_Leap),
		string(WindowsWifiEnterpriseEAPConfigurationEapType_Peap),
		string(WindowsWifiEnterpriseEAPConfigurationEapType_Teap),
	}
}

func (s *WindowsWifiEnterpriseEAPConfigurationEapType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseWindowsWifiEnterpriseEAPConfigurationEapType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseWindowsWifiEnterpriseEAPConfigurationEapType(input string) (*WindowsWifiEnterpriseEAPConfigurationEapType, error) {
	vals := map[string]WindowsWifiEnterpriseEAPConfigurationEapType{
		"eapfast": WindowsWifiEnterpriseEAPConfigurationEapType_EapFast,
		"eapsim":  WindowsWifiEnterpriseEAPConfigurationEapType_EapSim,
		"eaptls":  WindowsWifiEnterpriseEAPConfigurationEapType_EapTls,
		"eapttls": WindowsWifiEnterpriseEAPConfigurationEapType_EapTtls,
		"leap":    WindowsWifiEnterpriseEAPConfigurationEapType_Leap,
		"peap":    WindowsWifiEnterpriseEAPConfigurationEapType_Peap,
		"teap":    WindowsWifiEnterpriseEAPConfigurationEapType_Teap,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := WindowsWifiEnterpriseEAPConfigurationEapType(input)
	return &out, nil
}
