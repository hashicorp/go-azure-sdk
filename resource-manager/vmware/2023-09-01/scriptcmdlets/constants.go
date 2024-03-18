package scriptcmdlets

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

type ScriptCmdletAudience string

const (
	ScriptCmdletAudienceAny        ScriptCmdletAudience = "Any"
	ScriptCmdletAudienceAutomation ScriptCmdletAudience = "Automation"
)

func PossibleValuesForScriptCmdletAudience() []string {
	return []string{
		string(ScriptCmdletAudienceAny),
		string(ScriptCmdletAudienceAutomation),
	}
}

func (s *ScriptCmdletAudience) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseScriptCmdletAudience(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseScriptCmdletAudience(input string) (*ScriptCmdletAudience, error) {
	vals := map[string]ScriptCmdletAudience{
		"any":        ScriptCmdletAudienceAny,
		"automation": ScriptCmdletAudienceAutomation,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ScriptCmdletAudience(input)
	return &out, nil
}

type ScriptCmdletProvisioningState string

const (
	ScriptCmdletProvisioningStateCanceled  ScriptCmdletProvisioningState = "Canceled"
	ScriptCmdletProvisioningStateFailed    ScriptCmdletProvisioningState = "Failed"
	ScriptCmdletProvisioningStateSucceeded ScriptCmdletProvisioningState = "Succeeded"
)

func PossibleValuesForScriptCmdletProvisioningState() []string {
	return []string{
		string(ScriptCmdletProvisioningStateCanceled),
		string(ScriptCmdletProvisioningStateFailed),
		string(ScriptCmdletProvisioningStateSucceeded),
	}
}

func (s *ScriptCmdletProvisioningState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseScriptCmdletProvisioningState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseScriptCmdletProvisioningState(input string) (*ScriptCmdletProvisioningState, error) {
	vals := map[string]ScriptCmdletProvisioningState{
		"canceled":  ScriptCmdletProvisioningStateCanceled,
		"failed":    ScriptCmdletProvisioningStateFailed,
		"succeeded": ScriptCmdletProvisioningStateSucceeded,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ScriptCmdletProvisioningState(input)
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
