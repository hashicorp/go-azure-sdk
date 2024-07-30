package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceManagementConfigurationChoiceSettingDefinitionUxBehavior string

const (
	DeviceManagementConfigurationChoiceSettingDefinitionUxBehavior_ContextPane     DeviceManagementConfigurationChoiceSettingDefinitionUxBehavior = "contextPane"
	DeviceManagementConfigurationChoiceSettingDefinitionUxBehavior_Default         DeviceManagementConfigurationChoiceSettingDefinitionUxBehavior = "default"
	DeviceManagementConfigurationChoiceSettingDefinitionUxBehavior_Dropdown        DeviceManagementConfigurationChoiceSettingDefinitionUxBehavior = "dropdown"
	DeviceManagementConfigurationChoiceSettingDefinitionUxBehavior_LargeTextBox    DeviceManagementConfigurationChoiceSettingDefinitionUxBehavior = "largeTextBox"
	DeviceManagementConfigurationChoiceSettingDefinitionUxBehavior_MultiheaderGrid DeviceManagementConfigurationChoiceSettingDefinitionUxBehavior = "multiheaderGrid"
	DeviceManagementConfigurationChoiceSettingDefinitionUxBehavior_SmallTextBox    DeviceManagementConfigurationChoiceSettingDefinitionUxBehavior = "smallTextBox"
	DeviceManagementConfigurationChoiceSettingDefinitionUxBehavior_Toggle          DeviceManagementConfigurationChoiceSettingDefinitionUxBehavior = "toggle"
)

func PossibleValuesForDeviceManagementConfigurationChoiceSettingDefinitionUxBehavior() []string {
	return []string{
		string(DeviceManagementConfigurationChoiceSettingDefinitionUxBehavior_ContextPane),
		string(DeviceManagementConfigurationChoiceSettingDefinitionUxBehavior_Default),
		string(DeviceManagementConfigurationChoiceSettingDefinitionUxBehavior_Dropdown),
		string(DeviceManagementConfigurationChoiceSettingDefinitionUxBehavior_LargeTextBox),
		string(DeviceManagementConfigurationChoiceSettingDefinitionUxBehavior_MultiheaderGrid),
		string(DeviceManagementConfigurationChoiceSettingDefinitionUxBehavior_SmallTextBox),
		string(DeviceManagementConfigurationChoiceSettingDefinitionUxBehavior_Toggle),
	}
}

func (s *DeviceManagementConfigurationChoiceSettingDefinitionUxBehavior) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseDeviceManagementConfigurationChoiceSettingDefinitionUxBehavior(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseDeviceManagementConfigurationChoiceSettingDefinitionUxBehavior(input string) (*DeviceManagementConfigurationChoiceSettingDefinitionUxBehavior, error) {
	vals := map[string]DeviceManagementConfigurationChoiceSettingDefinitionUxBehavior{
		"contextpane":     DeviceManagementConfigurationChoiceSettingDefinitionUxBehavior_ContextPane,
		"default":         DeviceManagementConfigurationChoiceSettingDefinitionUxBehavior_Default,
		"dropdown":        DeviceManagementConfigurationChoiceSettingDefinitionUxBehavior_Dropdown,
		"largetextbox":    DeviceManagementConfigurationChoiceSettingDefinitionUxBehavior_LargeTextBox,
		"multiheadergrid": DeviceManagementConfigurationChoiceSettingDefinitionUxBehavior_MultiheaderGrid,
		"smalltextbox":    DeviceManagementConfigurationChoiceSettingDefinitionUxBehavior_SmallTextBox,
		"toggle":          DeviceManagementConfigurationChoiceSettingDefinitionUxBehavior_Toggle,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DeviceManagementConfigurationChoiceSettingDefinitionUxBehavior(input)
	return &out, nil
}
