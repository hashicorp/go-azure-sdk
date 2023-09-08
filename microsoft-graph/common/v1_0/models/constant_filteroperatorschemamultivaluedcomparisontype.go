package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type FilterOperatorSchemaMultivaluedComparisonType string

const (
	FilterOperatorSchemaMultivaluedComparisonTypeAll FilterOperatorSchemaMultivaluedComparisonType = "All"
	FilterOperatorSchemaMultivaluedComparisonTypeAny FilterOperatorSchemaMultivaluedComparisonType = "Any"
)

func PossibleValuesForFilterOperatorSchemaMultivaluedComparisonType() []string {
	return []string{
		string(FilterOperatorSchemaMultivaluedComparisonTypeAll),
		string(FilterOperatorSchemaMultivaluedComparisonTypeAny),
	}
}

func parseFilterOperatorSchemaMultivaluedComparisonType(input string) (*FilterOperatorSchemaMultivaluedComparisonType, error) {
	vals := map[string]FilterOperatorSchemaMultivaluedComparisonType{
		"all": FilterOperatorSchemaMultivaluedComparisonTypeAll,
		"any": FilterOperatorSchemaMultivaluedComparisonTypeAny,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := FilterOperatorSchemaMultivaluedComparisonType(input)
	return &out, nil
}
