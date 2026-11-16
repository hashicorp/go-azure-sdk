package entities

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DependenciesSignalGroupV2 struct {
	AggregationType    AggregationType  `json:"aggregationType"`
	DegradedThreshold  *float64         `json:"degradedThreshold,omitempty"`
	IgnoreUnknown      *bool            `json:"ignoreUnknown,omitempty"`
	UnhealthyThreshold *float64         `json:"unhealthyThreshold,omitempty"`
	Unit               *AggregationUnit `json:"unit,omitempty"`
}
