package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IosEnterpriseWiFiConfigurationProxySettings string

const (
	IosEnterpriseWiFiConfigurationProxySettingsautomatic IosEnterpriseWiFiConfigurationProxySettings = "Automatic"
	IosEnterpriseWiFiConfigurationProxySettingsmanual    IosEnterpriseWiFiConfigurationProxySettings = "Manual"
	IosEnterpriseWiFiConfigurationProxySettingsnone      IosEnterpriseWiFiConfigurationProxySettings = "None"
)

func PossibleValuesForIosEnterpriseWiFiConfigurationProxySettings() []string {
	return []string{
		string(IosEnterpriseWiFiConfigurationProxySettingsautomatic),
		string(IosEnterpriseWiFiConfigurationProxySettingsmanual),
		string(IosEnterpriseWiFiConfigurationProxySettingsnone),
	}
}

func parseIosEnterpriseWiFiConfigurationProxySettings(input string) (*IosEnterpriseWiFiConfigurationProxySettings, error) {
	vals := map[string]IosEnterpriseWiFiConfigurationProxySettings{
		"automatic": IosEnterpriseWiFiConfigurationProxySettingsautomatic,
		"manual":    IosEnterpriseWiFiConfigurationProxySettingsmanual,
		"none":      IosEnterpriseWiFiConfigurationProxySettingsnone,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := IosEnterpriseWiFiConfigurationProxySettings(input)
	return &out, nil
}
