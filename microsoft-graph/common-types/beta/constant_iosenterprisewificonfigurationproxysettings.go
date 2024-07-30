package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IosEnterpriseWiFiConfigurationProxySettings string

const (
	IosEnterpriseWiFiConfigurationProxySettings_Automatic IosEnterpriseWiFiConfigurationProxySettings = "automatic"
	IosEnterpriseWiFiConfigurationProxySettings_Manual    IosEnterpriseWiFiConfigurationProxySettings = "manual"
	IosEnterpriseWiFiConfigurationProxySettings_None      IosEnterpriseWiFiConfigurationProxySettings = "none"
)

func PossibleValuesForIosEnterpriseWiFiConfigurationProxySettings() []string {
	return []string{
		string(IosEnterpriseWiFiConfigurationProxySettings_Automatic),
		string(IosEnterpriseWiFiConfigurationProxySettings_Manual),
		string(IosEnterpriseWiFiConfigurationProxySettings_None),
	}
}

func (s *IosEnterpriseWiFiConfigurationProxySettings) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseIosEnterpriseWiFiConfigurationProxySettings(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseIosEnterpriseWiFiConfigurationProxySettings(input string) (*IosEnterpriseWiFiConfigurationProxySettings, error) {
	vals := map[string]IosEnterpriseWiFiConfigurationProxySettings{
		"automatic": IosEnterpriseWiFiConfigurationProxySettings_Automatic,
		"manual":    IosEnterpriseWiFiConfigurationProxySettings_Manual,
		"none":      IosEnterpriseWiFiConfigurationProxySettings_None,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := IosEnterpriseWiFiConfigurationProxySettings(input)
	return &out, nil
}
