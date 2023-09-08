package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceManagementConfigurationSimpleSettingDefinitionUxBehavior string

const (
	DeviceManagementConfigurationSimpleSettingDefinitionUxBehaviorcontextPane     DeviceManagementConfigurationSimpleSettingDefinitionUxBehavior = "ContextPane"
	DeviceManagementConfigurationSimpleSettingDefinitionUxBehaviordefault         DeviceManagementConfigurationSimpleSettingDefinitionUxBehavior = "Default"
	DeviceManagementConfigurationSimpleSettingDefinitionUxBehaviordropdown        DeviceManagementConfigurationSimpleSettingDefinitionUxBehavior = "Dropdown"
	DeviceManagementConfigurationSimpleSettingDefinitionUxBehaviorlargeTextBox    DeviceManagementConfigurationSimpleSettingDefinitionUxBehavior = "LargeTextBox"
	DeviceManagementConfigurationSimpleSettingDefinitionUxBehaviormultiheaderGrid DeviceManagementConfigurationSimpleSettingDefinitionUxBehavior = "MultiheaderGrid"
	DeviceManagementConfigurationSimpleSettingDefinitionUxBehaviorsmallTextBox    DeviceManagementConfigurationSimpleSettingDefinitionUxBehavior = "SmallTextBox"
	DeviceManagementConfigurationSimpleSettingDefinitionUxBehaviortoggle          DeviceManagementConfigurationSimpleSettingDefinitionUxBehavior = "Toggle"
)

func PossibleValuesForDeviceManagementConfigurationSimpleSettingDefinitionUxBehavior() []string {
	return []string{
		string(DeviceManagementConfigurationSimpleSettingDefinitionUxBehaviorcontextPane),
		string(DeviceManagementConfigurationSimpleSettingDefinitionUxBehaviordefault),
		string(DeviceManagementConfigurationSimpleSettingDefinitionUxBehaviordropdown),
		string(DeviceManagementConfigurationSimpleSettingDefinitionUxBehaviorlargeTextBox),
		string(DeviceManagementConfigurationSimpleSettingDefinitionUxBehaviormultiheaderGrid),
		string(DeviceManagementConfigurationSimpleSettingDefinitionUxBehaviorsmallTextBox),
		string(DeviceManagementConfigurationSimpleSettingDefinitionUxBehaviortoggle),
	}
}

func parseDeviceManagementConfigurationSimpleSettingDefinitionUxBehavior(input string) (*DeviceManagementConfigurationSimpleSettingDefinitionUxBehavior, error) {
	vals := map[string]DeviceManagementConfigurationSimpleSettingDefinitionUxBehavior{
		"contextpane":     DeviceManagementConfigurationSimpleSettingDefinitionUxBehaviorcontextPane,
		"default":         DeviceManagementConfigurationSimpleSettingDefinitionUxBehaviordefault,
		"dropdown":        DeviceManagementConfigurationSimpleSettingDefinitionUxBehaviordropdown,
		"largetextbox":    DeviceManagementConfigurationSimpleSettingDefinitionUxBehaviorlargeTextBox,
		"multiheadergrid": DeviceManagementConfigurationSimpleSettingDefinitionUxBehaviormultiheaderGrid,
		"smalltextbox":    DeviceManagementConfigurationSimpleSettingDefinitionUxBehaviorsmallTextBox,
		"toggle":          DeviceManagementConfigurationSimpleSettingDefinitionUxBehaviortoggle,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DeviceManagementConfigurationSimpleSettingDefinitionUxBehavior(input)
	return &out, nil
}
