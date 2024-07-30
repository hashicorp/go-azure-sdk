package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidGeneralDeviceConfigurationWebBrowserCookieSettings string

const (
	AndroidGeneralDeviceConfigurationWebBrowserCookieSettings_AllowAlways              AndroidGeneralDeviceConfigurationWebBrowserCookieSettings = "allowAlways"
	AndroidGeneralDeviceConfigurationWebBrowserCookieSettings_AllowCurrentWebSite      AndroidGeneralDeviceConfigurationWebBrowserCookieSettings = "allowCurrentWebSite"
	AndroidGeneralDeviceConfigurationWebBrowserCookieSettings_AllowFromWebsitesVisited AndroidGeneralDeviceConfigurationWebBrowserCookieSettings = "allowFromWebsitesVisited"
	AndroidGeneralDeviceConfigurationWebBrowserCookieSettings_BlockAlways              AndroidGeneralDeviceConfigurationWebBrowserCookieSettings = "blockAlways"
	AndroidGeneralDeviceConfigurationWebBrowserCookieSettings_BrowserDefault           AndroidGeneralDeviceConfigurationWebBrowserCookieSettings = "browserDefault"
)

func PossibleValuesForAndroidGeneralDeviceConfigurationWebBrowserCookieSettings() []string {
	return []string{
		string(AndroidGeneralDeviceConfigurationWebBrowserCookieSettings_AllowAlways),
		string(AndroidGeneralDeviceConfigurationWebBrowserCookieSettings_AllowCurrentWebSite),
		string(AndroidGeneralDeviceConfigurationWebBrowserCookieSettings_AllowFromWebsitesVisited),
		string(AndroidGeneralDeviceConfigurationWebBrowserCookieSettings_BlockAlways),
		string(AndroidGeneralDeviceConfigurationWebBrowserCookieSettings_BrowserDefault),
	}
}

func (s *AndroidGeneralDeviceConfigurationWebBrowserCookieSettings) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAndroidGeneralDeviceConfigurationWebBrowserCookieSettings(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAndroidGeneralDeviceConfigurationWebBrowserCookieSettings(input string) (*AndroidGeneralDeviceConfigurationWebBrowserCookieSettings, error) {
	vals := map[string]AndroidGeneralDeviceConfigurationWebBrowserCookieSettings{
		"allowalways":              AndroidGeneralDeviceConfigurationWebBrowserCookieSettings_AllowAlways,
		"allowcurrentwebsite":      AndroidGeneralDeviceConfigurationWebBrowserCookieSettings_AllowCurrentWebSite,
		"allowfromwebsitesvisited": AndroidGeneralDeviceConfigurationWebBrowserCookieSettings_AllowFromWebsitesVisited,
		"blockalways":              AndroidGeneralDeviceConfigurationWebBrowserCookieSettings_BlockAlways,
		"browserdefault":           AndroidGeneralDeviceConfigurationWebBrowserCookieSettings_BrowserDefault,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AndroidGeneralDeviceConfigurationWebBrowserCookieSettings(input)
	return &out, nil
}
