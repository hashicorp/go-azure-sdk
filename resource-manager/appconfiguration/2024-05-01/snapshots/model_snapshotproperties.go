package snapshots

import (
	"time"

	"github.com/hashicorp/go-azure-helpers/lang/dates"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SnapshotProperties struct {
	CompositionType   *CompositionType   `json:"compositionType,omitempty"`
	Created           *string            `json:"created,omitempty"`
	Etag              *string            `json:"etag,omitempty"`
	Expires           *string            `json:"expires,omitempty"`
	Filters           []KeyValueFilter   `json:"filters"`
	ItemsCount        *int64             `json:"itemsCount,omitempty"`
	ProvisioningState *ProvisioningState `json:"provisioningState,omitempty"`
	RetentionPeriod   *int64             `json:"retentionPeriod,omitempty"`
	Size              *int64             `json:"size,omitempty"`
	Status            *SnapshotStatus    `json:"status,omitempty"`
	Tags              *map[string]string `json:"tags,omitempty"`
}

func (o *SnapshotProperties) GetCreatedAsTime() (*time.Time, error) {
	if o.Created == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.Created, "2006-01-02T15:04:05Z07:00")
}

func (o *SnapshotProperties) SetCreatedAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.Created = &formatted
}

func (o *SnapshotProperties) GetExpiresAsTime() (*time.Time, error) {
	if o.Expires == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.Expires, "2006-01-02T15:04:05Z07:00")
}

func (o *SnapshotProperties) SetExpiresAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.Expires = &formatted
}
