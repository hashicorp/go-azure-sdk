package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type FilterOperatorSchemaArity string

const (
	FilterOperatorSchemaArity_Binary FilterOperatorSchemaArity = "Binary"
	FilterOperatorSchemaArity_Unary  FilterOperatorSchemaArity = "Unary"
)

func PossibleValuesForFilterOperatorSchemaArity() []string {
	return []string{
		string(FilterOperatorSchemaArity_Binary),
		string(FilterOperatorSchemaArity_Unary),
	}
}

func (s *FilterOperatorSchemaArity) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseFilterOperatorSchemaArity(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseFilterOperatorSchemaArity(input string) (*FilterOperatorSchemaArity, error) {
	vals := map[string]FilterOperatorSchemaArity{
		"binary": FilterOperatorSchemaArity_Binary,
		"unary":  FilterOperatorSchemaArity_Unary,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := FilterOperatorSchemaArity(input)
	return &out, nil
}
