package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MacOSWiFiConfigurationProxySettings string

const (
	MacOSWiFiConfigurationProxySettings_Automatic MacOSWiFiConfigurationProxySettings = "automatic"
	MacOSWiFiConfigurationProxySettings_Manual    MacOSWiFiConfigurationProxySettings = "manual"
	MacOSWiFiConfigurationProxySettings_None      MacOSWiFiConfigurationProxySettings = "none"
)

func PossibleValuesForMacOSWiFiConfigurationProxySettings() []string {
	return []string{
		string(MacOSWiFiConfigurationProxySettings_Automatic),
		string(MacOSWiFiConfigurationProxySettings_Manual),
		string(MacOSWiFiConfigurationProxySettings_None),
	}
}

func (s *MacOSWiFiConfigurationProxySettings) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseMacOSWiFiConfigurationProxySettings(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseMacOSWiFiConfigurationProxySettings(input string) (*MacOSWiFiConfigurationProxySettings, error) {
	vals := map[string]MacOSWiFiConfigurationProxySettings{
		"automatic": MacOSWiFiConfigurationProxySettings_Automatic,
		"manual":    MacOSWiFiConfigurationProxySettings_Manual,
		"none":      MacOSWiFiConfigurationProxySettings_None,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := MacOSWiFiConfigurationProxySettings(input)
	return &out, nil
}
