package pricings

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Code string

const (
	CodeFailed    Code = "Failed"
	CodeSucceeded Code = "Succeeded"
)

func PossibleValuesForCode() []string {
	return []string{
		string(CodeFailed),
		string(CodeSucceeded),
	}
}

func (s *Code) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseCode(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseCode(input string) (*Code, error) {
	vals := map[string]Code{
		"failed":    CodeFailed,
		"succeeded": CodeSucceeded,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Code(input)
	return &out, nil
}

type Enforce string

const (
	EnforceFalse Enforce = "False"
	EnforceTrue  Enforce = "True"
)

func PossibleValuesForEnforce() []string {
	return []string{
		string(EnforceFalse),
		string(EnforceTrue),
	}
}

func (s *Enforce) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseEnforce(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseEnforce(input string) (*Enforce, error) {
	vals := map[string]Enforce{
		"false": EnforceFalse,
		"true":  EnforceTrue,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Enforce(input)
	return &out, nil
}

type Inherited string

const (
	InheritedFalse Inherited = "False"
	InheritedTrue  Inherited = "True"
)

func PossibleValuesForInherited() []string {
	return []string{
		string(InheritedFalse),
		string(InheritedTrue),
	}
}

func (s *Inherited) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseInherited(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseInherited(input string) (*Inherited, error) {
	vals := map[string]Inherited{
		"false": InheritedFalse,
		"true":  InheritedTrue,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Inherited(input)
	return &out, nil
}

type IsEnabled string

const (
	IsEnabledFalse IsEnabled = "False"
	IsEnabledTrue  IsEnabled = "True"
)

func PossibleValuesForIsEnabled() []string {
	return []string{
		string(IsEnabledFalse),
		string(IsEnabledTrue),
	}
}

func (s *IsEnabled) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseIsEnabled(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseIsEnabled(input string) (*IsEnabled, error) {
	vals := map[string]IsEnabled{
		"false": IsEnabledFalse,
		"true":  IsEnabledTrue,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := IsEnabled(input)
	return &out, nil
}

type PricingTier string

const (
	PricingTierFree     PricingTier = "Free"
	PricingTierStandard PricingTier = "Standard"
)

func PossibleValuesForPricingTier() []string {
	return []string{
		string(PricingTierFree),
		string(PricingTierStandard),
	}
}

func (s *PricingTier) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parsePricingTier(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parsePricingTier(input string) (*PricingTier, error) {
	vals := map[string]PricingTier{
		"free":     PricingTierFree,
		"standard": PricingTierStandard,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := PricingTier(input)
	return &out, nil
}

type ResourcesCoverageStatus string

const (
	ResourcesCoverageStatusFullyCovered     ResourcesCoverageStatus = "FullyCovered"
	ResourcesCoverageStatusNotCovered       ResourcesCoverageStatus = "NotCovered"
	ResourcesCoverageStatusPartiallyCovered ResourcesCoverageStatus = "PartiallyCovered"
)

func PossibleValuesForResourcesCoverageStatus() []string {
	return []string{
		string(ResourcesCoverageStatusFullyCovered),
		string(ResourcesCoverageStatusNotCovered),
		string(ResourcesCoverageStatusPartiallyCovered),
	}
}

func (s *ResourcesCoverageStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseResourcesCoverageStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseResourcesCoverageStatus(input string) (*ResourcesCoverageStatus, error) {
	vals := map[string]ResourcesCoverageStatus{
		"fullycovered":     ResourcesCoverageStatusFullyCovered,
		"notcovered":       ResourcesCoverageStatusNotCovered,
		"partiallycovered": ResourcesCoverageStatusPartiallyCovered,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ResourcesCoverageStatus(input)
	return &out, nil
}
