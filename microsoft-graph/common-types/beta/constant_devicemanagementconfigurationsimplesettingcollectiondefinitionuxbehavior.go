package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceManagementConfigurationSimpleSettingCollectionDefinitionUxBehavior string

const (
	DeviceManagementConfigurationSimpleSettingCollectionDefinitionUxBehavior_ContextPane     DeviceManagementConfigurationSimpleSettingCollectionDefinitionUxBehavior = "contextPane"
	DeviceManagementConfigurationSimpleSettingCollectionDefinitionUxBehavior_Default         DeviceManagementConfigurationSimpleSettingCollectionDefinitionUxBehavior = "default"
	DeviceManagementConfigurationSimpleSettingCollectionDefinitionUxBehavior_Dropdown        DeviceManagementConfigurationSimpleSettingCollectionDefinitionUxBehavior = "dropdown"
	DeviceManagementConfigurationSimpleSettingCollectionDefinitionUxBehavior_LargeTextBox    DeviceManagementConfigurationSimpleSettingCollectionDefinitionUxBehavior = "largeTextBox"
	DeviceManagementConfigurationSimpleSettingCollectionDefinitionUxBehavior_MultiheaderGrid DeviceManagementConfigurationSimpleSettingCollectionDefinitionUxBehavior = "multiheaderGrid"
	DeviceManagementConfigurationSimpleSettingCollectionDefinitionUxBehavior_SmallTextBox    DeviceManagementConfigurationSimpleSettingCollectionDefinitionUxBehavior = "smallTextBox"
	DeviceManagementConfigurationSimpleSettingCollectionDefinitionUxBehavior_Toggle          DeviceManagementConfigurationSimpleSettingCollectionDefinitionUxBehavior = "toggle"
)

func PossibleValuesForDeviceManagementConfigurationSimpleSettingCollectionDefinitionUxBehavior() []string {
	return []string{
		string(DeviceManagementConfigurationSimpleSettingCollectionDefinitionUxBehavior_ContextPane),
		string(DeviceManagementConfigurationSimpleSettingCollectionDefinitionUxBehavior_Default),
		string(DeviceManagementConfigurationSimpleSettingCollectionDefinitionUxBehavior_Dropdown),
		string(DeviceManagementConfigurationSimpleSettingCollectionDefinitionUxBehavior_LargeTextBox),
		string(DeviceManagementConfigurationSimpleSettingCollectionDefinitionUxBehavior_MultiheaderGrid),
		string(DeviceManagementConfigurationSimpleSettingCollectionDefinitionUxBehavior_SmallTextBox),
		string(DeviceManagementConfigurationSimpleSettingCollectionDefinitionUxBehavior_Toggle),
	}
}

func (s *DeviceManagementConfigurationSimpleSettingCollectionDefinitionUxBehavior) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseDeviceManagementConfigurationSimpleSettingCollectionDefinitionUxBehavior(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseDeviceManagementConfigurationSimpleSettingCollectionDefinitionUxBehavior(input string) (*DeviceManagementConfigurationSimpleSettingCollectionDefinitionUxBehavior, error) {
	vals := map[string]DeviceManagementConfigurationSimpleSettingCollectionDefinitionUxBehavior{
		"contextpane":     DeviceManagementConfigurationSimpleSettingCollectionDefinitionUxBehavior_ContextPane,
		"default":         DeviceManagementConfigurationSimpleSettingCollectionDefinitionUxBehavior_Default,
		"dropdown":        DeviceManagementConfigurationSimpleSettingCollectionDefinitionUxBehavior_Dropdown,
		"largetextbox":    DeviceManagementConfigurationSimpleSettingCollectionDefinitionUxBehavior_LargeTextBox,
		"multiheadergrid": DeviceManagementConfigurationSimpleSettingCollectionDefinitionUxBehavior_MultiheaderGrid,
		"smalltextbox":    DeviceManagementConfigurationSimpleSettingCollectionDefinitionUxBehavior_SmallTextBox,
		"toggle":          DeviceManagementConfigurationSimpleSettingCollectionDefinitionUxBehavior_Toggle,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DeviceManagementConfigurationSimpleSettingCollectionDefinitionUxBehavior(input)
	return &out, nil
}
