package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ExternalConnectorsPropertyRuleOperation string

const (
	ExternalConnectorsPropertyRuleOperation_Contains    ExternalConnectorsPropertyRuleOperation = "contains"
	ExternalConnectorsPropertyRuleOperation_Equals      ExternalConnectorsPropertyRuleOperation = "equals"
	ExternalConnectorsPropertyRuleOperation_GreaterThan ExternalConnectorsPropertyRuleOperation = "greaterThan"
	ExternalConnectorsPropertyRuleOperation_LessThan    ExternalConnectorsPropertyRuleOperation = "lessThan"
	ExternalConnectorsPropertyRuleOperation_NotContains ExternalConnectorsPropertyRuleOperation = "notContains"
	ExternalConnectorsPropertyRuleOperation_NotEquals   ExternalConnectorsPropertyRuleOperation = "notEquals"
	ExternalConnectorsPropertyRuleOperation_Null        ExternalConnectorsPropertyRuleOperation = "null"
	ExternalConnectorsPropertyRuleOperation_StartsWith  ExternalConnectorsPropertyRuleOperation = "startsWith"
)

func PossibleValuesForExternalConnectorsPropertyRuleOperation() []string {
	return []string{
		string(ExternalConnectorsPropertyRuleOperation_Contains),
		string(ExternalConnectorsPropertyRuleOperation_Equals),
		string(ExternalConnectorsPropertyRuleOperation_GreaterThan),
		string(ExternalConnectorsPropertyRuleOperation_LessThan),
		string(ExternalConnectorsPropertyRuleOperation_NotContains),
		string(ExternalConnectorsPropertyRuleOperation_NotEquals),
		string(ExternalConnectorsPropertyRuleOperation_Null),
		string(ExternalConnectorsPropertyRuleOperation_StartsWith),
	}
}

func (s *ExternalConnectorsPropertyRuleOperation) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseExternalConnectorsPropertyRuleOperation(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseExternalConnectorsPropertyRuleOperation(input string) (*ExternalConnectorsPropertyRuleOperation, error) {
	vals := map[string]ExternalConnectorsPropertyRuleOperation{
		"contains":    ExternalConnectorsPropertyRuleOperation_Contains,
		"equals":      ExternalConnectorsPropertyRuleOperation_Equals,
		"greaterthan": ExternalConnectorsPropertyRuleOperation_GreaterThan,
		"lessthan":    ExternalConnectorsPropertyRuleOperation_LessThan,
		"notcontains": ExternalConnectorsPropertyRuleOperation_NotContains,
		"notequals":   ExternalConnectorsPropertyRuleOperation_NotEquals,
		"null":        ExternalConnectorsPropertyRuleOperation_Null,
		"startswith":  ExternalConnectorsPropertyRuleOperation_StartsWith,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ExternalConnectorsPropertyRuleOperation(input)
	return &out, nil
}
