package policystates

import (
	"time"

	"github.com/hashicorp/go-azure-helpers/lang/dates"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ComponentStateDetails struct {
	ComplianceState *string `json:"complianceState,omitempty"`
	Id              *string `json:"id,omitempty"`
	Name            *string `json:"name,omitempty"`
	Timestamp       *string `json:"timestamp,omitempty"`
	Type            *string `json:"type,omitempty"`
}

func (o *ComponentStateDetails) GetTimestampAsTime() (*time.Time, error) {
	if o.Timestamp == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.Timestamp, "2006-01-02T15:04:05Z07:00")
}

func (o *ComponentStateDetails) SetTimestampAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.Timestamp = &formatted
}
