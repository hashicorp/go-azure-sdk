package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MacOSWiFiConfigurationProxySettings string

const (
	MacOSWiFiConfigurationProxySettingsautomatic MacOSWiFiConfigurationProxySettings = "Automatic"
	MacOSWiFiConfigurationProxySettingsmanual    MacOSWiFiConfigurationProxySettings = "Manual"
	MacOSWiFiConfigurationProxySettingsnone      MacOSWiFiConfigurationProxySettings = "None"
)

func PossibleValuesForMacOSWiFiConfigurationProxySettings() []string {
	return []string{
		string(MacOSWiFiConfigurationProxySettingsautomatic),
		string(MacOSWiFiConfigurationProxySettingsmanual),
		string(MacOSWiFiConfigurationProxySettingsnone),
	}
}

func parseMacOSWiFiConfigurationProxySettings(input string) (*MacOSWiFiConfigurationProxySettings, error) {
	vals := map[string]MacOSWiFiConfigurationProxySettings{
		"automatic": MacOSWiFiConfigurationProxySettingsautomatic,
		"manual":    MacOSWiFiConfigurationProxySettingsmanual,
		"none":      MacOSWiFiConfigurationProxySettingsnone,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := MacOSWiFiConfigurationProxySettings(input)
	return &out, nil
}
