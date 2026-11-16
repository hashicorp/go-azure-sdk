package entities

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SignalAggregationGroup struct {
	AggregatedHealthState *HealthState     `json:"aggregatedHealthState,omitempty"`
	AggregationType       *AggregationType `json:"aggregationType,omitempty"`
	DegradedThreshold     *float64         `json:"degradedThreshold,omitempty"`
	DisplayName           *string          `json:"displayName,omitempty"`
	IgnoreUnknown         *bool            `json:"ignoreUnknown,omitempty"`
	Members               []string         `json:"members"`
	Name                  string           `json:"name"`
	UnhealthyThreshold    *float64         `json:"unhealthyThreshold,omitempty"`
	Unit                  *AggregationUnit `json:"unit,omitempty"`
	UnresolvedMembers     *[]string        `json:"unresolvedMembers,omitempty"`
}
