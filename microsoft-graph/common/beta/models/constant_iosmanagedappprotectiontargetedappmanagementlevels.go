package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IosManagedAppProtectionTargetedAppManagementLevels string

const (
	IosManagedAppProtectionTargetedAppManagementLevelsandroidEnterprise                                      IosManagedAppProtectionTargetedAppManagementLevels = "AndroidEnterprise"
	IosManagedAppProtectionTargetedAppManagementLevelsandroidEnterpriseDedicatedDevicesWithAzureAdSharedMode IosManagedAppProtectionTargetedAppManagementLevels = "AndroidEnterpriseDedicatedDevicesWithAzureAdSharedMode"
	IosManagedAppProtectionTargetedAppManagementLevelsandroidOpenSourceProjectUserAssociated                 IosManagedAppProtectionTargetedAppManagementLevels = "AndroidOpenSourceProjectUserAssociated"
	IosManagedAppProtectionTargetedAppManagementLevelsandroidOpenSourceProjectUserless                       IosManagedAppProtectionTargetedAppManagementLevels = "AndroidOpenSourceProjectUserless"
	IosManagedAppProtectionTargetedAppManagementLevelsmdm                                                    IosManagedAppProtectionTargetedAppManagementLevels = "Mdm"
	IosManagedAppProtectionTargetedAppManagementLevelsunmanaged                                              IosManagedAppProtectionTargetedAppManagementLevels = "Unmanaged"
	IosManagedAppProtectionTargetedAppManagementLevelsunspecified                                            IosManagedAppProtectionTargetedAppManagementLevels = "Unspecified"
)

func PossibleValuesForIosManagedAppProtectionTargetedAppManagementLevels() []string {
	return []string{
		string(IosManagedAppProtectionTargetedAppManagementLevelsandroidEnterprise),
		string(IosManagedAppProtectionTargetedAppManagementLevelsandroidEnterpriseDedicatedDevicesWithAzureAdSharedMode),
		string(IosManagedAppProtectionTargetedAppManagementLevelsandroidOpenSourceProjectUserAssociated),
		string(IosManagedAppProtectionTargetedAppManagementLevelsandroidOpenSourceProjectUserless),
		string(IosManagedAppProtectionTargetedAppManagementLevelsmdm),
		string(IosManagedAppProtectionTargetedAppManagementLevelsunmanaged),
		string(IosManagedAppProtectionTargetedAppManagementLevelsunspecified),
	}
}

func parseIosManagedAppProtectionTargetedAppManagementLevels(input string) (*IosManagedAppProtectionTargetedAppManagementLevels, error) {
	vals := map[string]IosManagedAppProtectionTargetedAppManagementLevels{
		"androidenterprise": IosManagedAppProtectionTargetedAppManagementLevelsandroidEnterprise,
		"androidenterprisededicateddeviceswithazureadsharedmode": IosManagedAppProtectionTargetedAppManagementLevelsandroidEnterpriseDedicatedDevicesWithAzureAdSharedMode,
		"androidopensourceprojectuserassociated":                 IosManagedAppProtectionTargetedAppManagementLevelsandroidOpenSourceProjectUserAssociated,
		"androidopensourceprojectuserless":                       IosManagedAppProtectionTargetedAppManagementLevelsandroidOpenSourceProjectUserless,
		"mdm":                                                    IosManagedAppProtectionTargetedAppManagementLevelsmdm,
		"unmanaged":                                              IosManagedAppProtectionTargetedAppManagementLevelsunmanaged,
		"unspecified":                                            IosManagedAppProtectionTargetedAppManagementLevelsunspecified,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := IosManagedAppProtectionTargetedAppManagementLevels(input)
	return &out, nil
}
