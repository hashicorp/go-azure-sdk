package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type FilterOperatorSchemaSupportedAttributeTypes string

const (
	FilterOperatorSchemaSupportedAttributeTypes_Binary    FilterOperatorSchemaSupportedAttributeTypes = "Binary"
	FilterOperatorSchemaSupportedAttributeTypes_Boolean   FilterOperatorSchemaSupportedAttributeTypes = "Boolean"
	FilterOperatorSchemaSupportedAttributeTypes_DateTime  FilterOperatorSchemaSupportedAttributeTypes = "DateTime"
	FilterOperatorSchemaSupportedAttributeTypes_Integer   FilterOperatorSchemaSupportedAttributeTypes = "Integer"
	FilterOperatorSchemaSupportedAttributeTypes_Reference FilterOperatorSchemaSupportedAttributeTypes = "Reference"
	FilterOperatorSchemaSupportedAttributeTypes_String    FilterOperatorSchemaSupportedAttributeTypes = "String"
)

func PossibleValuesForFilterOperatorSchemaSupportedAttributeTypes() []string {
	return []string{
		string(FilterOperatorSchemaSupportedAttributeTypes_Binary),
		string(FilterOperatorSchemaSupportedAttributeTypes_Boolean),
		string(FilterOperatorSchemaSupportedAttributeTypes_DateTime),
		string(FilterOperatorSchemaSupportedAttributeTypes_Integer),
		string(FilterOperatorSchemaSupportedAttributeTypes_Reference),
		string(FilterOperatorSchemaSupportedAttributeTypes_String),
	}
}

func (s *FilterOperatorSchemaSupportedAttributeTypes) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseFilterOperatorSchemaSupportedAttributeTypes(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseFilterOperatorSchemaSupportedAttributeTypes(input string) (*FilterOperatorSchemaSupportedAttributeTypes, error) {
	vals := map[string]FilterOperatorSchemaSupportedAttributeTypes{
		"binary":    FilterOperatorSchemaSupportedAttributeTypes_Binary,
		"boolean":   FilterOperatorSchemaSupportedAttributeTypes_Boolean,
		"datetime":  FilterOperatorSchemaSupportedAttributeTypes_DateTime,
		"integer":   FilterOperatorSchemaSupportedAttributeTypes_Integer,
		"reference": FilterOperatorSchemaSupportedAttributeTypes_Reference,
		"string":    FilterOperatorSchemaSupportedAttributeTypes_String,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := FilterOperatorSchemaSupportedAttributeTypes(input)
	return &out, nil
}
