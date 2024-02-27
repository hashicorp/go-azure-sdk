package smartgroups

import (
	"time"

	"github.com/hashicorp/go-azure-helpers/lang/dates"
)

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

func (o *SmartGroupProperties) GetLastModifiedDateTimeAsTime() (*time.Time, error) {
	if o.LastModifiedDateTime == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.LastModifiedDateTime, "2006-01-02T15:04:05Z07:00")
}

func (o *SmartGroupProperties) GetStartDateTimeAsTime() (*time.Time, error) {
	if o.StartDateTime == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.StartDateTime, "2006-01-02T15:04:05Z07:00")
}
