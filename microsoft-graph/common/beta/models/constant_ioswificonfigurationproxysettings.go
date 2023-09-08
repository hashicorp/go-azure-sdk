package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IosWiFiConfigurationProxySettings string

const (
	IosWiFiConfigurationProxySettingsautomatic IosWiFiConfigurationProxySettings = "Automatic"
	IosWiFiConfigurationProxySettingsmanual    IosWiFiConfigurationProxySettings = "Manual"
	IosWiFiConfigurationProxySettingsnone      IosWiFiConfigurationProxySettings = "None"
)

func PossibleValuesForIosWiFiConfigurationProxySettings() []string {
	return []string{
		string(IosWiFiConfigurationProxySettingsautomatic),
		string(IosWiFiConfigurationProxySettingsmanual),
		string(IosWiFiConfigurationProxySettingsnone),
	}
}

func parseIosWiFiConfigurationProxySettings(input string) (*IosWiFiConfigurationProxySettings, error) {
	vals := map[string]IosWiFiConfigurationProxySettings{
		"automatic": IosWiFiConfigurationProxySettingsautomatic,
		"manual":    IosWiFiConfigurationProxySettingsmanual,
		"none":      IosWiFiConfigurationProxySettingsnone,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := IosWiFiConfigurationProxySettings(input)
	return &out, nil
}
