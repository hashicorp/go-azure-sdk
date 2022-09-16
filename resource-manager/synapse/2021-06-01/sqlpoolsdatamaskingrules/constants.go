package sqlpoolsdatamaskingrules

import "strings"

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DataMaskingFunction string

const (
	DataMaskingFunctionCCN     DataMaskingFunction = "CCN"
	DataMaskingFunctionDefault DataMaskingFunction = "Default"
	DataMaskingFunctionEmail   DataMaskingFunction = "Email"
	DataMaskingFunctionNumber  DataMaskingFunction = "Number"
	DataMaskingFunctionSSN     DataMaskingFunction = "SSN"
	DataMaskingFunctionText    DataMaskingFunction = "Text"
)

func PossibleValuesForDataMaskingFunction() []string {
	return []string{
		string(DataMaskingFunctionCCN),
		string(DataMaskingFunctionDefault),
		string(DataMaskingFunctionEmail),
		string(DataMaskingFunctionNumber),
		string(DataMaskingFunctionSSN),
		string(DataMaskingFunctionText),
	}
}

func parseDataMaskingFunction(input string) (*DataMaskingFunction, error) {
	vals := map[string]DataMaskingFunction{
		"ccn":     DataMaskingFunctionCCN,
		"default": DataMaskingFunctionDefault,
		"email":   DataMaskingFunctionEmail,
		"number":  DataMaskingFunctionNumber,
		"ssn":     DataMaskingFunctionSSN,
		"text":    DataMaskingFunctionText,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DataMaskingFunction(input)
	return &out, nil
}

type DataMaskingRuleState string

const (
	DataMaskingRuleStateDisabled DataMaskingRuleState = "Disabled"
	DataMaskingRuleStateEnabled  DataMaskingRuleState = "Enabled"
)

func PossibleValuesForDataMaskingRuleState() []string {
	return []string{
		string(DataMaskingRuleStateDisabled),
		string(DataMaskingRuleStateEnabled),
	}
}

func parseDataMaskingRuleState(input string) (*DataMaskingRuleState, error) {
	vals := map[string]DataMaskingRuleState{
		"disabled": DataMaskingRuleStateDisabled,
		"enabled":  DataMaskingRuleStateEnabled,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DataMaskingRuleState(input)
	return &out, nil
}
