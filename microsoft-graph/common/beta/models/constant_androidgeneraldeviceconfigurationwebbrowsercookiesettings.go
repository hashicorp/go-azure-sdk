package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidGeneralDeviceConfigurationWebBrowserCookieSettings string

const (
	AndroidGeneralDeviceConfigurationWebBrowserCookieSettingsallowAlways              AndroidGeneralDeviceConfigurationWebBrowserCookieSettings = "AllowAlways"
	AndroidGeneralDeviceConfigurationWebBrowserCookieSettingsallowCurrentWebSite      AndroidGeneralDeviceConfigurationWebBrowserCookieSettings = "AllowCurrentWebSite"
	AndroidGeneralDeviceConfigurationWebBrowserCookieSettingsallowFromWebsitesVisited AndroidGeneralDeviceConfigurationWebBrowserCookieSettings = "AllowFromWebsitesVisited"
	AndroidGeneralDeviceConfigurationWebBrowserCookieSettingsblockAlways              AndroidGeneralDeviceConfigurationWebBrowserCookieSettings = "BlockAlways"
	AndroidGeneralDeviceConfigurationWebBrowserCookieSettingsbrowserDefault           AndroidGeneralDeviceConfigurationWebBrowserCookieSettings = "BrowserDefault"
)

func PossibleValuesForAndroidGeneralDeviceConfigurationWebBrowserCookieSettings() []string {
	return []string{
		string(AndroidGeneralDeviceConfigurationWebBrowserCookieSettingsallowAlways),
		string(AndroidGeneralDeviceConfigurationWebBrowserCookieSettingsallowCurrentWebSite),
		string(AndroidGeneralDeviceConfigurationWebBrowserCookieSettingsallowFromWebsitesVisited),
		string(AndroidGeneralDeviceConfigurationWebBrowserCookieSettingsblockAlways),
		string(AndroidGeneralDeviceConfigurationWebBrowserCookieSettingsbrowserDefault),
	}
}

func parseAndroidGeneralDeviceConfigurationWebBrowserCookieSettings(input string) (*AndroidGeneralDeviceConfigurationWebBrowserCookieSettings, error) {
	vals := map[string]AndroidGeneralDeviceConfigurationWebBrowserCookieSettings{
		"allowalways":              AndroidGeneralDeviceConfigurationWebBrowserCookieSettingsallowAlways,
		"allowcurrentwebsite":      AndroidGeneralDeviceConfigurationWebBrowserCookieSettingsallowCurrentWebSite,
		"allowfromwebsitesvisited": AndroidGeneralDeviceConfigurationWebBrowserCookieSettingsallowFromWebsitesVisited,
		"blockalways":              AndroidGeneralDeviceConfigurationWebBrowserCookieSettingsblockAlways,
		"browserdefault":           AndroidGeneralDeviceConfigurationWebBrowserCookieSettingsbrowserDefault,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AndroidGeneralDeviceConfigurationWebBrowserCookieSettings(input)
	return &out, nil
}
