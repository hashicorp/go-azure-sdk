package metricalerts

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AggregationTypeEnum string

const (
	AggregationTypeEnumAverage AggregationTypeEnum = "Average"
	AggregationTypeEnumCount   AggregationTypeEnum = "Count"
	AggregationTypeEnumMaximum AggregationTypeEnum = "Maximum"
	AggregationTypeEnumMinimum AggregationTypeEnum = "Minimum"
	AggregationTypeEnumTotal   AggregationTypeEnum = "Total"
)

func PossibleValuesForAggregationTypeEnum() []string {
	return []string{
		string(AggregationTypeEnumAverage),
		string(AggregationTypeEnumCount),
		string(AggregationTypeEnumMaximum),
		string(AggregationTypeEnumMinimum),
		string(AggregationTypeEnumTotal),
	}
}

type CriterionType string

const (
	CriterionTypeDynamicThresholdCriterion CriterionType = "DynamicThresholdCriterion"
	CriterionTypeStaticThresholdCriterion  CriterionType = "StaticThresholdCriterion"
)

func PossibleValuesForCriterionType() []string {
	return []string{
		string(CriterionTypeDynamicThresholdCriterion),
		string(CriterionTypeStaticThresholdCriterion),
	}
}

type DynamicThresholdOperator string

const (
	DynamicThresholdOperatorGreaterOrLessThan DynamicThresholdOperator = "GreaterOrLessThan"
	DynamicThresholdOperatorGreaterThan       DynamicThresholdOperator = "GreaterThan"
	DynamicThresholdOperatorLessThan          DynamicThresholdOperator = "LessThan"
)

func PossibleValuesForDynamicThresholdOperator() []string {
	return []string{
		string(DynamicThresholdOperatorGreaterOrLessThan),
		string(DynamicThresholdOperatorGreaterThan),
		string(DynamicThresholdOperatorLessThan),
	}
}

type DynamicThresholdSensitivity string

const (
	DynamicThresholdSensitivityHigh   DynamicThresholdSensitivity = "High"
	DynamicThresholdSensitivityLow    DynamicThresholdSensitivity = "Low"
	DynamicThresholdSensitivityMedium DynamicThresholdSensitivity = "Medium"
)

func PossibleValuesForDynamicThresholdSensitivity() []string {
	return []string{
		string(DynamicThresholdSensitivityHigh),
		string(DynamicThresholdSensitivityLow),
		string(DynamicThresholdSensitivityMedium),
	}
}

type Odatatype string

const (
	OdatatypeMicrosoftPointAzurePointMonitorPointMultipleResourceMultipleMetricCriteria Odatatype = "Microsoft.Azure.Monitor.MultipleResourceMultipleMetricCriteria"
	OdatatypeMicrosoftPointAzurePointMonitorPointSingleResourceMultipleMetricCriteria   Odatatype = "Microsoft.Azure.Monitor.SingleResourceMultipleMetricCriteria"
	OdatatypeMicrosoftPointAzurePointMonitorPointWebtestLocationAvailabilityCriteria    Odatatype = "Microsoft.Azure.Monitor.WebtestLocationAvailabilityCriteria"
)

func PossibleValuesForOdatatype() []string {
	return []string{
		string(OdatatypeMicrosoftPointAzurePointMonitorPointMultipleResourceMultipleMetricCriteria),
		string(OdatatypeMicrosoftPointAzurePointMonitorPointSingleResourceMultipleMetricCriteria),
		string(OdatatypeMicrosoftPointAzurePointMonitorPointWebtestLocationAvailabilityCriteria),
	}
}

type Operator string

const (
	OperatorEquals             Operator = "Equals"
	OperatorGreaterThan        Operator = "GreaterThan"
	OperatorGreaterThanOrEqual Operator = "GreaterThanOrEqual"
	OperatorLessThan           Operator = "LessThan"
	OperatorLessThanOrEqual    Operator = "LessThanOrEqual"
)

func PossibleValuesForOperator() []string {
	return []string{
		string(OperatorEquals),
		string(OperatorGreaterThan),
		string(OperatorGreaterThanOrEqual),
		string(OperatorLessThan),
		string(OperatorLessThanOrEqual),
	}
}
