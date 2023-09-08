package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidDeviceOwnerEnterpriseWiFiConfigurationProxySettings string

const (
	AndroidDeviceOwnerEnterpriseWiFiConfigurationProxySettingsautomatic AndroidDeviceOwnerEnterpriseWiFiConfigurationProxySettings = "Automatic"
	AndroidDeviceOwnerEnterpriseWiFiConfigurationProxySettingsmanual    AndroidDeviceOwnerEnterpriseWiFiConfigurationProxySettings = "Manual"
	AndroidDeviceOwnerEnterpriseWiFiConfigurationProxySettingsnone      AndroidDeviceOwnerEnterpriseWiFiConfigurationProxySettings = "None"
)

func PossibleValuesForAndroidDeviceOwnerEnterpriseWiFiConfigurationProxySettings() []string {
	return []string{
		string(AndroidDeviceOwnerEnterpriseWiFiConfigurationProxySettingsautomatic),
		string(AndroidDeviceOwnerEnterpriseWiFiConfigurationProxySettingsmanual),
		string(AndroidDeviceOwnerEnterpriseWiFiConfigurationProxySettingsnone),
	}
}

func parseAndroidDeviceOwnerEnterpriseWiFiConfigurationProxySettings(input string) (*AndroidDeviceOwnerEnterpriseWiFiConfigurationProxySettings, error) {
	vals := map[string]AndroidDeviceOwnerEnterpriseWiFiConfigurationProxySettings{
		"automatic": AndroidDeviceOwnerEnterpriseWiFiConfigurationProxySettingsautomatic,
		"manual":    AndroidDeviceOwnerEnterpriseWiFiConfigurationProxySettingsmanual,
		"none":      AndroidDeviceOwnerEnterpriseWiFiConfigurationProxySettingsnone,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AndroidDeviceOwnerEnterpriseWiFiConfigurationProxySettings(input)
	return &out, nil
}
