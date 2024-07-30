package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IosGeneralDeviceConfigurationSafariCookieSettings string

const (
	IosGeneralDeviceConfigurationSafariCookieSettings_AllowAlways              IosGeneralDeviceConfigurationSafariCookieSettings = "allowAlways"
	IosGeneralDeviceConfigurationSafariCookieSettings_AllowCurrentWebSite      IosGeneralDeviceConfigurationSafariCookieSettings = "allowCurrentWebSite"
	IosGeneralDeviceConfigurationSafariCookieSettings_AllowFromWebsitesVisited IosGeneralDeviceConfigurationSafariCookieSettings = "allowFromWebsitesVisited"
	IosGeneralDeviceConfigurationSafariCookieSettings_BlockAlways              IosGeneralDeviceConfigurationSafariCookieSettings = "blockAlways"
	IosGeneralDeviceConfigurationSafariCookieSettings_BrowserDefault           IosGeneralDeviceConfigurationSafariCookieSettings = "browserDefault"
)

func PossibleValuesForIosGeneralDeviceConfigurationSafariCookieSettings() []string {
	return []string{
		string(IosGeneralDeviceConfigurationSafariCookieSettings_AllowAlways),
		string(IosGeneralDeviceConfigurationSafariCookieSettings_AllowCurrentWebSite),
		string(IosGeneralDeviceConfigurationSafariCookieSettings_AllowFromWebsitesVisited),
		string(IosGeneralDeviceConfigurationSafariCookieSettings_BlockAlways),
		string(IosGeneralDeviceConfigurationSafariCookieSettings_BrowserDefault),
	}
}

func (s *IosGeneralDeviceConfigurationSafariCookieSettings) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseIosGeneralDeviceConfigurationSafariCookieSettings(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseIosGeneralDeviceConfigurationSafariCookieSettings(input string) (*IosGeneralDeviceConfigurationSafariCookieSettings, error) {
	vals := map[string]IosGeneralDeviceConfigurationSafariCookieSettings{
		"allowalways":              IosGeneralDeviceConfigurationSafariCookieSettings_AllowAlways,
		"allowcurrentwebsite":      IosGeneralDeviceConfigurationSafariCookieSettings_AllowCurrentWebSite,
		"allowfromwebsitesvisited": IosGeneralDeviceConfigurationSafariCookieSettings_AllowFromWebsitesVisited,
		"blockalways":              IosGeneralDeviceConfigurationSafariCookieSettings_BlockAlways,
		"browserdefault":           IosGeneralDeviceConfigurationSafariCookieSettings_BrowserDefault,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := IosGeneralDeviceConfigurationSafariCookieSettings(input)
	return &out, nil
}
