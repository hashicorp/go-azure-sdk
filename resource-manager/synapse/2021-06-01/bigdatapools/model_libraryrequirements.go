package bigdatapools

import (
	"time"

	"github.com/hashicorp/go-azure-helpers/lang/dates"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type LibraryRequirements struct {
	Content  *string `json:"content,omitempty"`
	Filename *string `json:"filename,omitempty"`
	Time     *string `json:"time,omitempty"`
}

func (o *LibraryRequirements) GetTimeAsTime() (*time.Time, error) {
	if o.Time == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.Time, "2006-01-02T15:04:05Z07:00")
}

func (o *LibraryRequirements) SetTimeAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.Time = &formatted
}
