package applyupdate

import (
	"time"

	"github.com/hashicorp/go-azure-helpers/lang/dates"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ApplyUpdateProperties struct {
	LastUpdateTime *string       `json:"lastUpdateTime,omitempty"`
	ResourceId     *string       `json:"resourceId,omitempty"`
	Status         *UpdateStatus `json:"status,omitempty"`
}

func (o *ApplyUpdateProperties) GetLastUpdateTimeAsTime() (*time.Time, error) {
	if o.LastUpdateTime == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.LastUpdateTime, "2006-01-02T15:04:05Z07:00")
}

func (o *ApplyUpdateProperties) SetLastUpdateTimeAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.LastUpdateTime = &formatted
}
