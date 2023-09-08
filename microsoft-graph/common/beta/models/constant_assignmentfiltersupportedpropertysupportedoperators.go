package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AssignmentFilterSupportedPropertySupportedOperators string

const (
	AssignmentFilterSupportedPropertySupportedOperatorscontains      AssignmentFilterSupportedPropertySupportedOperators = "Contains"
	AssignmentFilterSupportedPropertySupportedOperatorsendsWith      AssignmentFilterSupportedPropertySupportedOperators = "EndsWith"
	AssignmentFilterSupportedPropertySupportedOperatorsequals        AssignmentFilterSupportedPropertySupportedOperators = "Equals"
	AssignmentFilterSupportedPropertySupportedOperatorsin            AssignmentFilterSupportedPropertySupportedOperators = "In"
	AssignmentFilterSupportedPropertySupportedOperatorsnotContains   AssignmentFilterSupportedPropertySupportedOperators = "NotContains"
	AssignmentFilterSupportedPropertySupportedOperatorsnotEndsWith   AssignmentFilterSupportedPropertySupportedOperators = "NotEndsWith"
	AssignmentFilterSupportedPropertySupportedOperatorsnotEquals     AssignmentFilterSupportedPropertySupportedOperators = "NotEquals"
	AssignmentFilterSupportedPropertySupportedOperatorsnotIn         AssignmentFilterSupportedPropertySupportedOperators = "NotIn"
	AssignmentFilterSupportedPropertySupportedOperatorsnotSet        AssignmentFilterSupportedPropertySupportedOperators = "NotSet"
	AssignmentFilterSupportedPropertySupportedOperatorsnotStartsWith AssignmentFilterSupportedPropertySupportedOperators = "NotStartsWith"
	AssignmentFilterSupportedPropertySupportedOperatorsstartsWith    AssignmentFilterSupportedPropertySupportedOperators = "StartsWith"
)

func PossibleValuesForAssignmentFilterSupportedPropertySupportedOperators() []string {
	return []string{
		string(AssignmentFilterSupportedPropertySupportedOperatorscontains),
		string(AssignmentFilterSupportedPropertySupportedOperatorsendsWith),
		string(AssignmentFilterSupportedPropertySupportedOperatorsequals),
		string(AssignmentFilterSupportedPropertySupportedOperatorsin),
		string(AssignmentFilterSupportedPropertySupportedOperatorsnotContains),
		string(AssignmentFilterSupportedPropertySupportedOperatorsnotEndsWith),
		string(AssignmentFilterSupportedPropertySupportedOperatorsnotEquals),
		string(AssignmentFilterSupportedPropertySupportedOperatorsnotIn),
		string(AssignmentFilterSupportedPropertySupportedOperatorsnotSet),
		string(AssignmentFilterSupportedPropertySupportedOperatorsnotStartsWith),
		string(AssignmentFilterSupportedPropertySupportedOperatorsstartsWith),
	}
}

func parseAssignmentFilterSupportedPropertySupportedOperators(input string) (*AssignmentFilterSupportedPropertySupportedOperators, error) {
	vals := map[string]AssignmentFilterSupportedPropertySupportedOperators{
		"contains":      AssignmentFilterSupportedPropertySupportedOperatorscontains,
		"endswith":      AssignmentFilterSupportedPropertySupportedOperatorsendsWith,
		"equals":        AssignmentFilterSupportedPropertySupportedOperatorsequals,
		"in":            AssignmentFilterSupportedPropertySupportedOperatorsin,
		"notcontains":   AssignmentFilterSupportedPropertySupportedOperatorsnotContains,
		"notendswith":   AssignmentFilterSupportedPropertySupportedOperatorsnotEndsWith,
		"notequals":     AssignmentFilterSupportedPropertySupportedOperatorsnotEquals,
		"notin":         AssignmentFilterSupportedPropertySupportedOperatorsnotIn,
		"notset":        AssignmentFilterSupportedPropertySupportedOperatorsnotSet,
		"notstartswith": AssignmentFilterSupportedPropertySupportedOperatorsnotStartsWith,
		"startswith":    AssignmentFilterSupportedPropertySupportedOperatorsstartsWith,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AssignmentFilterSupportedPropertySupportedOperators(input)
	return &out, nil
}
