package globalparameters

import "strings"

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GlobalParameterType string

const (
	GlobalParameterTypeArray  GlobalParameterType = "Array"
	GlobalParameterTypeBool   GlobalParameterType = "Bool"
	GlobalParameterTypeFloat  GlobalParameterType = "Float"
	GlobalParameterTypeInt    GlobalParameterType = "Int"
	GlobalParameterTypeObject GlobalParameterType = "Object"
	GlobalParameterTypeString GlobalParameterType = "String"
)

func PossibleValuesForGlobalParameterType() []string {
	return []string{
		string(GlobalParameterTypeArray),
		string(GlobalParameterTypeBool),
		string(GlobalParameterTypeFloat),
		string(GlobalParameterTypeInt),
		string(GlobalParameterTypeObject),
		string(GlobalParameterTypeString),
	}
}

func parseGlobalParameterType(input string) (*GlobalParameterType, error) {
	vals := map[string]GlobalParameterType{
		"array":  GlobalParameterTypeArray,
		"bool":   GlobalParameterTypeBool,
		"float":  GlobalParameterTypeFloat,
		"int":    GlobalParameterTypeInt,
		"object": GlobalParameterTypeObject,
		"string": GlobalParameterTypeString,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := GlobalParameterType(input)
	return &out, nil
}
