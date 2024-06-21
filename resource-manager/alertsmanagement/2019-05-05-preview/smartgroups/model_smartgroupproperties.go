package smartgroups

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SmartGroupProperties struct {
	AlertSeverities      *[]SmartGroupAggregatedProperty `json:"alertSeverities,omitempty"`
	AlertStates          *[]SmartGroupAggregatedProperty `json:"alertStates,omitempty"`
	AlertsCount          *int64                          `json:"alertsCount,omitempty"`
	LastModifiedDateTime *string                         `json:"lastModifiedDateTime,omitempty"`
	LastModifiedUserName *string                         `json:"lastModifiedUserName,omitempty"`
	MonitorConditions    *[]SmartGroupAggregatedProperty `json:"monitorConditions,omitempty"`
	MonitorServices      *[]SmartGroupAggregatedProperty `json:"monitorServices,omitempty"`
	NextLink             *string                         `json:"nextLink,omitempty"`
	ResourceGroups       *[]SmartGroupAggregatedProperty `json:"resourceGroups,omitempty"`
	ResourceTypes        *[]SmartGroupAggregatedProperty `json:"resourceTypes,omitempty"`
	Resources            *[]SmartGroupAggregatedProperty `json:"resources,omitempty"`
	Severity             *Severity                       `json:"severity,omitempty"`
	SmartGroupState      *State                          `json:"smartGroupState,omitempty"`
	StartDateTime        *string                         `json:"startDateTime,omitempty"`
}
