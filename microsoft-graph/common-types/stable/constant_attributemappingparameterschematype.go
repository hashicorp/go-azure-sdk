package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AttributeMappingParameterSchemaType string

const (
	AttributeMappingParameterSchemaType_Binary    AttributeMappingParameterSchemaType = "Binary"
	AttributeMappingParameterSchemaType_Boolean   AttributeMappingParameterSchemaType = "Boolean"
	AttributeMappingParameterSchemaType_DateTime  AttributeMappingParameterSchemaType = "DateTime"
	AttributeMappingParameterSchemaType_Integer   AttributeMappingParameterSchemaType = "Integer"
	AttributeMappingParameterSchemaType_Reference AttributeMappingParameterSchemaType = "Reference"
	AttributeMappingParameterSchemaType_String    AttributeMappingParameterSchemaType = "String"
)

func PossibleValuesForAttributeMappingParameterSchemaType() []string {
	return []string{
		string(AttributeMappingParameterSchemaType_Binary),
		string(AttributeMappingParameterSchemaType_Boolean),
		string(AttributeMappingParameterSchemaType_DateTime),
		string(AttributeMappingParameterSchemaType_Integer),
		string(AttributeMappingParameterSchemaType_Reference),
		string(AttributeMappingParameterSchemaType_String),
	}
}

func (s *AttributeMappingParameterSchemaType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAttributeMappingParameterSchemaType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAttributeMappingParameterSchemaType(input string) (*AttributeMappingParameterSchemaType, error) {
	vals := map[string]AttributeMappingParameterSchemaType{
		"binary":    AttributeMappingParameterSchemaType_Binary,
		"boolean":   AttributeMappingParameterSchemaType_Boolean,
		"datetime":  AttributeMappingParameterSchemaType_DateTime,
		"integer":   AttributeMappingParameterSchemaType_Integer,
		"reference": AttributeMappingParameterSchemaType_Reference,
		"string":    AttributeMappingParameterSchemaType_String,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AttributeMappingParameterSchemaType(input)
	return &out, nil
}
