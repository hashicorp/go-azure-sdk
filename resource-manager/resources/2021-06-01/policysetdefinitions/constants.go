package policysetdefinitions

import "strings"

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ParameterType string

const (
	ParameterTypeArray    ParameterType = "Array"
	ParameterTypeBoolean  ParameterType = "Boolean"
	ParameterTypeDateTime ParameterType = "DateTime"
	ParameterTypeFloat    ParameterType = "Float"
	ParameterTypeInteger  ParameterType = "Integer"
	ParameterTypeObject   ParameterType = "Object"
	ParameterTypeString   ParameterType = "String"
)

func PossibleValuesForParameterType() []string {
	return []string{
		string(ParameterTypeArray),
		string(ParameterTypeBoolean),
		string(ParameterTypeDateTime),
		string(ParameterTypeFloat),
		string(ParameterTypeInteger),
		string(ParameterTypeObject),
		string(ParameterTypeString),
	}
}

func parseParameterType(input string) (*ParameterType, error) {
	vals := map[string]ParameterType{
		"array":    ParameterTypeArray,
		"boolean":  ParameterTypeBoolean,
		"datetime": ParameterTypeDateTime,
		"float":    ParameterTypeFloat,
		"integer":  ParameterTypeInteger,
		"object":   ParameterTypeObject,
		"string":   ParameterTypeString,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ParameterType(input)
	return &out, nil
}

type PolicyType string

const (
	PolicyTypeBuiltIn      PolicyType = "BuiltIn"
	PolicyTypeCustom       PolicyType = "Custom"
	PolicyTypeNotSpecified PolicyType = "NotSpecified"
	PolicyTypeStatic       PolicyType = "Static"
)

func PossibleValuesForPolicyType() []string {
	return []string{
		string(PolicyTypeBuiltIn),
		string(PolicyTypeCustom),
		string(PolicyTypeNotSpecified),
		string(PolicyTypeStatic),
	}
}

func parsePolicyType(input string) (*PolicyType, error) {
	vals := map[string]PolicyType{
		"builtin":      PolicyTypeBuiltIn,
		"custom":       PolicyTypeCustom,
		"notspecified": PolicyTypeNotSpecified,
		"static":       PolicyTypeStatic,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := PolicyType(input)
	return &out, nil
}
