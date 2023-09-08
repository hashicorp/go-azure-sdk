package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceManagementConfigurationSettingDefinitionUxBehavior string

const (
	DeviceManagementConfigurationSettingDefinitionUxBehaviorcontextPane     DeviceManagementConfigurationSettingDefinitionUxBehavior = "ContextPane"
	DeviceManagementConfigurationSettingDefinitionUxBehaviordefault         DeviceManagementConfigurationSettingDefinitionUxBehavior = "Default"
	DeviceManagementConfigurationSettingDefinitionUxBehaviordropdown        DeviceManagementConfigurationSettingDefinitionUxBehavior = "Dropdown"
	DeviceManagementConfigurationSettingDefinitionUxBehaviorlargeTextBox    DeviceManagementConfigurationSettingDefinitionUxBehavior = "LargeTextBox"
	DeviceManagementConfigurationSettingDefinitionUxBehaviormultiheaderGrid DeviceManagementConfigurationSettingDefinitionUxBehavior = "MultiheaderGrid"
	DeviceManagementConfigurationSettingDefinitionUxBehaviorsmallTextBox    DeviceManagementConfigurationSettingDefinitionUxBehavior = "SmallTextBox"
	DeviceManagementConfigurationSettingDefinitionUxBehaviortoggle          DeviceManagementConfigurationSettingDefinitionUxBehavior = "Toggle"
)

func PossibleValuesForDeviceManagementConfigurationSettingDefinitionUxBehavior() []string {
	return []string{
		string(DeviceManagementConfigurationSettingDefinitionUxBehaviorcontextPane),
		string(DeviceManagementConfigurationSettingDefinitionUxBehaviordefault),
		string(DeviceManagementConfigurationSettingDefinitionUxBehaviordropdown),
		string(DeviceManagementConfigurationSettingDefinitionUxBehaviorlargeTextBox),
		string(DeviceManagementConfigurationSettingDefinitionUxBehaviormultiheaderGrid),
		string(DeviceManagementConfigurationSettingDefinitionUxBehaviorsmallTextBox),
		string(DeviceManagementConfigurationSettingDefinitionUxBehaviortoggle),
	}
}

func parseDeviceManagementConfigurationSettingDefinitionUxBehavior(input string) (*DeviceManagementConfigurationSettingDefinitionUxBehavior, error) {
	vals := map[string]DeviceManagementConfigurationSettingDefinitionUxBehavior{
		"contextpane":     DeviceManagementConfigurationSettingDefinitionUxBehaviorcontextPane,
		"default":         DeviceManagementConfigurationSettingDefinitionUxBehaviordefault,
		"dropdown":        DeviceManagementConfigurationSettingDefinitionUxBehaviordropdown,
		"largetextbox":    DeviceManagementConfigurationSettingDefinitionUxBehaviorlargeTextBox,
		"multiheadergrid": DeviceManagementConfigurationSettingDefinitionUxBehaviormultiheaderGrid,
		"smalltextbox":    DeviceManagementConfigurationSettingDefinitionUxBehaviorsmallTextBox,
		"toggle":          DeviceManagementConfigurationSettingDefinitionUxBehaviortoggle,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DeviceManagementConfigurationSettingDefinitionUxBehavior(input)
	return &out, nil
}
