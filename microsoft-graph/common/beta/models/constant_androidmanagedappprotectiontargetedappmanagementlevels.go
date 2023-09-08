package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidManagedAppProtectionTargetedAppManagementLevels string

const (
	AndroidManagedAppProtectionTargetedAppManagementLevelsandroidEnterprise                                      AndroidManagedAppProtectionTargetedAppManagementLevels = "AndroidEnterprise"
	AndroidManagedAppProtectionTargetedAppManagementLevelsandroidEnterpriseDedicatedDevicesWithAzureAdSharedMode AndroidManagedAppProtectionTargetedAppManagementLevels = "AndroidEnterpriseDedicatedDevicesWithAzureAdSharedMode"
	AndroidManagedAppProtectionTargetedAppManagementLevelsandroidOpenSourceProjectUserAssociated                 AndroidManagedAppProtectionTargetedAppManagementLevels = "AndroidOpenSourceProjectUserAssociated"
	AndroidManagedAppProtectionTargetedAppManagementLevelsandroidOpenSourceProjectUserless                       AndroidManagedAppProtectionTargetedAppManagementLevels = "AndroidOpenSourceProjectUserless"
	AndroidManagedAppProtectionTargetedAppManagementLevelsmdm                                                    AndroidManagedAppProtectionTargetedAppManagementLevels = "Mdm"
	AndroidManagedAppProtectionTargetedAppManagementLevelsunmanaged                                              AndroidManagedAppProtectionTargetedAppManagementLevels = "Unmanaged"
	AndroidManagedAppProtectionTargetedAppManagementLevelsunspecified                                            AndroidManagedAppProtectionTargetedAppManagementLevels = "Unspecified"
)

func PossibleValuesForAndroidManagedAppProtectionTargetedAppManagementLevels() []string {
	return []string{
		string(AndroidManagedAppProtectionTargetedAppManagementLevelsandroidEnterprise),
		string(AndroidManagedAppProtectionTargetedAppManagementLevelsandroidEnterpriseDedicatedDevicesWithAzureAdSharedMode),
		string(AndroidManagedAppProtectionTargetedAppManagementLevelsandroidOpenSourceProjectUserAssociated),
		string(AndroidManagedAppProtectionTargetedAppManagementLevelsandroidOpenSourceProjectUserless),
		string(AndroidManagedAppProtectionTargetedAppManagementLevelsmdm),
		string(AndroidManagedAppProtectionTargetedAppManagementLevelsunmanaged),
		string(AndroidManagedAppProtectionTargetedAppManagementLevelsunspecified),
	}
}

func parseAndroidManagedAppProtectionTargetedAppManagementLevels(input string) (*AndroidManagedAppProtectionTargetedAppManagementLevels, error) {
	vals := map[string]AndroidManagedAppProtectionTargetedAppManagementLevels{
		"androidenterprise": AndroidManagedAppProtectionTargetedAppManagementLevelsandroidEnterprise,
		"androidenterprisededicateddeviceswithazureadsharedmode": AndroidManagedAppProtectionTargetedAppManagementLevelsandroidEnterpriseDedicatedDevicesWithAzureAdSharedMode,
		"androidopensourceprojectuserassociated":                 AndroidManagedAppProtectionTargetedAppManagementLevelsandroidOpenSourceProjectUserAssociated,
		"androidopensourceprojectuserless":                       AndroidManagedAppProtectionTargetedAppManagementLevelsandroidOpenSourceProjectUserless,
		"mdm":                                                    AndroidManagedAppProtectionTargetedAppManagementLevelsmdm,
		"unmanaged":                                              AndroidManagedAppProtectionTargetedAppManagementLevelsunmanaged,
		"unspecified":                                            AndroidManagedAppProtectionTargetedAppManagementLevelsunspecified,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AndroidManagedAppProtectionTargetedAppManagementLevels(input)
	return &out, nil
}
