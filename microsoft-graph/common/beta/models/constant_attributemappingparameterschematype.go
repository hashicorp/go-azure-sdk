package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AttributeMappingParameterSchemaType string

const (
	AttributeMappingParameterSchemaTypeBinary    AttributeMappingParameterSchemaType = "Binary"
	AttributeMappingParameterSchemaTypeBoolean   AttributeMappingParameterSchemaType = "Boolean"
	AttributeMappingParameterSchemaTypeDateTime  AttributeMappingParameterSchemaType = "DateTime"
	AttributeMappingParameterSchemaTypeInteger   AttributeMappingParameterSchemaType = "Integer"
	AttributeMappingParameterSchemaTypeReference AttributeMappingParameterSchemaType = "Reference"
	AttributeMappingParameterSchemaTypeString    AttributeMappingParameterSchemaType = "String"
)

func PossibleValuesForAttributeMappingParameterSchemaType() []string {
	return []string{
		string(AttributeMappingParameterSchemaTypeBinary),
		string(AttributeMappingParameterSchemaTypeBoolean),
		string(AttributeMappingParameterSchemaTypeDateTime),
		string(AttributeMappingParameterSchemaTypeInteger),
		string(AttributeMappingParameterSchemaTypeReference),
		string(AttributeMappingParameterSchemaTypeString),
	}
}

func parseAttributeMappingParameterSchemaType(input string) (*AttributeMappingParameterSchemaType, error) {
	vals := map[string]AttributeMappingParameterSchemaType{
		"binary":    AttributeMappingParameterSchemaTypeBinary,
		"boolean":   AttributeMappingParameterSchemaTypeBoolean,
		"datetime":  AttributeMappingParameterSchemaTypeDateTime,
		"integer":   AttributeMappingParameterSchemaTypeInteger,
		"reference": AttributeMappingParameterSchemaTypeReference,
		"string":    AttributeMappingParameterSchemaTypeString,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AttributeMappingParameterSchemaType(input)
	return &out, nil
}
