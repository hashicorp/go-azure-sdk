package restorepoints

import (
	"time"

	"github.com/hashicorp/go-azure-helpers/lang/dates"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RestorePointProperties struct {
	EarliestRestoreDate      *string           `json:"earliestRestoreDate,omitempty"`
	RestorePointCreationDate *string           `json:"restorePointCreationDate,omitempty"`
	RestorePointLabel        *string           `json:"restorePointLabel,omitempty"`
	RestorePointType         *RestorePointType `json:"restorePointType,omitempty"`
}

func (o *RestorePointProperties) GetEarliestRestoreDateAsTime() (*time.Time, error) {
	if o.EarliestRestoreDate == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.EarliestRestoreDate, "2006-01-02T15:04:05Z07:00")
}

func (o *RestorePointProperties) SetEarliestRestoreDateAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.EarliestRestoreDate = &formatted
}

func (o *RestorePointProperties) GetRestorePointCreationDateAsTime() (*time.Time, error) {
	if o.RestorePointCreationDate == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.RestorePointCreationDate, "2006-01-02T15:04:05Z07:00")
}

func (o *RestorePointProperties) SetRestorePointCreationDateAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.RestorePointCreationDate = &formatted
}
