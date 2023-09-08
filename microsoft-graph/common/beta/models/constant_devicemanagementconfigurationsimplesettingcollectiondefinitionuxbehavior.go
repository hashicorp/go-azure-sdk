package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceManagementConfigurationSimpleSettingCollectionDefinitionUxBehavior string

const (
	DeviceManagementConfigurationSimpleSettingCollectionDefinitionUxBehaviorcontextPane     DeviceManagementConfigurationSimpleSettingCollectionDefinitionUxBehavior = "ContextPane"
	DeviceManagementConfigurationSimpleSettingCollectionDefinitionUxBehaviordefault         DeviceManagementConfigurationSimpleSettingCollectionDefinitionUxBehavior = "Default"
	DeviceManagementConfigurationSimpleSettingCollectionDefinitionUxBehaviordropdown        DeviceManagementConfigurationSimpleSettingCollectionDefinitionUxBehavior = "Dropdown"
	DeviceManagementConfigurationSimpleSettingCollectionDefinitionUxBehaviorlargeTextBox    DeviceManagementConfigurationSimpleSettingCollectionDefinitionUxBehavior = "LargeTextBox"
	DeviceManagementConfigurationSimpleSettingCollectionDefinitionUxBehaviormultiheaderGrid DeviceManagementConfigurationSimpleSettingCollectionDefinitionUxBehavior = "MultiheaderGrid"
	DeviceManagementConfigurationSimpleSettingCollectionDefinitionUxBehaviorsmallTextBox    DeviceManagementConfigurationSimpleSettingCollectionDefinitionUxBehavior = "SmallTextBox"
	DeviceManagementConfigurationSimpleSettingCollectionDefinitionUxBehaviortoggle          DeviceManagementConfigurationSimpleSettingCollectionDefinitionUxBehavior = "Toggle"
)

func PossibleValuesForDeviceManagementConfigurationSimpleSettingCollectionDefinitionUxBehavior() []string {
	return []string{
		string(DeviceManagementConfigurationSimpleSettingCollectionDefinitionUxBehaviorcontextPane),
		string(DeviceManagementConfigurationSimpleSettingCollectionDefinitionUxBehaviordefault),
		string(DeviceManagementConfigurationSimpleSettingCollectionDefinitionUxBehaviordropdown),
		string(DeviceManagementConfigurationSimpleSettingCollectionDefinitionUxBehaviorlargeTextBox),
		string(DeviceManagementConfigurationSimpleSettingCollectionDefinitionUxBehaviormultiheaderGrid),
		string(DeviceManagementConfigurationSimpleSettingCollectionDefinitionUxBehaviorsmallTextBox),
		string(DeviceManagementConfigurationSimpleSettingCollectionDefinitionUxBehaviortoggle),
	}
}

func parseDeviceManagementConfigurationSimpleSettingCollectionDefinitionUxBehavior(input string) (*DeviceManagementConfigurationSimpleSettingCollectionDefinitionUxBehavior, error) {
	vals := map[string]DeviceManagementConfigurationSimpleSettingCollectionDefinitionUxBehavior{
		"contextpane":     DeviceManagementConfigurationSimpleSettingCollectionDefinitionUxBehaviorcontextPane,
		"default":         DeviceManagementConfigurationSimpleSettingCollectionDefinitionUxBehaviordefault,
		"dropdown":        DeviceManagementConfigurationSimpleSettingCollectionDefinitionUxBehaviordropdown,
		"largetextbox":    DeviceManagementConfigurationSimpleSettingCollectionDefinitionUxBehaviorlargeTextBox,
		"multiheadergrid": DeviceManagementConfigurationSimpleSettingCollectionDefinitionUxBehaviormultiheaderGrid,
		"smalltextbox":    DeviceManagementConfigurationSimpleSettingCollectionDefinitionUxBehaviorsmallTextBox,
		"toggle":          DeviceManagementConfigurationSimpleSettingCollectionDefinitionUxBehaviortoggle,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DeviceManagementConfigurationSimpleSettingCollectionDefinitionUxBehavior(input)
	return &out, nil
}
