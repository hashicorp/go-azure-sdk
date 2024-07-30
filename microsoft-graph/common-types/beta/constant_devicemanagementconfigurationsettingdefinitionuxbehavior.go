package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceManagementConfigurationSettingDefinitionUxBehavior string

const (
	DeviceManagementConfigurationSettingDefinitionUxBehavior_ContextPane     DeviceManagementConfigurationSettingDefinitionUxBehavior = "contextPane"
	DeviceManagementConfigurationSettingDefinitionUxBehavior_Default         DeviceManagementConfigurationSettingDefinitionUxBehavior = "default"
	DeviceManagementConfigurationSettingDefinitionUxBehavior_Dropdown        DeviceManagementConfigurationSettingDefinitionUxBehavior = "dropdown"
	DeviceManagementConfigurationSettingDefinitionUxBehavior_LargeTextBox    DeviceManagementConfigurationSettingDefinitionUxBehavior = "largeTextBox"
	DeviceManagementConfigurationSettingDefinitionUxBehavior_MultiheaderGrid DeviceManagementConfigurationSettingDefinitionUxBehavior = "multiheaderGrid"
	DeviceManagementConfigurationSettingDefinitionUxBehavior_SmallTextBox    DeviceManagementConfigurationSettingDefinitionUxBehavior = "smallTextBox"
	DeviceManagementConfigurationSettingDefinitionUxBehavior_Toggle          DeviceManagementConfigurationSettingDefinitionUxBehavior = "toggle"
)

func PossibleValuesForDeviceManagementConfigurationSettingDefinitionUxBehavior() []string {
	return []string{
		string(DeviceManagementConfigurationSettingDefinitionUxBehavior_ContextPane),
		string(DeviceManagementConfigurationSettingDefinitionUxBehavior_Default),
		string(DeviceManagementConfigurationSettingDefinitionUxBehavior_Dropdown),
		string(DeviceManagementConfigurationSettingDefinitionUxBehavior_LargeTextBox),
		string(DeviceManagementConfigurationSettingDefinitionUxBehavior_MultiheaderGrid),
		string(DeviceManagementConfigurationSettingDefinitionUxBehavior_SmallTextBox),
		string(DeviceManagementConfigurationSettingDefinitionUxBehavior_Toggle),
	}
}

func (s *DeviceManagementConfigurationSettingDefinitionUxBehavior) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseDeviceManagementConfigurationSettingDefinitionUxBehavior(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseDeviceManagementConfigurationSettingDefinitionUxBehavior(input string) (*DeviceManagementConfigurationSettingDefinitionUxBehavior, error) {
	vals := map[string]DeviceManagementConfigurationSettingDefinitionUxBehavior{
		"contextpane":     DeviceManagementConfigurationSettingDefinitionUxBehavior_ContextPane,
		"default":         DeviceManagementConfigurationSettingDefinitionUxBehavior_Default,
		"dropdown":        DeviceManagementConfigurationSettingDefinitionUxBehavior_Dropdown,
		"largetextbox":    DeviceManagementConfigurationSettingDefinitionUxBehavior_LargeTextBox,
		"multiheadergrid": DeviceManagementConfigurationSettingDefinitionUxBehavior_MultiheaderGrid,
		"smalltextbox":    DeviceManagementConfigurationSettingDefinitionUxBehavior_SmallTextBox,
		"toggle":          DeviceManagementConfigurationSettingDefinitionUxBehavior_Toggle,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DeviceManagementConfigurationSettingDefinitionUxBehavior(input)
	return &out, nil
}
