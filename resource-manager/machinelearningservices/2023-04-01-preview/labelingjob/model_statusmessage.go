package labelingjob

import (
	"time"

	"github.com/hashicorp/go-azure-helpers/lang/dates"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type StatusMessage struct {
	Code            *string             `json:"code,omitempty"`
	CreatedDateTime *string             `json:"createdDateTime,omitempty"`
	Level           *StatusMessageLevel `json:"level,omitempty"`
	Message         *string             `json:"message,omitempty"`
}

func (o *StatusMessage) GetCreatedDateTimeAsTime() (*time.Time, error) {
	if o.CreatedDateTime == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.CreatedDateTime, "2006-01-02T15:04:05Z07:00")
}

func (o *StatusMessage) SetCreatedDateTimeAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.CreatedDateTime = &formatted
}
