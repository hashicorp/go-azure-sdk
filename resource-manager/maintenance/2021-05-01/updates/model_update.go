package updates

import (
	"time"

	"github.com/hashicorp/go-azure-helpers/lang/dates"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Update struct {
	ImpactDurationInSec *int64            `json:"impactDurationInSec,omitempty"`
	ImpactType          *ImpactType       `json:"impactType,omitempty"`
	MaintenanceScope    *MaintenanceScope `json:"maintenanceScope,omitempty"`
	NotBefore           *string           `json:"notBefore,omitempty"`
	Properties          *UpdateProperties `json:"properties"`
	Status              *UpdateStatus     `json:"status,omitempty"`
}

func (o *Update) GetNotBeforeAsTime() (*time.Time, error) {
	if o.NotBefore == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.NotBefore, "2006-01-02T15:04:05Z07:00")
}

func (o *Update) SetNotBeforeAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.NotBefore = &formatted
}
