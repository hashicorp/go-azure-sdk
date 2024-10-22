package costallocationrules

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CostAllocationPolicyType string

const (
	CostAllocationPolicyTypeFixedProportion CostAllocationPolicyType = "FixedProportion"
)

func PossibleValuesForCostAllocationPolicyType() []string {
	return []string{
		string(CostAllocationPolicyTypeFixedProportion),
	}
}

func (s *CostAllocationPolicyType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseCostAllocationPolicyType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseCostAllocationPolicyType(input string) (*CostAllocationPolicyType, error) {
	vals := map[string]CostAllocationPolicyType{
		"fixedproportion": CostAllocationPolicyTypeFixedProportion,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := CostAllocationPolicyType(input)
	return &out, nil
}

type CostAllocationResourceType string

const (
	CostAllocationResourceTypeDimension CostAllocationResourceType = "Dimension"
	CostAllocationResourceTypeTag       CostAllocationResourceType = "Tag"
)

func PossibleValuesForCostAllocationResourceType() []string {
	return []string{
		string(CostAllocationResourceTypeDimension),
		string(CostAllocationResourceTypeTag),
	}
}

func (s *CostAllocationResourceType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseCostAllocationResourceType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseCostAllocationResourceType(input string) (*CostAllocationResourceType, error) {
	vals := map[string]CostAllocationResourceType{
		"dimension": CostAllocationResourceTypeDimension,
		"tag":       CostAllocationResourceTypeTag,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := CostAllocationResourceType(input)
	return &out, nil
}

type Reason string

const (
	ReasonAlreadyExists Reason = "AlreadyExists"
	ReasonInvalid       Reason = "Invalid"
	ReasonValid         Reason = "Valid"
)

func PossibleValuesForReason() []string {
	return []string{
		string(ReasonAlreadyExists),
		string(ReasonInvalid),
		string(ReasonValid),
	}
}

func (s *Reason) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseReason(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseReason(input string) (*Reason, error) {
	vals := map[string]Reason{
		"alreadyexists": ReasonAlreadyExists,
		"invalid":       ReasonInvalid,
		"valid":         ReasonValid,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Reason(input)
	return &out, nil
}

type RuleStatus string

const (
	RuleStatusActive     RuleStatus = "Active"
	RuleStatusNotActive  RuleStatus = "NotActive"
	RuleStatusProcessing RuleStatus = "Processing"
)

func PossibleValuesForRuleStatus() []string {
	return []string{
		string(RuleStatusActive),
		string(RuleStatusNotActive),
		string(RuleStatusProcessing),
	}
}

func (s *RuleStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseRuleStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseRuleStatus(input string) (*RuleStatus, error) {
	vals := map[string]RuleStatus{
		"active":     RuleStatusActive,
		"notactive":  RuleStatusNotActive,
		"processing": RuleStatusProcessing,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := RuleStatus(input)
	return &out, nil
}
