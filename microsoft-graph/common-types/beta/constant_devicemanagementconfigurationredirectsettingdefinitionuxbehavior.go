package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceManagementConfigurationRedirectSettingDefinitionUxBehavior string

const (
	DeviceManagementConfigurationRedirectSettingDefinitionUxBehavior_ContextPane     DeviceManagementConfigurationRedirectSettingDefinitionUxBehavior = "contextPane"
	DeviceManagementConfigurationRedirectSettingDefinitionUxBehavior_Default         DeviceManagementConfigurationRedirectSettingDefinitionUxBehavior = "default"
	DeviceManagementConfigurationRedirectSettingDefinitionUxBehavior_Dropdown        DeviceManagementConfigurationRedirectSettingDefinitionUxBehavior = "dropdown"
	DeviceManagementConfigurationRedirectSettingDefinitionUxBehavior_LargeTextBox    DeviceManagementConfigurationRedirectSettingDefinitionUxBehavior = "largeTextBox"
	DeviceManagementConfigurationRedirectSettingDefinitionUxBehavior_MultiheaderGrid DeviceManagementConfigurationRedirectSettingDefinitionUxBehavior = "multiheaderGrid"
	DeviceManagementConfigurationRedirectSettingDefinitionUxBehavior_SmallTextBox    DeviceManagementConfigurationRedirectSettingDefinitionUxBehavior = "smallTextBox"
	DeviceManagementConfigurationRedirectSettingDefinitionUxBehavior_Toggle          DeviceManagementConfigurationRedirectSettingDefinitionUxBehavior = "toggle"
)

func PossibleValuesForDeviceManagementConfigurationRedirectSettingDefinitionUxBehavior() []string {
	return []string{
		string(DeviceManagementConfigurationRedirectSettingDefinitionUxBehavior_ContextPane),
		string(DeviceManagementConfigurationRedirectSettingDefinitionUxBehavior_Default),
		string(DeviceManagementConfigurationRedirectSettingDefinitionUxBehavior_Dropdown),
		string(DeviceManagementConfigurationRedirectSettingDefinitionUxBehavior_LargeTextBox),
		string(DeviceManagementConfigurationRedirectSettingDefinitionUxBehavior_MultiheaderGrid),
		string(DeviceManagementConfigurationRedirectSettingDefinitionUxBehavior_SmallTextBox),
		string(DeviceManagementConfigurationRedirectSettingDefinitionUxBehavior_Toggle),
	}
}

func (s *DeviceManagementConfigurationRedirectSettingDefinitionUxBehavior) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseDeviceManagementConfigurationRedirectSettingDefinitionUxBehavior(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseDeviceManagementConfigurationRedirectSettingDefinitionUxBehavior(input string) (*DeviceManagementConfigurationRedirectSettingDefinitionUxBehavior, error) {
	vals := map[string]DeviceManagementConfigurationRedirectSettingDefinitionUxBehavior{
		"contextpane":     DeviceManagementConfigurationRedirectSettingDefinitionUxBehavior_ContextPane,
		"default":         DeviceManagementConfigurationRedirectSettingDefinitionUxBehavior_Default,
		"dropdown":        DeviceManagementConfigurationRedirectSettingDefinitionUxBehavior_Dropdown,
		"largetextbox":    DeviceManagementConfigurationRedirectSettingDefinitionUxBehavior_LargeTextBox,
		"multiheadergrid": DeviceManagementConfigurationRedirectSettingDefinitionUxBehavior_MultiheaderGrid,
		"smalltextbox":    DeviceManagementConfigurationRedirectSettingDefinitionUxBehavior_SmallTextBox,
		"toggle":          DeviceManagementConfigurationRedirectSettingDefinitionUxBehavior_Toggle,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DeviceManagementConfigurationRedirectSettingDefinitionUxBehavior(input)
	return &out, nil
}
