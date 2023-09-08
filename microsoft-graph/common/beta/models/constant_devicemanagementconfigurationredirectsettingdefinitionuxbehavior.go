package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceManagementConfigurationRedirectSettingDefinitionUxBehavior string

const (
	DeviceManagementConfigurationRedirectSettingDefinitionUxBehaviorcontextPane     DeviceManagementConfigurationRedirectSettingDefinitionUxBehavior = "ContextPane"
	DeviceManagementConfigurationRedirectSettingDefinitionUxBehaviordefault         DeviceManagementConfigurationRedirectSettingDefinitionUxBehavior = "Default"
	DeviceManagementConfigurationRedirectSettingDefinitionUxBehaviordropdown        DeviceManagementConfigurationRedirectSettingDefinitionUxBehavior = "Dropdown"
	DeviceManagementConfigurationRedirectSettingDefinitionUxBehaviorlargeTextBox    DeviceManagementConfigurationRedirectSettingDefinitionUxBehavior = "LargeTextBox"
	DeviceManagementConfigurationRedirectSettingDefinitionUxBehaviormultiheaderGrid DeviceManagementConfigurationRedirectSettingDefinitionUxBehavior = "MultiheaderGrid"
	DeviceManagementConfigurationRedirectSettingDefinitionUxBehaviorsmallTextBox    DeviceManagementConfigurationRedirectSettingDefinitionUxBehavior = "SmallTextBox"
	DeviceManagementConfigurationRedirectSettingDefinitionUxBehaviortoggle          DeviceManagementConfigurationRedirectSettingDefinitionUxBehavior = "Toggle"
)

func PossibleValuesForDeviceManagementConfigurationRedirectSettingDefinitionUxBehavior() []string {
	return []string{
		string(DeviceManagementConfigurationRedirectSettingDefinitionUxBehaviorcontextPane),
		string(DeviceManagementConfigurationRedirectSettingDefinitionUxBehaviordefault),
		string(DeviceManagementConfigurationRedirectSettingDefinitionUxBehaviordropdown),
		string(DeviceManagementConfigurationRedirectSettingDefinitionUxBehaviorlargeTextBox),
		string(DeviceManagementConfigurationRedirectSettingDefinitionUxBehaviormultiheaderGrid),
		string(DeviceManagementConfigurationRedirectSettingDefinitionUxBehaviorsmallTextBox),
		string(DeviceManagementConfigurationRedirectSettingDefinitionUxBehaviortoggle),
	}
}

func parseDeviceManagementConfigurationRedirectSettingDefinitionUxBehavior(input string) (*DeviceManagementConfigurationRedirectSettingDefinitionUxBehavior, error) {
	vals := map[string]DeviceManagementConfigurationRedirectSettingDefinitionUxBehavior{
		"contextpane":     DeviceManagementConfigurationRedirectSettingDefinitionUxBehaviorcontextPane,
		"default":         DeviceManagementConfigurationRedirectSettingDefinitionUxBehaviordefault,
		"dropdown":        DeviceManagementConfigurationRedirectSettingDefinitionUxBehaviordropdown,
		"largetextbox":    DeviceManagementConfigurationRedirectSettingDefinitionUxBehaviorlargeTextBox,
		"multiheadergrid": DeviceManagementConfigurationRedirectSettingDefinitionUxBehaviormultiheaderGrid,
		"smalltextbox":    DeviceManagementConfigurationRedirectSettingDefinitionUxBehaviorsmallTextBox,
		"toggle":          DeviceManagementConfigurationRedirectSettingDefinitionUxBehaviortoggle,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DeviceManagementConfigurationRedirectSettingDefinitionUxBehavior(input)
	return &out, nil
}
