package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WindowsWifiEnterpriseEAPConfigurationProxySetting string

const (
	WindowsWifiEnterpriseEAPConfigurationProxySettingautomatic WindowsWifiEnterpriseEAPConfigurationProxySetting = "Automatic"
	WindowsWifiEnterpriseEAPConfigurationProxySettingmanual    WindowsWifiEnterpriseEAPConfigurationProxySetting = "Manual"
	WindowsWifiEnterpriseEAPConfigurationProxySettingnone      WindowsWifiEnterpriseEAPConfigurationProxySetting = "None"
)

func PossibleValuesForWindowsWifiEnterpriseEAPConfigurationProxySetting() []string {
	return []string{
		string(WindowsWifiEnterpriseEAPConfigurationProxySettingautomatic),
		string(WindowsWifiEnterpriseEAPConfigurationProxySettingmanual),
		string(WindowsWifiEnterpriseEAPConfigurationProxySettingnone),
	}
}

func parseWindowsWifiEnterpriseEAPConfigurationProxySetting(input string) (*WindowsWifiEnterpriseEAPConfigurationProxySetting, error) {
	vals := map[string]WindowsWifiEnterpriseEAPConfigurationProxySetting{
		"automatic": WindowsWifiEnterpriseEAPConfigurationProxySettingautomatic,
		"manual":    WindowsWifiEnterpriseEAPConfigurationProxySettingmanual,
		"none":      WindowsWifiEnterpriseEAPConfigurationProxySettingnone,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := WindowsWifiEnterpriseEAPConfigurationProxySetting(input)
	return &out, nil
}
