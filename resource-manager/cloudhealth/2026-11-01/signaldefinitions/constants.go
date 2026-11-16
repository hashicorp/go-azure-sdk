package signaldefinitions

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

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

func (s *DynamicThresholdSensitivity) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseDynamicThresholdSensitivity(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseDynamicThresholdSensitivity(input string) (*DynamicThresholdSensitivity, error) {
	vals := map[string]DynamicThresholdSensitivity{
		"high":   DynamicThresholdSensitivityHigh,
		"low":    DynamicThresholdSensitivityLow,
		"medium": DynamicThresholdSensitivityMedium,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DynamicThresholdSensitivity(input)
	return &out, nil
}

type HealthModelProvisioningState string

const (
	HealthModelProvisioningStateCanceled  HealthModelProvisioningState = "Canceled"
	HealthModelProvisioningStateCreating  HealthModelProvisioningState = "Creating"
	HealthModelProvisioningStateDeleting  HealthModelProvisioningState = "Deleting"
	HealthModelProvisioningStateFailed    HealthModelProvisioningState = "Failed"
	HealthModelProvisioningStateSucceeded HealthModelProvisioningState = "Succeeded"
)

func PossibleValuesForHealthModelProvisioningState() []string {
	return []string{
		string(HealthModelProvisioningStateCanceled),
		string(HealthModelProvisioningStateCreating),
		string(HealthModelProvisioningStateDeleting),
		string(HealthModelProvisioningStateFailed),
		string(HealthModelProvisioningStateSucceeded),
	}
}

func (s *HealthModelProvisioningState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseHealthModelProvisioningState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseHealthModelProvisioningState(input string) (*HealthModelProvisioningState, error) {
	vals := map[string]HealthModelProvisioningState{
		"canceled":  HealthModelProvisioningStateCanceled,
		"creating":  HealthModelProvisioningStateCreating,
		"deleting":  HealthModelProvisioningStateDeleting,
		"failed":    HealthModelProvisioningStateFailed,
		"succeeded": HealthModelProvisioningStateSucceeded,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := HealthModelProvisioningState(input)
	return &out, nil
}

type MetricAggregationType string

const (
	MetricAggregationTypeAverage MetricAggregationType = "Average"
	MetricAggregationTypeCount   MetricAggregationType = "Count"
	MetricAggregationTypeMaximum MetricAggregationType = "Maximum"
	MetricAggregationTypeMinimum MetricAggregationType = "Minimum"
	MetricAggregationTypeNone    MetricAggregationType = "None"
	MetricAggregationTypeTotal   MetricAggregationType = "Total"
)

func PossibleValuesForMetricAggregationType() []string {
	return []string{
		string(MetricAggregationTypeAverage),
		string(MetricAggregationTypeCount),
		string(MetricAggregationTypeMaximum),
		string(MetricAggregationTypeMinimum),
		string(MetricAggregationTypeNone),
		string(MetricAggregationTypeTotal),
	}
}

func (s *MetricAggregationType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseMetricAggregationType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseMetricAggregationType(input string) (*MetricAggregationType, error) {
	vals := map[string]MetricAggregationType{
		"average": MetricAggregationTypeAverage,
		"count":   MetricAggregationTypeCount,
		"maximum": MetricAggregationTypeMaximum,
		"minimum": MetricAggregationTypeMinimum,
		"none":    MetricAggregationTypeNone,
		"total":   MetricAggregationTypeTotal,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := MetricAggregationType(input)
	return &out, nil
}

type RefreshInterval string

const (
	RefreshIntervalPTFiveM      RefreshInterval = "PT5M"
	RefreshIntervalPTOneFiveM   RefreshInterval = "PT15M"
	RefreshIntervalPTOneH       RefreshInterval = "PT1H"
	RefreshIntervalPTOneM       RefreshInterval = "PT1M"
	RefreshIntervalPTOneZeroM   RefreshInterval = "PT10M"
	RefreshIntervalPTThreeZeroM RefreshInterval = "PT30M"
	RefreshIntervalPTTwoH       RefreshInterval = "PT2H"
)

func PossibleValuesForRefreshInterval() []string {
	return []string{
		string(RefreshIntervalPTFiveM),
		string(RefreshIntervalPTOneFiveM),
		string(RefreshIntervalPTOneH),
		string(RefreshIntervalPTOneM),
		string(RefreshIntervalPTOneZeroM),
		string(RefreshIntervalPTThreeZeroM),
		string(RefreshIntervalPTTwoH),
	}
}

func (s *RefreshInterval) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseRefreshInterval(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseRefreshInterval(input string) (*RefreshInterval, error) {
	vals := map[string]RefreshInterval{
		"pt5m":  RefreshIntervalPTFiveM,
		"pt15m": RefreshIntervalPTOneFiveM,
		"pt1h":  RefreshIntervalPTOneH,
		"pt1m":  RefreshIntervalPTOneM,
		"pt10m": RefreshIntervalPTOneZeroM,
		"pt30m": RefreshIntervalPTThreeZeroM,
		"pt2h":  RefreshIntervalPTTwoH,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := RefreshInterval(input)
	return &out, nil
}

type SignalKind string

const (
	SignalKindAzureResourceMetric    SignalKind = "AzureResourceMetric"
	SignalKindExternal               SignalKind = "External"
	SignalKindLogAnalyticsQuery      SignalKind = "LogAnalyticsQuery"
	SignalKindPrometheusMetricsQuery SignalKind = "PrometheusMetricsQuery"
)

func PossibleValuesForSignalKind() []string {
	return []string{
		string(SignalKindAzureResourceMetric),
		string(SignalKindExternal),
		string(SignalKindLogAnalyticsQuery),
		string(SignalKindPrometheusMetricsQuery),
	}
}

func (s *SignalKind) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSignalKind(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSignalKind(input string) (*SignalKind, error) {
	vals := map[string]SignalKind{
		"azureresourcemetric":    SignalKindAzureResourceMetric,
		"external":               SignalKindExternal,
		"loganalyticsquery":      SignalKindLogAnalyticsQuery,
		"prometheusmetricsquery": SignalKindPrometheusMetricsQuery,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SignalKind(input)
	return &out, nil
}

type SignalOperator string

const (
	SignalOperatorDynamic            SignalOperator = "Dynamic"
	SignalOperatorEqual              SignalOperator = "Equal"
	SignalOperatorGreaterThan        SignalOperator = "GreaterThan"
	SignalOperatorGreaterThanOrEqual SignalOperator = "GreaterThanOrEqual"
	SignalOperatorLessThan           SignalOperator = "LessThan"
	SignalOperatorLessThanOrEqual    SignalOperator = "LessThanOrEqual"
	SignalOperatorNotEqual           SignalOperator = "NotEqual"
)

func PossibleValuesForSignalOperator() []string {
	return []string{
		string(SignalOperatorDynamic),
		string(SignalOperatorEqual),
		string(SignalOperatorGreaterThan),
		string(SignalOperatorGreaterThanOrEqual),
		string(SignalOperatorLessThan),
		string(SignalOperatorLessThanOrEqual),
		string(SignalOperatorNotEqual),
	}
}

func (s *SignalOperator) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSignalOperator(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSignalOperator(input string) (*SignalOperator, error) {
	vals := map[string]SignalOperator{
		"dynamic":            SignalOperatorDynamic,
		"equal":              SignalOperatorEqual,
		"greaterthan":        SignalOperatorGreaterThan,
		"greaterthanorequal": SignalOperatorGreaterThanOrEqual,
		"lessthan":           SignalOperatorLessThan,
		"lessthanorequal":    SignalOperatorLessThanOrEqual,
		"notequal":           SignalOperatorNotEqual,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SignalOperator(input)
	return &out, nil
}
