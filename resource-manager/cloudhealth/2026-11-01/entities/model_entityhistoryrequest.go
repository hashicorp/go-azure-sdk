package entities

import (
	"time"

	"github.com/hashicorp/go-azure-helpers/lang/dates"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EntityHistoryRequest struct {
	EndAt      *string `json:"endAt,omitempty"`
	NextMarker *string `json:"nextMarker,omitempty"`
	StartAt    *string `json:"startAt,omitempty"`
	Top        *int64  `json:"top,omitempty"`
}

func (o *EntityHistoryRequest) GetEndAtAsTime() (*time.Time, error) {
	if o.EndAt == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.EndAt, "2006-01-02T15:04:05Z07:00")
}

func (o *EntityHistoryRequest) SetEndAtAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.EndAt = &formatted
}

func (o *EntityHistoryRequest) GetStartAtAsTime() (*time.Time, error) {
	if o.StartAt == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.StartAt, "2006-01-02T15:04:05Z07:00")
}

func (o *EntityHistoryRequest) SetStartAtAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.StartAt = &formatted
}
