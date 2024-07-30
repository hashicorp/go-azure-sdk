package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceManagementConfigurationSettingGroupDefinitionUxBehavior string

const (
	DeviceManagementConfigurationSettingGroupDefinitionUxBehavior_ContextPane     DeviceManagementConfigurationSettingGroupDefinitionUxBehavior = "contextPane"
	DeviceManagementConfigurationSettingGroupDefinitionUxBehavior_Default         DeviceManagementConfigurationSettingGroupDefinitionUxBehavior = "default"
	DeviceManagementConfigurationSettingGroupDefinitionUxBehavior_Dropdown        DeviceManagementConfigurationSettingGroupDefinitionUxBehavior = "dropdown"
	DeviceManagementConfigurationSettingGroupDefinitionUxBehavior_LargeTextBox    DeviceManagementConfigurationSettingGroupDefinitionUxBehavior = "largeTextBox"
	DeviceManagementConfigurationSettingGroupDefinitionUxBehavior_MultiheaderGrid DeviceManagementConfigurationSettingGroupDefinitionUxBehavior = "multiheaderGrid"
	DeviceManagementConfigurationSettingGroupDefinitionUxBehavior_SmallTextBox    DeviceManagementConfigurationSettingGroupDefinitionUxBehavior = "smallTextBox"
	DeviceManagementConfigurationSettingGroupDefinitionUxBehavior_Toggle          DeviceManagementConfigurationSettingGroupDefinitionUxBehavior = "toggle"
)

func PossibleValuesForDeviceManagementConfigurationSettingGroupDefinitionUxBehavior() []string {
	return []string{
		string(DeviceManagementConfigurationSettingGroupDefinitionUxBehavior_ContextPane),
		string(DeviceManagementConfigurationSettingGroupDefinitionUxBehavior_Default),
		string(DeviceManagementConfigurationSettingGroupDefinitionUxBehavior_Dropdown),
		string(DeviceManagementConfigurationSettingGroupDefinitionUxBehavior_LargeTextBox),
		string(DeviceManagementConfigurationSettingGroupDefinitionUxBehavior_MultiheaderGrid),
		string(DeviceManagementConfigurationSettingGroupDefinitionUxBehavior_SmallTextBox),
		string(DeviceManagementConfigurationSettingGroupDefinitionUxBehavior_Toggle),
	}
}

func (s *DeviceManagementConfigurationSettingGroupDefinitionUxBehavior) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseDeviceManagementConfigurationSettingGroupDefinitionUxBehavior(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseDeviceManagementConfigurationSettingGroupDefinitionUxBehavior(input string) (*DeviceManagementConfigurationSettingGroupDefinitionUxBehavior, error) {
	vals := map[string]DeviceManagementConfigurationSettingGroupDefinitionUxBehavior{
		"contextpane":     DeviceManagementConfigurationSettingGroupDefinitionUxBehavior_ContextPane,
		"default":         DeviceManagementConfigurationSettingGroupDefinitionUxBehavior_Default,
		"dropdown":        DeviceManagementConfigurationSettingGroupDefinitionUxBehavior_Dropdown,
		"largetextbox":    DeviceManagementConfigurationSettingGroupDefinitionUxBehavior_LargeTextBox,
		"multiheadergrid": DeviceManagementConfigurationSettingGroupDefinitionUxBehavior_MultiheaderGrid,
		"smalltextbox":    DeviceManagementConfigurationSettingGroupDefinitionUxBehavior_SmallTextBox,
		"toggle":          DeviceManagementConfigurationSettingGroupDefinitionUxBehavior_Toggle,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DeviceManagementConfigurationSettingGroupDefinitionUxBehavior(input)
	return &out, nil
}
