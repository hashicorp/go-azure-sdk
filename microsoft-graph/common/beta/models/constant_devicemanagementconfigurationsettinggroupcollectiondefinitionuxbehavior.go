package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceManagementConfigurationSettingGroupCollectionDefinitionUxBehavior string

const (
	DeviceManagementConfigurationSettingGroupCollectionDefinitionUxBehaviorcontextPane     DeviceManagementConfigurationSettingGroupCollectionDefinitionUxBehavior = "ContextPane"
	DeviceManagementConfigurationSettingGroupCollectionDefinitionUxBehaviordefault         DeviceManagementConfigurationSettingGroupCollectionDefinitionUxBehavior = "Default"
	DeviceManagementConfigurationSettingGroupCollectionDefinitionUxBehaviordropdown        DeviceManagementConfigurationSettingGroupCollectionDefinitionUxBehavior = "Dropdown"
	DeviceManagementConfigurationSettingGroupCollectionDefinitionUxBehaviorlargeTextBox    DeviceManagementConfigurationSettingGroupCollectionDefinitionUxBehavior = "LargeTextBox"
	DeviceManagementConfigurationSettingGroupCollectionDefinitionUxBehaviormultiheaderGrid DeviceManagementConfigurationSettingGroupCollectionDefinitionUxBehavior = "MultiheaderGrid"
	DeviceManagementConfigurationSettingGroupCollectionDefinitionUxBehaviorsmallTextBox    DeviceManagementConfigurationSettingGroupCollectionDefinitionUxBehavior = "SmallTextBox"
	DeviceManagementConfigurationSettingGroupCollectionDefinitionUxBehaviortoggle          DeviceManagementConfigurationSettingGroupCollectionDefinitionUxBehavior = "Toggle"
)

func PossibleValuesForDeviceManagementConfigurationSettingGroupCollectionDefinitionUxBehavior() []string {
	return []string{
		string(DeviceManagementConfigurationSettingGroupCollectionDefinitionUxBehaviorcontextPane),
		string(DeviceManagementConfigurationSettingGroupCollectionDefinitionUxBehaviordefault),
		string(DeviceManagementConfigurationSettingGroupCollectionDefinitionUxBehaviordropdown),
		string(DeviceManagementConfigurationSettingGroupCollectionDefinitionUxBehaviorlargeTextBox),
		string(DeviceManagementConfigurationSettingGroupCollectionDefinitionUxBehaviormultiheaderGrid),
		string(DeviceManagementConfigurationSettingGroupCollectionDefinitionUxBehaviorsmallTextBox),
		string(DeviceManagementConfigurationSettingGroupCollectionDefinitionUxBehaviortoggle),
	}
}

func parseDeviceManagementConfigurationSettingGroupCollectionDefinitionUxBehavior(input string) (*DeviceManagementConfigurationSettingGroupCollectionDefinitionUxBehavior, error) {
	vals := map[string]DeviceManagementConfigurationSettingGroupCollectionDefinitionUxBehavior{
		"contextpane":     DeviceManagementConfigurationSettingGroupCollectionDefinitionUxBehaviorcontextPane,
		"default":         DeviceManagementConfigurationSettingGroupCollectionDefinitionUxBehaviordefault,
		"dropdown":        DeviceManagementConfigurationSettingGroupCollectionDefinitionUxBehaviordropdown,
		"largetextbox":    DeviceManagementConfigurationSettingGroupCollectionDefinitionUxBehaviorlargeTextBox,
		"multiheadergrid": DeviceManagementConfigurationSettingGroupCollectionDefinitionUxBehaviormultiheaderGrid,
		"smalltextbox":    DeviceManagementConfigurationSettingGroupCollectionDefinitionUxBehaviorsmallTextBox,
		"toggle":          DeviceManagementConfigurationSettingGroupCollectionDefinitionUxBehaviortoggle,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DeviceManagementConfigurationSettingGroupCollectionDefinitionUxBehavior(input)
	return &out, nil
}
