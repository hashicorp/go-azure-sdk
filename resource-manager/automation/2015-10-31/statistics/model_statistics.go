package statistics

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Statistics struct {
	CounterProperty *string `json:"counterProperty,omitempty"`
	CounterValue    *int64  `json:"counterValue,omitempty"`
	EndTime         *string `json:"endTime,omitempty"`
	Id              *string `json:"id,omitempty"`
	StartTime       *string `json:"startTime,omitempty"`
}
