package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidDeviceOwnerGeneralDeviceConfigurationKioskCustomizationStatusBar string

const (
	AndroidDeviceOwnerGeneralDeviceConfigurationKioskCustomizationStatusBarnotConfigured                     AndroidDeviceOwnerGeneralDeviceConfigurationKioskCustomizationStatusBar = "NotConfigured"
	AndroidDeviceOwnerGeneralDeviceConfigurationKioskCustomizationStatusBarnotificationsAndSystemInfoEnabled AndroidDeviceOwnerGeneralDeviceConfigurationKioskCustomizationStatusBar = "NotificationsAndSystemInfoEnabled"
	AndroidDeviceOwnerGeneralDeviceConfigurationKioskCustomizationStatusBarsystemInfoOnly                    AndroidDeviceOwnerGeneralDeviceConfigurationKioskCustomizationStatusBar = "SystemInfoOnly"
)

func PossibleValuesForAndroidDeviceOwnerGeneralDeviceConfigurationKioskCustomizationStatusBar() []string {
	return []string{
		string(AndroidDeviceOwnerGeneralDeviceConfigurationKioskCustomizationStatusBarnotConfigured),
		string(AndroidDeviceOwnerGeneralDeviceConfigurationKioskCustomizationStatusBarnotificationsAndSystemInfoEnabled),
		string(AndroidDeviceOwnerGeneralDeviceConfigurationKioskCustomizationStatusBarsystemInfoOnly),
	}
}

func parseAndroidDeviceOwnerGeneralDeviceConfigurationKioskCustomizationStatusBar(input string) (*AndroidDeviceOwnerGeneralDeviceConfigurationKioskCustomizationStatusBar, error) {
	vals := map[string]AndroidDeviceOwnerGeneralDeviceConfigurationKioskCustomizationStatusBar{
		"notconfigured":                     AndroidDeviceOwnerGeneralDeviceConfigurationKioskCustomizationStatusBarnotConfigured,
		"notificationsandsysteminfoenabled": AndroidDeviceOwnerGeneralDeviceConfigurationKioskCustomizationStatusBarnotificationsAndSystemInfoEnabled,
		"systeminfoonly":                    AndroidDeviceOwnerGeneralDeviceConfigurationKioskCustomizationStatusBarsystemInfoOnly,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AndroidDeviceOwnerGeneralDeviceConfigurationKioskCustomizationStatusBar(input)
	return &out, nil
}
