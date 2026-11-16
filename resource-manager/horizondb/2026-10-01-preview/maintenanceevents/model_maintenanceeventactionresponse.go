package maintenanceevents

import (
	"time"

	"github.com/hashicorp/go-azure-helpers/lang/dates"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MaintenanceEventActionResponse struct {
	AppliedNow          *bool                   `json:"appliedNow,omitempty"`
	LastUpdatedTimeUtc  *string                 `json:"lastUpdatedTimeUtc,omitempty"`
	MaintenanceEventId  *string                 `json:"maintenanceEventId,omitempty"`
	PlannedEndTimeUtc   *string                 `json:"plannedEndTimeUtc,omitempty"`
	PlannedStartTimeUtc *string                 `json:"plannedStartTimeUtc,omitempty"`
	PoolResourceId      *string                 `json:"poolResourceId,omitempty"`
	Status              *MaintenanceEventStatus `json:"status,omitempty"`
}

func (o *MaintenanceEventActionResponse) GetLastUpdatedTimeUtcAsTime() (*time.Time, error) {
	if o.LastUpdatedTimeUtc == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.LastUpdatedTimeUtc, "2006-01-02T15:04:05Z07:00")
}

func (o *MaintenanceEventActionResponse) SetLastUpdatedTimeUtcAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.LastUpdatedTimeUtc = &formatted
}

func (o *MaintenanceEventActionResponse) GetPlannedEndTimeUtcAsTime() (*time.Time, error) {
	if o.PlannedEndTimeUtc == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.PlannedEndTimeUtc, "2006-01-02T15:04:05Z07:00")
}

func (o *MaintenanceEventActionResponse) SetPlannedEndTimeUtcAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.PlannedEndTimeUtc = &formatted
}

func (o *MaintenanceEventActionResponse) GetPlannedStartTimeUtcAsTime() (*time.Time, error) {
	if o.PlannedStartTimeUtc == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.PlannedStartTimeUtc, "2006-01-02T15:04:05Z07:00")
}

func (o *MaintenanceEventActionResponse) SetPlannedStartTimeUtcAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.PlannedStartTimeUtc = &formatted
}
