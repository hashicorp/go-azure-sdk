package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ManagedTenantsSettingValueType string

const (
	ManagedTenantsSettingValueType_Boolean           ManagedTenantsSettingValueType = "boolean"
	ManagedTenantsSettingValueType_BooleanCollection ManagedTenantsSettingValueType = "booleanCollection"
	ManagedTenantsSettingValueType_Guid              ManagedTenantsSettingValueType = "guid"
	ManagedTenantsSettingValueType_GuidCollection    ManagedTenantsSettingValueType = "guidCollection"
	ManagedTenantsSettingValueType_Integer           ManagedTenantsSettingValueType = "integer"
	ManagedTenantsSettingValueType_IntegerCollection ManagedTenantsSettingValueType = "integerCollection"
	ManagedTenantsSettingValueType_String            ManagedTenantsSettingValueType = "string"
	ManagedTenantsSettingValueType_StringCollection  ManagedTenantsSettingValueType = "stringCollection"
)

func PossibleValuesForManagedTenantsSettingValueType() []string {
	return []string{
		string(ManagedTenantsSettingValueType_Boolean),
		string(ManagedTenantsSettingValueType_BooleanCollection),
		string(ManagedTenantsSettingValueType_Guid),
		string(ManagedTenantsSettingValueType_GuidCollection),
		string(ManagedTenantsSettingValueType_Integer),
		string(ManagedTenantsSettingValueType_IntegerCollection),
		string(ManagedTenantsSettingValueType_String),
		string(ManagedTenantsSettingValueType_StringCollection),
	}
}

func (s *ManagedTenantsSettingValueType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseManagedTenantsSettingValueType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseManagedTenantsSettingValueType(input string) (*ManagedTenantsSettingValueType, error) {
	vals := map[string]ManagedTenantsSettingValueType{
		"boolean":           ManagedTenantsSettingValueType_Boolean,
		"booleancollection": ManagedTenantsSettingValueType_BooleanCollection,
		"guid":              ManagedTenantsSettingValueType_Guid,
		"guidcollection":    ManagedTenantsSettingValueType_GuidCollection,
		"integer":           ManagedTenantsSettingValueType_Integer,
		"integercollection": ManagedTenantsSettingValueType_IntegerCollection,
		"string":            ManagedTenantsSettingValueType_String,
		"stringcollection":  ManagedTenantsSettingValueType_StringCollection,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ManagedTenantsSettingValueType(input)
	return &out, nil
}
