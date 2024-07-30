package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidDeviceOwnerWiFiConfigurationProxySettings string

const (
	AndroidDeviceOwnerWiFiConfigurationProxySettings_Automatic AndroidDeviceOwnerWiFiConfigurationProxySettings = "automatic"
	AndroidDeviceOwnerWiFiConfigurationProxySettings_Manual    AndroidDeviceOwnerWiFiConfigurationProxySettings = "manual"
	AndroidDeviceOwnerWiFiConfigurationProxySettings_None      AndroidDeviceOwnerWiFiConfigurationProxySettings = "none"
)

func PossibleValuesForAndroidDeviceOwnerWiFiConfigurationProxySettings() []string {
	return []string{
		string(AndroidDeviceOwnerWiFiConfigurationProxySettings_Automatic),
		string(AndroidDeviceOwnerWiFiConfigurationProxySettings_Manual),
		string(AndroidDeviceOwnerWiFiConfigurationProxySettings_None),
	}
}

func (s *AndroidDeviceOwnerWiFiConfigurationProxySettings) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAndroidDeviceOwnerWiFiConfigurationProxySettings(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAndroidDeviceOwnerWiFiConfigurationProxySettings(input string) (*AndroidDeviceOwnerWiFiConfigurationProxySettings, error) {
	vals := map[string]AndroidDeviceOwnerWiFiConfigurationProxySettings{
		"automatic": AndroidDeviceOwnerWiFiConfigurationProxySettings_Automatic,
		"manual":    AndroidDeviceOwnerWiFiConfigurationProxySettings_Manual,
		"none":      AndroidDeviceOwnerWiFiConfigurationProxySettings_None,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AndroidDeviceOwnerWiFiConfigurationProxySettings(input)
	return &out, nil
}
