package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidWorkProfileGeneralDeviceConfigurationWorkProfileDefaultAppPermissionPolicy string

const (
	AndroidWorkProfileGeneralDeviceConfigurationWorkProfileDefaultAppPermissionPolicyautoDeny      AndroidWorkProfileGeneralDeviceConfigurationWorkProfileDefaultAppPermissionPolicy = "AutoDeny"
	AndroidWorkProfileGeneralDeviceConfigurationWorkProfileDefaultAppPermissionPolicyautoGrant     AndroidWorkProfileGeneralDeviceConfigurationWorkProfileDefaultAppPermissionPolicy = "AutoGrant"
	AndroidWorkProfileGeneralDeviceConfigurationWorkProfileDefaultAppPermissionPolicydeviceDefault AndroidWorkProfileGeneralDeviceConfigurationWorkProfileDefaultAppPermissionPolicy = "DeviceDefault"
	AndroidWorkProfileGeneralDeviceConfigurationWorkProfileDefaultAppPermissionPolicyprompt        AndroidWorkProfileGeneralDeviceConfigurationWorkProfileDefaultAppPermissionPolicy = "Prompt"
)

func PossibleValuesForAndroidWorkProfileGeneralDeviceConfigurationWorkProfileDefaultAppPermissionPolicy() []string {
	return []string{
		string(AndroidWorkProfileGeneralDeviceConfigurationWorkProfileDefaultAppPermissionPolicyautoDeny),
		string(AndroidWorkProfileGeneralDeviceConfigurationWorkProfileDefaultAppPermissionPolicyautoGrant),
		string(AndroidWorkProfileGeneralDeviceConfigurationWorkProfileDefaultAppPermissionPolicydeviceDefault),
		string(AndroidWorkProfileGeneralDeviceConfigurationWorkProfileDefaultAppPermissionPolicyprompt),
	}
}

func parseAndroidWorkProfileGeneralDeviceConfigurationWorkProfileDefaultAppPermissionPolicy(input string) (*AndroidWorkProfileGeneralDeviceConfigurationWorkProfileDefaultAppPermissionPolicy, error) {
	vals := map[string]AndroidWorkProfileGeneralDeviceConfigurationWorkProfileDefaultAppPermissionPolicy{
		"autodeny":      AndroidWorkProfileGeneralDeviceConfigurationWorkProfileDefaultAppPermissionPolicyautoDeny,
		"autogrant":     AndroidWorkProfileGeneralDeviceConfigurationWorkProfileDefaultAppPermissionPolicyautoGrant,
		"devicedefault": AndroidWorkProfileGeneralDeviceConfigurationWorkProfileDefaultAppPermissionPolicydeviceDefault,
		"prompt":        AndroidWorkProfileGeneralDeviceConfigurationWorkProfileDefaultAppPermissionPolicyprompt,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AndroidWorkProfileGeneralDeviceConfigurationWorkProfileDefaultAppPermissionPolicy(input)
	return &out, nil
}
