package metricdefinitions

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MetricDefinition struct {
	Id                     *string               `json:"id,omitempty"`
	MetricAvailabilities   *[]MetricAvailability `json:"metricAvailabilities,omitempty"`
	Name                   *LocalizableString    `json:"name,omitempty"`
	PrimaryAggregationType *AggregationType      `json:"primaryAggregationType,omitempty"`
	ResourceId             *string               `json:"resourceId,omitempty"`
	Unit                   *Unit                 `json:"unit,omitempty"`
}
