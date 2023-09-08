package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type FilterOperatorSchemaSupportedAttributeTypes string

const (
	FilterOperatorSchemaSupportedAttributeTypesBinary    FilterOperatorSchemaSupportedAttributeTypes = "Binary"
	FilterOperatorSchemaSupportedAttributeTypesBoolean   FilterOperatorSchemaSupportedAttributeTypes = "Boolean"
	FilterOperatorSchemaSupportedAttributeTypesDateTime  FilterOperatorSchemaSupportedAttributeTypes = "DateTime"
	FilterOperatorSchemaSupportedAttributeTypesInteger   FilterOperatorSchemaSupportedAttributeTypes = "Integer"
	FilterOperatorSchemaSupportedAttributeTypesReference FilterOperatorSchemaSupportedAttributeTypes = "Reference"
	FilterOperatorSchemaSupportedAttributeTypesString    FilterOperatorSchemaSupportedAttributeTypes = "String"
)

func PossibleValuesForFilterOperatorSchemaSupportedAttributeTypes() []string {
	return []string{
		string(FilterOperatorSchemaSupportedAttributeTypesBinary),
		string(FilterOperatorSchemaSupportedAttributeTypesBoolean),
		string(FilterOperatorSchemaSupportedAttributeTypesDateTime),
		string(FilterOperatorSchemaSupportedAttributeTypesInteger),
		string(FilterOperatorSchemaSupportedAttributeTypesReference),
		string(FilterOperatorSchemaSupportedAttributeTypesString),
	}
}

func parseFilterOperatorSchemaSupportedAttributeTypes(input string) (*FilterOperatorSchemaSupportedAttributeTypes, error) {
	vals := map[string]FilterOperatorSchemaSupportedAttributeTypes{
		"binary":    FilterOperatorSchemaSupportedAttributeTypesBinary,
		"boolean":   FilterOperatorSchemaSupportedAttributeTypesBoolean,
		"datetime":  FilterOperatorSchemaSupportedAttributeTypesDateTime,
		"integer":   FilterOperatorSchemaSupportedAttributeTypesInteger,
		"reference": FilterOperatorSchemaSupportedAttributeTypesReference,
		"string":    FilterOperatorSchemaSupportedAttributeTypesString,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := FilterOperatorSchemaSupportedAttributeTypes(input)
	return &out, nil
}
