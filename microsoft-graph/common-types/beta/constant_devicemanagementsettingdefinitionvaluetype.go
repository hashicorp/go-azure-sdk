package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceManagementSettingDefinitionValueType string

const (
	DeviceManagementSettingDefinitionValueType_AbstractComplex DeviceManagementSettingDefinitionValueType = "abstractComplex"
	DeviceManagementSettingDefinitionValueType_Boolean         DeviceManagementSettingDefinitionValueType = "boolean"
	DeviceManagementSettingDefinitionValueType_Collection      DeviceManagementSettingDefinitionValueType = "collection"
	DeviceManagementSettingDefinitionValueType_Complex         DeviceManagementSettingDefinitionValueType = "complex"
	DeviceManagementSettingDefinitionValueType_Integer         DeviceManagementSettingDefinitionValueType = "integer"
	DeviceManagementSettingDefinitionValueType_String          DeviceManagementSettingDefinitionValueType = "string"
)

func PossibleValuesForDeviceManagementSettingDefinitionValueType() []string {
	return []string{
		string(DeviceManagementSettingDefinitionValueType_AbstractComplex),
		string(DeviceManagementSettingDefinitionValueType_Boolean),
		string(DeviceManagementSettingDefinitionValueType_Collection),
		string(DeviceManagementSettingDefinitionValueType_Complex),
		string(DeviceManagementSettingDefinitionValueType_Integer),
		string(DeviceManagementSettingDefinitionValueType_String),
	}
}

func (s *DeviceManagementSettingDefinitionValueType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseDeviceManagementSettingDefinitionValueType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseDeviceManagementSettingDefinitionValueType(input string) (*DeviceManagementSettingDefinitionValueType, error) {
	vals := map[string]DeviceManagementSettingDefinitionValueType{
		"abstractcomplex": DeviceManagementSettingDefinitionValueType_AbstractComplex,
		"boolean":         DeviceManagementSettingDefinitionValueType_Boolean,
		"collection":      DeviceManagementSettingDefinitionValueType_Collection,
		"complex":         DeviceManagementSettingDefinitionValueType_Complex,
		"integer":         DeviceManagementSettingDefinitionValueType_Integer,
		"string":          DeviceManagementSettingDefinitionValueType_String,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DeviceManagementSettingDefinitionValueType(input)
	return &out, nil
}
