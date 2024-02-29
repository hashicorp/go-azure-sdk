package alertrules

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ConditionOperator string

const (
	ConditionOperatorGreaterThan        ConditionOperator = "GreaterThan"
	ConditionOperatorGreaterThanOrEqual ConditionOperator = "GreaterThanOrEqual"
	ConditionOperatorLessThan           ConditionOperator = "LessThan"
	ConditionOperatorLessThanOrEqual    ConditionOperator = "LessThanOrEqual"
)

func PossibleValuesForConditionOperator() []string {
	return []string{
		string(ConditionOperatorGreaterThan),
		string(ConditionOperatorGreaterThanOrEqual),
		string(ConditionOperatorLessThan),
		string(ConditionOperatorLessThanOrEqual),
	}
}

func (s *ConditionOperator) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseConditionOperator(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseConditionOperator(input string) (*ConditionOperator, error) {
	vals := map[string]ConditionOperator{
		"greaterthan":        ConditionOperatorGreaterThan,
		"greaterthanorequal": ConditionOperatorGreaterThanOrEqual,
		"lessthan":           ConditionOperatorLessThan,
		"lessthanorequal":    ConditionOperatorLessThanOrEqual,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ConditionOperator(input)
	return &out, nil
}

type TimeAggregationOperator string

const (
	TimeAggregationOperatorAverage TimeAggregationOperator = "Average"
	TimeAggregationOperatorLast    TimeAggregationOperator = "Last"
	TimeAggregationOperatorMaximum TimeAggregationOperator = "Maximum"
	TimeAggregationOperatorMinimum TimeAggregationOperator = "Minimum"
	TimeAggregationOperatorTotal   TimeAggregationOperator = "Total"
)

func PossibleValuesForTimeAggregationOperator() []string {
	return []string{
		string(TimeAggregationOperatorAverage),
		string(TimeAggregationOperatorLast),
		string(TimeAggregationOperatorMaximum),
		string(TimeAggregationOperatorMinimum),
		string(TimeAggregationOperatorTotal),
	}
}

func (s *TimeAggregationOperator) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseTimeAggregationOperator(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseTimeAggregationOperator(input string) (*TimeAggregationOperator, error) {
	vals := map[string]TimeAggregationOperator{
		"average": TimeAggregationOperatorAverage,
		"last":    TimeAggregationOperatorLast,
		"maximum": TimeAggregationOperatorMaximum,
		"minimum": TimeAggregationOperatorMinimum,
		"total":   TimeAggregationOperatorTotal,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := TimeAggregationOperator(input)
	return &out, nil
}
