package summaryrules

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ProvisioningStateEnum string

const (
	ProvisioningStateEnumCanceled  ProvisioningStateEnum = "Canceled"
	ProvisioningStateEnumDeleting  ProvisioningStateEnum = "Deleting"
	ProvisioningStateEnumFailed    ProvisioningStateEnum = "Failed"
	ProvisioningStateEnumSucceeded ProvisioningStateEnum = "Succeeded"
	ProvisioningStateEnumUpdating  ProvisioningStateEnum = "Updating"
)

func PossibleValuesForProvisioningStateEnum() []string {
	return []string{
		string(ProvisioningStateEnumCanceled),
		string(ProvisioningStateEnumDeleting),
		string(ProvisioningStateEnumFailed),
		string(ProvisioningStateEnumSucceeded),
		string(ProvisioningStateEnumUpdating),
	}
}

func (s *ProvisioningStateEnum) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseProvisioningStateEnum(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseProvisioningStateEnum(input string) (*ProvisioningStateEnum, error) {
	vals := map[string]ProvisioningStateEnum{
		"canceled":  ProvisioningStateEnumCanceled,
		"deleting":  ProvisioningStateEnumDeleting,
		"failed":    ProvisioningStateEnumFailed,
		"succeeded": ProvisioningStateEnumSucceeded,
		"updating":  ProvisioningStateEnumUpdating,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ProvisioningStateEnum(input)
	return &out, nil
}

type RuleTypeEnum string

const (
	RuleTypeEnumUser RuleTypeEnum = "User"
)

func PossibleValuesForRuleTypeEnum() []string {
	return []string{
		string(RuleTypeEnumUser),
	}
}

func (s *RuleTypeEnum) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseRuleTypeEnum(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseRuleTypeEnum(input string) (*RuleTypeEnum, error) {
	vals := map[string]RuleTypeEnum{
		"user": RuleTypeEnumUser,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := RuleTypeEnum(input)
	return &out, nil
}

type StatusCodeEnum string

const (
	StatusCodeEnumDataPlaneError StatusCodeEnum = "DataPlaneError"
	StatusCodeEnumUserAction     StatusCodeEnum = "UserAction"
)

func PossibleValuesForStatusCodeEnum() []string {
	return []string{
		string(StatusCodeEnumDataPlaneError),
		string(StatusCodeEnumUserAction),
	}
}

func (s *StatusCodeEnum) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseStatusCodeEnum(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseStatusCodeEnum(input string) (*StatusCodeEnum, error) {
	vals := map[string]StatusCodeEnum{
		"dataplaneerror": StatusCodeEnumDataPlaneError,
		"useraction":     StatusCodeEnumUserAction,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := StatusCodeEnum(input)
	return &out, nil
}

type TimeSelectorEnum string

const (
	TimeSelectorEnumTimeGenerated TimeSelectorEnum = "TimeGenerated"
)

func PossibleValuesForTimeSelectorEnum() []string {
	return []string{
		string(TimeSelectorEnumTimeGenerated),
	}
}

func (s *TimeSelectorEnum) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseTimeSelectorEnum(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseTimeSelectorEnum(input string) (*TimeSelectorEnum, error) {
	vals := map[string]TimeSelectorEnum{
		"timegenerated": TimeSelectorEnumTimeGenerated,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := TimeSelectorEnum(input)
	return &out, nil
}
