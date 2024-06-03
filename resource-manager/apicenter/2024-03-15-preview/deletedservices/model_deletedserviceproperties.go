package deletedservices

import (
	"time"

	"github.com/hashicorp/go-azure-helpers/lang/dates"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeletedServiceProperties struct {
	ScheduledPurgeDate *string `json:"scheduledPurgeDate,omitempty"`
	SoftDeletionDate   *string `json:"softDeletionDate,omitempty"`
}

func (o *DeletedServiceProperties) GetScheduledPurgeDateAsTime() (*time.Time, error) {
	if o.ScheduledPurgeDate == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.ScheduledPurgeDate, "2006-01-02T15:04:05Z07:00")
}

func (o *DeletedServiceProperties) SetScheduledPurgeDateAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.ScheduledPurgeDate = &formatted
}

func (o *DeletedServiceProperties) GetSoftDeletionDateAsTime() (*time.Time, error) {
	if o.SoftDeletionDate == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.SoftDeletionDate, "2006-01-02T15:04:05Z07:00")
}

func (o *DeletedServiceProperties) SetSoftDeletionDateAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.SoftDeletionDate = &formatted
}
