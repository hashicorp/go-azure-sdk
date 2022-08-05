package waitstatistics

import (
	"time"

	"github.com/hashicorp/go-azure-helpers/lang/dates"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WaitStatisticProperties struct {
	Count         *int64   `json:"count,omitempty"`
	DatabaseName  *string  `json:"databaseName,omitempty"`
	EndTime       *string  `json:"endTime,omitempty"`
	EventName     *string  `json:"eventName,omitempty"`
	EventTypeName *string  `json:"eventTypeName,omitempty"`
	QueryId       *int64   `json:"queryId,omitempty"`
	StartTime     *string  `json:"startTime,omitempty"`
	TotalTimeInMs *float64 `json:"totalTimeInMs,omitempty"`
	UserId        *int64   `json:"userId,omitempty"`
}

func (o *WaitStatisticProperties) GetEndTimeAsTime() (*time.Time, error) {
	if o.EndTime == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.EndTime, "2006-01-02T15:04:05Z07:00")
}

func (o *WaitStatisticProperties) SetEndTimeAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.EndTime = &formatted
}

func (o *WaitStatisticProperties) GetStartTimeAsTime() (*time.Time, error) {
	if o.StartTime == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.StartTime, "2006-01-02T15:04:05Z07:00")
}

func (o *WaitStatisticProperties) SetStartTimeAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.StartTime = &formatted
}
