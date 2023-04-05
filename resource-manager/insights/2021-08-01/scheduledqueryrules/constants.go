package scheduledqueryrules

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AlertSeverity int64

const (
	AlertSeverityFour  AlertSeverity = 4
	AlertSeverityOne   AlertSeverity = 1
	AlertSeverityThree AlertSeverity = 3
	AlertSeverityTwo   AlertSeverity = 2
	AlertSeverityZero  AlertSeverity = 0
)

func PossibleValuesForAlertSeverity() []int64 {
	return []int64{
		int64(AlertSeverityFour),
		int64(AlertSeverityOne),
		int64(AlertSeverityThree),
		int64(AlertSeverityTwo),
		int64(AlertSeverityZero),
	}
}

type ConditionOperator string

const (
	ConditionOperatorEquals             ConditionOperator = "Equals"
	ConditionOperatorGreaterThan        ConditionOperator = "GreaterThan"
	ConditionOperatorGreaterThanOrEqual ConditionOperator = "GreaterThanOrEqual"
	ConditionOperatorLessThan           ConditionOperator = "LessThan"
	ConditionOperatorLessThanOrEqual    ConditionOperator = "LessThanOrEqual"
)

func PossibleValuesForConditionOperator() []string {
	return []string{
		string(ConditionOperatorEquals),
		string(ConditionOperatorGreaterThan),
		string(ConditionOperatorGreaterThanOrEqual),
		string(ConditionOperatorLessThan),
		string(ConditionOperatorLessThanOrEqual),
	}
}

type DimensionOperator string

const (
	DimensionOperatorExclude DimensionOperator = "Exclude"
	DimensionOperatorInclude DimensionOperator = "Include"
)

func PossibleValuesForDimensionOperator() []string {
	return []string{
		string(DimensionOperatorExclude),
		string(DimensionOperatorInclude),
	}
}

type Kind string

const (
	KindLogAlert    Kind = "LogAlert"
	KindLogToMetric Kind = "LogToMetric"
)

func PossibleValuesForKind() []string {
	return []string{
		string(KindLogAlert),
		string(KindLogToMetric),
	}
}

type TimeAggregation string

const (
	TimeAggregationAverage TimeAggregation = "Average"
	TimeAggregationCount   TimeAggregation = "Count"
	TimeAggregationMaximum TimeAggregation = "Maximum"
	TimeAggregationMinimum TimeAggregation = "Minimum"
	TimeAggregationTotal   TimeAggregation = "Total"
)

func PossibleValuesForTimeAggregation() []string {
	return []string{
		string(TimeAggregationAverage),
		string(TimeAggregationCount),
		string(TimeAggregationMaximum),
		string(TimeAggregationMinimum),
		string(TimeAggregationTotal),
	}
}
