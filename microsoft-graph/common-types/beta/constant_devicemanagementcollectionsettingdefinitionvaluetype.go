package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceManagementCollectionSettingDefinitionValueType string

const (
	DeviceManagementCollectionSettingDefinitionValueType_AbstractComplex DeviceManagementCollectionSettingDefinitionValueType = "abstractComplex"
	DeviceManagementCollectionSettingDefinitionValueType_Boolean         DeviceManagementCollectionSettingDefinitionValueType = "boolean"
	DeviceManagementCollectionSettingDefinitionValueType_Collection      DeviceManagementCollectionSettingDefinitionValueType = "collection"
	DeviceManagementCollectionSettingDefinitionValueType_Complex         DeviceManagementCollectionSettingDefinitionValueType = "complex"
	DeviceManagementCollectionSettingDefinitionValueType_Integer         DeviceManagementCollectionSettingDefinitionValueType = "integer"
	DeviceManagementCollectionSettingDefinitionValueType_String          DeviceManagementCollectionSettingDefinitionValueType = "string"
)

func PossibleValuesForDeviceManagementCollectionSettingDefinitionValueType() []string {
	return []string{
		string(DeviceManagementCollectionSettingDefinitionValueType_AbstractComplex),
		string(DeviceManagementCollectionSettingDefinitionValueType_Boolean),
		string(DeviceManagementCollectionSettingDefinitionValueType_Collection),
		string(DeviceManagementCollectionSettingDefinitionValueType_Complex),
		string(DeviceManagementCollectionSettingDefinitionValueType_Integer),
		string(DeviceManagementCollectionSettingDefinitionValueType_String),
	}
}

func (s *DeviceManagementCollectionSettingDefinitionValueType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseDeviceManagementCollectionSettingDefinitionValueType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseDeviceManagementCollectionSettingDefinitionValueType(input string) (*DeviceManagementCollectionSettingDefinitionValueType, error) {
	vals := map[string]DeviceManagementCollectionSettingDefinitionValueType{
		"abstractcomplex": DeviceManagementCollectionSettingDefinitionValueType_AbstractComplex,
		"boolean":         DeviceManagementCollectionSettingDefinitionValueType_Boolean,
		"collection":      DeviceManagementCollectionSettingDefinitionValueType_Collection,
		"complex":         DeviceManagementCollectionSettingDefinitionValueType_Complex,
		"integer":         DeviceManagementCollectionSettingDefinitionValueType_Integer,
		"string":          DeviceManagementCollectionSettingDefinitionValueType_String,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DeviceManagementCollectionSettingDefinitionValueType(input)
	return &out, nil
}
