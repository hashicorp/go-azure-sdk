package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceManagementConfigurationSettingGroupCollectionDefinitionUxBehavior string

const (
	DeviceManagementConfigurationSettingGroupCollectionDefinitionUxBehavior_ContextPane     DeviceManagementConfigurationSettingGroupCollectionDefinitionUxBehavior = "contextPane"
	DeviceManagementConfigurationSettingGroupCollectionDefinitionUxBehavior_Default         DeviceManagementConfigurationSettingGroupCollectionDefinitionUxBehavior = "default"
	DeviceManagementConfigurationSettingGroupCollectionDefinitionUxBehavior_Dropdown        DeviceManagementConfigurationSettingGroupCollectionDefinitionUxBehavior = "dropdown"
	DeviceManagementConfigurationSettingGroupCollectionDefinitionUxBehavior_LargeTextBox    DeviceManagementConfigurationSettingGroupCollectionDefinitionUxBehavior = "largeTextBox"
	DeviceManagementConfigurationSettingGroupCollectionDefinitionUxBehavior_MultiheaderGrid DeviceManagementConfigurationSettingGroupCollectionDefinitionUxBehavior = "multiheaderGrid"
	DeviceManagementConfigurationSettingGroupCollectionDefinitionUxBehavior_SmallTextBox    DeviceManagementConfigurationSettingGroupCollectionDefinitionUxBehavior = "smallTextBox"
	DeviceManagementConfigurationSettingGroupCollectionDefinitionUxBehavior_Toggle          DeviceManagementConfigurationSettingGroupCollectionDefinitionUxBehavior = "toggle"
)

func PossibleValuesForDeviceManagementConfigurationSettingGroupCollectionDefinitionUxBehavior() []string {
	return []string{
		string(DeviceManagementConfigurationSettingGroupCollectionDefinitionUxBehavior_ContextPane),
		string(DeviceManagementConfigurationSettingGroupCollectionDefinitionUxBehavior_Default),
		string(DeviceManagementConfigurationSettingGroupCollectionDefinitionUxBehavior_Dropdown),
		string(DeviceManagementConfigurationSettingGroupCollectionDefinitionUxBehavior_LargeTextBox),
		string(DeviceManagementConfigurationSettingGroupCollectionDefinitionUxBehavior_MultiheaderGrid),
		string(DeviceManagementConfigurationSettingGroupCollectionDefinitionUxBehavior_SmallTextBox),
		string(DeviceManagementConfigurationSettingGroupCollectionDefinitionUxBehavior_Toggle),
	}
}

func (s *DeviceManagementConfigurationSettingGroupCollectionDefinitionUxBehavior) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseDeviceManagementConfigurationSettingGroupCollectionDefinitionUxBehavior(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseDeviceManagementConfigurationSettingGroupCollectionDefinitionUxBehavior(input string) (*DeviceManagementConfigurationSettingGroupCollectionDefinitionUxBehavior, error) {
	vals := map[string]DeviceManagementConfigurationSettingGroupCollectionDefinitionUxBehavior{
		"contextpane":     DeviceManagementConfigurationSettingGroupCollectionDefinitionUxBehavior_ContextPane,
		"default":         DeviceManagementConfigurationSettingGroupCollectionDefinitionUxBehavior_Default,
		"dropdown":        DeviceManagementConfigurationSettingGroupCollectionDefinitionUxBehavior_Dropdown,
		"largetextbox":    DeviceManagementConfigurationSettingGroupCollectionDefinitionUxBehavior_LargeTextBox,
		"multiheadergrid": DeviceManagementConfigurationSettingGroupCollectionDefinitionUxBehavior_MultiheaderGrid,
		"smalltextbox":    DeviceManagementConfigurationSettingGroupCollectionDefinitionUxBehavior_SmallTextBox,
		"toggle":          DeviceManagementConfigurationSettingGroupCollectionDefinitionUxBehavior_Toggle,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DeviceManagementConfigurationSettingGroupCollectionDefinitionUxBehavior(input)
	return &out, nil
}
