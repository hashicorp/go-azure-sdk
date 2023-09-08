package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidDeviceOwnerGeneralDeviceConfigurationKioskModeUseManagedHomeScreenApp string

const (
	AndroidDeviceOwnerGeneralDeviceConfigurationKioskModeUseManagedHomeScreenAppmultiAppMode  AndroidDeviceOwnerGeneralDeviceConfigurationKioskModeUseManagedHomeScreenApp = "MultiAppMode"
	AndroidDeviceOwnerGeneralDeviceConfigurationKioskModeUseManagedHomeScreenAppnotConfigured AndroidDeviceOwnerGeneralDeviceConfigurationKioskModeUseManagedHomeScreenApp = "NotConfigured"
	AndroidDeviceOwnerGeneralDeviceConfigurationKioskModeUseManagedHomeScreenAppsingleAppMode AndroidDeviceOwnerGeneralDeviceConfigurationKioskModeUseManagedHomeScreenApp = "SingleAppMode"
)

func PossibleValuesForAndroidDeviceOwnerGeneralDeviceConfigurationKioskModeUseManagedHomeScreenApp() []string {
	return []string{
		string(AndroidDeviceOwnerGeneralDeviceConfigurationKioskModeUseManagedHomeScreenAppmultiAppMode),
		string(AndroidDeviceOwnerGeneralDeviceConfigurationKioskModeUseManagedHomeScreenAppnotConfigured),
		string(AndroidDeviceOwnerGeneralDeviceConfigurationKioskModeUseManagedHomeScreenAppsingleAppMode),
	}
}

func parseAndroidDeviceOwnerGeneralDeviceConfigurationKioskModeUseManagedHomeScreenApp(input string) (*AndroidDeviceOwnerGeneralDeviceConfigurationKioskModeUseManagedHomeScreenApp, error) {
	vals := map[string]AndroidDeviceOwnerGeneralDeviceConfigurationKioskModeUseManagedHomeScreenApp{
		"multiappmode":  AndroidDeviceOwnerGeneralDeviceConfigurationKioskModeUseManagedHomeScreenAppmultiAppMode,
		"notconfigured": AndroidDeviceOwnerGeneralDeviceConfigurationKioskModeUseManagedHomeScreenAppnotConfigured,
		"singleappmode": AndroidDeviceOwnerGeneralDeviceConfigurationKioskModeUseManagedHomeScreenAppsingleAppMode,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AndroidDeviceOwnerGeneralDeviceConfigurationKioskModeUseManagedHomeScreenApp(input)
	return &out, nil
}
