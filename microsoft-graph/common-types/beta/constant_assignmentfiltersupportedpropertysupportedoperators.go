package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AssignmentFilterSupportedPropertySupportedOperators string

const (
	AssignmentFilterSupportedPropertySupportedOperators_Contains      AssignmentFilterSupportedPropertySupportedOperators = "contains"
	AssignmentFilterSupportedPropertySupportedOperators_EndsWith      AssignmentFilterSupportedPropertySupportedOperators = "endsWith"
	AssignmentFilterSupportedPropertySupportedOperators_Equals        AssignmentFilterSupportedPropertySupportedOperators = "equals"
	AssignmentFilterSupportedPropertySupportedOperators_In            AssignmentFilterSupportedPropertySupportedOperators = "in"
	AssignmentFilterSupportedPropertySupportedOperators_NotContains   AssignmentFilterSupportedPropertySupportedOperators = "notContains"
	AssignmentFilterSupportedPropertySupportedOperators_NotEndsWith   AssignmentFilterSupportedPropertySupportedOperators = "notEndsWith"
	AssignmentFilterSupportedPropertySupportedOperators_NotEquals     AssignmentFilterSupportedPropertySupportedOperators = "notEquals"
	AssignmentFilterSupportedPropertySupportedOperators_NotIn         AssignmentFilterSupportedPropertySupportedOperators = "notIn"
	AssignmentFilterSupportedPropertySupportedOperators_NotSet        AssignmentFilterSupportedPropertySupportedOperators = "notSet"
	AssignmentFilterSupportedPropertySupportedOperators_NotStartsWith AssignmentFilterSupportedPropertySupportedOperators = "notStartsWith"
	AssignmentFilterSupportedPropertySupportedOperators_StartsWith    AssignmentFilterSupportedPropertySupportedOperators = "startsWith"
)

func PossibleValuesForAssignmentFilterSupportedPropertySupportedOperators() []string {
	return []string{
		string(AssignmentFilterSupportedPropertySupportedOperators_Contains),
		string(AssignmentFilterSupportedPropertySupportedOperators_EndsWith),
		string(AssignmentFilterSupportedPropertySupportedOperators_Equals),
		string(AssignmentFilterSupportedPropertySupportedOperators_In),
		string(AssignmentFilterSupportedPropertySupportedOperators_NotContains),
		string(AssignmentFilterSupportedPropertySupportedOperators_NotEndsWith),
		string(AssignmentFilterSupportedPropertySupportedOperators_NotEquals),
		string(AssignmentFilterSupportedPropertySupportedOperators_NotIn),
		string(AssignmentFilterSupportedPropertySupportedOperators_NotSet),
		string(AssignmentFilterSupportedPropertySupportedOperators_NotStartsWith),
		string(AssignmentFilterSupportedPropertySupportedOperators_StartsWith),
	}
}

func (s *AssignmentFilterSupportedPropertySupportedOperators) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAssignmentFilterSupportedPropertySupportedOperators(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAssignmentFilterSupportedPropertySupportedOperators(input string) (*AssignmentFilterSupportedPropertySupportedOperators, error) {
	vals := map[string]AssignmentFilterSupportedPropertySupportedOperators{
		"contains":      AssignmentFilterSupportedPropertySupportedOperators_Contains,
		"endswith":      AssignmentFilterSupportedPropertySupportedOperators_EndsWith,
		"equals":        AssignmentFilterSupportedPropertySupportedOperators_Equals,
		"in":            AssignmentFilterSupportedPropertySupportedOperators_In,
		"notcontains":   AssignmentFilterSupportedPropertySupportedOperators_NotContains,
		"notendswith":   AssignmentFilterSupportedPropertySupportedOperators_NotEndsWith,
		"notequals":     AssignmentFilterSupportedPropertySupportedOperators_NotEquals,
		"notin":         AssignmentFilterSupportedPropertySupportedOperators_NotIn,
		"notset":        AssignmentFilterSupportedPropertySupportedOperators_NotSet,
		"notstartswith": AssignmentFilterSupportedPropertySupportedOperators_NotStartsWith,
		"startswith":    AssignmentFilterSupportedPropertySupportedOperators_StartsWith,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AssignmentFilterSupportedPropertySupportedOperators(input)
	return &out, nil
}
