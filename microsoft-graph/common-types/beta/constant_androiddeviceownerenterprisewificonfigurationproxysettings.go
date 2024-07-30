package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidDeviceOwnerEnterpriseWiFiConfigurationProxySettings string

const (
	AndroidDeviceOwnerEnterpriseWiFiConfigurationProxySettings_Automatic AndroidDeviceOwnerEnterpriseWiFiConfigurationProxySettings = "automatic"
	AndroidDeviceOwnerEnterpriseWiFiConfigurationProxySettings_Manual    AndroidDeviceOwnerEnterpriseWiFiConfigurationProxySettings = "manual"
	AndroidDeviceOwnerEnterpriseWiFiConfigurationProxySettings_None      AndroidDeviceOwnerEnterpriseWiFiConfigurationProxySettings = "none"
)

func PossibleValuesForAndroidDeviceOwnerEnterpriseWiFiConfigurationProxySettings() []string {
	return []string{
		string(AndroidDeviceOwnerEnterpriseWiFiConfigurationProxySettings_Automatic),
		string(AndroidDeviceOwnerEnterpriseWiFiConfigurationProxySettings_Manual),
		string(AndroidDeviceOwnerEnterpriseWiFiConfigurationProxySettings_None),
	}
}

func (s *AndroidDeviceOwnerEnterpriseWiFiConfigurationProxySettings) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAndroidDeviceOwnerEnterpriseWiFiConfigurationProxySettings(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAndroidDeviceOwnerEnterpriseWiFiConfigurationProxySettings(input string) (*AndroidDeviceOwnerEnterpriseWiFiConfigurationProxySettings, error) {
	vals := map[string]AndroidDeviceOwnerEnterpriseWiFiConfigurationProxySettings{
		"automatic": AndroidDeviceOwnerEnterpriseWiFiConfigurationProxySettings_Automatic,
		"manual":    AndroidDeviceOwnerEnterpriseWiFiConfigurationProxySettings_Manual,
		"none":      AndroidDeviceOwnerEnterpriseWiFiConfigurationProxySettings_None,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AndroidDeviceOwnerEnterpriseWiFiConfigurationProxySettings(input)
	return &out, nil
}
