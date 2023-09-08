package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AttributeDefinitionType string

const (
	AttributeDefinitionTypeBinary    AttributeDefinitionType = "Binary"
	AttributeDefinitionTypeBoolean   AttributeDefinitionType = "Boolean"
	AttributeDefinitionTypeDateTime  AttributeDefinitionType = "DateTime"
	AttributeDefinitionTypeInteger   AttributeDefinitionType = "Integer"
	AttributeDefinitionTypeReference AttributeDefinitionType = "Reference"
	AttributeDefinitionTypeString    AttributeDefinitionType = "String"
)

func PossibleValuesForAttributeDefinitionType() []string {
	return []string{
		string(AttributeDefinitionTypeBinary),
		string(AttributeDefinitionTypeBoolean),
		string(AttributeDefinitionTypeDateTime),
		string(AttributeDefinitionTypeInteger),
		string(AttributeDefinitionTypeReference),
		string(AttributeDefinitionTypeString),
	}
}

func parseAttributeDefinitionType(input string) (*AttributeDefinitionType, error) {
	vals := map[string]AttributeDefinitionType{
		"binary":    AttributeDefinitionTypeBinary,
		"boolean":   AttributeDefinitionTypeBoolean,
		"datetime":  AttributeDefinitionTypeDateTime,
		"integer":   AttributeDefinitionTypeInteger,
		"reference": AttributeDefinitionTypeReference,
		"string":    AttributeDefinitionTypeString,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AttributeDefinitionType(input)
	return &out, nil
}
