package waitstatistics

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
