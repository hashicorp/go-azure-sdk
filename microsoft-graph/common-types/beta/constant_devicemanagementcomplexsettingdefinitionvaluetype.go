package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceManagementComplexSettingDefinitionValueType string

const (
	DeviceManagementComplexSettingDefinitionValueType_AbstractComplex DeviceManagementComplexSettingDefinitionValueType = "abstractComplex"
	DeviceManagementComplexSettingDefinitionValueType_Boolean         DeviceManagementComplexSettingDefinitionValueType = "boolean"
	DeviceManagementComplexSettingDefinitionValueType_Collection      DeviceManagementComplexSettingDefinitionValueType = "collection"
	DeviceManagementComplexSettingDefinitionValueType_Complex         DeviceManagementComplexSettingDefinitionValueType = "complex"
	DeviceManagementComplexSettingDefinitionValueType_Integer         DeviceManagementComplexSettingDefinitionValueType = "integer"
	DeviceManagementComplexSettingDefinitionValueType_String          DeviceManagementComplexSettingDefinitionValueType = "string"
)

func PossibleValuesForDeviceManagementComplexSettingDefinitionValueType() []string {
	return []string{
		string(DeviceManagementComplexSettingDefinitionValueType_AbstractComplex),
		string(DeviceManagementComplexSettingDefinitionValueType_Boolean),
		string(DeviceManagementComplexSettingDefinitionValueType_Collection),
		string(DeviceManagementComplexSettingDefinitionValueType_Complex),
		string(DeviceManagementComplexSettingDefinitionValueType_Integer),
		string(DeviceManagementComplexSettingDefinitionValueType_String),
	}
}

func (s *DeviceManagementComplexSettingDefinitionValueType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseDeviceManagementComplexSettingDefinitionValueType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseDeviceManagementComplexSettingDefinitionValueType(input string) (*DeviceManagementComplexSettingDefinitionValueType, error) {
	vals := map[string]DeviceManagementComplexSettingDefinitionValueType{
		"abstractcomplex": DeviceManagementComplexSettingDefinitionValueType_AbstractComplex,
		"boolean":         DeviceManagementComplexSettingDefinitionValueType_Boolean,
		"collection":      DeviceManagementComplexSettingDefinitionValueType_Collection,
		"complex":         DeviceManagementComplexSettingDefinitionValueType_Complex,
		"integer":         DeviceManagementComplexSettingDefinitionValueType_Integer,
		"string":          DeviceManagementComplexSettingDefinitionValueType_String,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DeviceManagementComplexSettingDefinitionValueType(input)
	return &out, nil
}
