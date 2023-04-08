package jitrequests

import (
	"time"

	"github.com/hashicorp/go-azure-helpers/lang/dates"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type JitSchedulingPolicy struct {
	Duration  string            `json:"duration"`
	StartTime string            `json:"startTime"`
	Type      JitSchedulingType `json:"type"`
}

func (o *JitSchedulingPolicy) GetStartTimeAsTime() (*time.Time, error) {
	return dates.ParseAsFormat(&o.StartTime, "2006-01-02T15:04:05Z07:00")
}

func (o *JitSchedulingPolicy) SetStartTimeAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.StartTime = formatted
}
