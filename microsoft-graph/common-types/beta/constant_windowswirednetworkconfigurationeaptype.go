package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WindowsWiredNetworkConfigurationEapType string

const (
	WindowsWiredNetworkConfigurationEapType_EapFast WindowsWiredNetworkConfigurationEapType = "eapFast"
	WindowsWiredNetworkConfigurationEapType_EapSim  WindowsWiredNetworkConfigurationEapType = "eapSim"
	WindowsWiredNetworkConfigurationEapType_EapTls  WindowsWiredNetworkConfigurationEapType = "eapTls"
	WindowsWiredNetworkConfigurationEapType_EapTtls WindowsWiredNetworkConfigurationEapType = "eapTtls"
	WindowsWiredNetworkConfigurationEapType_Leap    WindowsWiredNetworkConfigurationEapType = "leap"
	WindowsWiredNetworkConfigurationEapType_Peap    WindowsWiredNetworkConfigurationEapType = "peap"
	WindowsWiredNetworkConfigurationEapType_Teap    WindowsWiredNetworkConfigurationEapType = "teap"
)

func PossibleValuesForWindowsWiredNetworkConfigurationEapType() []string {
	return []string{
		string(WindowsWiredNetworkConfigurationEapType_EapFast),
		string(WindowsWiredNetworkConfigurationEapType_EapSim),
		string(WindowsWiredNetworkConfigurationEapType_EapTls),
		string(WindowsWiredNetworkConfigurationEapType_EapTtls),
		string(WindowsWiredNetworkConfigurationEapType_Leap),
		string(WindowsWiredNetworkConfigurationEapType_Peap),
		string(WindowsWiredNetworkConfigurationEapType_Teap),
	}
}

func (s *WindowsWiredNetworkConfigurationEapType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseWindowsWiredNetworkConfigurationEapType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseWindowsWiredNetworkConfigurationEapType(input string) (*WindowsWiredNetworkConfigurationEapType, error) {
	vals := map[string]WindowsWiredNetworkConfigurationEapType{
		"eapfast": WindowsWiredNetworkConfigurationEapType_EapFast,
		"eapsim":  WindowsWiredNetworkConfigurationEapType_EapSim,
		"eaptls":  WindowsWiredNetworkConfigurationEapType_EapTls,
		"eapttls": WindowsWiredNetworkConfigurationEapType_EapTtls,
		"leap":    WindowsWiredNetworkConfigurationEapType_Leap,
		"peap":    WindowsWiredNetworkConfigurationEapType_Peap,
		"teap":    WindowsWiredNetworkConfigurationEapType_Teap,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := WindowsWiredNetworkConfigurationEapType(input)
	return &out, nil
}
