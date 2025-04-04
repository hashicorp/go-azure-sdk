package clusters

import (
	"time"

	"github.com/hashicorp/go-azure-helpers/lang/dates"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ClusterReplicationProperties struct {
	CreatedDate                *string                  `json:"createdDate,omitempty"`
	Enabled                    *bool                    `json:"enabled,omitempty"`
	IsAvailabilityZonesEnabled *bool                    `json:"isAvailabilityZonesEnabled,omitempty"`
	LastModifiedDate           *string                  `json:"lastModifiedDate,omitempty"`
	Location                   *string                  `json:"location,omitempty"`
	ProvisioningState          *ClusterReplicationState `json:"provisioningState,omitempty"`
}

func (o *ClusterReplicationProperties) GetCreatedDateAsTime() (*time.Time, error) {
	if o.CreatedDate == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.CreatedDate, "2006-01-02T15:04:05Z07:00")
}

func (o *ClusterReplicationProperties) SetCreatedDateAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.CreatedDate = &formatted
}

func (o *ClusterReplicationProperties) GetLastModifiedDateAsTime() (*time.Time, error) {
	if o.LastModifiedDate == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.LastModifiedDate, "2006-01-02T15:04:05Z07:00")
}

func (o *ClusterReplicationProperties) SetLastModifiedDateAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.LastModifiedDate = &formatted
}
