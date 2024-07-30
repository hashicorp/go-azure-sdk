package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MacOSWiredNetworkConfigurationEapType string

const (
	MacOSWiredNetworkConfigurationEapType_EapFast MacOSWiredNetworkConfigurationEapType = "eapFast"
	MacOSWiredNetworkConfigurationEapType_EapSim  MacOSWiredNetworkConfigurationEapType = "eapSim"
	MacOSWiredNetworkConfigurationEapType_EapTls  MacOSWiredNetworkConfigurationEapType = "eapTls"
	MacOSWiredNetworkConfigurationEapType_EapTtls MacOSWiredNetworkConfigurationEapType = "eapTtls"
	MacOSWiredNetworkConfigurationEapType_Leap    MacOSWiredNetworkConfigurationEapType = "leap"
	MacOSWiredNetworkConfigurationEapType_Peap    MacOSWiredNetworkConfigurationEapType = "peap"
	MacOSWiredNetworkConfigurationEapType_Teap    MacOSWiredNetworkConfigurationEapType = "teap"
)

func PossibleValuesForMacOSWiredNetworkConfigurationEapType() []string {
	return []string{
		string(MacOSWiredNetworkConfigurationEapType_EapFast),
		string(MacOSWiredNetworkConfigurationEapType_EapSim),
		string(MacOSWiredNetworkConfigurationEapType_EapTls),
		string(MacOSWiredNetworkConfigurationEapType_EapTtls),
		string(MacOSWiredNetworkConfigurationEapType_Leap),
		string(MacOSWiredNetworkConfigurationEapType_Peap),
		string(MacOSWiredNetworkConfigurationEapType_Teap),
	}
}

func (s *MacOSWiredNetworkConfigurationEapType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseMacOSWiredNetworkConfigurationEapType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseMacOSWiredNetworkConfigurationEapType(input string) (*MacOSWiredNetworkConfigurationEapType, error) {
	vals := map[string]MacOSWiredNetworkConfigurationEapType{
		"eapfast": MacOSWiredNetworkConfigurationEapType_EapFast,
		"eapsim":  MacOSWiredNetworkConfigurationEapType_EapSim,
		"eaptls":  MacOSWiredNetworkConfigurationEapType_EapTls,
		"eapttls": MacOSWiredNetworkConfigurationEapType_EapTtls,
		"leap":    MacOSWiredNetworkConfigurationEapType_Leap,
		"peap":    MacOSWiredNetworkConfigurationEapType_Peap,
		"teap":    MacOSWiredNetworkConfigurationEapType_Teap,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := MacOSWiredNetworkConfigurationEapType(input)
	return &out, nil
}
