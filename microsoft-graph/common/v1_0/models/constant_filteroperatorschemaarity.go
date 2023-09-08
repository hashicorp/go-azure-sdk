package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type FilterOperatorSchemaArity string

const (
	FilterOperatorSchemaArityBinary FilterOperatorSchemaArity = "Binary"
	FilterOperatorSchemaArityUnary  FilterOperatorSchemaArity = "Unary"
)

func PossibleValuesForFilterOperatorSchemaArity() []string {
	return []string{
		string(FilterOperatorSchemaArityBinary),
		string(FilterOperatorSchemaArityUnary),
	}
}

func parseFilterOperatorSchemaArity(input string) (*FilterOperatorSchemaArity, error) {
	vals := map[string]FilterOperatorSchemaArity{
		"binary": FilterOperatorSchemaArityBinary,
		"unary":  FilterOperatorSchemaArityUnary,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := FilterOperatorSchemaArity(input)
	return &out, nil
}
