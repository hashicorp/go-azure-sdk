package sqlpoolsdatamaskingrules

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DataMaskingFunction string

const (
	DataMaskingFunctionCCN     DataMaskingFunction = "CCN"
	DataMaskingFunctionDefault DataMaskingFunction = "Default"
	DataMaskingFunctionEmail   DataMaskingFunction = "Email"
	DataMaskingFunctionNumber  DataMaskingFunction = "Number"
	DataMaskingFunctionSSN     DataMaskingFunction = "SSN"
	DataMaskingFunctionText    DataMaskingFunction = "Text"
)

func PossibleValuesForDataMaskingFunction() []string {
	return []string{
		string(DataMaskingFunctionCCN),
		string(DataMaskingFunctionDefault),
		string(DataMaskingFunctionEmail),
		string(DataMaskingFunctionNumber),
		string(DataMaskingFunctionSSN),
		string(DataMaskingFunctionText),
	}
}

func (s *DataMaskingFunction) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseDataMaskingFunction(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseDataMaskingFunction(input string) (*DataMaskingFunction, error) {
	vals := map[string]DataMaskingFunction{
		"ccn":     DataMaskingFunctionCCN,
		"default": DataMaskingFunctionDefault,
		"email":   DataMaskingFunctionEmail,
		"number":  DataMaskingFunctionNumber,
		"ssn":     DataMaskingFunctionSSN,
		"text":    DataMaskingFunctionText,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DataMaskingFunction(input)
	return &out, nil
}

type DataMaskingRuleState string

const (
	DataMaskingRuleStateDisabled DataMaskingRuleState = "Disabled"
	DataMaskingRuleStateEnabled  DataMaskingRuleState = "Enabled"
)

func PossibleValuesForDataMaskingRuleState() []string {
	return []string{
		string(DataMaskingRuleStateDisabled),
		string(DataMaskingRuleStateEnabled),
	}
}

func (s *DataMaskingRuleState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseDataMaskingRuleState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseDataMaskingRuleState(input string) (*DataMaskingRuleState, error) {
	vals := map[string]DataMaskingRuleState{
		"disabled": DataMaskingRuleStateDisabled,
		"enabled":  DataMaskingRuleStateEnabled,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DataMaskingRuleState(input)
	return &out, nil
}
