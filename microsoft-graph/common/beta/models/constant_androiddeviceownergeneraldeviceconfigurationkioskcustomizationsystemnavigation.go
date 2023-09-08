package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidDeviceOwnerGeneralDeviceConfigurationKioskCustomizationSystemNavigation string

const (
	AndroidDeviceOwnerGeneralDeviceConfigurationKioskCustomizationSystemNavigationhomeButtonOnly    AndroidDeviceOwnerGeneralDeviceConfigurationKioskCustomizationSystemNavigation = "HomeButtonOnly"
	AndroidDeviceOwnerGeneralDeviceConfigurationKioskCustomizationSystemNavigationnavigationEnabled AndroidDeviceOwnerGeneralDeviceConfigurationKioskCustomizationSystemNavigation = "NavigationEnabled"
	AndroidDeviceOwnerGeneralDeviceConfigurationKioskCustomizationSystemNavigationnotConfigured     AndroidDeviceOwnerGeneralDeviceConfigurationKioskCustomizationSystemNavigation = "NotConfigured"
)

func PossibleValuesForAndroidDeviceOwnerGeneralDeviceConfigurationKioskCustomizationSystemNavigation() []string {
	return []string{
		string(AndroidDeviceOwnerGeneralDeviceConfigurationKioskCustomizationSystemNavigationhomeButtonOnly),
		string(AndroidDeviceOwnerGeneralDeviceConfigurationKioskCustomizationSystemNavigationnavigationEnabled),
		string(AndroidDeviceOwnerGeneralDeviceConfigurationKioskCustomizationSystemNavigationnotConfigured),
	}
}

func parseAndroidDeviceOwnerGeneralDeviceConfigurationKioskCustomizationSystemNavigation(input string) (*AndroidDeviceOwnerGeneralDeviceConfigurationKioskCustomizationSystemNavigation, error) {
	vals := map[string]AndroidDeviceOwnerGeneralDeviceConfigurationKioskCustomizationSystemNavigation{
		"homebuttononly":    AndroidDeviceOwnerGeneralDeviceConfigurationKioskCustomizationSystemNavigationhomeButtonOnly,
		"navigationenabled": AndroidDeviceOwnerGeneralDeviceConfigurationKioskCustomizationSystemNavigationnavigationEnabled,
		"notconfigured":     AndroidDeviceOwnerGeneralDeviceConfigurationKioskCustomizationSystemNavigationnotConfigured,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AndroidDeviceOwnerGeneralDeviceConfigurationKioskCustomizationSystemNavigation(input)
	return &out, nil
}
