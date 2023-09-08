package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MacOSEnterpriseWiFiConfigurationProxySettings string

const (
	MacOSEnterpriseWiFiConfigurationProxySettingsautomatic MacOSEnterpriseWiFiConfigurationProxySettings = "Automatic"
	MacOSEnterpriseWiFiConfigurationProxySettingsmanual    MacOSEnterpriseWiFiConfigurationProxySettings = "Manual"
	MacOSEnterpriseWiFiConfigurationProxySettingsnone      MacOSEnterpriseWiFiConfigurationProxySettings = "None"
)

func PossibleValuesForMacOSEnterpriseWiFiConfigurationProxySettings() []string {
	return []string{
		string(MacOSEnterpriseWiFiConfigurationProxySettingsautomatic),
		string(MacOSEnterpriseWiFiConfigurationProxySettingsmanual),
		string(MacOSEnterpriseWiFiConfigurationProxySettingsnone),
	}
}

func parseMacOSEnterpriseWiFiConfigurationProxySettings(input string) (*MacOSEnterpriseWiFiConfigurationProxySettings, error) {
	vals := map[string]MacOSEnterpriseWiFiConfigurationProxySettings{
		"automatic": MacOSEnterpriseWiFiConfigurationProxySettingsautomatic,
		"manual":    MacOSEnterpriseWiFiConfigurationProxySettingsmanual,
		"none":      MacOSEnterpriseWiFiConfigurationProxySettingsnone,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := MacOSEnterpriseWiFiConfigurationProxySettings(input)
	return &out, nil
}
