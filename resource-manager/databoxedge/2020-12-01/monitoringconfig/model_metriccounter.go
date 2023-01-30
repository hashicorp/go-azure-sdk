package monitoringconfig

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MetricCounter struct {
	AdditionalDimensions *[]MetricDimension `json:"additionalDimensions,omitempty"`
	DimensionFilter      *[]MetricDimension `json:"dimensionFilter,omitempty"`
	Instance             *string            `json:"instance,omitempty"`
	Name                 string             `json:"name"`
}
