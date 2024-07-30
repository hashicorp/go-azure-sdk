package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ManagedTenantsTemplateParameterValueType string

const (
	ManagedTenantsTemplateParameterValueType_Boolean           ManagedTenantsTemplateParameterValueType = "boolean"
	ManagedTenantsTemplateParameterValueType_BooleanCollection ManagedTenantsTemplateParameterValueType = "booleanCollection"
	ManagedTenantsTemplateParameterValueType_Guid              ManagedTenantsTemplateParameterValueType = "guid"
	ManagedTenantsTemplateParameterValueType_GuidCollection    ManagedTenantsTemplateParameterValueType = "guidCollection"
	ManagedTenantsTemplateParameterValueType_Integer           ManagedTenantsTemplateParameterValueType = "integer"
	ManagedTenantsTemplateParameterValueType_IntegerCollection ManagedTenantsTemplateParameterValueType = "integerCollection"
	ManagedTenantsTemplateParameterValueType_String            ManagedTenantsTemplateParameterValueType = "string"
	ManagedTenantsTemplateParameterValueType_StringCollection  ManagedTenantsTemplateParameterValueType = "stringCollection"
)

func PossibleValuesForManagedTenantsTemplateParameterValueType() []string {
	return []string{
		string(ManagedTenantsTemplateParameterValueType_Boolean),
		string(ManagedTenantsTemplateParameterValueType_BooleanCollection),
		string(ManagedTenantsTemplateParameterValueType_Guid),
		string(ManagedTenantsTemplateParameterValueType_GuidCollection),
		string(ManagedTenantsTemplateParameterValueType_Integer),
		string(ManagedTenantsTemplateParameterValueType_IntegerCollection),
		string(ManagedTenantsTemplateParameterValueType_String),
		string(ManagedTenantsTemplateParameterValueType_StringCollection),
	}
}

func (s *ManagedTenantsTemplateParameterValueType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseManagedTenantsTemplateParameterValueType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseManagedTenantsTemplateParameterValueType(input string) (*ManagedTenantsTemplateParameterValueType, error) {
	vals := map[string]ManagedTenantsTemplateParameterValueType{
		"boolean":           ManagedTenantsTemplateParameterValueType_Boolean,
		"booleancollection": ManagedTenantsTemplateParameterValueType_BooleanCollection,
		"guid":              ManagedTenantsTemplateParameterValueType_Guid,
		"guidcollection":    ManagedTenantsTemplateParameterValueType_GuidCollection,
		"integer":           ManagedTenantsTemplateParameterValueType_Integer,
		"integercollection": ManagedTenantsTemplateParameterValueType_IntegerCollection,
		"string":            ManagedTenantsTemplateParameterValueType_String,
		"stringcollection":  ManagedTenantsTemplateParameterValueType_StringCollection,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ManagedTenantsTemplateParameterValueType(input)
	return &out, nil
}
