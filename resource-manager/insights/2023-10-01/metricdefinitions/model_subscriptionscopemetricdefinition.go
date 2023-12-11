package metricdefinitions

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SubscriptionScopeMetricDefinition struct {
	Category                  *string                  `json:"category,omitempty"`
	Dimensions                *[]LocalizableString     `json:"dimensions,omitempty"`
	DisplayDescription        *string                  `json:"displayDescription,omitempty"`
	Id                        *string                  `json:"id,omitempty"`
	IsDimensionRequired       *bool                    `json:"isDimensionRequired,omitempty"`
	MetricAvailabilities      *[]MetricAvailability    `json:"metricAvailabilities,omitempty"`
	MetricClass               *MetricClass             `json:"metricClass,omitempty"`
	Name                      *LocalizableString       `json:"name,omitempty"`
	Namespace                 *string                  `json:"namespace,omitempty"`
	PrimaryAggregationType    *MetricAggregationType   `json:"primaryAggregationType,omitempty"`
	ResourceId                *string                  `json:"resourceId,omitempty"`
	SupportedAggregationTypes *[]MetricAggregationType `json:"supportedAggregationTypes,omitempty"`
	Unit                      *MetricUnit              `json:"unit,omitempty"`
}
