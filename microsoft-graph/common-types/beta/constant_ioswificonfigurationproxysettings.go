package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IosWiFiConfigurationProxySettings string

const (
	IosWiFiConfigurationProxySettings_Automatic IosWiFiConfigurationProxySettings = "automatic"
	IosWiFiConfigurationProxySettings_Manual    IosWiFiConfigurationProxySettings = "manual"
	IosWiFiConfigurationProxySettings_None      IosWiFiConfigurationProxySettings = "none"
)

func PossibleValuesForIosWiFiConfigurationProxySettings() []string {
	return []string{
		string(IosWiFiConfigurationProxySettings_Automatic),
		string(IosWiFiConfigurationProxySettings_Manual),
		string(IosWiFiConfigurationProxySettings_None),
	}
}

func (s *IosWiFiConfigurationProxySettings) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseIosWiFiConfigurationProxySettings(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseIosWiFiConfigurationProxySettings(input string) (*IosWiFiConfigurationProxySettings, error) {
	vals := map[string]IosWiFiConfigurationProxySettings{
		"automatic": IosWiFiConfigurationProxySettings_Automatic,
		"manual":    IosWiFiConfigurationProxySettings_Manual,
		"none":      IosWiFiConfigurationProxySettings_None,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := IosWiFiConfigurationProxySettings(input)
	return &out, nil
}
