package scripts

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type OptionalParamEnum string

const (
	OptionalParamEnumOptional OptionalParamEnum = "Optional"
	OptionalParamEnumRequired OptionalParamEnum = "Required"
)

func PossibleValuesForOptionalParamEnum() []string {
	return []string{
		string(OptionalParamEnumOptional),
		string(OptionalParamEnumRequired),
	}
}

func (s *OptionalParamEnum) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseOptionalParamEnum(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseOptionalParamEnum(input string) (*OptionalParamEnum, error) {
	vals := map[string]OptionalParamEnum{
		"optional": OptionalParamEnumOptional,
		"required": OptionalParamEnumRequired,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := OptionalParamEnum(input)
	return &out, nil
}

type ScriptExecutionParameterType string

const (
	ScriptExecutionParameterTypeCredential  ScriptExecutionParameterType = "Credential"
	ScriptExecutionParameterTypeSecureValue ScriptExecutionParameterType = "SecureValue"
	ScriptExecutionParameterTypeValue       ScriptExecutionParameterType = "Value"
)

func PossibleValuesForScriptExecutionParameterType() []string {
	return []string{
		string(ScriptExecutionParameterTypeCredential),
		string(ScriptExecutionParameterTypeSecureValue),
		string(ScriptExecutionParameterTypeValue),
	}
}

func (s *ScriptExecutionParameterType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseScriptExecutionParameterType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseScriptExecutionParameterType(input string) (*ScriptExecutionParameterType, error) {
	vals := map[string]ScriptExecutionParameterType{
		"credential":  ScriptExecutionParameterTypeCredential,
		"securevalue": ScriptExecutionParameterTypeSecureValue,
		"value":       ScriptExecutionParameterTypeValue,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ScriptExecutionParameterType(input)
	return &out, nil
}

type ScriptExecutionProvisioningState string

const (
	ScriptExecutionProvisioningStateCanceled   ScriptExecutionProvisioningState = "Canceled"
	ScriptExecutionProvisioningStateCancelled  ScriptExecutionProvisioningState = "Cancelled"
	ScriptExecutionProvisioningStateCancelling ScriptExecutionProvisioningState = "Cancelling"
	ScriptExecutionProvisioningStateDeleting   ScriptExecutionProvisioningState = "Deleting"
	ScriptExecutionProvisioningStateFailed     ScriptExecutionProvisioningState = "Failed"
	ScriptExecutionProvisioningStatePending    ScriptExecutionProvisioningState = "Pending"
	ScriptExecutionProvisioningStateRunning    ScriptExecutionProvisioningState = "Running"
	ScriptExecutionProvisioningStateSucceeded  ScriptExecutionProvisioningState = "Succeeded"
)

func PossibleValuesForScriptExecutionProvisioningState() []string {
	return []string{
		string(ScriptExecutionProvisioningStateCanceled),
		string(ScriptExecutionProvisioningStateCancelled),
		string(ScriptExecutionProvisioningStateCancelling),
		string(ScriptExecutionProvisioningStateDeleting),
		string(ScriptExecutionProvisioningStateFailed),
		string(ScriptExecutionProvisioningStatePending),
		string(ScriptExecutionProvisioningStateRunning),
		string(ScriptExecutionProvisioningStateSucceeded),
	}
}

func (s *ScriptExecutionProvisioningState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseScriptExecutionProvisioningState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseScriptExecutionProvisioningState(input string) (*ScriptExecutionProvisioningState, error) {
	vals := map[string]ScriptExecutionProvisioningState{
		"canceled":   ScriptExecutionProvisioningStateCanceled,
		"cancelled":  ScriptExecutionProvisioningStateCancelled,
		"cancelling": ScriptExecutionProvisioningStateCancelling,
		"deleting":   ScriptExecutionProvisioningStateDeleting,
		"failed":     ScriptExecutionProvisioningStateFailed,
		"pending":    ScriptExecutionProvisioningStatePending,
		"running":    ScriptExecutionProvisioningStateRunning,
		"succeeded":  ScriptExecutionProvisioningStateSucceeded,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ScriptExecutionProvisioningState(input)
	return &out, nil
}

type ScriptOutputStreamType string

const (
	ScriptOutputStreamTypeError       ScriptOutputStreamType = "Error"
	ScriptOutputStreamTypeInformation ScriptOutputStreamType = "Information"
	ScriptOutputStreamTypeOutput      ScriptOutputStreamType = "Output"
	ScriptOutputStreamTypeWarning     ScriptOutputStreamType = "Warning"
)

func PossibleValuesForScriptOutputStreamType() []string {
	return []string{
		string(ScriptOutputStreamTypeError),
		string(ScriptOutputStreamTypeInformation),
		string(ScriptOutputStreamTypeOutput),
		string(ScriptOutputStreamTypeWarning),
	}
}

func (s *ScriptOutputStreamType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseScriptOutputStreamType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseScriptOutputStreamType(input string) (*ScriptOutputStreamType, error) {
	vals := map[string]ScriptOutputStreamType{
		"error":       ScriptOutputStreamTypeError,
		"information": ScriptOutputStreamTypeInformation,
		"output":      ScriptOutputStreamTypeOutput,
		"warning":     ScriptOutputStreamTypeWarning,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ScriptOutputStreamType(input)
	return &out, nil
}

type ScriptParameterTypes string

const (
	ScriptParameterTypesBool         ScriptParameterTypes = "Bool"
	ScriptParameterTypesCredential   ScriptParameterTypes = "Credential"
	ScriptParameterTypesFloat        ScriptParameterTypes = "Float"
	ScriptParameterTypesInt          ScriptParameterTypes = "Int"
	ScriptParameterTypesSecureString ScriptParameterTypes = "SecureString"
	ScriptParameterTypesString       ScriptParameterTypes = "String"
)

func PossibleValuesForScriptParameterTypes() []string {
	return []string{
		string(ScriptParameterTypesBool),
		string(ScriptParameterTypesCredential),
		string(ScriptParameterTypesFloat),
		string(ScriptParameterTypesInt),
		string(ScriptParameterTypesSecureString),
		string(ScriptParameterTypesString),
	}
}

func (s *ScriptParameterTypes) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseScriptParameterTypes(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseScriptParameterTypes(input string) (*ScriptParameterTypes, error) {
	vals := map[string]ScriptParameterTypes{
		"bool":         ScriptParameterTypesBool,
		"credential":   ScriptParameterTypesCredential,
		"float":        ScriptParameterTypesFloat,
		"int":          ScriptParameterTypesInt,
		"securestring": ScriptParameterTypesSecureString,
		"string":       ScriptParameterTypesString,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ScriptParameterTypes(input)
	return &out, nil
}

type VisibilityParameterEnum string

const (
	VisibilityParameterEnumHidden  VisibilityParameterEnum = "Hidden"
	VisibilityParameterEnumVisible VisibilityParameterEnum = "Visible"
)

func PossibleValuesForVisibilityParameterEnum() []string {
	return []string{
		string(VisibilityParameterEnumHidden),
		string(VisibilityParameterEnumVisible),
	}
}

func (s *VisibilityParameterEnum) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseVisibilityParameterEnum(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseVisibilityParameterEnum(input string) (*VisibilityParameterEnum, error) {
	vals := map[string]VisibilityParameterEnum{
		"hidden":  VisibilityParameterEnumHidden,
		"visible": VisibilityParameterEnumVisible,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := VisibilityParameterEnum(input)
	return &out, nil
}
