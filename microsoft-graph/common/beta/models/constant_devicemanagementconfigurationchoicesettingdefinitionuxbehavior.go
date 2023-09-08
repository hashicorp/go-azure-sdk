package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceManagementConfigurationChoiceSettingDefinitionUxBehavior string

const (
	DeviceManagementConfigurationChoiceSettingDefinitionUxBehaviorcontextPane     DeviceManagementConfigurationChoiceSettingDefinitionUxBehavior = "ContextPane"
	DeviceManagementConfigurationChoiceSettingDefinitionUxBehaviordefault         DeviceManagementConfigurationChoiceSettingDefinitionUxBehavior = "Default"
	DeviceManagementConfigurationChoiceSettingDefinitionUxBehaviordropdown        DeviceManagementConfigurationChoiceSettingDefinitionUxBehavior = "Dropdown"
	DeviceManagementConfigurationChoiceSettingDefinitionUxBehaviorlargeTextBox    DeviceManagementConfigurationChoiceSettingDefinitionUxBehavior = "LargeTextBox"
	DeviceManagementConfigurationChoiceSettingDefinitionUxBehaviormultiheaderGrid DeviceManagementConfigurationChoiceSettingDefinitionUxBehavior = "MultiheaderGrid"
	DeviceManagementConfigurationChoiceSettingDefinitionUxBehaviorsmallTextBox    DeviceManagementConfigurationChoiceSettingDefinitionUxBehavior = "SmallTextBox"
	DeviceManagementConfigurationChoiceSettingDefinitionUxBehaviortoggle          DeviceManagementConfigurationChoiceSettingDefinitionUxBehavior = "Toggle"
)

func PossibleValuesForDeviceManagementConfigurationChoiceSettingDefinitionUxBehavior() []string {
	return []string{
		string(DeviceManagementConfigurationChoiceSettingDefinitionUxBehaviorcontextPane),
		string(DeviceManagementConfigurationChoiceSettingDefinitionUxBehaviordefault),
		string(DeviceManagementConfigurationChoiceSettingDefinitionUxBehaviordropdown),
		string(DeviceManagementConfigurationChoiceSettingDefinitionUxBehaviorlargeTextBox),
		string(DeviceManagementConfigurationChoiceSettingDefinitionUxBehaviormultiheaderGrid),
		string(DeviceManagementConfigurationChoiceSettingDefinitionUxBehaviorsmallTextBox),
		string(DeviceManagementConfigurationChoiceSettingDefinitionUxBehaviortoggle),
	}
}

func parseDeviceManagementConfigurationChoiceSettingDefinitionUxBehavior(input string) (*DeviceManagementConfigurationChoiceSettingDefinitionUxBehavior, error) {
	vals := map[string]DeviceManagementConfigurationChoiceSettingDefinitionUxBehavior{
		"contextpane":     DeviceManagementConfigurationChoiceSettingDefinitionUxBehaviorcontextPane,
		"default":         DeviceManagementConfigurationChoiceSettingDefinitionUxBehaviordefault,
		"dropdown":        DeviceManagementConfigurationChoiceSettingDefinitionUxBehaviordropdown,
		"largetextbox":    DeviceManagementConfigurationChoiceSettingDefinitionUxBehaviorlargeTextBox,
		"multiheadergrid": DeviceManagementConfigurationChoiceSettingDefinitionUxBehaviormultiheaderGrid,
		"smalltextbox":    DeviceManagementConfigurationChoiceSettingDefinitionUxBehaviorsmallTextBox,
		"toggle":          DeviceManagementConfigurationChoiceSettingDefinitionUxBehaviortoggle,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DeviceManagementConfigurationChoiceSettingDefinitionUxBehavior(input)
	return &out, nil
}
