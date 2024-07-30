package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AttributeDefinitionType string

const (
	AttributeDefinitionType_Binary    AttributeDefinitionType = "Binary"
	AttributeDefinitionType_Boolean   AttributeDefinitionType = "Boolean"
	AttributeDefinitionType_DateTime  AttributeDefinitionType = "DateTime"
	AttributeDefinitionType_Integer   AttributeDefinitionType = "Integer"
	AttributeDefinitionType_Reference AttributeDefinitionType = "Reference"
	AttributeDefinitionType_String    AttributeDefinitionType = "String"
)

func PossibleValuesForAttributeDefinitionType() []string {
	return []string{
		string(AttributeDefinitionType_Binary),
		string(AttributeDefinitionType_Boolean),
		string(AttributeDefinitionType_DateTime),
		string(AttributeDefinitionType_Integer),
		string(AttributeDefinitionType_Reference),
		string(AttributeDefinitionType_String),
	}
}

func (s *AttributeDefinitionType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAttributeDefinitionType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAttributeDefinitionType(input string) (*AttributeDefinitionType, error) {
	vals := map[string]AttributeDefinitionType{
		"binary":    AttributeDefinitionType_Binary,
		"boolean":   AttributeDefinitionType_Boolean,
		"datetime":  AttributeDefinitionType_DateTime,
		"integer":   AttributeDefinitionType_Integer,
		"reference": AttributeDefinitionType_Reference,
		"string":    AttributeDefinitionType_String,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AttributeDefinitionType(input)
	return &out, nil
}
