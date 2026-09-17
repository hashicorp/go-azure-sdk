package deletedbackupvaults

import (
	"time"

	"github.com/hashicorp/go-azure-helpers/lang/dates"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ResourceDeletionInfo struct {
	DeleteActivityId   *string `json:"deleteActivityId,omitempty"`
	DeletionTime       *string `json:"deletionTime,omitempty"`
	ScheduledPurgeTime *string `json:"scheduledPurgeTime,omitempty"`
}

func (o *ResourceDeletionInfo) GetDeletionTimeAsTime() (*time.Time, error) {
	if o.DeletionTime == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.DeletionTime, "2006-01-02T15:04:05Z07:00")
}

func (o *ResourceDeletionInfo) SetDeletionTimeAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.DeletionTime = &formatted
}

func (o *ResourceDeletionInfo) GetScheduledPurgeTimeAsTime() (*time.Time, error) {
	if o.ScheduledPurgeTime == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.ScheduledPurgeTime, "2006-01-02T15:04:05Z07:00")
}

func (o *ResourceDeletionInfo) SetScheduledPurgeTimeAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.ScheduledPurgeTime = &formatted
}
