package maintenanceevents

import (
	"time"

	"github.com/hashicorp/go-azure-helpers/lang/dates"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MaintenanceEventActionResponse struct {
	AppliedNow         *bool                   `json:"appliedNow,omitempty"`
	LastUpdatedTime    *string                 `json:"lastUpdatedTime,omitempty"`
	MaintenanceEventId *string                 `json:"maintenanceEventId,omitempty"`
	PlannedEndTime     *string                 `json:"plannedEndTime,omitempty"`
	PlannedStartTime   *string                 `json:"plannedStartTime,omitempty"`
	ServerId           *string                 `json:"serverId,omitempty"`
	Status             *MaintenanceEventStatus `json:"status,omitempty"`
}

func (o *MaintenanceEventActionResponse) GetLastUpdatedTimeAsTime() (*time.Time, error) {
	if o.LastUpdatedTime == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.LastUpdatedTime, "2006-01-02T15:04:05Z07:00")
}

func (o *MaintenanceEventActionResponse) SetLastUpdatedTimeAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.LastUpdatedTime = &formatted
}

func (o *MaintenanceEventActionResponse) GetPlannedEndTimeAsTime() (*time.Time, error) {
	if o.PlannedEndTime == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.PlannedEndTime, "2006-01-02T15:04:05Z07:00")
}

func (o *MaintenanceEventActionResponse) SetPlannedEndTimeAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.PlannedEndTime = &formatted
}

func (o *MaintenanceEventActionResponse) GetPlannedStartTimeAsTime() (*time.Time, error) {
	if o.PlannedStartTime == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.PlannedStartTime, "2006-01-02T15:04:05Z07:00")
}

func (o *MaintenanceEventActionResponse) SetPlannedStartTimeAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.PlannedStartTime = &formatted
}
