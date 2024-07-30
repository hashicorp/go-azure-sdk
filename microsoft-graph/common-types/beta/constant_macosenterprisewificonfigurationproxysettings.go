package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MacOSEnterpriseWiFiConfigurationProxySettings string

const (
	MacOSEnterpriseWiFiConfigurationProxySettings_Automatic MacOSEnterpriseWiFiConfigurationProxySettings = "automatic"
	MacOSEnterpriseWiFiConfigurationProxySettings_Manual    MacOSEnterpriseWiFiConfigurationProxySettings = "manual"
	MacOSEnterpriseWiFiConfigurationProxySettings_None      MacOSEnterpriseWiFiConfigurationProxySettings = "none"
)

func PossibleValuesForMacOSEnterpriseWiFiConfigurationProxySettings() []string {
	return []string{
		string(MacOSEnterpriseWiFiConfigurationProxySettings_Automatic),
		string(MacOSEnterpriseWiFiConfigurationProxySettings_Manual),
		string(MacOSEnterpriseWiFiConfigurationProxySettings_None),
	}
}

func (s *MacOSEnterpriseWiFiConfigurationProxySettings) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseMacOSEnterpriseWiFiConfigurationProxySettings(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseMacOSEnterpriseWiFiConfigurationProxySettings(input string) (*MacOSEnterpriseWiFiConfigurationProxySettings, error) {
	vals := map[string]MacOSEnterpriseWiFiConfigurationProxySettings{
		"automatic": MacOSEnterpriseWiFiConfigurationProxySettings_Automatic,
		"manual":    MacOSEnterpriseWiFiConfigurationProxySettings_Manual,
		"none":      MacOSEnterpriseWiFiConfigurationProxySettings_None,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := MacOSEnterpriseWiFiConfigurationProxySettings(input)
	return &out, nil
}
