package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceManagementConfigurationPolicyTemplateLifecycleState string

const (
	DeviceManagementConfigurationPolicyTemplateLifecycleStateactive     DeviceManagementConfigurationPolicyTemplateLifecycleState = "Active"
	DeviceManagementConfigurationPolicyTemplateLifecycleStatedeprecated DeviceManagementConfigurationPolicyTemplateLifecycleState = "Deprecated"
	DeviceManagementConfigurationPolicyTemplateLifecycleStatedraft      DeviceManagementConfigurationPolicyTemplateLifecycleState = "Draft"
	DeviceManagementConfigurationPolicyTemplateLifecycleStateinvalid    DeviceManagementConfigurationPolicyTemplateLifecycleState = "Invalid"
	DeviceManagementConfigurationPolicyTemplateLifecycleStateretired    DeviceManagementConfigurationPolicyTemplateLifecycleState = "Retired"
	DeviceManagementConfigurationPolicyTemplateLifecycleStatesuperseded DeviceManagementConfigurationPolicyTemplateLifecycleState = "Superseded"
)

func PossibleValuesForDeviceManagementConfigurationPolicyTemplateLifecycleState() []string {
	return []string{
		string(DeviceManagementConfigurationPolicyTemplateLifecycleStateactive),
		string(DeviceManagementConfigurationPolicyTemplateLifecycleStatedeprecated),
		string(DeviceManagementConfigurationPolicyTemplateLifecycleStatedraft),
		string(DeviceManagementConfigurationPolicyTemplateLifecycleStateinvalid),
		string(DeviceManagementConfigurationPolicyTemplateLifecycleStateretired),
		string(DeviceManagementConfigurationPolicyTemplateLifecycleStatesuperseded),
	}
}

func parseDeviceManagementConfigurationPolicyTemplateLifecycleState(input string) (*DeviceManagementConfigurationPolicyTemplateLifecycleState, error) {
	vals := map[string]DeviceManagementConfigurationPolicyTemplateLifecycleState{
		"active":     DeviceManagementConfigurationPolicyTemplateLifecycleStateactive,
		"deprecated": DeviceManagementConfigurationPolicyTemplateLifecycleStatedeprecated,
		"draft":      DeviceManagementConfigurationPolicyTemplateLifecycleStatedraft,
		"invalid":    DeviceManagementConfigurationPolicyTemplateLifecycleStateinvalid,
		"retired":    DeviceManagementConfigurationPolicyTemplateLifecycleStateretired,
		"superseded": DeviceManagementConfigurationPolicyTemplateLifecycleStatesuperseded,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DeviceManagementConfigurationPolicyTemplateLifecycleState(input)
	return &out, nil
}
