package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidDeviceOwnerGeneralDeviceConfigurationAppsDefaultPermissionPolicy string

const (
	AndroidDeviceOwnerGeneralDeviceConfigurationAppsDefaultPermissionPolicyautoDeny      AndroidDeviceOwnerGeneralDeviceConfigurationAppsDefaultPermissionPolicy = "AutoDeny"
	AndroidDeviceOwnerGeneralDeviceConfigurationAppsDefaultPermissionPolicyautoGrant     AndroidDeviceOwnerGeneralDeviceConfigurationAppsDefaultPermissionPolicy = "AutoGrant"
	AndroidDeviceOwnerGeneralDeviceConfigurationAppsDefaultPermissionPolicydeviceDefault AndroidDeviceOwnerGeneralDeviceConfigurationAppsDefaultPermissionPolicy = "DeviceDefault"
	AndroidDeviceOwnerGeneralDeviceConfigurationAppsDefaultPermissionPolicyprompt        AndroidDeviceOwnerGeneralDeviceConfigurationAppsDefaultPermissionPolicy = "Prompt"
)

func PossibleValuesForAndroidDeviceOwnerGeneralDeviceConfigurationAppsDefaultPermissionPolicy() []string {
	return []string{
		string(AndroidDeviceOwnerGeneralDeviceConfigurationAppsDefaultPermissionPolicyautoDeny),
		string(AndroidDeviceOwnerGeneralDeviceConfigurationAppsDefaultPermissionPolicyautoGrant),
		string(AndroidDeviceOwnerGeneralDeviceConfigurationAppsDefaultPermissionPolicydeviceDefault),
		string(AndroidDeviceOwnerGeneralDeviceConfigurationAppsDefaultPermissionPolicyprompt),
	}
}

func parseAndroidDeviceOwnerGeneralDeviceConfigurationAppsDefaultPermissionPolicy(input string) (*AndroidDeviceOwnerGeneralDeviceConfigurationAppsDefaultPermissionPolicy, error) {
	vals := map[string]AndroidDeviceOwnerGeneralDeviceConfigurationAppsDefaultPermissionPolicy{
		"autodeny":      AndroidDeviceOwnerGeneralDeviceConfigurationAppsDefaultPermissionPolicyautoDeny,
		"autogrant":     AndroidDeviceOwnerGeneralDeviceConfigurationAppsDefaultPermissionPolicyautoGrant,
		"devicedefault": AndroidDeviceOwnerGeneralDeviceConfigurationAppsDefaultPermissionPolicydeviceDefault,
		"prompt":        AndroidDeviceOwnerGeneralDeviceConfigurationAppsDefaultPermissionPolicyprompt,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AndroidDeviceOwnerGeneralDeviceConfigurationAppsDefaultPermissionPolicy(input)
	return &out, nil
}
