package triggers

import (
	"time"

	"github.com/hashicorp/go-azure-helpers/lang/dates"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PeriodicTimerSourceInfo struct {
	Schedule  string  `json:"schedule"`
	StartTime string  `json:"startTime"`
	Topic     *string `json:"topic,omitempty"`
}

func (o *PeriodicTimerSourceInfo) GetStartTimeAsTime() (*time.Time, error) {
	return dates.ParseAsFormat(&o.StartTime, "2006-01-02T15:04:05Z07:00")
}

func (o *PeriodicTimerSourceInfo) SetStartTimeAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.StartTime = formatted
}
