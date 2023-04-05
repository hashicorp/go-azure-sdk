package scheduledqueryrules

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AlertSeverity string

const (
	AlertSeverityFour  AlertSeverity = "4"
	AlertSeverityOne   AlertSeverity = "1"
	AlertSeverityThree AlertSeverity = "3"
	AlertSeverityTwo   AlertSeverity = "2"
	AlertSeverityZero  AlertSeverity = "0"
)

func PossibleValuesForAlertSeverity() []string {
	return []string{
		string(AlertSeverityFour),
		string(AlertSeverityOne),
		string(AlertSeverityThree),
		string(AlertSeverityTwo),
		string(AlertSeverityZero),
	}
}

type ConditionalOperator string

const (
	ConditionalOperatorEqual              ConditionalOperator = "Equal"
	ConditionalOperatorGreaterThan        ConditionalOperator = "GreaterThan"
	ConditionalOperatorGreaterThanOrEqual ConditionalOperator = "GreaterThanOrEqual"
	ConditionalOperatorLessThan           ConditionalOperator = "LessThan"
	ConditionalOperatorLessThanOrEqual    ConditionalOperator = "LessThanOrEqual"
)

func PossibleValuesForConditionalOperator() []string {
	return []string{
		string(ConditionalOperatorEqual),
		string(ConditionalOperatorGreaterThan),
		string(ConditionalOperatorGreaterThanOrEqual),
		string(ConditionalOperatorLessThan),
		string(ConditionalOperatorLessThanOrEqual),
	}
}

type Enabled string

const (
	EnabledFalse Enabled = "false"
	EnabledTrue  Enabled = "true"
)

func PossibleValuesForEnabled() []string {
	return []string{
		string(EnabledFalse),
		string(EnabledTrue),
	}
}

type MetricTriggerType string

const (
	MetricTriggerTypeConsecutive MetricTriggerType = "Consecutive"
	MetricTriggerTypeTotal       MetricTriggerType = "Total"
)

func PossibleValuesForMetricTriggerType() []string {
	return []string{
		string(MetricTriggerTypeConsecutive),
		string(MetricTriggerTypeTotal),
	}
}

type Operator string

const (
	OperatorInclude Operator = "Include"
)

func PossibleValuesForOperator() []string {
	return []string{
		string(OperatorInclude),
	}
}

type ProvisioningState string

const (
	ProvisioningStateCanceled  ProvisioningState = "Canceled"
	ProvisioningStateDeploying ProvisioningState = "Deploying"
	ProvisioningStateFailed    ProvisioningState = "Failed"
	ProvisioningStateSucceeded ProvisioningState = "Succeeded"
)

func PossibleValuesForProvisioningState() []string {
	return []string{
		string(ProvisioningStateCanceled),
		string(ProvisioningStateDeploying),
		string(ProvisioningStateFailed),
		string(ProvisioningStateSucceeded),
	}
}

type QueryType string

const (
	QueryTypeResultCount QueryType = "ResultCount"
)

func PossibleValuesForQueryType() []string {
	return []string{
		string(QueryTypeResultCount),
	}
}
