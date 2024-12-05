package incidenttasks

import (
	"time"

	"github.com/hashicorp/go-azure-helpers/lang/dates"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IncidentTaskProperties struct {
	CreatedBy           *ClientInfo        `json:"createdBy,omitempty"`
	CreatedTimeUtc      *string            `json:"createdTimeUtc,omitempty"`
	Description         *string            `json:"description,omitempty"`
	LastModifiedBy      *ClientInfo        `json:"lastModifiedBy,omitempty"`
	LastModifiedTimeUtc *string            `json:"lastModifiedTimeUtc,omitempty"`
	Status              IncidentTaskStatus `json:"status"`
	Title               string             `json:"title"`
}

func (o *IncidentTaskProperties) GetCreatedTimeUtcAsTime() (*time.Time, error) {
	if o.CreatedTimeUtc == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.CreatedTimeUtc, "2006-01-02T15:04:05Z07:00")
}

func (o *IncidentTaskProperties) SetCreatedTimeUtcAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.CreatedTimeUtc = &formatted
}

func (o *IncidentTaskProperties) GetLastModifiedTimeUtcAsTime() (*time.Time, error) {
	if o.LastModifiedTimeUtc == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.LastModifiedTimeUtc, "2006-01-02T15:04:05Z07:00")
}

func (o *IncidentTaskProperties) SetLastModifiedTimeUtcAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.LastModifiedTimeUtc = &formatted
}
