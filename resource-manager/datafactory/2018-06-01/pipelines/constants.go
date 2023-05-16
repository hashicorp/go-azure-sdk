package pipelines

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DependencyCondition string

const (
	DependencyConditionCompleted DependencyCondition = "Completed"
	DependencyConditionFailed    DependencyCondition = "Failed"
	DependencyConditionSkipped   DependencyCondition = "Skipped"
	DependencyConditionSucceeded DependencyCondition = "Succeeded"
)

func PossibleValuesForDependencyCondition() []string {
	return []string{
		string(DependencyConditionCompleted),
		string(DependencyConditionFailed),
		string(DependencyConditionSkipped),
		string(DependencyConditionSucceeded),
	}
}

func (s *DependencyCondition) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseDependencyCondition(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseDependencyCondition(input string) (*DependencyCondition, error) {
	vals := map[string]DependencyCondition{
		"completed": DependencyConditionCompleted,
		"failed":    DependencyConditionFailed,
		"skipped":   DependencyConditionSkipped,
		"succeeded": DependencyConditionSucceeded,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DependencyCondition(input)
	return &out, nil
}

type ParameterType string

const (
	ParameterTypeArray        ParameterType = "Array"
	ParameterTypeBool         ParameterType = "Bool"
	ParameterTypeFloat        ParameterType = "Float"
	ParameterTypeInt          ParameterType = "Int"
	ParameterTypeObject       ParameterType = "Object"
	ParameterTypeSecureString ParameterType = "SecureString"
	ParameterTypeString       ParameterType = "String"
)

func PossibleValuesForParameterType() []string {
	return []string{
		string(ParameterTypeArray),
		string(ParameterTypeBool),
		string(ParameterTypeFloat),
		string(ParameterTypeInt),
		string(ParameterTypeObject),
		string(ParameterTypeSecureString),
		string(ParameterTypeString),
	}
}

func (s *ParameterType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseParameterType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseParameterType(input string) (*ParameterType, error) {
	vals := map[string]ParameterType{
		"array":        ParameterTypeArray,
		"bool":         ParameterTypeBool,
		"float":        ParameterTypeFloat,
		"int":          ParameterTypeInt,
		"object":       ParameterTypeObject,
		"securestring": ParameterTypeSecureString,
		"string":       ParameterTypeString,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ParameterType(input)
	return &out, nil
}

type VariableType string

const (
	VariableTypeArray  VariableType = "Array"
	VariableTypeBool   VariableType = "Bool"
	VariableTypeString VariableType = "String"
)

func PossibleValuesForVariableType() []string {
	return []string{
		string(VariableTypeArray),
		string(VariableTypeBool),
		string(VariableTypeString),
	}
}

func (s *VariableType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseVariableType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseVariableType(input string) (*VariableType, error) {
	vals := map[string]VariableType{
		"array":  VariableTypeArray,
		"bool":   VariableTypeBool,
		"string": VariableTypeString,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := VariableType(input)
	return &out, nil
}
