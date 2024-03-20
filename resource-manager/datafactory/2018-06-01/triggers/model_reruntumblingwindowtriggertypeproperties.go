package triggers

import (
	"time"

	"github.com/hashicorp/go-azure-helpers/lang/dates"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RerunTumblingWindowTriggerTypeProperties struct {
	ParentTrigger      interface{} `json:"parentTrigger"`
	RequestedEndTime   string      `json:"requestedEndTime"`
	RequestedStartTime string      `json:"requestedStartTime"`
	RerunConcurrency   int64       `json:"rerunConcurrency"`
}

func (o *RerunTumblingWindowTriggerTypeProperties) GetRequestedEndTimeAsTime() (*time.Time, error) {
	return dates.ParseAsFormat(&o.RequestedEndTime, "2006-01-02T15:04:05Z07:00")
}

func (o *RerunTumblingWindowTriggerTypeProperties) SetRequestedEndTimeAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.RequestedEndTime = formatted
}

func (o *RerunTumblingWindowTriggerTypeProperties) GetRequestedStartTimeAsTime() (*time.Time, error) {
	return dates.ParseAsFormat(&o.RequestedStartTime, "2006-01-02T15:04:05Z07:00")
}

func (o *RerunTumblingWindowTriggerTypeProperties) SetRequestedStartTimeAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.RequestedStartTime = formatted
}
