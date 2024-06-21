package alertsmanagements

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Essentials struct {
	ActionStatus                     *ActionStatus     `json:"actionStatus,omitempty"`
	AlertRule                        *string           `json:"alertRule,omitempty"`
	AlertState                       *AlertState       `json:"alertState,omitempty"`
	Description                      *string           `json:"description,omitempty"`
	LastModifiedDateTime             *string           `json:"lastModifiedDateTime,omitempty"`
	LastModifiedUserName             *string           `json:"lastModifiedUserName,omitempty"`
	MonitorCondition                 *MonitorCondition `json:"monitorCondition,omitempty"`
	MonitorConditionResolvedDateTime *string           `json:"monitorConditionResolvedDateTime,omitempty"`
	MonitorService                   *MonitorService   `json:"monitorService,omitempty"`
	Severity                         *Severity         `json:"severity,omitempty"`
	SignalType                       *SignalType       `json:"signalType,omitempty"`
	SmartGroupId                     *string           `json:"smartGroupId,omitempty"`
	SmartGroupingReason              *string           `json:"smartGroupingReason,omitempty"`
	SourceCreatedId                  *string           `json:"sourceCreatedId,omitempty"`
	StartDateTime                    *string           `json:"startDateTime,omitempty"`
	TargetResource                   *string           `json:"targetResource,omitempty"`
	TargetResourceGroup              *string           `json:"targetResourceGroup,omitempty"`
	TargetResourceName               *string           `json:"targetResourceName,omitempty"`
	TargetResourceType               *string           `json:"targetResourceType,omitempty"`
}
