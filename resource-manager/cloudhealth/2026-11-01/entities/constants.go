package entities

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AggregationType string

const (
	AggregationTypeBestOf        AggregationType = "BestOf"
	AggregationTypeMaxNotHealthy AggregationType = "MaxNotHealthy"
	AggregationTypeMinHealthy    AggregationType = "MinHealthy"
	AggregationTypeWorstOf       AggregationType = "WorstOf"
)

func PossibleValuesForAggregationType() []string {
	return []string{
		string(AggregationTypeBestOf),
		string(AggregationTypeMaxNotHealthy),
		string(AggregationTypeMinHealthy),
		string(AggregationTypeWorstOf),
	}
}

func (s *AggregationType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAggregationType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAggregationType(input string) (*AggregationType, error) {
	vals := map[string]AggregationType{
		"bestof":        AggregationTypeBestOf,
		"maxnothealthy": AggregationTypeMaxNotHealthy,
		"minhealthy":    AggregationTypeMinHealthy,
		"worstof":       AggregationTypeWorstOf,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AggregationType(input)
	return &out, nil
}

type AggregationUnit string

const (
	AggregationUnitAbsolute   AggregationUnit = "Absolute"
	AggregationUnitPercentage AggregationUnit = "Percentage"
)

func PossibleValuesForAggregationUnit() []string {
	return []string{
		string(AggregationUnitAbsolute),
		string(AggregationUnitPercentage),
	}
}

func (s *AggregationUnit) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAggregationUnit(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAggregationUnit(input string) (*AggregationUnit, error) {
	vals := map[string]AggregationUnit{
		"absolute":   AggregationUnitAbsolute,
		"percentage": AggregationUnitPercentage,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AggregationUnit(input)
	return &out, nil
}

type AlertSeverity string

const (
	AlertSeveritySevFour  AlertSeverity = "Sev4"
	AlertSeveritySevOne   AlertSeverity = "Sev1"
	AlertSeveritySevThree AlertSeverity = "Sev3"
	AlertSeveritySevTwo   AlertSeverity = "Sev2"
	AlertSeveritySevZero  AlertSeverity = "Sev0"
)

func PossibleValuesForAlertSeverity() []string {
	return []string{
		string(AlertSeveritySevFour),
		string(AlertSeveritySevOne),
		string(AlertSeveritySevThree),
		string(AlertSeveritySevTwo),
		string(AlertSeveritySevZero),
	}
}

func (s *AlertSeverity) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAlertSeverity(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAlertSeverity(input string) (*AlertSeverity, error) {
	vals := map[string]AlertSeverity{
		"sev4": AlertSeveritySevFour,
		"sev1": AlertSeveritySevOne,
		"sev3": AlertSeveritySevThree,
		"sev2": AlertSeveritySevTwo,
		"sev0": AlertSeveritySevZero,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AlertSeverity(input)
	return &out, nil
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

type EntityImpact string

const (
	EntityImpactLimited    EntityImpact = "Limited"
	EntityImpactStandard   EntityImpact = "Standard"
	EntityImpactSuppressed EntityImpact = "Suppressed"
)

func PossibleValuesForEntityImpact() []string {
	return []string{
		string(EntityImpactLimited),
		string(EntityImpactStandard),
		string(EntityImpactSuppressed),
	}
}

func (s *EntityImpact) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseEntityImpact(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseEntityImpact(input string) (*EntityImpact, error) {
	vals := map[string]EntityImpact{
		"limited":    EntityImpactLimited,
		"standard":   EntityImpactStandard,
		"suppressed": EntityImpactSuppressed,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := EntityImpact(input)
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

type HealthState string

const (
	HealthStateDegraded  HealthState = "Degraded"
	HealthStateHealthy   HealthState = "Healthy"
	HealthStateUnhealthy HealthState = "Unhealthy"
	HealthStateUnknown   HealthState = "Unknown"
)

func PossibleValuesForHealthState() []string {
	return []string{
		string(HealthStateDegraded),
		string(HealthStateHealthy),
		string(HealthStateUnhealthy),
		string(HealthStateUnknown),
	}
}

func (s *HealthState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseHealthState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseHealthState(input string) (*HealthState, error) {
	vals := map[string]HealthState{
		"degraded":  HealthStateDegraded,
		"healthy":   HealthStateHealthy,
		"unhealthy": HealthStateUnhealthy,
		"unknown":   HealthStateUnknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := HealthState(input)
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

type ResourceHealthAvailabilityState string

const (
	ResourceHealthAvailabilityStateAvailable   ResourceHealthAvailabilityState = "Available"
	ResourceHealthAvailabilityStateDegraded    ResourceHealthAvailabilityState = "Degraded"
	ResourceHealthAvailabilityStateUnavailable ResourceHealthAvailabilityState = "Unavailable"
	ResourceHealthAvailabilityStateUnknown     ResourceHealthAvailabilityState = "Unknown"
)

func PossibleValuesForResourceHealthAvailabilityState() []string {
	return []string{
		string(ResourceHealthAvailabilityStateAvailable),
		string(ResourceHealthAvailabilityStateDegraded),
		string(ResourceHealthAvailabilityStateUnavailable),
		string(ResourceHealthAvailabilityStateUnknown),
	}
}

func (s *ResourceHealthAvailabilityState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseResourceHealthAvailabilityState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseResourceHealthAvailabilityState(input string) (*ResourceHealthAvailabilityState, error) {
	vals := map[string]ResourceHealthAvailabilityState{
		"available":   ResourceHealthAvailabilityStateAvailable,
		"degraded":    ResourceHealthAvailabilityStateDegraded,
		"unavailable": ResourceHealthAvailabilityStateUnavailable,
		"unknown":     ResourceHealthAvailabilityStateUnknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ResourceHealthAvailabilityState(input)
	return &out, nil
}

type ResourceHealthAvailabilityStateSignalBehavior string

const (
	ResourceHealthAvailabilityStateSignalBehaviorDisabled ResourceHealthAvailabilityStateSignalBehavior = "Disabled"
	ResourceHealthAvailabilityStateSignalBehaviorEnabled  ResourceHealthAvailabilityStateSignalBehavior = "Enabled"
)

func PossibleValuesForResourceHealthAvailabilityStateSignalBehavior() []string {
	return []string{
		string(ResourceHealthAvailabilityStateSignalBehaviorDisabled),
		string(ResourceHealthAvailabilityStateSignalBehaviorEnabled),
	}
}

func (s *ResourceHealthAvailabilityStateSignalBehavior) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseResourceHealthAvailabilityStateSignalBehavior(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseResourceHealthAvailabilityStateSignalBehavior(input string) (*ResourceHealthAvailabilityStateSignalBehavior, error) {
	vals := map[string]ResourceHealthAvailabilityStateSignalBehavior{
		"disabled": ResourceHealthAvailabilityStateSignalBehaviorDisabled,
		"enabled":  ResourceHealthAvailabilityStateSignalBehaviorEnabled,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ResourceHealthAvailabilityStateSignalBehavior(input)
	return &out, nil
}

type ResourceHealthCategory string

const (
	ResourceHealthCategoryPlanned   ResourceHealthCategory = "Planned"
	ResourceHealthCategoryUnplanned ResourceHealthCategory = "Unplanned"
)

func PossibleValuesForResourceHealthCategory() []string {
	return []string{
		string(ResourceHealthCategoryPlanned),
		string(ResourceHealthCategoryUnplanned),
	}
}

func (s *ResourceHealthCategory) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseResourceHealthCategory(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseResourceHealthCategory(input string) (*ResourceHealthCategory, error) {
	vals := map[string]ResourceHealthCategory{
		"planned":   ResourceHealthCategoryPlanned,
		"unplanned": ResourceHealthCategoryUnplanned,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ResourceHealthCategory(input)
	return &out, nil
}

type ResourceHealthReasonChronicity string

const (
	ResourceHealthReasonChronicityPersistent ResourceHealthReasonChronicity = "Persistent"
	ResourceHealthReasonChronicityTransient  ResourceHealthReasonChronicity = "Transient"
)

func PossibleValuesForResourceHealthReasonChronicity() []string {
	return []string{
		string(ResourceHealthReasonChronicityPersistent),
		string(ResourceHealthReasonChronicityTransient),
	}
}

func (s *ResourceHealthReasonChronicity) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseResourceHealthReasonChronicity(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseResourceHealthReasonChronicity(input string) (*ResourceHealthReasonChronicity, error) {
	vals := map[string]ResourceHealthReasonChronicity{
		"persistent": ResourceHealthReasonChronicityPersistent,
		"transient":  ResourceHealthReasonChronicityTransient,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ResourceHealthReasonChronicity(input)
	return &out, nil
}

type ResourceHealthReasonType string

const (
	ResourceHealthReasonTypePlanned       ResourceHealthReasonType = "Planned"
	ResourceHealthReasonTypeUnplanned     ResourceHealthReasonType = "Unplanned"
	ResourceHealthReasonTypeUserInitiated ResourceHealthReasonType = "UserInitiated"
)

func PossibleValuesForResourceHealthReasonType() []string {
	return []string{
		string(ResourceHealthReasonTypePlanned),
		string(ResourceHealthReasonTypeUnplanned),
		string(ResourceHealthReasonTypeUserInitiated),
	}
}

func (s *ResourceHealthReasonType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseResourceHealthReasonType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseResourceHealthReasonType(input string) (*ResourceHealthReasonType, error) {
	vals := map[string]ResourceHealthReasonType{
		"planned":       ResourceHealthReasonTypePlanned,
		"unplanned":     ResourceHealthReasonTypeUnplanned,
		"userinitiated": ResourceHealthReasonTypeUserInitiated,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ResourceHealthReasonType(input)
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

type SignalRecommendationKind string

const (
	SignalRecommendationKindAzureResourceMetric    SignalRecommendationKind = "AzureResourceMetric"
	SignalRecommendationKindLogAnalyticsQuery      SignalRecommendationKind = "LogAnalyticsQuery"
	SignalRecommendationKindPrometheusMetricsQuery SignalRecommendationKind = "PrometheusMetricsQuery"
)

func PossibleValuesForSignalRecommendationKind() []string {
	return []string{
		string(SignalRecommendationKindAzureResourceMetric),
		string(SignalRecommendationKindLogAnalyticsQuery),
		string(SignalRecommendationKindPrometheusMetricsQuery),
	}
}

func (s *SignalRecommendationKind) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSignalRecommendationKind(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSignalRecommendationKind(input string) (*SignalRecommendationKind, error) {
	vals := map[string]SignalRecommendationKind{
		"azureresourcemetric":    SignalRecommendationKindAzureResourceMetric,
		"loganalyticsquery":      SignalRecommendationKindLogAnalyticsQuery,
		"prometheusmetricsquery": SignalRecommendationKindPrometheusMetricsQuery,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SignalRecommendationKind(input)
	return &out, nil
}
