package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceManagementConfigurationSimpleSettingDefinitionUxBehavior string

const (
	DeviceManagementConfigurationSimpleSettingDefinitionUxBehavior_ContextPane     DeviceManagementConfigurationSimpleSettingDefinitionUxBehavior = "contextPane"
	DeviceManagementConfigurationSimpleSettingDefinitionUxBehavior_Default         DeviceManagementConfigurationSimpleSettingDefinitionUxBehavior = "default"
	DeviceManagementConfigurationSimpleSettingDefinitionUxBehavior_Dropdown        DeviceManagementConfigurationSimpleSettingDefinitionUxBehavior = "dropdown"
	DeviceManagementConfigurationSimpleSettingDefinitionUxBehavior_LargeTextBox    DeviceManagementConfigurationSimpleSettingDefinitionUxBehavior = "largeTextBox"
	DeviceManagementConfigurationSimpleSettingDefinitionUxBehavior_MultiheaderGrid DeviceManagementConfigurationSimpleSettingDefinitionUxBehavior = "multiheaderGrid"
	DeviceManagementConfigurationSimpleSettingDefinitionUxBehavior_SmallTextBox    DeviceManagementConfigurationSimpleSettingDefinitionUxBehavior = "smallTextBox"
	DeviceManagementConfigurationSimpleSettingDefinitionUxBehavior_Toggle          DeviceManagementConfigurationSimpleSettingDefinitionUxBehavior = "toggle"
)

func PossibleValuesForDeviceManagementConfigurationSimpleSettingDefinitionUxBehavior() []string {
	return []string{
		string(DeviceManagementConfigurationSimpleSettingDefinitionUxBehavior_ContextPane),
		string(DeviceManagementConfigurationSimpleSettingDefinitionUxBehavior_Default),
		string(DeviceManagementConfigurationSimpleSettingDefinitionUxBehavior_Dropdown),
		string(DeviceManagementConfigurationSimpleSettingDefinitionUxBehavior_LargeTextBox),
		string(DeviceManagementConfigurationSimpleSettingDefinitionUxBehavior_MultiheaderGrid),
		string(DeviceManagementConfigurationSimpleSettingDefinitionUxBehavior_SmallTextBox),
		string(DeviceManagementConfigurationSimpleSettingDefinitionUxBehavior_Toggle),
	}
}

func (s *DeviceManagementConfigurationSimpleSettingDefinitionUxBehavior) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseDeviceManagementConfigurationSimpleSettingDefinitionUxBehavior(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseDeviceManagementConfigurationSimpleSettingDefinitionUxBehavior(input string) (*DeviceManagementConfigurationSimpleSettingDefinitionUxBehavior, error) {
	vals := map[string]DeviceManagementConfigurationSimpleSettingDefinitionUxBehavior{
		"contextpane":     DeviceManagementConfigurationSimpleSettingDefinitionUxBehavior_ContextPane,
		"default":         DeviceManagementConfigurationSimpleSettingDefinitionUxBehavior_Default,
		"dropdown":        DeviceManagementConfigurationSimpleSettingDefinitionUxBehavior_Dropdown,
		"largetextbox":    DeviceManagementConfigurationSimpleSettingDefinitionUxBehavior_LargeTextBox,
		"multiheadergrid": DeviceManagementConfigurationSimpleSettingDefinitionUxBehavior_MultiheaderGrid,
		"smalltextbox":    DeviceManagementConfigurationSimpleSettingDefinitionUxBehavior_SmallTextBox,
		"toggle":          DeviceManagementConfigurationSimpleSettingDefinitionUxBehavior_Toggle,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DeviceManagementConfigurationSimpleSettingDefinitionUxBehavior(input)
	return &out, nil
}
