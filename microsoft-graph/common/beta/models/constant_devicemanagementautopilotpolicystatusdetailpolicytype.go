package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceManagementAutopilotPolicyStatusDetailPolicyType string

const (
	DeviceManagementAutopilotPolicyStatusDetailPolicyTypeappModel            DeviceManagementAutopilotPolicyStatusDetailPolicyType = "AppModel"
	DeviceManagementAutopilotPolicyStatusDetailPolicyTypeapplication         DeviceManagementAutopilotPolicyStatusDetailPolicyType = "Application"
	DeviceManagementAutopilotPolicyStatusDetailPolicyTypeconfigurationPolicy DeviceManagementAutopilotPolicyStatusDetailPolicyType = "ConfigurationPolicy"
	DeviceManagementAutopilotPolicyStatusDetailPolicyTypeunknown             DeviceManagementAutopilotPolicyStatusDetailPolicyType = "Unknown"
)

func PossibleValuesForDeviceManagementAutopilotPolicyStatusDetailPolicyType() []string {
	return []string{
		string(DeviceManagementAutopilotPolicyStatusDetailPolicyTypeappModel),
		string(DeviceManagementAutopilotPolicyStatusDetailPolicyTypeapplication),
		string(DeviceManagementAutopilotPolicyStatusDetailPolicyTypeconfigurationPolicy),
		string(DeviceManagementAutopilotPolicyStatusDetailPolicyTypeunknown),
	}
}

func parseDeviceManagementAutopilotPolicyStatusDetailPolicyType(input string) (*DeviceManagementAutopilotPolicyStatusDetailPolicyType, error) {
	vals := map[string]DeviceManagementAutopilotPolicyStatusDetailPolicyType{
		"appmodel":            DeviceManagementAutopilotPolicyStatusDetailPolicyTypeappModel,
		"application":         DeviceManagementAutopilotPolicyStatusDetailPolicyTypeapplication,
		"configurationpolicy": DeviceManagementAutopilotPolicyStatusDetailPolicyTypeconfigurationPolicy,
		"unknown":             DeviceManagementAutopilotPolicyStatusDetailPolicyTypeunknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DeviceManagementAutopilotPolicyStatusDetailPolicyType(input)
	return &out, nil
}
