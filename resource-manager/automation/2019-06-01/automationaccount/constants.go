package automationaccount

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AutomationAccountState string

const (
	AutomationAccountStateOk          AutomationAccountState = "Ok"
	AutomationAccountStateSuspended   AutomationAccountState = "Suspended"
	AutomationAccountStateUnavailable AutomationAccountState = "Unavailable"
)

func PossibleValuesForAutomationAccountState() []string {
	return []string{
		string(AutomationAccountStateOk),
		string(AutomationAccountStateSuspended),
		string(AutomationAccountStateUnavailable),
	}
}

func (s *AutomationAccountState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAutomationAccountState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAutomationAccountState(input string) (*AutomationAccountState, error) {
	vals := map[string]AutomationAccountState{
		"ok":          AutomationAccountStateOk,
		"suspended":   AutomationAccountStateSuspended,
		"unavailable": AutomationAccountStateUnavailable,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AutomationAccountState(input)
	return &out, nil
}

type SkuNameEnum string

const (
	SkuNameEnumBasic SkuNameEnum = "Basic"
	SkuNameEnumFree  SkuNameEnum = "Free"
)

func PossibleValuesForSkuNameEnum() []string {
	return []string{
		string(SkuNameEnumBasic),
		string(SkuNameEnumFree),
	}
}

func (s *SkuNameEnum) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSkuNameEnum(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSkuNameEnum(input string) (*SkuNameEnum, error) {
	vals := map[string]SkuNameEnum{
		"basic": SkuNameEnumBasic,
		"free":  SkuNameEnumFree,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SkuNameEnum(input)
	return &out, nil
}
