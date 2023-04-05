package autoscalesettings

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ComparisonOperationType string

const (
	ComparisonOperationTypeEquals             ComparisonOperationType = "Equals"
	ComparisonOperationTypeGreaterThan        ComparisonOperationType = "GreaterThan"
	ComparisonOperationTypeGreaterThanOrEqual ComparisonOperationType = "GreaterThanOrEqual"
	ComparisonOperationTypeLessThan           ComparisonOperationType = "LessThan"
	ComparisonOperationTypeLessThanOrEqual    ComparisonOperationType = "LessThanOrEqual"
	ComparisonOperationTypeNotEquals          ComparisonOperationType = "NotEquals"
)

func PossibleValuesForComparisonOperationType() []string {
	return []string{
		string(ComparisonOperationTypeEquals),
		string(ComparisonOperationTypeGreaterThan),
		string(ComparisonOperationTypeGreaterThanOrEqual),
		string(ComparisonOperationTypeLessThan),
		string(ComparisonOperationTypeLessThanOrEqual),
		string(ComparisonOperationTypeNotEquals),
	}
}

type MetricStatisticType string

const (
	MetricStatisticTypeAverage MetricStatisticType = "Average"
	MetricStatisticTypeCount   MetricStatisticType = "Count"
	MetricStatisticTypeMax     MetricStatisticType = "Max"
	MetricStatisticTypeMin     MetricStatisticType = "Min"
	MetricStatisticTypeSum     MetricStatisticType = "Sum"
)

func PossibleValuesForMetricStatisticType() []string {
	return []string{
		string(MetricStatisticTypeAverage),
		string(MetricStatisticTypeCount),
		string(MetricStatisticTypeMax),
		string(MetricStatisticTypeMin),
		string(MetricStatisticTypeSum),
	}
}

type OperationType string

const (
	OperationTypeScale OperationType = "Scale"
)

func PossibleValuesForOperationType() []string {
	return []string{
		string(OperationTypeScale),
	}
}

type PredictiveAutoscalePolicyScaleMode string

const (
	PredictiveAutoscalePolicyScaleModeDisabled     PredictiveAutoscalePolicyScaleMode = "Disabled"
	PredictiveAutoscalePolicyScaleModeEnabled      PredictiveAutoscalePolicyScaleMode = "Enabled"
	PredictiveAutoscalePolicyScaleModeForecastOnly PredictiveAutoscalePolicyScaleMode = "ForecastOnly"
)

func PossibleValuesForPredictiveAutoscalePolicyScaleMode() []string {
	return []string{
		string(PredictiveAutoscalePolicyScaleModeDisabled),
		string(PredictiveAutoscalePolicyScaleModeEnabled),
		string(PredictiveAutoscalePolicyScaleModeForecastOnly),
	}
}

type RecurrenceFrequency string

const (
	RecurrenceFrequencyDay    RecurrenceFrequency = "Day"
	RecurrenceFrequencyHour   RecurrenceFrequency = "Hour"
	RecurrenceFrequencyMinute RecurrenceFrequency = "Minute"
	RecurrenceFrequencyMonth  RecurrenceFrequency = "Month"
	RecurrenceFrequencyNone   RecurrenceFrequency = "None"
	RecurrenceFrequencySecond RecurrenceFrequency = "Second"
	RecurrenceFrequencyWeek   RecurrenceFrequency = "Week"
	RecurrenceFrequencyYear   RecurrenceFrequency = "Year"
)

func PossibleValuesForRecurrenceFrequency() []string {
	return []string{
		string(RecurrenceFrequencyDay),
		string(RecurrenceFrequencyHour),
		string(RecurrenceFrequencyMinute),
		string(RecurrenceFrequencyMonth),
		string(RecurrenceFrequencyNone),
		string(RecurrenceFrequencySecond),
		string(RecurrenceFrequencyWeek),
		string(RecurrenceFrequencyYear),
	}
}

type ScaleDirection string

const (
	ScaleDirectionDecrease ScaleDirection = "Decrease"
	ScaleDirectionIncrease ScaleDirection = "Increase"
	ScaleDirectionNone     ScaleDirection = "None"
)

func PossibleValuesForScaleDirection() []string {
	return []string{
		string(ScaleDirectionDecrease),
		string(ScaleDirectionIncrease),
		string(ScaleDirectionNone),
	}
}

type ScaleRuleMetricDimensionOperationType string

const (
	ScaleRuleMetricDimensionOperationTypeEquals    ScaleRuleMetricDimensionOperationType = "Equals"
	ScaleRuleMetricDimensionOperationTypeNotEquals ScaleRuleMetricDimensionOperationType = "NotEquals"
)

func PossibleValuesForScaleRuleMetricDimensionOperationType() []string {
	return []string{
		string(ScaleRuleMetricDimensionOperationTypeEquals),
		string(ScaleRuleMetricDimensionOperationTypeNotEquals),
	}
}

type ScaleType string

const (
	ScaleTypeChangeCount             ScaleType = "ChangeCount"
	ScaleTypeExactCount              ScaleType = "ExactCount"
	ScaleTypePercentChangeCount      ScaleType = "PercentChangeCount"
	ScaleTypeServiceAllowedNextValue ScaleType = "ServiceAllowedNextValue"
)

func PossibleValuesForScaleType() []string {
	return []string{
		string(ScaleTypeChangeCount),
		string(ScaleTypeExactCount),
		string(ScaleTypePercentChangeCount),
		string(ScaleTypeServiceAllowedNextValue),
	}
}

type TimeAggregationType string

const (
	TimeAggregationTypeAverage TimeAggregationType = "Average"
	TimeAggregationTypeCount   TimeAggregationType = "Count"
	TimeAggregationTypeLast    TimeAggregationType = "Last"
	TimeAggregationTypeMaximum TimeAggregationType = "Maximum"
	TimeAggregationTypeMinimum TimeAggregationType = "Minimum"
	TimeAggregationTypeTotal   TimeAggregationType = "Total"
)

func PossibleValuesForTimeAggregationType() []string {
	return []string{
		string(TimeAggregationTypeAverage),
		string(TimeAggregationTypeCount),
		string(TimeAggregationTypeLast),
		string(TimeAggregationTypeMaximum),
		string(TimeAggregationTypeMinimum),
		string(TimeAggregationTypeTotal),
	}
}
