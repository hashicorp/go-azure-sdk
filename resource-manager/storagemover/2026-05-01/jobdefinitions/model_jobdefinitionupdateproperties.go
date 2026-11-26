package jobdefinitions

import (
	"time"

	"github.com/hashicorp/go-azure-helpers/lang/dates"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type JobDefinitionUpdateProperties struct {
	AgentName               *string                  `json:"agentName,omitempty"`
	Connections             *[]string                `json:"connections,omitempty"`
	CopyMode                *CopyMode                `json:"copyMode,omitempty"`
	DataIntegrityValidation *DataIntegrityValidation `json:"dataIntegrityValidation,omitempty"`
	Description             *string                  `json:"description,omitempty"`
	MoverSyncedUntil        *string                  `json:"moverSyncedUntil,omitempty"`
	Schedule                *ScheduleInfo            `json:"schedule,omitempty"`
	SyncMode                *string                  `json:"syncMode,omitempty"`
}

func (o *JobDefinitionUpdateProperties) GetMoverSyncedUntilAsTime() (*time.Time, error) {
	if o.MoverSyncedUntil == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.MoverSyncedUntil, "2006-01-02T15:04:05Z07:00")
}

func (o *JobDefinitionUpdateProperties) SetMoverSyncedUntilAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.MoverSyncedUntil = &formatted
}
