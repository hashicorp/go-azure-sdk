package containerappslabelhistory

import (
	"time"

	"github.com/hashicorp/go-azure-helpers/lang/dates"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type LabelHistoryRecordItem struct {
	Revision *string `json:"revision,omitempty"`
	Start    *string `json:"start,omitempty"`
	Status   *Status `json:"status,omitempty"`
	Stop     *string `json:"stop,omitempty"`
}

func (o *LabelHistoryRecordItem) GetStartAsTime() (*time.Time, error) {
	if o.Start == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.Start, "2006-01-02T15:04:05Z07:00")
}

func (o *LabelHistoryRecordItem) SetStartAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.Start = &formatted
}

func (o *LabelHistoryRecordItem) GetStopAsTime() (*time.Time, error) {
	if o.Stop == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.Stop, "2006-01-02T15:04:05Z07:00")
}

func (o *LabelHistoryRecordItem) SetStopAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.Stop = &formatted
}
