package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidWorkProfileEnterpriseWiFiConfigurationProxySettings string

const (
	AndroidWorkProfileEnterpriseWiFiConfigurationProxySettingsautomatic AndroidWorkProfileEnterpriseWiFiConfigurationProxySettings = "Automatic"
	AndroidWorkProfileEnterpriseWiFiConfigurationProxySettingsmanual    AndroidWorkProfileEnterpriseWiFiConfigurationProxySettings = "Manual"
	AndroidWorkProfileEnterpriseWiFiConfigurationProxySettingsnone      AndroidWorkProfileEnterpriseWiFiConfigurationProxySettings = "None"
)

func PossibleValuesForAndroidWorkProfileEnterpriseWiFiConfigurationProxySettings() []string {
	return []string{
		string(AndroidWorkProfileEnterpriseWiFiConfigurationProxySettingsautomatic),
		string(AndroidWorkProfileEnterpriseWiFiConfigurationProxySettingsmanual),
		string(AndroidWorkProfileEnterpriseWiFiConfigurationProxySettingsnone),
	}
}

func parseAndroidWorkProfileEnterpriseWiFiConfigurationProxySettings(input string) (*AndroidWorkProfileEnterpriseWiFiConfigurationProxySettings, error) {
	vals := map[string]AndroidWorkProfileEnterpriseWiFiConfigurationProxySettings{
		"automatic": AndroidWorkProfileEnterpriseWiFiConfigurationProxySettingsautomatic,
		"manual":    AndroidWorkProfileEnterpriseWiFiConfigurationProxySettingsmanual,
		"none":      AndroidWorkProfileEnterpriseWiFiConfigurationProxySettingsnone,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AndroidWorkProfileEnterpriseWiFiConfigurationProxySettings(input)
	return &out, nil
}
