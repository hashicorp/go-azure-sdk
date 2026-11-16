package entities

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EntityProperties struct {
	Alerts                  *EntityAlerts                 `json:"alerts,omitempty"`
	CanvasPosition          *EntityCoordinates            `json:"canvasPosition,omitempty"`
	DiscoveredBy            *string                       `json:"discoveredBy,omitempty"`
	DisplayName             *string                       `json:"displayName,omitempty"`
	HealthObjective         *float64                      `json:"healthObjective,omitempty"`
	HealthState             *HealthState                  `json:"healthState,omitempty"`
	Icon                    *IconDefinition               `json:"icon,omitempty"`
	Impact                  *EntityImpact                 `json:"impact,omitempty"`
	ProvisioningState       *HealthModelProvisioningState `json:"provisioningState,omitempty"`
	SignalAggregationGroups *[]SignalAggregationGroup     `json:"signalAggregationGroups,omitempty"`
	SignalGroups            *SignalGroups                 `json:"signalGroups,omitempty"`
	Tags                    *map[string]string            `json:"tags,omitempty"`
}
