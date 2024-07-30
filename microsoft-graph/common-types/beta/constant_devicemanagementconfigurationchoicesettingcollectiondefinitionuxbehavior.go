package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceManagementConfigurationChoiceSettingCollectionDefinitionUxBehavior string

const (
	DeviceManagementConfigurationChoiceSettingCollectionDefinitionUxBehavior_ContextPane     DeviceManagementConfigurationChoiceSettingCollectionDefinitionUxBehavior = "contextPane"
	DeviceManagementConfigurationChoiceSettingCollectionDefinitionUxBehavior_Default         DeviceManagementConfigurationChoiceSettingCollectionDefinitionUxBehavior = "default"
	DeviceManagementConfigurationChoiceSettingCollectionDefinitionUxBehavior_Dropdown        DeviceManagementConfigurationChoiceSettingCollectionDefinitionUxBehavior = "dropdown"
	DeviceManagementConfigurationChoiceSettingCollectionDefinitionUxBehavior_LargeTextBox    DeviceManagementConfigurationChoiceSettingCollectionDefinitionUxBehavior = "largeTextBox"
	DeviceManagementConfigurationChoiceSettingCollectionDefinitionUxBehavior_MultiheaderGrid DeviceManagementConfigurationChoiceSettingCollectionDefinitionUxBehavior = "multiheaderGrid"
	DeviceManagementConfigurationChoiceSettingCollectionDefinitionUxBehavior_SmallTextBox    DeviceManagementConfigurationChoiceSettingCollectionDefinitionUxBehavior = "smallTextBox"
	DeviceManagementConfigurationChoiceSettingCollectionDefinitionUxBehavior_Toggle          DeviceManagementConfigurationChoiceSettingCollectionDefinitionUxBehavior = "toggle"
)

func PossibleValuesForDeviceManagementConfigurationChoiceSettingCollectionDefinitionUxBehavior() []string {
	return []string{
		string(DeviceManagementConfigurationChoiceSettingCollectionDefinitionUxBehavior_ContextPane),
		string(DeviceManagementConfigurationChoiceSettingCollectionDefinitionUxBehavior_Default),
		string(DeviceManagementConfigurationChoiceSettingCollectionDefinitionUxBehavior_Dropdown),
		string(DeviceManagementConfigurationChoiceSettingCollectionDefinitionUxBehavior_LargeTextBox),
		string(DeviceManagementConfigurationChoiceSettingCollectionDefinitionUxBehavior_MultiheaderGrid),
		string(DeviceManagementConfigurationChoiceSettingCollectionDefinitionUxBehavior_SmallTextBox),
		string(DeviceManagementConfigurationChoiceSettingCollectionDefinitionUxBehavior_Toggle),
	}
}

func (s *DeviceManagementConfigurationChoiceSettingCollectionDefinitionUxBehavior) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseDeviceManagementConfigurationChoiceSettingCollectionDefinitionUxBehavior(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseDeviceManagementConfigurationChoiceSettingCollectionDefinitionUxBehavior(input string) (*DeviceManagementConfigurationChoiceSettingCollectionDefinitionUxBehavior, error) {
	vals := map[string]DeviceManagementConfigurationChoiceSettingCollectionDefinitionUxBehavior{
		"contextpane":     DeviceManagementConfigurationChoiceSettingCollectionDefinitionUxBehavior_ContextPane,
		"default":         DeviceManagementConfigurationChoiceSettingCollectionDefinitionUxBehavior_Default,
		"dropdown":        DeviceManagementConfigurationChoiceSettingCollectionDefinitionUxBehavior_Dropdown,
		"largetextbox":    DeviceManagementConfigurationChoiceSettingCollectionDefinitionUxBehavior_LargeTextBox,
		"multiheadergrid": DeviceManagementConfigurationChoiceSettingCollectionDefinitionUxBehavior_MultiheaderGrid,
		"smalltextbox":    DeviceManagementConfigurationChoiceSettingCollectionDefinitionUxBehavior_SmallTextBox,
		"toggle":          DeviceManagementConfigurationChoiceSettingCollectionDefinitionUxBehavior_Toggle,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DeviceManagementConfigurationChoiceSettingCollectionDefinitionUxBehavior(input)
	return &out, nil
}
