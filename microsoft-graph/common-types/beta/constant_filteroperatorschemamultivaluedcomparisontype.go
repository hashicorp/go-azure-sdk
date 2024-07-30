package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type FilterOperatorSchemaMultivaluedComparisonType string

const (
	FilterOperatorSchemaMultivaluedComparisonType_All FilterOperatorSchemaMultivaluedComparisonType = "All"
	FilterOperatorSchemaMultivaluedComparisonType_Any FilterOperatorSchemaMultivaluedComparisonType = "Any"
)

func PossibleValuesForFilterOperatorSchemaMultivaluedComparisonType() []string {
	return []string{
		string(FilterOperatorSchemaMultivaluedComparisonType_All),
		string(FilterOperatorSchemaMultivaluedComparisonType_Any),
	}
}

func (s *FilterOperatorSchemaMultivaluedComparisonType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseFilterOperatorSchemaMultivaluedComparisonType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseFilterOperatorSchemaMultivaluedComparisonType(input string) (*FilterOperatorSchemaMultivaluedComparisonType, error) {
	vals := map[string]FilterOperatorSchemaMultivaluedComparisonType{
		"all": FilterOperatorSchemaMultivaluedComparisonType_All,
		"any": FilterOperatorSchemaMultivaluedComparisonType_Any,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := FilterOperatorSchemaMultivaluedComparisonType(input)
	return &out, nil
}
