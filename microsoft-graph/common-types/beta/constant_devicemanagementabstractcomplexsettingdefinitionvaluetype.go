package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceManagementAbstractComplexSettingDefinitionValueType string

const (
	DeviceManagementAbstractComplexSettingDefinitionValueType_AbstractComplex DeviceManagementAbstractComplexSettingDefinitionValueType = "abstractComplex"
	DeviceManagementAbstractComplexSettingDefinitionValueType_Boolean         DeviceManagementAbstractComplexSettingDefinitionValueType = "boolean"
	DeviceManagementAbstractComplexSettingDefinitionValueType_Collection      DeviceManagementAbstractComplexSettingDefinitionValueType = "collection"
	DeviceManagementAbstractComplexSettingDefinitionValueType_Complex         DeviceManagementAbstractComplexSettingDefinitionValueType = "complex"
	DeviceManagementAbstractComplexSettingDefinitionValueType_Integer         DeviceManagementAbstractComplexSettingDefinitionValueType = "integer"
	DeviceManagementAbstractComplexSettingDefinitionValueType_String          DeviceManagementAbstractComplexSettingDefinitionValueType = "string"
)

func PossibleValuesForDeviceManagementAbstractComplexSettingDefinitionValueType() []string {
	return []string{
		string(DeviceManagementAbstractComplexSettingDefinitionValueType_AbstractComplex),
		string(DeviceManagementAbstractComplexSettingDefinitionValueType_Boolean),
		string(DeviceManagementAbstractComplexSettingDefinitionValueType_Collection),
		string(DeviceManagementAbstractComplexSettingDefinitionValueType_Complex),
		string(DeviceManagementAbstractComplexSettingDefinitionValueType_Integer),
		string(DeviceManagementAbstractComplexSettingDefinitionValueType_String),
	}
}

func (s *DeviceManagementAbstractComplexSettingDefinitionValueType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseDeviceManagementAbstractComplexSettingDefinitionValueType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseDeviceManagementAbstractComplexSettingDefinitionValueType(input string) (*DeviceManagementAbstractComplexSettingDefinitionValueType, error) {
	vals := map[string]DeviceManagementAbstractComplexSettingDefinitionValueType{
		"abstractcomplex": DeviceManagementAbstractComplexSettingDefinitionValueType_AbstractComplex,
		"boolean":         DeviceManagementAbstractComplexSettingDefinitionValueType_Boolean,
		"collection":      DeviceManagementAbstractComplexSettingDefinitionValueType_Collection,
		"complex":         DeviceManagementAbstractComplexSettingDefinitionValueType_Complex,
		"integer":         DeviceManagementAbstractComplexSettingDefinitionValueType_Integer,
		"string":          DeviceManagementAbstractComplexSettingDefinitionValueType_String,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DeviceManagementAbstractComplexSettingDefinitionValueType(input)
	return &out, nil
}
