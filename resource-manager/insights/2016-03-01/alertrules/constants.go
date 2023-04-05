package alertrules

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
