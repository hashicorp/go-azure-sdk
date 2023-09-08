package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidDeviceOwnerGeneralDeviceConfigurationPersonalProfilePlayStoreMode string

const (
	AndroidDeviceOwnerGeneralDeviceConfigurationPersonalProfilePlayStoreModeallowedApps   AndroidDeviceOwnerGeneralDeviceConfigurationPersonalProfilePlayStoreMode = "AllowedApps"
	AndroidDeviceOwnerGeneralDeviceConfigurationPersonalProfilePlayStoreModeblockedApps   AndroidDeviceOwnerGeneralDeviceConfigurationPersonalProfilePlayStoreMode = "BlockedApps"
	AndroidDeviceOwnerGeneralDeviceConfigurationPersonalProfilePlayStoreModenotConfigured AndroidDeviceOwnerGeneralDeviceConfigurationPersonalProfilePlayStoreMode = "NotConfigured"
)

func PossibleValuesForAndroidDeviceOwnerGeneralDeviceConfigurationPersonalProfilePlayStoreMode() []string {
	return []string{
		string(AndroidDeviceOwnerGeneralDeviceConfigurationPersonalProfilePlayStoreModeallowedApps),
		string(AndroidDeviceOwnerGeneralDeviceConfigurationPersonalProfilePlayStoreModeblockedApps),
		string(AndroidDeviceOwnerGeneralDeviceConfigurationPersonalProfilePlayStoreModenotConfigured),
	}
}

func parseAndroidDeviceOwnerGeneralDeviceConfigurationPersonalProfilePlayStoreMode(input string) (*AndroidDeviceOwnerGeneralDeviceConfigurationPersonalProfilePlayStoreMode, error) {
	vals := map[string]AndroidDeviceOwnerGeneralDeviceConfigurationPersonalProfilePlayStoreMode{
		"allowedapps":   AndroidDeviceOwnerGeneralDeviceConfigurationPersonalProfilePlayStoreModeallowedApps,
		"blockedapps":   AndroidDeviceOwnerGeneralDeviceConfigurationPersonalProfilePlayStoreModeblockedApps,
		"notconfigured": AndroidDeviceOwnerGeneralDeviceConfigurationPersonalProfilePlayStoreModenotConfigured,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AndroidDeviceOwnerGeneralDeviceConfigurationPersonalProfilePlayStoreMode(input)
	return &out, nil
}
