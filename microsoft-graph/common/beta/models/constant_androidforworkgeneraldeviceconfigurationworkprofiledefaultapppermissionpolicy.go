package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidForWorkGeneralDeviceConfigurationWorkProfileDefaultAppPermissionPolicy string

const (
	AndroidForWorkGeneralDeviceConfigurationWorkProfileDefaultAppPermissionPolicyautoDeny      AndroidForWorkGeneralDeviceConfigurationWorkProfileDefaultAppPermissionPolicy = "AutoDeny"
	AndroidForWorkGeneralDeviceConfigurationWorkProfileDefaultAppPermissionPolicyautoGrant     AndroidForWorkGeneralDeviceConfigurationWorkProfileDefaultAppPermissionPolicy = "AutoGrant"
	AndroidForWorkGeneralDeviceConfigurationWorkProfileDefaultAppPermissionPolicydeviceDefault AndroidForWorkGeneralDeviceConfigurationWorkProfileDefaultAppPermissionPolicy = "DeviceDefault"
	AndroidForWorkGeneralDeviceConfigurationWorkProfileDefaultAppPermissionPolicyprompt        AndroidForWorkGeneralDeviceConfigurationWorkProfileDefaultAppPermissionPolicy = "Prompt"
)

func PossibleValuesForAndroidForWorkGeneralDeviceConfigurationWorkProfileDefaultAppPermissionPolicy() []string {
	return []string{
		string(AndroidForWorkGeneralDeviceConfigurationWorkProfileDefaultAppPermissionPolicyautoDeny),
		string(AndroidForWorkGeneralDeviceConfigurationWorkProfileDefaultAppPermissionPolicyautoGrant),
		string(AndroidForWorkGeneralDeviceConfigurationWorkProfileDefaultAppPermissionPolicydeviceDefault),
		string(AndroidForWorkGeneralDeviceConfigurationWorkProfileDefaultAppPermissionPolicyprompt),
	}
}

func parseAndroidForWorkGeneralDeviceConfigurationWorkProfileDefaultAppPermissionPolicy(input string) (*AndroidForWorkGeneralDeviceConfigurationWorkProfileDefaultAppPermissionPolicy, error) {
	vals := map[string]AndroidForWorkGeneralDeviceConfigurationWorkProfileDefaultAppPermissionPolicy{
		"autodeny":      AndroidForWorkGeneralDeviceConfigurationWorkProfileDefaultAppPermissionPolicyautoDeny,
		"autogrant":     AndroidForWorkGeneralDeviceConfigurationWorkProfileDefaultAppPermissionPolicyautoGrant,
		"devicedefault": AndroidForWorkGeneralDeviceConfigurationWorkProfileDefaultAppPermissionPolicydeviceDefault,
		"prompt":        AndroidForWorkGeneralDeviceConfigurationWorkProfileDefaultAppPermissionPolicyprompt,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AndroidForWorkGeneralDeviceConfigurationWorkProfileDefaultAppPermissionPolicy(input)
	return &out, nil
}
