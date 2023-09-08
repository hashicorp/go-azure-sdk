package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ExternalConnectorsPropertyRuleOperation string

const (
	ExternalConnectorsPropertyRuleOperationcontains    ExternalConnectorsPropertyRuleOperation = "Contains"
	ExternalConnectorsPropertyRuleOperationequals      ExternalConnectorsPropertyRuleOperation = "Equals"
	ExternalConnectorsPropertyRuleOperationgreaterThan ExternalConnectorsPropertyRuleOperation = "GreaterThan"
	ExternalConnectorsPropertyRuleOperationlessThan    ExternalConnectorsPropertyRuleOperation = "LessThan"
	ExternalConnectorsPropertyRuleOperationnotContains ExternalConnectorsPropertyRuleOperation = "NotContains"
	ExternalConnectorsPropertyRuleOperationnotEquals   ExternalConnectorsPropertyRuleOperation = "NotEquals"
	ExternalConnectorsPropertyRuleOperationnull        ExternalConnectorsPropertyRuleOperation = "Null"
	ExternalConnectorsPropertyRuleOperationstartsWith  ExternalConnectorsPropertyRuleOperation = "StartsWith"
)

func PossibleValuesForExternalConnectorsPropertyRuleOperation() []string {
	return []string{
		string(ExternalConnectorsPropertyRuleOperationcontains),
		string(ExternalConnectorsPropertyRuleOperationequals),
		string(ExternalConnectorsPropertyRuleOperationgreaterThan),
		string(ExternalConnectorsPropertyRuleOperationlessThan),
		string(ExternalConnectorsPropertyRuleOperationnotContains),
		string(ExternalConnectorsPropertyRuleOperationnotEquals),
		string(ExternalConnectorsPropertyRuleOperationnull),
		string(ExternalConnectorsPropertyRuleOperationstartsWith),
	}
}

func parseExternalConnectorsPropertyRuleOperation(input string) (*ExternalConnectorsPropertyRuleOperation, error) {
	vals := map[string]ExternalConnectorsPropertyRuleOperation{
		"contains":    ExternalConnectorsPropertyRuleOperationcontains,
		"equals":      ExternalConnectorsPropertyRuleOperationequals,
		"greaterthan": ExternalConnectorsPropertyRuleOperationgreaterThan,
		"lessthan":    ExternalConnectorsPropertyRuleOperationlessThan,
		"notcontains": ExternalConnectorsPropertyRuleOperationnotContains,
		"notequals":   ExternalConnectorsPropertyRuleOperationnotEquals,
		"null":        ExternalConnectorsPropertyRuleOperationnull,
		"startswith":  ExternalConnectorsPropertyRuleOperationstartsWith,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ExternalConnectorsPropertyRuleOperation(input)
	return &out, nil
}
