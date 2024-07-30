package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidWorkProfileEnterpriseWiFiConfigurationProxySettings string

const (
	AndroidWorkProfileEnterpriseWiFiConfigurationProxySettings_Automatic AndroidWorkProfileEnterpriseWiFiConfigurationProxySettings = "automatic"
	AndroidWorkProfileEnterpriseWiFiConfigurationProxySettings_Manual    AndroidWorkProfileEnterpriseWiFiConfigurationProxySettings = "manual"
	AndroidWorkProfileEnterpriseWiFiConfigurationProxySettings_None      AndroidWorkProfileEnterpriseWiFiConfigurationProxySettings = "none"
)

func PossibleValuesForAndroidWorkProfileEnterpriseWiFiConfigurationProxySettings() []string {
	return []string{
		string(AndroidWorkProfileEnterpriseWiFiConfigurationProxySettings_Automatic),
		string(AndroidWorkProfileEnterpriseWiFiConfigurationProxySettings_Manual),
		string(AndroidWorkProfileEnterpriseWiFiConfigurationProxySettings_None),
	}
}

func (s *AndroidWorkProfileEnterpriseWiFiConfigurationProxySettings) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAndroidWorkProfileEnterpriseWiFiConfigurationProxySettings(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAndroidWorkProfileEnterpriseWiFiConfigurationProxySettings(input string) (*AndroidWorkProfileEnterpriseWiFiConfigurationProxySettings, error) {
	vals := map[string]AndroidWorkProfileEnterpriseWiFiConfigurationProxySettings{
		"automatic": AndroidWorkProfileEnterpriseWiFiConfigurationProxySettings_Automatic,
		"manual":    AndroidWorkProfileEnterpriseWiFiConfigurationProxySettings_Manual,
		"none":      AndroidWorkProfileEnterpriseWiFiConfigurationProxySettings_None,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AndroidWorkProfileEnterpriseWiFiConfigurationProxySettings(input)
	return &out, nil
}
