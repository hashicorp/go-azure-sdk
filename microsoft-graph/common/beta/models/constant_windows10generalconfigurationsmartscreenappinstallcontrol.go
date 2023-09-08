package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Windows10GeneralConfigurationSmartScreenAppInstallControl string

const (
	Windows10GeneralConfigurationSmartScreenAppInstallControlanywhere        Windows10GeneralConfigurationSmartScreenAppInstallControl = "Anywhere"
	Windows10GeneralConfigurationSmartScreenAppInstallControlnotConfigured   Windows10GeneralConfigurationSmartScreenAppInstallControl = "NotConfigured"
	Windows10GeneralConfigurationSmartScreenAppInstallControlpreferStore     Windows10GeneralConfigurationSmartScreenAppInstallControl = "PreferStore"
	Windows10GeneralConfigurationSmartScreenAppInstallControlrecommendations Windows10GeneralConfigurationSmartScreenAppInstallControl = "Recommendations"
	Windows10GeneralConfigurationSmartScreenAppInstallControlstoreOnly       Windows10GeneralConfigurationSmartScreenAppInstallControl = "StoreOnly"
)

func PossibleValuesForWindows10GeneralConfigurationSmartScreenAppInstallControl() []string {
	return []string{
		string(Windows10GeneralConfigurationSmartScreenAppInstallControlanywhere),
		string(Windows10GeneralConfigurationSmartScreenAppInstallControlnotConfigured),
		string(Windows10GeneralConfigurationSmartScreenAppInstallControlpreferStore),
		string(Windows10GeneralConfigurationSmartScreenAppInstallControlrecommendations),
		string(Windows10GeneralConfigurationSmartScreenAppInstallControlstoreOnly),
	}
}

func parseWindows10GeneralConfigurationSmartScreenAppInstallControl(input string) (*Windows10GeneralConfigurationSmartScreenAppInstallControl, error) {
	vals := map[string]Windows10GeneralConfigurationSmartScreenAppInstallControl{
		"anywhere":        Windows10GeneralConfigurationSmartScreenAppInstallControlanywhere,
		"notconfigured":   Windows10GeneralConfigurationSmartScreenAppInstallControlnotConfigured,
		"preferstore":     Windows10GeneralConfigurationSmartScreenAppInstallControlpreferStore,
		"recommendations": Windows10GeneralConfigurationSmartScreenAppInstallControlrecommendations,
		"storeonly":       Windows10GeneralConfigurationSmartScreenAppInstallControlstoreOnly,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Windows10GeneralConfigurationSmartScreenAppInstallControl(input)
	return &out, nil
}
