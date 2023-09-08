package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IosGeneralDeviceConfigurationSafariCookieSettings string

const (
	IosGeneralDeviceConfigurationSafariCookieSettingsallowAlways              IosGeneralDeviceConfigurationSafariCookieSettings = "AllowAlways"
	IosGeneralDeviceConfigurationSafariCookieSettingsallowCurrentWebSite      IosGeneralDeviceConfigurationSafariCookieSettings = "AllowCurrentWebSite"
	IosGeneralDeviceConfigurationSafariCookieSettingsallowFromWebsitesVisited IosGeneralDeviceConfigurationSafariCookieSettings = "AllowFromWebsitesVisited"
	IosGeneralDeviceConfigurationSafariCookieSettingsblockAlways              IosGeneralDeviceConfigurationSafariCookieSettings = "BlockAlways"
	IosGeneralDeviceConfigurationSafariCookieSettingsbrowserDefault           IosGeneralDeviceConfigurationSafariCookieSettings = "BrowserDefault"
)

func PossibleValuesForIosGeneralDeviceConfigurationSafariCookieSettings() []string {
	return []string{
		string(IosGeneralDeviceConfigurationSafariCookieSettingsallowAlways),
		string(IosGeneralDeviceConfigurationSafariCookieSettingsallowCurrentWebSite),
		string(IosGeneralDeviceConfigurationSafariCookieSettingsallowFromWebsitesVisited),
		string(IosGeneralDeviceConfigurationSafariCookieSettingsblockAlways),
		string(IosGeneralDeviceConfigurationSafariCookieSettingsbrowserDefault),
	}
}

func parseIosGeneralDeviceConfigurationSafariCookieSettings(input string) (*IosGeneralDeviceConfigurationSafariCookieSettings, error) {
	vals := map[string]IosGeneralDeviceConfigurationSafariCookieSettings{
		"allowalways":              IosGeneralDeviceConfigurationSafariCookieSettingsallowAlways,
		"allowcurrentwebsite":      IosGeneralDeviceConfigurationSafariCookieSettingsallowCurrentWebSite,
		"allowfromwebsitesvisited": IosGeneralDeviceConfigurationSafariCookieSettingsallowFromWebsitesVisited,
		"blockalways":              IosGeneralDeviceConfigurationSafariCookieSettingsblockAlways,
		"browserdefault":           IosGeneralDeviceConfigurationSafariCookieSettingsbrowserDefault,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := IosGeneralDeviceConfigurationSafariCookieSettings(input)
	return &out, nil
}
