package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceManagementConfigurationSettingGroupDefinitionUxBehavior string

const (
	DeviceManagementConfigurationSettingGroupDefinitionUxBehaviorcontextPane     DeviceManagementConfigurationSettingGroupDefinitionUxBehavior = "ContextPane"
	DeviceManagementConfigurationSettingGroupDefinitionUxBehaviordefault         DeviceManagementConfigurationSettingGroupDefinitionUxBehavior = "Default"
	DeviceManagementConfigurationSettingGroupDefinitionUxBehaviordropdown        DeviceManagementConfigurationSettingGroupDefinitionUxBehavior = "Dropdown"
	DeviceManagementConfigurationSettingGroupDefinitionUxBehaviorlargeTextBox    DeviceManagementConfigurationSettingGroupDefinitionUxBehavior = "LargeTextBox"
	DeviceManagementConfigurationSettingGroupDefinitionUxBehaviormultiheaderGrid DeviceManagementConfigurationSettingGroupDefinitionUxBehavior = "MultiheaderGrid"
	DeviceManagementConfigurationSettingGroupDefinitionUxBehaviorsmallTextBox    DeviceManagementConfigurationSettingGroupDefinitionUxBehavior = "SmallTextBox"
	DeviceManagementConfigurationSettingGroupDefinitionUxBehaviortoggle          DeviceManagementConfigurationSettingGroupDefinitionUxBehavior = "Toggle"
)

func PossibleValuesForDeviceManagementConfigurationSettingGroupDefinitionUxBehavior() []string {
	return []string{
		string(DeviceManagementConfigurationSettingGroupDefinitionUxBehaviorcontextPane),
		string(DeviceManagementConfigurationSettingGroupDefinitionUxBehaviordefault),
		string(DeviceManagementConfigurationSettingGroupDefinitionUxBehaviordropdown),
		string(DeviceManagementConfigurationSettingGroupDefinitionUxBehaviorlargeTextBox),
		string(DeviceManagementConfigurationSettingGroupDefinitionUxBehaviormultiheaderGrid),
		string(DeviceManagementConfigurationSettingGroupDefinitionUxBehaviorsmallTextBox),
		string(DeviceManagementConfigurationSettingGroupDefinitionUxBehaviortoggle),
	}
}

func parseDeviceManagementConfigurationSettingGroupDefinitionUxBehavior(input string) (*DeviceManagementConfigurationSettingGroupDefinitionUxBehavior, error) {
	vals := map[string]DeviceManagementConfigurationSettingGroupDefinitionUxBehavior{
		"contextpane":     DeviceManagementConfigurationSettingGroupDefinitionUxBehaviorcontextPane,
		"default":         DeviceManagementConfigurationSettingGroupDefinitionUxBehaviordefault,
		"dropdown":        DeviceManagementConfigurationSettingGroupDefinitionUxBehaviordropdown,
		"largetextbox":    DeviceManagementConfigurationSettingGroupDefinitionUxBehaviorlargeTextBox,
		"multiheadergrid": DeviceManagementConfigurationSettingGroupDefinitionUxBehaviormultiheaderGrid,
		"smalltextbox":    DeviceManagementConfigurationSettingGroupDefinitionUxBehaviorsmallTextBox,
		"toggle":          DeviceManagementConfigurationSettingGroupDefinitionUxBehaviortoggle,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DeviceManagementConfigurationSettingGroupDefinitionUxBehavior(input)
	return &out, nil
}
