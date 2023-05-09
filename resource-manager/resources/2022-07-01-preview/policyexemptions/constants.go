package policyexemptions

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AssignmentScopeValidation string

const (
	AssignmentScopeValidationDefault       AssignmentScopeValidation = "Default"
	AssignmentScopeValidationDoNotValidate AssignmentScopeValidation = "DoNotValidate"
)

func PossibleValuesForAssignmentScopeValidation() []string {
	return []string{
		string(AssignmentScopeValidationDefault),
		string(AssignmentScopeValidationDoNotValidate),
	}
}

func (s *AssignmentScopeValidation) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAssignmentScopeValidation(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAssignmentScopeValidation(input string) (*AssignmentScopeValidation, error) {
	vals := map[string]AssignmentScopeValidation{
		"default":       AssignmentScopeValidationDefault,
		"donotvalidate": AssignmentScopeValidationDoNotValidate,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AssignmentScopeValidation(input)
	return &out, nil
}

type ExemptionCategory string

const (
	ExemptionCategoryMitigated ExemptionCategory = "Mitigated"
	ExemptionCategoryWaiver    ExemptionCategory = "Waiver"
)

func PossibleValuesForExemptionCategory() []string {
	return []string{
		string(ExemptionCategoryMitigated),
		string(ExemptionCategoryWaiver),
	}
}

func (s *ExemptionCategory) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseExemptionCategory(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseExemptionCategory(input string) (*ExemptionCategory, error) {
	vals := map[string]ExemptionCategory{
		"mitigated": ExemptionCategoryMitigated,
		"waiver":    ExemptionCategoryWaiver,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ExemptionCategory(input)
	return &out, nil
}

type SelectorKind string

const (
	SelectorKindPolicyDefinitionReferenceId SelectorKind = "policyDefinitionReferenceId"
	SelectorKindResourceLocation            SelectorKind = "resourceLocation"
	SelectorKindResourceType                SelectorKind = "resourceType"
	SelectorKindResourceWithoutLocation     SelectorKind = "resourceWithoutLocation"
)

func PossibleValuesForSelectorKind() []string {
	return []string{
		string(SelectorKindPolicyDefinitionReferenceId),
		string(SelectorKindResourceLocation),
		string(SelectorKindResourceType),
		string(SelectorKindResourceWithoutLocation),
	}
}

func (s *SelectorKind) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSelectorKind(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSelectorKind(input string) (*SelectorKind, error) {
	vals := map[string]SelectorKind{
		"policydefinitionreferenceid": SelectorKindPolicyDefinitionReferenceId,
		"resourcelocation":            SelectorKindResourceLocation,
		"resourcetype":                SelectorKindResourceType,
		"resourcewithoutlocation":     SelectorKindResourceWithoutLocation,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SelectorKind(input)
	return &out, nil
}
