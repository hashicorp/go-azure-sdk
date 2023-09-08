package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceManagementConfigurationChoiceSettingCollectionDefinitionUxBehavior string

const (
	DeviceManagementConfigurationChoiceSettingCollectionDefinitionUxBehaviorcontextPane     DeviceManagementConfigurationChoiceSettingCollectionDefinitionUxBehavior = "ContextPane"
	DeviceManagementConfigurationChoiceSettingCollectionDefinitionUxBehaviordefault         DeviceManagementConfigurationChoiceSettingCollectionDefinitionUxBehavior = "Default"
	DeviceManagementConfigurationChoiceSettingCollectionDefinitionUxBehaviordropdown        DeviceManagementConfigurationChoiceSettingCollectionDefinitionUxBehavior = "Dropdown"
	DeviceManagementConfigurationChoiceSettingCollectionDefinitionUxBehaviorlargeTextBox    DeviceManagementConfigurationChoiceSettingCollectionDefinitionUxBehavior = "LargeTextBox"
	DeviceManagementConfigurationChoiceSettingCollectionDefinitionUxBehaviormultiheaderGrid DeviceManagementConfigurationChoiceSettingCollectionDefinitionUxBehavior = "MultiheaderGrid"
	DeviceManagementConfigurationChoiceSettingCollectionDefinitionUxBehaviorsmallTextBox    DeviceManagementConfigurationChoiceSettingCollectionDefinitionUxBehavior = "SmallTextBox"
	DeviceManagementConfigurationChoiceSettingCollectionDefinitionUxBehaviortoggle          DeviceManagementConfigurationChoiceSettingCollectionDefinitionUxBehavior = "Toggle"
)

func PossibleValuesForDeviceManagementConfigurationChoiceSettingCollectionDefinitionUxBehavior() []string {
	return []string{
		string(DeviceManagementConfigurationChoiceSettingCollectionDefinitionUxBehaviorcontextPane),
		string(DeviceManagementConfigurationChoiceSettingCollectionDefinitionUxBehaviordefault),
		string(DeviceManagementConfigurationChoiceSettingCollectionDefinitionUxBehaviordropdown),
		string(DeviceManagementConfigurationChoiceSettingCollectionDefinitionUxBehaviorlargeTextBox),
		string(DeviceManagementConfigurationChoiceSettingCollectionDefinitionUxBehaviormultiheaderGrid),
		string(DeviceManagementConfigurationChoiceSettingCollectionDefinitionUxBehaviorsmallTextBox),
		string(DeviceManagementConfigurationChoiceSettingCollectionDefinitionUxBehaviortoggle),
	}
}

func parseDeviceManagementConfigurationChoiceSettingCollectionDefinitionUxBehavior(input string) (*DeviceManagementConfigurationChoiceSettingCollectionDefinitionUxBehavior, error) {
	vals := map[string]DeviceManagementConfigurationChoiceSettingCollectionDefinitionUxBehavior{
		"contextpane":     DeviceManagementConfigurationChoiceSettingCollectionDefinitionUxBehaviorcontextPane,
		"default":         DeviceManagementConfigurationChoiceSettingCollectionDefinitionUxBehaviordefault,
		"dropdown":        DeviceManagementConfigurationChoiceSettingCollectionDefinitionUxBehaviordropdown,
		"largetextbox":    DeviceManagementConfigurationChoiceSettingCollectionDefinitionUxBehaviorlargeTextBox,
		"multiheadergrid": DeviceManagementConfigurationChoiceSettingCollectionDefinitionUxBehaviormultiheaderGrid,
		"smalltextbox":    DeviceManagementConfigurationChoiceSettingCollectionDefinitionUxBehaviorsmallTextBox,
		"toggle":          DeviceManagementConfigurationChoiceSettingCollectionDefinitionUxBehaviortoggle,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DeviceManagementConfigurationChoiceSettingCollectionDefinitionUxBehavior(input)
	return &out, nil
}
