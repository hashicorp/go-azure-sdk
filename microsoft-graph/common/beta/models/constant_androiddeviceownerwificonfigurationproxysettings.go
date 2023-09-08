package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidDeviceOwnerWiFiConfigurationProxySettings string

const (
	AndroidDeviceOwnerWiFiConfigurationProxySettingsautomatic AndroidDeviceOwnerWiFiConfigurationProxySettings = "Automatic"
	AndroidDeviceOwnerWiFiConfigurationProxySettingsmanual    AndroidDeviceOwnerWiFiConfigurationProxySettings = "Manual"
	AndroidDeviceOwnerWiFiConfigurationProxySettingsnone      AndroidDeviceOwnerWiFiConfigurationProxySettings = "None"
)

func PossibleValuesForAndroidDeviceOwnerWiFiConfigurationProxySettings() []string {
	return []string{
		string(AndroidDeviceOwnerWiFiConfigurationProxySettingsautomatic),
		string(AndroidDeviceOwnerWiFiConfigurationProxySettingsmanual),
		string(AndroidDeviceOwnerWiFiConfigurationProxySettingsnone),
	}
}

func parseAndroidDeviceOwnerWiFiConfigurationProxySettings(input string) (*AndroidDeviceOwnerWiFiConfigurationProxySettings, error) {
	vals := map[string]AndroidDeviceOwnerWiFiConfigurationProxySettings{
		"automatic": AndroidDeviceOwnerWiFiConfigurationProxySettingsautomatic,
		"manual":    AndroidDeviceOwnerWiFiConfigurationProxySettingsmanual,
		"none":      AndroidDeviceOwnerWiFiConfigurationProxySettingsnone,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AndroidDeviceOwnerWiFiConfigurationProxySettings(input)
	return &out, nil
}
