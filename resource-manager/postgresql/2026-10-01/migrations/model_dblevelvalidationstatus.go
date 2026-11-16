package migrations

import (
	"time"

	"github.com/hashicorp/go-azure-helpers/lang/dates"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DbLevelValidationStatus struct {
	DatabaseName *string                  `json:"databaseName,omitempty"`
	EndedOn      *string                  `json:"endedOn,omitempty"`
	StartedOn    *string                  `json:"startedOn,omitempty"`
	Summary      *[]ValidationSummaryItem `json:"summary,omitempty"`
}

func (o *DbLevelValidationStatus) GetEndedOnAsTime() (*time.Time, error) {
	if o.EndedOn == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.EndedOn, "2006-01-02T15:04:05Z07:00")
}

func (o *DbLevelValidationStatus) SetEndedOnAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.EndedOn = &formatted
}

func (o *DbLevelValidationStatus) GetStartedOnAsTime() (*time.Time, error) {
	if o.StartedOn == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.StartedOn, "2006-01-02T15:04:05Z07:00")
}

func (o *DbLevelValidationStatus) SetStartedOnAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.StartedOn = &formatted
}
